package env

import (
	"fmt"
	"os"
)

type Env struct {
	missingVars []string
}

func RequireEnv(varName string) *Env {
	_, exists := os.LookupEnv(varName)

	env := &Env{
		missingVars: []string{},
	}

	if !exists {
		env.missingVars = append(env.missingVars, varName)
	}

	return env
}

func (e *Env) RequireEnv(varName string) *Env {
	_, exists := os.LookupEnv(varName)

	if !exists {
		e.missingVars = append(e.missingVars, varName)
	}

	return e
}

func (e *Env) GetError() error {
	if len(e.missingVars) > 0 {
		builtError := "The following required environment variables are missing:\n"

		for _, name := range e.missingVars {
			builtError += "  - " + name + "\n"
		}

		builtError += "Please set them then restart the application."

		return fmt.Errorf(builtError)
	}

	return nil
}
