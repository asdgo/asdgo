package asdgo

import (
	"github.com/asdgo/asdgo/database"
	"github.com/asdgo/asdgo/decoder"
	"github.com/asdgo/asdgo/router"
	"github.com/asdgo/asdgo/session"
	"github.com/asdgo/asdgo/template"
	"github.com/asdgo/asdgo/validator"
	"github.com/joho/godotenv"
)

func New() {
	godotenv.Load()

	database.New()
	router.New()
	template.New()
	session.New()

	decoder.New()
	validator.New()
}

func DB() *database.Database {
	return database.Instance
}

func Router() *router.Router {
	return router.Instance
}

func Template() *template.Template {
	return template.Instance
}

func Session() *session.Session {
	return session.Instance
}

func Validator() *validator.Validator {
	return validator.Instance
}

func Decoder() *decoder.Decoder {
	return decoder.Instance
}
