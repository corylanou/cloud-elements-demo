package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

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

	client := cloudElements.NewClient(credentials)

	googleFiles, err := client.CloudFiles.Contents("/", cloudElements.GOOGLE_DRIVE)
	if err != nil {
		spew.Dump(err)
		return
	}
	writeCSV(googleFiles, "google")

	dropboxFiles, err := client.CloudFiles.Contents("/", cloudElements.DROPBOX)
	if err != nil {
		spew.Dump(err)
		return
	}
	writeCSV(dropboxFiles, "dropbox")

}

// Write the contents out nicely in columns
func writeCSV(files []cloudElements.CloudFile, source string) {
	fmt.Println()
	columnNames := []string{"source", "id", "name", "directory"}
	rows := []string{}

	rows = append(rows, strings.Join(columnNames, "\t"))

	for _, f := range files {
		row := []string{source, f.Id, f.Name, fmt.Sprintf("%v", f.Directory)}
		rows = append(rows, strings.Join(row, "\t"))
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	for _, r := range rows {
		fmt.Fprintln(w, r)
	}
	w.Flush()

}
