package main

import (
	_ "proxima/app/svc-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/svc-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
