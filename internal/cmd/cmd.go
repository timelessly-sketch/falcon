package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gtag"
)

var FC = cFC{}

type cFC struct {
	g.Meta `name:"falcon" ad:"{cFCAd}"`
}

const (
	cFCAd = `
ADDITIONAL
    Use "falcon COMMAND -h" for details about a command.
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cFCAd`: cFCAd,
	})
}

type cFCInput struct {
	g.Meta  `name:"falcon"`
	Version bool `short:"v" name:"version" brief:"show version information of current binary"   orphan:"true"`
	Debug   bool `short:"d" name:"debug"   brief:"show internal detailed debugging information" orphan:"true"`
}

type cFCOutput struct{}

func (c cFC) Index(ctx context.Context, in cFCInput) (out *cFCOutput, err error) {
	if in.Version {
		_, err = Version.Index(ctx, cVersionInput{})
		return
	}
	gcmd.CommandFromCtx(ctx).Print()
	return
}
