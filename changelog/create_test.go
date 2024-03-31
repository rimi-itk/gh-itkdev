package changelog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://cs.opensource.google/go/go/+/refs/tags/go1.22.1:src/testing/run_example.go

func TestCreate(t *testing.T) {
	// @see https://stackoverflow.com/a/10476304
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	name := strings.TrimRight(os.TempDir(), "/") + "/changelog_test.md"
	defer os.Remove(name)
	Create(name)
	assert.FileExists(t, name)
	os.Remove(name)

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	assert.Equal(t, fmt.Sprintf("New changelog written to %s\n", name), out)
}

// func TestCreateAlreadyExists(t *testing.T) {
// 	// @see https://stackoverflow.com/a/10476304
// 	old := os.Stdout // keep backup of the real stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	name := os.TempDir() + "changelog_test.md"
// 	os.Create(name)
// 	defer os.Remove(name)
// 	Create(name)
// 	assert.FileExists(t, name)
// 	os.Remove(name)

// 	outC := make(chan string)
// 	// copy the output in a separate goroutine so printing can't block indefinitely
// 	go func() {
// 		var buf bytes.Buffer
// 		io.Copy(&buf, r)
// 		outC <- buf.String()
// 	}()

// 	// back to normal state
// 	w.Close()
// 	os.Stdout = old // restoring the real stdout
// 	out := <-outC

// 	assert.Equal(t, fmt.Sprintf("New changelog written to %s\n", name), out)
// }
