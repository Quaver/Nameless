package utils

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/db"
	"io"
	"io/fs"
	"io/ioutil"
	"net/url"
	"os"
)

type AzureStorageClient struct {
	accountName string
	accountKey string
	credential azblob.Credential
	pipe pipeline.Pipeline
}

const accountName string = ""
const accountKey string = ""

// AzureClient Global storage client used throughout the application.
// Must call InitializeAzure to create
var AzureClient AzureStorageClient

// InitializeAzure Initializes the azure storage client
func InitializeAzure() {
	if AzureClient != (AzureStorageClient{}) {
		return
	}
	
	var err error
	
	client := AzureStorageClient{
		accountName: accountName,
		accountKey: accountKey,
	}
	
	client.credential, err = azblob.NewSharedKeyCredential(accountName, accountKey)
	
	if err != nil {
		panic(err)
	}

	client.pipe = azblob.NewPipeline(client.credential, azblob.PipelineOptions{})
	AzureClient = client
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
func (c *AzureStorageClient) DownloadFile(container string, name string, path string) (bytes.Buffer, error){
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
	const folder string = "C:\\Users\\Swan\\go\\src\\Nameless\\data\\maps"
	err := os.MkdirAll(folder, os.ModePerm)

	if err != nil {
		fmt.Println(err.Error())
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

	// Do a final hash check on the downloaded file
	if _, err := os.Stat(path); err == nil {
		md5, err := GetFileMD5(path)

		if err != nil {
			return "", err
		}

		if md5 != m.MD5 {
			return "", fmt.Errorf("md5 hash mismatch `%v` vs `%v`", md5, m.MD5)
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
