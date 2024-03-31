package changelog

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/cli/go-gh"
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
	repository, err := gh.CurrentRepository()
	if err != nil {
		return "", fmt.Errorf("error getting repository: %s", err)
	}

	url := fmt.Sprintf("https://%s/%s/%s", repository.Host(), repository.Owner(), repository.Name())

	return url, nil
}

func createChangelog(name string, repositoryUrl string) (string, error) {
	tmpl, err := template.New("changelog").Parse(changelogTemplate)
	if err != nil {
		return "", fmt.Errorf("error reading changelog template: %s", err)
	}

	file, err := os.Create(name)
	if err != nil {
		return "", fmt.Errorf("cannot create file %s (%s)", name, err)
	}
	defer file.Close()

	err = tmpl.Execute(file, map[string]string{
		"RepositoryUrl": repositoryUrl,
	})
	if err != nil {
		return "", fmt.Errorf("error rendering changelog template: %s", err)
	}

	return fmt.Sprintf("New changelog written to %s", name), nil
}

func Create(name string) {
	if _, err := os.Stat(name); err == nil {
		log.Fatalf("File %s already exist", name)
	}

	repositoryUrl, err := getRepositoryUrl()
	if err != nil {
		log.Fatalf("error getting repository url: %s", err)
	}

	message, err := createChangelog(name, repositoryUrl)
	if err != nil {
		log.Fatalf("error creating changelog %s: %s", name, err)
	}

	fmt.Println(message)
}
