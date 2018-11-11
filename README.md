
# slideshare-downloader 
An easy to deploy slideshare downloader in golang built on top of [gin](https://github.com/gin-gonic/gin) and [slideshare-go](https://github.com/mohan3d/slideshare-go).

## Installation
```sh
$ go get -u github.com/mohan3d/slideshare-downloader
```

## Run Locally
Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ cd $GOPATH/src/github.com/mohan3d/slideshare-downloader
$ heroku local web
```

## Deployment
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/mohan3d/slideshare-downloader)  

## Manual deployment
```sh
$ heroku create
$ git push heroku master
$ heroku open
```