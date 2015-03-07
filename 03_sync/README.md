# Syncing Files with Could Elemnents API

As before, I'm using the environment variables to load my credentials in the CLI.

To run this and list out files on root of Google Drive and Dropbox:

```sh
source dev.env
go run main.go -path=cloudelements
```


This will create a folder on google and drop box if they don't exist and sync files between them.  It will NOT overwrite.


Currently failing with the following error when posting a file to the `/files` endpoint:

```sh
$ go run main.go -path=cloudelements
2015-03-07T13:23:58-07:00 Syncing file "foo.txt" from DropBox to Google Drive...2015/03/07 13:23:59 
failed to write file foo.txt to DropBox:  &{Message:Failed to update file RequestId:54fb5d68e4b0e889a45c6eea StatusCode:500}
exit status 1
```
