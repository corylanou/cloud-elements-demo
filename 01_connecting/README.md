# Basic Connection

This file creates a basic connection, and lists the files.

It uses a local file to set up sensitive information, like tokens and secrets.

To do that, create a file called `dev.env` and add the following contents:

```sh
export CE_GOOGLEDRIVE="<token>"
export CE_DROPBOX="<token>"
export CE_USER="<token>"
export CE_ORGANIZATION="<token>"
```

Then run this set of commands:

```sh
source dev.env
go run main.go
```
