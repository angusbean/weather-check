package config

import (
	"html/template"
	"log"

	"github.com/angusbean/weather-check/models"
	"github.com/go-redis/redis/v7"
)

//AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Cities        models.Cities
	RedisClient   redis.Client
}
