package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	flag "github.com/ogier/pflag"
)

type Command struct {
	GitURL  string
	GitArgs []string
	Follow  bool
}

func (c *Command) MakeGitDir() error {
	return os.MkdirAll(ParseGitPath(c.GitURL), 0755)
}

func (c *Command) Clone() error {
	_, err := exec.LookPath("git")
	if err != nil {
		return errors.New("git must be installed")
	}

	args := []string{"clone"}
	args = append(args, c.GitArgs...)
	args = append(args, c.GitURL, ParseGitPath(c.GitURL))

	cmd := exec.Command("git", args...)
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
	flag.BoolVarP(&cmd.Follow, "follow", "f", false, followDesc)
	flag.Parse()

	args := flag.Args()
	cmd.GitArgs = args[:len(args)-1]
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
