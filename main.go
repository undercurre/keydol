package main

import (
	_ "keydol/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"keydol/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
