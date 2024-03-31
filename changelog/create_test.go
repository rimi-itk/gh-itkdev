package changelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://cs.opensource.google/go/go/+/refs/tags/go1.22.1:src/testing/run_example.go

func TestCreate(t *testing.T) {
	repositoryUrl := "https://example.com"
	expected := `# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

[Unreleased]: https://example.com
`
	actual, _ := createChangelog(repositoryUrl)

	assert.Equal(t, expected, actual)
}
