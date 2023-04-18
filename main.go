package main

import (
	"flag"
	"os"

	"github.com/STRockefeller/langdida-server/configs"
	"github.com/STRockefeller/langdida-server/delivery/ginserver"
	"github.com/STRockefeller/langdida-server/internal/logger"
	"github.com/STRockefeller/langdida-server/storage"
	"github.com/STRockefeller/langdida-server/storage/sqlite"
)

func main() {
	/* ---------------------------- initialize logger --------------------------- */
	logger.InitialLogger()
	/* ------------------------------- parse flags ------------------------------ */
	configPath := flag.String("configs", "./assets/configs.yaml", "specify the config file, refer to configs.md")
	flag.Parse()

	/* ----------------------------- decode configs ----------------------------- */
	r, err := os.Open(*configPath)
	if err != nil {
		panic(err)
	}
	conf, err := configs.DecodeYaml(r)
	if err != nil {
		panic(err)
	}

	/* -------------------------------------------------------------------------- */
	/*                                  main part                                 */
	/* -------------------------------------------------------------------------- */

	/* --------------------------- initialize storage --------------------------- */
	var storageInstance storage.Storage
	switch {
	case conf.UseSqlite:
		storageInstance = sqlite.NewStorage(conf.SqliteSettings.DBPath, true)
	default:
		panic("unspecified storage")
	}

	/* ------------------------------ run delivery ------------------------------ */
	if conf.UseGinServer {
		ginserver.Run(conf.GinServerSettings.Port, storageInstance)
	}
}
