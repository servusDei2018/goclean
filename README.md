# goclean
Aimed at go developers, goclean sifts through complex directory structures and removes useless files to save space.

## About
Whenever `go get` is used to grab a package, there are typical `LICENSE`, `AUTHORS`, `.travis.yml` files and the like which take up space but have no impact on the go package itself. goclean searches for these files and removes them, freeing up disk space.

## Installation

`go install github.com/servusDei2018/goclean`

## Usage

`goclean` ~ 
