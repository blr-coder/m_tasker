package configs

import "github.com/ilyakaznacheev/cleanenv"

type MongoDatabaseConfig struct {
	Server     string `env:"DB_SERVER" env-default:"mongodb://localhost:27017"`
	Name       string `env:"DB_NAME" env-default:"learn_mongo"`
	Collection string `env:"DB_COLLECTION" env-default:"tasks"`
}

func NewConfig() *MongoDatabaseConfig {
	cfg := new(MongoDatabaseConfig)
	_ = cleanenv.ReadEnv(cfg)
	return cfg

}

