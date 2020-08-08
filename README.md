# Shorty

![Go](https://github.com/farhaanbukhsh/shorty/workflows/Go/badge.svg?branch=master&event=pull_request)


URL shortner for the masses. This is the most basic URL shortner written, which probably shouldn't be used in production but can be run for your small server at your `localhost`.

## Usage 

``
go get github.com/farhaanbukhsh/shorty
``

If `GOPATH` is set properly  and added to `PATH`, we should have a utility called `shorty` in the path.

By default it uses `sqlite` as the database.

This is still a work in progress.

## Config

``
{
    "hostname": "http://localhost:8080/",
    "database": "shorty.db",
    "port": ":8080"
}
``

This is a config demo file that can be injected to modify the parameters.


*P.S*: Will be updating this repo with more information and features