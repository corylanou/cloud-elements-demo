# Cloud Elements Demo

The purpose of this project is to show how to leverage Go to consume API's, specifically, in this case, the [Cloud Elements](http://cloud-elements.com/) API.

## Demo

### Seting up Cloud Elements

Write something about how easy this was

### First pass at a basic connection

When using any api, the first thing I like to do is get my basic connection dance down.
I chose a very simple end point for this (`/folders`) to see if I could just list out the files
in my google drive.

Let's take a look at the [code](https://github.com/corylanou/cloud-elements-demo/blob/master/01_connecting/).

*WARNING*: It's UGLY!

### Creating a Library

#### The CloudFiles Library

#### Consuming the File Library

[code](https://github.com/corylanou/cloud-elements-demo/blob/master/02_consuming/).

#### Syncing different sources

## Learning Notes

### Cloud Elements Site

Whenever I use a new API, I always find it helpful to talk about things that I found challenging, and I think others would too.
These points are not meant to put down the API I'm using, but are to be used for me to give feedback, and maybe even contribute
to their system to makie it easier to adopt for future developers.  As an open source developer, this type of feedback
is invaluable.

- I wasn't able to find a "quick start" guide.  As someone writing an API, this creates a high barrier to entry.  This turned out to be due to me working from the consule, which does not provide links to the developer docs.
- I took me a while to find teh base url.  It was at the bottom of the API endpoint documentation for a specific endpoint.
- The session timeout of the site is too short.  As a developer, you are trying things constantly.  I kept getting logged out while I was developing.
- Finding information on how to authorized took quite some time. Most developers are looking for a curl example right away, and it was buried about 6 pages deep.
- Wrong curl example [here](http://cloud-elements.com/developer/full-api-documentation/):
```sh
curl -X POST 
-H 'Authorization: Element <INSERT ELEMENT TOKEN>, User <INSERT_USER_SECRET>' 
-F file=@test.txt 
'https://api.cloud-elements.com/elements/api-v2/hubs/documents/files?path=/test.txt&description=Test%20file&tags%5B%5D=test'
```
Appears we now need the organization as part of the authorization header as well.

No direction on what to post for creating a folder.  Error messages were very unclear as to what exactly they needed.
ex: Folder name provided and folder name in path do not match

Another fairly useless message encountered when trying to create a file:
Message:Failed to update file RequestId:54fb5e3de4b06d3806bf3809
Looking at the logs provided no more information.
