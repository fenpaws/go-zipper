package downloader

import (
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadFile downloads a file from the specified URL, saves it to the specified download folder
// with the specified file name, and returns an error if any occurred.
func DownloadFile(url string, downloadFolder string, fileName string) error {
	Debug := true //TODO: acces global config to determine debug status

	// Get the file from the specified URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Log the download process if debugging is enabled
	if Debug {
		log.Printf("Downloading file: %s in %s", fileName, downloadFolder)
	}

	// Open or create a file handle for the file in the specified download folder
	fileHandle, err := os.OpenFile(downloadFolder+"/"+fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	// Copy the file from the HTTP response body to the file handle
	_, err = io.Copy(fileHandle, resp.Body)
	if err != nil {
		return err
	}

	// Log the download completion if debugging is enabled
	if Debug {
		log.Printf("Download of file: %s in folder %s finished!", fileName, downloadFolder)
	}
	return nil
}
