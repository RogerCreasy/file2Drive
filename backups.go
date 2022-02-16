package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var config Config

// data struct for config json
type Config struct {
	FolderID, ServiceAccount, ServiceAccountAuth string
	Files                                        []string
}

// populate config var from the system configuration
func setConfig() {
	content, err := ioutil.ReadFile("/etc/backup2Drive/config.json")
	if err != nil {
		log.Fatal("Error opening config ", err)
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// for testing only
	log.Printf("folderID: %s\n", config.FolderID)
}

// function to write the file
func writeFile(service *drive.Service, name string, mimeType string, content io.Reader, driveID string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{driveID},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

func main() {
	ctx := context.Background()

	//populate the Global varable "config" with json data
	setConfig()

	srv, err := drive.NewService(ctx, option.WithCredentialsFile("/etc/backup2Drive/credentials.json"))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	// open the file
	for _, file := range config.Files {
		f, err := os.Open(file)
		if err != nil {
			panic(fmt.Sprintf("Unable to open file: %v", err))
		}

		defer f.Close()

		folderID := config.FolderID
		fileName := filepath.Base(f.Name())
		uploadedFile, err := writeFile(srv, fileName, "application/octet-stream", f, folderID)
		if err != nil {
			panic(fmt.Sprintf("Unable to write file to drive: %v\n", err))
		}

		fmt.Printf("File '%s' successfully uploaded", uploadedFile.Name)
	}
}
