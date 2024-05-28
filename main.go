package main

import (
	"flag"
	"fmt"
	"os"

	questionnaire "github.com/STRockefeller/config-questionnaire"
	"github.com/STRockefeller/langdida-server/configs"
	"github.com/STRockefeller/langdida-server/delivery/ginserver"
	"github.com/STRockefeller/langdida-server/internal/logger"
	"github.com/STRockefeller/langdida-server/storage"
	"github.com/STRockefeller/langdida-server/storage/sqlite"
	"gopkg.in/yaml.v3"
)

func main() {
	/* ---------------------------- initialize logger --------------------------- */
	logger.InitialLogger()
	/* ------------------------------- parse flags ------------------------------ */
	configPath := flag.String("configs", "./assets/configs.yaml", "specify the config file, refer to configs.md")
	flag.Parse()

	/* ----------------------------- decode configs ----------------------------- */
	var conf configs.YamlConfigs
	r, err := os.Open(*configPath)
	if err != nil {
		fmt.Println("configs file not found, generating...")
		conf, err = questionnaire.GenerateAndRunQuestionnaire[configs.YamlConfigs]()
		if err != nil {
			panic(err)
		}
		// save the config file to the specified path
		f, err := os.Create(*configPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		encoder := yaml.NewEncoder(f)
		encoder.SetIndent(2)
		err = encoder.Encode(conf)
		if err != nil {
			panic(err)
		}
		fmt.Println("configs file generated")
	} else {
		conf, err = configs.DecodeYaml(r)
		if err != nil {
			panic(err)
		}
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
