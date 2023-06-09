package openapi

type openAPIJSONType struct {
	Type   string
	Format string
}

var acceptedTypeMap = map[string]openAPIJSONType{
	"string": openAPIJSONType{
		Type: "string",
	},
	"int": openAPIJSONType{
		Type:   "integer",
		Format: "int32",
	},
	"int8": openAPIJSONType{
		Type:   "integer",
		Format: "int8",
	},
	"int16": openAPIJSONType{
		Type:   "integer",
		Format: "int16",
	},
	"int32": openAPIJSONType{
		Type:   "integer",
		Format: "int32",
	},
	"int64": openAPIJSONType{
		Type:   "integer",
		Format: "int64",
	},
	"uint": openAPIJSONType{
		Type:   "integer",
		Format: "uint",
	},
	"uint8": openAPIJSONType{
		Type:   "integer",
		Format: "uint8",
	},
	"uint16": openAPIJSONType{
		Type:   "integer",
		Format: "uint16",
	},
	"uint32": openAPIJSONType{
		Type:   "integer",
		Format: "uint32",
	},
	"uint64": openAPIJSONType{
		Type:   "integer",
		Format: "uint64",
	},
	"float": openAPIJSONType{
		Type:   "number",
		Format: "float",
	},
	"float32": openAPIJSONType{
		Type:   "number",
		Format: "float32",
	},
	"float64": openAPIJSONType{
		Type:   "number",
		Format: "float64",
	},
	"bool": openAPIJSONType{
		Type: "boolean",
	},
	"byte": openAPIJSONType{
		Type:   "string",
		Format: "byte",
	},
	"rune": openAPIJSONType{
		Type:   "string",
		Format: "rune",
	},
	"struct": openAPIJSONType{
		Type: "object",
	},
	"map": openAPIJSONType{
		Type: "object",
	},
	"slice": openAPIJSONType{
		Type: "array",
	},
	"any": openAPIJSONType{
		Type: "",
	},
	"nil": openAPIJSONType{
		Type: "",
	},
}

func mapAcceptedType(acceptedType string) Schema {
	if acceptedType, ok := acceptedTypeMap[acceptedType]; ok {
		if acceptedType.Type == "" {
			return Schema{}
		}
		return Schema{
			Type: acceptedType.Type,
		}
	}

	return Schema{}
}
