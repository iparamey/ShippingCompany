package rest

import (
	"cont"
	"db"
)

func GetContext() *cont.AppContext {
	config := &db.Config{}
	config.Default()

	context := &cont.AppContext{
		DB:    &db.DB{},
		Repo:  &cont.Repositories{},
	}

	context.DB.Initialize(config)
	context.Repo.Initialize(context.DB)
	return context
}

