package changelog

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/cli/go-gh"
)

type PullRequest struct {
	Number int
	Title  string
	Url    string
}

func isChanged(name string) bool {
	cmd := exec.Command("git", "diff", "--exit-code", name)

	return cmd.Run() != nil
}

func getPullRequest() (PullRequest, error) {
	var pr PullRequest

	content, _, err := gh.Exec("pr", "view", "--json", "number,title,url")
	if err != nil {
		return pr, err
	}

	err = json.Unmarshal(content.Bytes(), &pr)

	return pr, err
}

func addPullRequest(name string, pr PullRequest, itemTemplate string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	headerPattern := regexp.MustCompile(`(?i)^\#+ +\[unreleased\]`)
	unreleasedHeaderIndex := -1
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if headerPattern.MatchString(line) {
			unreleasedHeaderIndex = len(lines)
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()

	if unreleasedHeaderIndex < 0 {
		log.Fatalf("cannot find \"Unreleased\" header in %s", name)
	}

	// Make sure that we have a blank line after the header
	if !(unreleasedHeaderIndex < len(lines)-2) {
		lines = append(lines, "")
	}
	insertIndex := unreleasedHeaderIndex + 2
	tmpl, err := template.New("item").Parse(itemTemplate)
	if err != nil {
		log.Fatalf("cannot parse item template: %s", err)
	}
	var builder strings.Builder
	tmpl.Execute(&builder, pr)
	item := builder.String()

	// If content right after insertion point is a header or a link, we insert a blank line
	if insertIndex < len(lines) && regexp.MustCompile(`^[\[$]`).MatchString(lines[insertIndex]) {
		item += "\n"
	}

	lines = slices.Concat(
		lines[0:insertIndex],
		[]string{item},
		lines[insertIndex:],
	)

	content := strings.Join(lines, "\n") + "\n"
	if err := os.WriteFile(name, []byte(content), 0644); err != nil {
		log.Fatal(err)
	}

	return item, nil
}

func FuckingChangelog(name string, itemTemplate string) {
	if _, err := os.Stat(name); err != nil {
		log.Fatalf("File %s does not exist", name)
	}
	if isChanged(name) {
		log.Fatalf("File %s is changed", name)
	}

	pr, err := getPullRequest()
	if err != nil {
		log.Fatalf("error getting pull request: %s\n", err)
	}

	content, err := addPullRequest(name, pr, itemTemplate)
	if err != nil {
		log.Fatalf("error adding pull request: %s\n", err)
	}

	fmt.Printf("Content\n\n%s\n\nadded to %s.\n", strings.TrimSpace(content), name)
}
