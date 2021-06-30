package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
)

// GetFileMD5 Returns the MD5 hash of a file
func GetFileMD5(path string) (string, error) {
	if _, err := os.Stat(path); err != nil {
		return "", err
	} 
	
	file, err := os.Open(path)
	
	if err != nil {
		return "", err
	}
	
	defer func(file *os.File) {
		err := file.Close()
		
		if err != nil {
			return
		}
	}(file)

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		return "", err
	}
	
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// DownloadFile Downloads a file from a URL
func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()

		if err != nil {
			return
		}
	}(resp.Body)

	out, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer func(out *os.File) {
		err := out.Close()

		if err != nil {
			return
		}
	}(out)

	_, err = io.Copy(out, resp.Body)
	return err
}
