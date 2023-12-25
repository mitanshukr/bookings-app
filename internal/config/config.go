package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	IsDev bool
	TemplateCache map[string]*template.Template
}

