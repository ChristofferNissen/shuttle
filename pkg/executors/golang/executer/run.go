package executer

import (
	"context"

	"github.com/lunarway/shuttle/pkg/config"
	"github.com/lunarway/shuttle/pkg/ui"
)

func Run(
	ctx context.Context,
	ui *ui.UI,
	c *config.ShuttleProjectContext,
	path string,
	args ...string,
) error {
	if !isActionsEnabled() {
		ui.Verboseln("shuttle golang actions disabled")
		return nil
	}

	binaries, err := prepare(ctx, ui, path, c)
	if err != nil {
		ui.Errorln("failed to run command: %v", err)
		return err
	}

	ui.Verboseln("executing shuttle golang actions")
	if err := executeAction(ctx, binaries, args...); err != nil {
		return err
	}

	return nil
}
