package utils

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/Swan/Nameless/common"
	"github.com/Swan/Nameless/config"
	"github.com/Swan/Nameless/db"
	log "github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"io/ioutil"
	"net/url"
	"os"
)

type AzureStorageClient struct {
	accountName string
	accountKey  string
	credential  azblob.Credential
	pipe        pipeline.Pipeline
}

// AzureClient Global storage client used throughout the application.
// Must call InitializeAzure to create
var AzureClient AzureStorageClient
var ErrAzureMismatchedMD5 = errors.New("MD5 hash of cached file mismatches database")

// InitializeAzure Initializes the azure storage client
func InitializeAzure() {
	if AzureClient != (AzureStorageClient{}) {
		return
	}

	var err error

	client := AzureStorageClient{
		accountName: config.Data.Azure.AccountName,
		accountKey:  config.Data.Azure.AccountKey,
	}

	client.credential, err = azblob.NewSharedKeyCredential(client.accountName, client.accountKey)

	if err != nil {
		panic(err)
	}

	client.pipe = azblob.NewPipeline(client.credential, azblob.PipelineOptions{})
	AzureClient = client

	log.Info("Successfully created Azure client!")
}

// Create a ContainerURL object to be able to make requests on that container
func (c *AzureStorageClient) createContainerURL(container string) azblob.ContainerURL {
	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", c.accountName, container))
	return azblob.NewContainerURL(*URL, c.pipe)
}

// UploadFile Uploads a file to a given container
func (c *AzureStorageClient) UploadFile(container string, fileName string, data []byte) error {
	containerURL := c.createContainerURL(container)
	blobURL := containerURL.NewBlockBlobURL(fileName)
	ctx := context.Background()

	_, err := azblob.UploadBufferToBlockBlob(ctx, data, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16,
	})

	if err != nil {
		return err
	}

	return nil
}

// DownloadFile Downloads a blob from a given container
func (c *AzureStorageClient) DownloadFile(container string, name string, path string) (bytes.Buffer, error) {
	containerURL := c.createContainerURL(container)
	blobURL := containerURL.NewBlockBlobURL(name)
	ctx := context.Background()

	resp, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{},
		false, azblob.ClientProvidedKeyOptions{})

	if err != nil {
		return bytes.Buffer{}, err
	}

	bodyStream := resp.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

	defer func(bodyStream io.ReadCloser) {
		err := bodyStream.Close()
		if err != nil {
			return
		}
	}(bodyStream)

	// Read the body into a buffer
	downloadedData := bytes.Buffer{}
	_, err = downloadedData.ReadFrom(bodyStream)

	if err != nil {
		return bytes.Buffer{}, err
	}

	err = ioutil.WriteFile(path, downloadedData.Bytes(), fs.ModePerm)

	if err != nil {
		return bytes.Buffer{}, err
	}

	return downloadedData, nil
}

// CacheQuaFile Downloads a .qua file from the API to disk
func CacheQuaFile(m db.Map) (string, error) {
	folder := config.Data.MapCacheDir
	err := os.MkdirAll(folder, os.ModePerm)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to cache qua file #%v - %v", m.Id, err))
		return "", err
	}

	path := fmt.Sprintf("%v/%v.qua", folder, m.Id)
	needsDownload := false

	// If the file exists, check if the MD5 hash matches the DB
	if _, err := os.Stat(path); err == nil {
		md5, err := GetFileMD5(path)

		if err != nil {
			return "", err
		}

		if md5 != m.MD5 {
			needsDownload = true
		}
	} else {
		needsDownload = true
	}

	// Attempt to download the file from azure
	if needsDownload {
		buffer, err := AzureClient.DownloadFile("maps", fmt.Sprintf("%v.qua", m.Id), path)

		if err != nil {
			return "", err
		}

		// Maps are that are uploaded by donors are gzipped, so we need to unpack and rewrite the file
		if m.RankedStatus == common.StatusNotSubmitted {
			err = ungzipFile(&buffer, path)

			if err != nil {
				return "", err
			}
		}
	}

	if m.RankedStatus == common.StatusNotSubmitted {
		return path, nil
	}

	// Do a final Md5 hash check on ranked maps.
	if _, err := os.Stat(path); err == nil {
		md5, err := GetFileMD5(path)

		if err != nil {
			return "", err
		}

		if md5 != m.MD5 {
			return "", ErrAzureMismatchedMD5
		}

		return path, nil
	}

	return "", fmt.Errorf("failed to cache `%v.qua`", m.Id)
}

// Ungzips a file and rewrites it
func ungzipFile(buffer *bytes.Buffer, path string) error {
	reader, err := gzip.NewReader(buffer)

	defer func(reader *gzip.Reader) {
		_ = reader.Close()
	}(reader)

	if err != nil {
		return err
	}

	var unpacked bytes.Buffer
	_, err = unpacked.ReadFrom(reader)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, unpacked.Bytes(), fs.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
