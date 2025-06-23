package main

import (
	_ "shu/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shu/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
