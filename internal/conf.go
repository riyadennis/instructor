package internal

import (
	"github.com/alexflint/go-arg"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type Config struct {
	OpenAIKey     string `arg:"env:OPENAI_API_KEY" validate:"required"`
	APIMaxRetries int    `arg:"env:MAX_RETRIES" default:"3"`
}

// NewConfig return a new instance of Config
func NewConfig() (Config, error) {
	config := Config{}
	arg.MustParse(&config)

	validate := validator.New()
	err := validate.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return Config{}, err
	}
	// perform argument validation and return an error if any argument fails validation.
	return config, nil
}
