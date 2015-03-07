package main

// Parse parameters

// Retreive list of files

// Compare to see what needs to be synced

// Actually sync the files

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/corylanou/cloud-elements"
)

type SyncCommands []SyncCommand
type SyncCommand struct {
	source      cloudElements.Provider
	destination cloudElements.Provider
	file        cloudElements.File
}

func main() {
	credentials := cloudElements.Credentials{
		User:         os.Getenv("CE_USER"),
		Organization: os.Getenv("CE_ORGANIZATION"),
		Elements:     make(map[cloudElements.Provider]string),
	}
	credentials.Elements[cloudElements.GOOGLE_DRIVE] = os.Getenv("CE_GOOGLEDRIVE")
	credentials.Elements[cloudElements.DROPBOX] = os.Getenv("CE_DROPBOX")

	client := cloudElements.NewClient(credentials)

	var path, sources string
	flag.StringVar(&path, "path", "/", "the path to sync files")
	flag.StringVar(&sources, "sources", "google,dropbox", "the sources to sync files, comma delimited.")
	flag.Parse()

	googleFiles, err := listAndCreateFolder(path, cloudElements.GOOGLE_DRIVE, client)
	if err != nil {
		log.Fatalf("failed to create/list folder %s - %s:  %+v", path, cloudElements.GOOGLE_DRIVE, err)
	}

	dropboxFiles, err := listAndCreateFolder(path, cloudElements.DROPBOX, client)
	if err != nil {
		log.Fatalf("failed to create/list folder %s - %s:  %+v", path, cloudElements.DROPBOX, err)
	}

	var sync SyncCommands
	for _, f := range googleFiles {
		ff := dropboxFiles.FindByName(f.Name)
		if ff == nil {
			sync = append(sync, SyncCommand{source: cloudElements.GOOGLE_DRIVE, destination: cloudElements.DROPBOX, file: f})
		}
	}

	for _, f := range dropboxFiles {
		ff := googleFiles.FindByName(f.Name)
		if ff == nil {
			sync = append(sync, SyncCommand{source: cloudElements.DROPBOX, destination: cloudElements.GOOGLE_DRIVE, file: f})
		}
	}

	for _, s := range sync {
		fmt.Printf("%s Syncing file %q from %s to %s...", time.Now().Format(time.RFC3339), s.file.Name, s.source, s.destination)
		r, err := client.Folders.GetFileById(s.file.Id, s.source)
		if err != nil {
			log.Fatalf("\nfailed to retrieve file %s from %s:  %+v", s.file.Name, s.source, err)
		}
		_, err = client.Folders.CreateFile(s.file, false, r, s.destination)
		if err != nil {
			log.Fatalf("\nfailed to write file %s to %s:  %+v", s.file.Name, s.destination, err)
		}

		fmt.Printf(" DONE!\n")
	}
}

func listAndCreateFolder(path string, provider cloudElements.Provider, client *cloudElements.Client) (cloudElements.Files, *cloudElements.Error) {
	files, err := client.Folders.Contents(path, provider)
	if err != nil {
		if err.StatusCode == http.StatusNotFound {
			// The folder doesn't exist, so we need to create it
			_, err := client.Folders.CreateFolder(path, cloudElements.DROPBOX)
			if err != nil {
				return cloudElements.Files{}, err
			}
		} else {
			return cloudElements.Files{}, err
		}
	}
	return files, nil
}
