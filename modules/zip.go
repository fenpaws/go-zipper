package modules

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type Zipper interface {
	Zip(folderPath string, password string) error
}

type goZipper struct {
}

type sevenZipper struct {
}

func NewSevenZipper() Zipper {
	return sevenZipper{}
}

func (s sevenZipper) Zip(folderPath string, password string) error {

	argString := []string{"a", folderPath + ".zip", "-r", "-p" + password, folderPath + "/*"}

	progPath := "7z"

	if runtime.GOOS == "windows" {
		progPath = "C:\\Program Files\\7-Zip\\7z.exe"
	}

	cmd := exec.Command(progPath, argString...)

	err := cmd.Run()

	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)

	}
	return err
}

func NewGoZipper() Zipper {
	return goZipper{}
}

func (g goZipper) Zip(folderPath string, password string) error {

	// creates a folder with the selected archive format
	file, err := os.Create(folderPath + ".zip")
	if err != nil {
		return err
	}
	defer file.Close() // makes sure that the zip file is closed at the end of the function

	writer := zip.NewWriter(file)
	defer writer.Close()

	// getting all files in the folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	// looping thought all files
	for _, file := range files {
		log.Printf("opening file %s...", file.Name())

		// reads the raw data of the file
		readerFile, err := os.Open(folderPath + "/" + file.Name())
		if err != nil {
			return err
		}
		defer readerFile.Close() // makes sure that the file is closed at the end of the function

		// creating file inside the archive
		log.Printf("creating file  %s in the archive...", file.Name())
		writerZip, err := writer.Create(file.Name())
		if err != nil {
			return nil
		}

		// copy the data to the archive
		log.Printf("copy file data of  %s to the archive...", file.Name())
		_, err = io.Copy(writerZip, readerFile)
		if err != nil {
			return err
		}

	}
	return nil
}
