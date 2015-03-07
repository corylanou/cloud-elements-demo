package main

import (
	"os"

	"github.com/corylanou/cloud-elements"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	credentials := cloudElements.Credentials{
		User:         os.Getenv("CE_USER"),
		Organization: os.Getenv("CE_ORGANIZATION"),
		Elements:     make(map[int]string),
	}
	credentials.Elements[cloudElements.GOOGLE_DRIVE] = os.Getenv("CE_GOOGLEDRIVE")
	credentials.Elements[cloudElements.DROPBOX] = os.Getenv("CE_DROPBOX")
	spew.Dump(credentials)

	client := cloudElements.NewClient(credentials)

	googleFiles, err := client.CloudFiles.Contents("/", cloudElements.GOOGLE_DRIVE)
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(googleFiles)

	dropboxFiles, err := client.CloudFiles.Contents("/", cloudElements.DROPBOX)
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(dropboxFiles)

}
