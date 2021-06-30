package utils

import (
	"fmt"
	"github.com/Swan/Nameless/src/db"
	"os"
)

const folder string = "C:\\Users\\Swan\\go\\src\\Nameless\\data\\maps"
const apiUrl string = "https://api.quavergame.com"

// CacheQuaFile Downloads a .qua file from the API to disk
func CacheQuaFile(m db.Map) (string, error) {
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

	// Attempt to download the file from the API if needed
	if needsDownload {
		url := fmt.Sprintf("%v/d/web/map/%v", apiUrl, m.Id)
		err = DownloadFile(path, url)

		if err != nil {
			return "", err
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
