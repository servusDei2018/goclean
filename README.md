# goclean ![](https://img.shields.io/github/downloads/servusdei2018/goclean/total?style=flat-square) ![](https://img.shields.io/github/issues-raw/servusdei2018/goclean?style=flat-square) ![](https://img.shields.io/badge/go-1.15-critical?style=flat-square&logo=go)
Aimed at go developers, goclean sifts through complex directory structures and removes useless files to save space.

## About
Whenever `go get` is used to grab a package, there are typical `LICENSE`, `AUTHORS`, `.travis.yml` files and the like which take up space but have no impact on the go package itself. goclean searches for these files and removes them, freeing up disk space.

## Installation

Head over to the ![Releases page](https://github.com/servusDei2018/goclean/releases/latest) and download the latest build or run the following command below:

`go install github.com/servusDei2018/goclean`

## Usage

`goclean` ~ Automatically clean up $GOPATH

### Flags

##### --path
The `--path=<string>` flag may be used to specify the path to clean, like so:

`goclean --path="/home/me/go"` ~ Specify path to clean

##### --dry-run
The `--dry-run` flag may be used to preview files that would be deleted (but do nothing), like so:

`goclean --dry-run`


