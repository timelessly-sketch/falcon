package cmd

import (
	"bytes"
	"context"
	"falcon/internal/utility/mlog"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

var (
	Version = cVersion{}
)

const (
	defaultIndent = "{{indent}}"
)

type cVersion struct {
	g.Meta `name:"version" brief:"show version information of current binary"`
}

type cVersionInput struct {
	g.Meta `name:"version"`
}

type cVersionOutput struct{}

func (c cVersion) Index(ctx context.Context, in cVersionInput) (*cVersionOutput, error) {
	detailBuffer := &detailBuffer{}
	detailBuffer.WriteString("1.1.1")

	detailBuffer.appendLine(0, "Welcome to GoFrame!")
	mlog.Print(detailBuffer.replaceAllIndent("  "))
	return nil, nil
}

// detailBuffer is a buffer for detail information.
type detailBuffer struct {
	bytes.Buffer
}

// appendLine appends a line to the buffer with given indent level.
func (d *detailBuffer) appendLine(indentLevel int, line string) {
	d.WriteString(fmt.Sprintf("\n%s%s", strings.Repeat(defaultIndent, indentLevel), line))
}

// replaceAllIndent replaces the tab with given indent string and prints the buffer content.
func (d *detailBuffer) replaceAllIndent(indentStr string) string {
	return strings.ReplaceAll(d.String(), defaultIndent, indentStr)
}
