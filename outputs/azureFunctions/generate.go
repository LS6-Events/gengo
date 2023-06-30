package azureFunctions

import (
	"encoding/json"
	"github.com/ls6-events/gengo"
	"os"
	"strings"
)

type AzureFunctionsBinding struct {
	Type      string   `json:"type"`
	Direction string   `json:"direction"`
	Name      string   `json:"name"`
	Methods   []string `json:"methods,omitempty"`
	Route     string   `json:"route,omitempty"`
}

type AzureFunctionsOutput struct {
	Bindings []AzureFunctionsBinding `json:"bindings"`
}

func generate(directoryPath string) gengo.GenerateFunction {
	return func(s *gengo.Service) error {
		s.Log.Debug().Msg("Generating Azure Functions output")

		for _, route := range s.Routes {
			output := AzureFunctionsOutput{
				Bindings: []AzureFunctionsBinding{
					{
						Type:      "httpTrigger",
						Direction: "in",
						Name:      "req",
						Methods: []string{
							strings.ToLower(route.Method),
							"options",
						},
						Route: convertRoute(route),
					},
					{
						Type:      "http",
						Direction: "out",
						Name:      "res",
					},
				},
			}

			functionDirectoryPath := directoryPath + "/" + route.FuncName + "/"
			filePath := functionDirectoryPath + "function.json"

			err := os.Mkdir(functionDirectoryPath, 0755)
			if err != nil {
				s.Log.Error().Err(err).Msg("Failed to create directory")
				return err
			}

			s.Log.Debug().Msg("Generated JSON output")
			file, err := json.MarshalIndent(output, "", "  ")
			if err != nil {
				s.Log.Error().Err(err).Msg("Failed to marshal JSON output")
				return err
			}

			s.Log.Debug().Str("filePath", filePath).Msg("Writing JSON output to file")
			err = os.WriteFile(filePath, file, 0644)
			if err != nil {
				s.Log.Error().Err(err).Msg("Failed to write JSON output to file")
				return err
			}
		}

		s.Log.Info().Msg("Generated JSON output")
		return nil
	}
}
