package main_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kusold/goclone"
)

var expectedGitPath = strings.Join([]string{os.Getenv("GOPATH"), "src", "github.com", "kusold", "goclone"}, "/")
var parseGitPathTests = []struct {
	in  string
	out string
}{
	{"git@github.com:kusold/goclone.git", expectedGitPath},
	{"ssh://git@github.com:22/kusold/goclone.git", expectedGitPath},
	{"ssh://git@github.com:22/kusold/goclone.git/", expectedGitPath},
	{"git://github.com:22/kusold/goclone.git", expectedGitPath},
	{"http://github.com:22/kusold/goclone.git", expectedGitPath},
	{"https://github.com:22/kusold/goclone.git", expectedGitPath},
	{"ftp://github.com:22/kusold/goclone.git", expectedGitPath},
	{"ftps://github.com:22/kusold/goclone.git", expectedGitPath},
	{"git@github.com:kusold/goclone.git", expectedGitPath},
}

func TestParseGitPath(t *testing.T) {
	for _, test := range parseGitPathTests {
		actual := main.ParseGitPath(test.in)
		if actual != test.out {
			t.Error("Expected:", test.out, "Actual:", actual)
		}
	}
}
