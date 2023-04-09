package main

import (
	_ "github.com/halalala222/gf-learning/internal/logic"
	_ "github.com/halalala222/gf-learning/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/halalala222/gf-learning/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
