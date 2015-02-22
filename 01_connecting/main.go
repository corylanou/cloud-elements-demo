package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/davecgh/go-spew/spew"
)

const (
	baseURL = "console.cloud-elements.com/elements/api-v2"

	GOOGLE_DRIVE = iota
	DROPBOX
)

func main() {
	credentials := Credentials{
		User:         os.Getenv("CE_USER"),
		Organization: os.Getenv("CE_ORGANIZATION"),
		Elements:     make(map[int]string),
	}
	credentials.Elements[GOOGLE_DRIVE] = os.Getenv("CE_GOOGLEDRIVE")
	credentials.Elements[DROPBOX] = os.Getenv("CE_DROPBOX")
	spew.Dump(credentials)

	u := mustEndpoint("/hubs/documents/folders/contents")
	query := url.Values{}
	query.Set("path", "/")
	u.RawQuery = query.Encode()

	client := http.Client{}
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Authorization", credentials.Authorization(GOOGLE_DRIVE))

	response, err := client.Do(req)

	if err != nil {
		log.Println("Error retreiving url:", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		var e Error
		if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
			panic(err)
		}
		spew.Dump(e)

		return
	}
	var cloudFiles CloudFiles

	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&cloudFiles); err != nil {
		panic(err)
	}
	spew.Dump(cloudFiles)

}

func mustEndpoint(endPath string) url.URL {
	endPath = path.Join(baseURL, endPath)
	return url.URL{
		Path:   endPath,
		Scheme: "https",
	}
}

// Models

type Credentials struct {
	Elements     map[int]string
	Organization string
	User         string // User Secret
}

// Authorization is formatted as:
// User <secret>, Organization <token>, Element <token>
func (c Credentials) Authorization(element int) string {
	return fmt.Sprintf(`User %s, Organization %s, Element %s`, c.User, c.Organization, c.Elements[element])
}

type CloudFiles []CloudFile
type CloudFile struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	CreateDate   string   `json:"createDate"`
	Directory    bool     `json:"directory"`
	ModifiedDate string   `json:"modifiedDate"`
	Size         int      `json:"size"`
	Tags         []string `json:"tags"`
}

type Error struct {
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}
