# goclone
Automatically clones a repo to the appropriate location for your $GOPATH.

## Usage
```
$ go get github.com/kusold/goclone
$ goclone git@repourl
```

## Example
```
~GOPATH/src/github.com/kusold
❯ goclone git@github.com:kusold/goclone.git
Cloning into '/home/mike/Development/go/src/github.com/kusold/goclone'...
remote: Counting objects: 9, done.
remote: Compressing objects: 100% (7/7), done.
remote: Total 9 (delta 0), reused 6 (delta 0), pack-reused 0
Receiving objects: 100% (9/9), done.
Checking connectivity... done.

~GOPATH/src/github.com/kusold
❯ ls $GOPATH/src/github.com/kusold/goclone
LICENSE  main.go
```
