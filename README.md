# goclone
Automatically clones a repo to the appropriate location for your $GOPATH.

## Usage
```
$ go get github.com/kusold/goclone
$ goclone git@repourl
```

To pass options to the `git clone` just put them after a `--`:
```
$ goclone -- --quiet git@repourl
```

## Example
```
~
❯ goclone git@github.com:kusold/goclone.git
Cloning into '/home/mike/Development/go/src/github.com/kusold/goclone'...
remote: Counting objects: 18, done.
remote: Compressing objects: 100% (13/13), done.
remote: Total 18 (delta 5), reused 13 (delta 3), pack-reused 0
Receiving objects: 100% (18/18), done.
Resolving deltas: 100% (5/5), done.
Checking connectivity... done.

~
❯ ls $GOPATH/src/github.com/kusold/goclone
LICENSE  main.go  README.md

~
❯ rm -rf $GOPATH/src/github.com/kusold/goclone

~
❯ goclone -- -q git@github.com:kusold/goclone.git

~
❯ 
```
