package modules

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(url string, downloadFolder string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("Downloading file: %s in %s", fileName, downloadFolder)

	fileHandle, err := os.OpenFile(downloadFolder+"/"+fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	_, err = io.Copy(fileHandle, resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Download of file: %s in folder %s finished!", fileName, downloadFolder)
	return nil
}
