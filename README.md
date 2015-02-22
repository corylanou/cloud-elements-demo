# Cloud Elements Demo

The purpose of this project is to show how to leverage Go to consume API's, specifically, in this case, the [Cloud Elements](http://cloud-elements.com/) API.


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
