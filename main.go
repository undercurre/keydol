package main

import (
	_ "keybol/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"keybol/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
