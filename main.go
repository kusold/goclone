package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	GitURL string
	Follow bool
}

func (c *Command) MakeGitDir() error {
	return os.MkdirAll(ParseGitPath(c.GitURL), 0755)
}

func (c *Command) Clone() error {
	_, err := exec.LookPath("git")
	if err != nil {
		return errors.New("git must be installed")
	}

	cmd := exec.Command("git", "clone", c.GitURL, ParseGitPath(c.GitURL))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func (c *Command) Chdir() error {
	// Currently doesn't work. Somehow I need to jump to the parent process
	return os.Chdir(ParseGitPath(c.GitURL))
}

func ParseGitPath(path string) string {
	// git@github.com:Kusold/goclone.git
	path = strings.TrimPrefix(path, "git@")
	path = strings.Replace(path, ":", "/", 1)
	path = strings.TrimSuffix(path, ".git")

	gopath := os.Getenv("GOPATH")

	combine := []string{gopath, "src", path}
	path = strings.Join(combine, "/")
	return path
}

func main() {
	cmd := Command{}

	followDesc := "(WIP) After the clone, change to the new directory"
	flag.BoolVar(&cmd.Follow, "follow", false, followDesc)
	flag.BoolVar(&cmd.Follow, "f", false, followDesc+" (shorthand)")
	flag.Parse()

	args := flag.Args()
	cmd.GitURL = args[len(args)-1]

	if err := cmd.MakeGitDir(); err != nil {
		log.Fatal("Problem creating the directory", err)
	}

	if err := cmd.Clone(); err != nil {
		log.Fatal("Problem cloning the repo: ", err)
	}

	if cmd.Follow {
		err := cmd.Chdir()
		if err != nil {
			log.Fatal("Unable to follow: ", err)
		}
	}

}
