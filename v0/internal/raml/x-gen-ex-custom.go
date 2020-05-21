package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

// NewCustomExample returns a new internal implementation of the
// raml.CustomExample interface.
//
// Generated @ 2020-05-20T21:46:00.638880955-04:00
func NewCustomExample() *CustomExample {
	return &CustomExample{
		annotations: NewAnnotationMap(),
		extra:       NewAnyMap(),
	}
}

// CustomExample is a generated internal implementation of the
// raml.CustomExample interface.
type CustomExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *CustomExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *CustomExample) SetDisplayName(name string) raml.CustomExample {
	e.displayName = &name
	return e
}

func (e *CustomExample) UnsetDisplayName() raml.CustomExample {
	e.displayName = nil
	return e
}

func (e *CustomExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *CustomExample) SetDescription(desc string) raml.CustomExample {
	e.description = &desc
	return e
}

func (e *CustomExample) UnsetDescription() raml.CustomExample {
	e.description = nil
	return e
}

func (e *CustomExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *CustomExample) SetAnnotations(ann raml.AnnotationMap) raml.CustomExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *CustomExample) UnsetAnnotations() raml.CustomExample {
	e.annotations = NewAnnotationMap()
	return e
}

func (e *CustomExample) Value() interface{} {
	return e.value
}

func (e *CustomExample) SetValue(val interface{}) raml.CustomExample {
	e.value = val
	return e
}

func (e *CustomExample) Strict() bool {
	return e.strict
}

func (e *CustomExample) SetStrict(b bool) raml.CustomExample {
	e.strict = b
	return e
}

func (e *CustomExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *CustomExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.ForEachMap(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *CustomExample) MarshalRAML(out raml.AnyMap) (bool, error) {
	if e.expand() {
		out.PutNonNil(rmeta.KeyDisplayName, e.displayName).
			PutNonNil(rmeta.KeyDescription, e.description).
			Put(rmeta.KeyValue, e.value)

		if e.strict != rmeta.ExampleDefaultStrict {
			out.Put(rmeta.KeyStrict, e.strict)
		}

		e.annotations.ForEach(func(k string, v raml.Annotation) { out.Put(k, v) })
		e.extra.ForEach(func(k interface{}, v interface{}) { out.Put(k, v) })

		return false, nil
	}

	out.Put("", e.value)
	return true, nil
}

func (e *CustomExample) assign(key, val *yaml.Node) error {
	if !xyml.IsString(key) {
		if ver, err := xyml.CastYmlTypeToScalar(key); err != nil {
			return err
		} else {
			e.extra.Put(ver, val)
		}
		return nil
	}

	if key.Value[0] == '(' {
		tmp := NewAnnotation()
		if err := tmp.UnmarshalRAML(val); err != nil {
			return err
		}
		e.annotations.Put(key.Value, tmp)
		return nil
	}

	switch key.Value {
	case rmeta.KeyDisplayName:
		return assign.AsStringPtr(val, &e.displayName)
	case rmeta.KeyDescription:
		return assign.AsStringPtr(val, &e.description)
	case rmeta.KeyStrict:
		return assign.AsBool(val, &e.strict)
	case rmeta.KeyValue:
		return e.assignVal(val)
	}

	if ver, err := xyml.CastYmlTypeToScalar(key); err != nil {
		return err
	} else {
		e.extra.Put(ver, val)
	}

	return nil
}

func (e *CustomExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *CustomExample) assignVal(val *yaml.Node) error {
	e.value = val

	return nil
}
