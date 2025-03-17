package main

import (
	"falcon/root"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx          = gctx.GetInitCtx()
		command, err = root.GetCommand(ctx)
	)

	if err != nil {
		g.Log().Fatalf(ctx, "%+v", err)
	}
	if command == nil {
		panic(gerror.New(`retrieve root command failed for "falcon"`))
	}
	command.Run(ctx)
}
