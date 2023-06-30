package openapi

type openAPIJSONType struct {
	Type   string
	Format string
}

var acceptedTypeMap = map[string]openAPIJSONType{
	"string": {
		Type: "string",
	},
	"int": {
		Type:   "integer",
		Format: "int32",
	},
	"int8": {
		Type:   "integer",
		Format: "int8",
	},
	"int16": {
		Type:   "integer",
		Format: "int16",
	},
	"int32": {
		Type:   "integer",
		Format: "int32",
	},
	"int64": {
		Type:   "integer",
		Format: "int64",
	},
	"uint": {
		Type:   "integer",
		Format: "uint",
	},
	"uint8": {
		Type:   "integer",
		Format: "uint8",
	},
	"uint16": {
		Type:   "integer",
		Format: "uint16",
	},
	"uint32": {
		Type:   "integer",
		Format: "uint32",
	},
	"uint64": {
		Type:   "integer",
		Format: "uint64",
	},
	"float": {
		Type:   "number",
		Format: "float",
	},
	"float32": {
		Type:   "number",
		Format: "float32",
	},
	"float64": {
		Type:   "number",
		Format: "float64",
	},
	"bool": {
		Type: "boolean",
	},
	"byte": {
		Type:   "string",
		Format: "byte",
	},
	"time.Time": {
		Type:   "string",
		Format: "date-time",
	},
	"uuid.UUID": {
		Type:   "string",
		Format: "uuid",
	},
	"rune": {
		Type:   "string",
		Format: "rune",
	},
	"struct": {
		Type: "object",
	},
	"map": {
		Type: "object",
	},
	"slice": {
		Type: "array",
	},
	"any": {
		Type: "",
	},
	"nil": {
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
