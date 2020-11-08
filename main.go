package main

import (
	"github.com/zorbyte/whiskey/cmds"
	"github.com/zorbyte/whiskey/handlers"
	"github.com/zorbyte/whiskey/lib"
)

func main() {
	w := lib.NewWhiskey()
	handlers.RegisterHandlers(w)
	cmds.RegisterCmds(w)
	w.Start()
}