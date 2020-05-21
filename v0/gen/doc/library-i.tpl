{{define "library-interface" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

// Library represents the contents of a Library type RAML fragment.
type Library interface {
	yaml.Unmarshaler
	yaml.Marshaler
	{{template "hasAnnotations" "Library"}}
	{{template "hasAnnotationTypes" "Library"}}
	{{template "hasFacets" "Library"}}
	{{template "hasImports" "Library"}}
	{{template "hasResourceTypes" "Library"}}
	{{template "hasSecuritySchemes" "Library"}}
	{{template "hasTraits" "Library"}}
	{{template "hasTypes" "Library"}}
	{{template "hasUsage" "Library"}}
}
{{- end}}