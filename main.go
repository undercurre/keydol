package main

import (
	"keybol/internal/cmd"
	_ "keybol/internal/logic"
	_ "keybol/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
