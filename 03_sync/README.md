# Syncing Files with Could Elemnents API

As before, I'm using the environment variables to load my credentials in the CLI.

To run this and list out files on root of Google Drive and Dropbox:

```sh
source dev.env
go run main.go -path=cloudelements
```


This will create a folder on google and drop box if they don't exist and sync files between them.  It will NOT overwrite.
