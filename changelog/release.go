package changelog

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strings"
	"time"
)

func createRelease(release string, base string) (string, error) {
	branch := "release/" + release
	cmd := exec.Command("git", "checkout", "-b", branch, base)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %s", output, err)
	}

	return branch, nil
}

func updateChangelog(name string, release string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	headerPattern := regexp.MustCompile(`(?i)^\#+ +\[unreleased\]`)
	unreleasedHeaderIndex := -1
	linkUnreleasedPattern := regexp.MustCompile(`(?i)^\[unreleased\]: (.+)`)
	linkUnreleasedIndex := -1

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if headerPattern.MatchString(line) {
			unreleasedHeaderIndex = len(lines)
		} else if linkUnreleasedPattern.MatchString(line) {
			linkUnreleasedIndex = len(lines)
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()

	if unreleasedHeaderIndex < 0 {
		log.Fatalf("cannot find \"Unreleased\" header in %s", name)
	} else if linkUnreleasedIndex < 0 {
		log.Fatalf("cannot find \"Unreleased\" link in %s", name)
	} else if linkUnreleasedIndex < unreleasedHeaderIndex {
		log.Fatalf("\"Unreleased\" link must come after header in %s", name)
	}

	// ----------------------------------------------------------------

	insertIndex := linkUnreleasedIndex + 1

	match := linkUnreleasedPattern.FindStringSubmatch(lines[linkUnreleasedIndex])
	url, err := url.Parse(match[1])
	if err != nil {
		log.Fatalf("invalid url: %s", match[1])
	}
	url.Path = fmt.Sprintf("compare/%s...HEAD", release)
	unreleasedLink := fmt.Sprintf("[Unreleased]: %s", url)

	if insertIndex < len(lines) {
		line := lines[insertIndex]
		var previousRelease string
		if index := strings.LastIndex(line, "..."); index > -1 {
			previousRelease = line[index+3:]
		} else if index := strings.LastIndex(line, "/"); index > -1 {
			previousRelease = line[index+1:]
		} else {
			log.Fatalf("cannot find previous version from %q", line)
		}
		url.Path = fmt.Sprintf("compare/%s...%s", previousRelease, release)
	} else {
		url.Path = fmt.Sprintf("releases/tag/%s", release)
	}

	releaseLink := fmt.Sprintf("[%s]: %s", release, url)

	lines = slices.Concat(
		lines[0:insertIndex-1],
		[]string{
			unreleasedLink,
			releaseLink,
		},
		lines[insertIndex:],
	)

	// ----------------------------------------------------------------

	insertIndex = unreleasedHeaderIndex + 2

	item := fmt.Sprintf("## [%s] - %s\n", release, time.Now().Format("2006-01-02"))

	lines = slices.Concat(
		lines[0:insertIndex],
		[]string{item},
		lines[insertIndex:],
	)

	// ----------------------------------------------------------------

	content := strings.Join(lines, "\n") + "\n"
	if err := os.WriteFile(name, []byte(content), 0644); err != nil {
		log.Fatal(err)
	}

	return item, nil
}

func Release(release string, base string, name string) {
	branch, err := createRelease(release, base)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Branch %s created\n", branch)

	updateChangelog(name, release)

	cmd := exec.Command("git", "diff", name)
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}
