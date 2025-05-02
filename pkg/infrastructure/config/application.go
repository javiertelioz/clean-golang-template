package config

type Application struct {
	// Name is the name of the application
	Name string `mapstructure:"app_name" validate:"required"`
}
