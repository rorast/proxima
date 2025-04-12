package main

import (
	_ "proxima/app/api-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/api-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
