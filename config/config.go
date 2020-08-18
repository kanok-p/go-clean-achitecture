package config

import "github.com/caarlos0/env/v6"

type Config struct {
	MongoDBEndpoint string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName     string `env:"MONGODB_NAME,required"`
	MongoDBCollUser string `env:"MONGODB_COLL_USERS,required"`
}

func Get() *Config {
	conf := &Config{}
	if err := env.Parse(conf); err != nil {
		panic(err)
	}

	return conf
}
