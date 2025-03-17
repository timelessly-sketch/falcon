package root

import (
	"context"
	"falcon/internal/cmd"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

type Command struct {
	*gcmd.Command
}

func (c *Command) Run(ctx context.Context) {
	defer func() {
		if exception := recover(); exception != nil {
			if err, ok := exception.(error); ok {
				g.Log().Info(ctx, err.Error())
			} else {
				panic(gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception))
			}
		}
	}()

	// CLI configuration, using the `hack/config.yaml` in priority.
	//if path, _ := gfile.Search(cliFolderName); path != "" {
	//	if adapter, ok := g.Cfg().GetAdapter().(*gcfg.AdapterFile); ok {
	//		if err := adapter.SetPath(path); err != nil {
	//			mlog.Fatal(err)
	//		}
	//	}
	//}

	// just run.
	if err := c.RunWithError(ctx); err != nil {
		g.Log().Fatalf(ctx, "%+v", err)
	}
}

func GetCommand(ctx context.Context) (*Command, error) {
	root, err := gcmd.NewFromObject(cmd.FC)
	if err != nil {
		return nil, err
	}
	err = root.AddObject()
	if err != nil {
		return nil, err
	}
	command := &Command{
		root,
	}
	return command, nil
}
