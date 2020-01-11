package extension

import (
	"testing"

	"github.com/anytypeio/goldmark"
	"github.com/anytypeio/goldmark/renderer/html"
	"github.com/anytypeio/goldmark/testutil"
)

func TestTaskList(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			TaskList,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/tasklist.txt", t)
}
