package changelog

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// Cf. https://raw.githubusercontent.com/olivierlacan/keep-a-changelog/main/CHANGELOG.md
const changelogTemplate = `
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

[Unreleased]: {{ .RepositoryUrl }}
`

func getRepositoryUrl() (string, error) {
	remote := "origin"
	cmd := exec.Command("git", "remote", "get-url", remote)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error getting git remote %s (%s)", remote, err)
	}

	// Remove user from URL
	url, err := url.Parse(strings.TrimSpace(string(output)))
	if err != nil {
		return "", err
	}
	url.User = nil

	return url.String(), nil
}

func Create(name string) {
	if _, err := os.Stat(name); err == nil {
		log.Fatalf("File %s already exist", name)
	}

	remote, err := getRepositoryUrl()
	if err != nil {
		log.Fatalf("error getting repository url: %s", err)
	}

	tmpl, err := template.New("changelog").Parse(changelogTemplate)
	if err != nil {
		log.Fatalf("error reading changelog template: %s", err)
	}

	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("cannot create file %s (%s)", name, err)
	}

	err = tmpl.Execute(file, map[string]string{
		"RepositoryUrl": remote,
	})
	if err != nil {
		log.Fatalf("error rendering changelog template: %s", err)
	}

	fmt.Printf("New changelog written to %s\n", name)
}
