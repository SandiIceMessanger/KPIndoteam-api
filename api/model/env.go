package model

import "chatnews-api/lib/db"

type (
	EnvConfig struct {
		Host  string
		Port  int
		Mongo db.MongoConfig
	}
)
