package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/halalala222/gf-learning/internal/logic"
	_ "github.com/halalala222/gf-learning/internal/packed"

	"github.com/halalala222/gf-learning/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
