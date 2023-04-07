package main

import (
	"github.com/STRockefeller/langdida-server/delivery/ginserver"
	"github.com/STRockefeller/langdida-server/storage/sqlite"
)

// todo: specify sql from flags or configs
func main() {
	ginserver.Run(80, sqlite.NewStorage("./assets/sqlite.db"))
}
