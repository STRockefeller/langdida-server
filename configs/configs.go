package configs

import (
	"io"

	"gopkg.in/yaml.v3"
)

type YamlConfigs struct {
	/* ------------------------ choose one of delivery methods ------------------------ */
	UseGinServer bool `yaml:"useGinServer"`
	/* ------------------------- choose one of storage methods ------------------------ */
	UseSqlite   bool `yaml:"useSqlite"`
	UsePostgres bool `yaml:"usePostgres"`
	/* ---------------------------- delivery settings --------------------------- */
	GinServerSettings GinServerSettings `yaml:"ginServerSettings"`
	/* ---------------------------- storage settings ---------------------------- */
	SqliteSettings   SqliteSettings   `yaml:"sqliteSettings"`
	PostgresSettings PostgresSettings `yaml:"postgresSettings"`
}

type GinServerSettings struct {
	Port int `yaml:"port"`
}

type SqliteSettings struct {
	DBPath string `yaml:"dbPath"`
}

type PostgresSettings struct {
	ConnectionString string `yaml:"connectionString"`
}

func DecodeYaml(file io.Reader) (YamlConfigs, error) {
	var configs YamlConfigs
	if err := yaml.NewDecoder(file).Decode(&configs); err != nil {
		return YamlConfigs{}, err
	}
	return configs, nil
}
