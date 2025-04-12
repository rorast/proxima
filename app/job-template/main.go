package main

import (
	_ "proxima/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
