package gengo

type GenerateFunction func(service *Service) error

type Output struct {
	Mode OutputMode

	Generate GenerateFunction
}

type OutputMode string

const (
	OutputModeAzureFunctions OutputMode = "azureFunctions"
	OutputModeJSON           OutputMode = "json"
	OutputModeOpenAPI        OutputMode = "openapi"
)
