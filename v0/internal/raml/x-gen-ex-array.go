package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"gopkg.in/yaml.v3"
)

// NewArrayExample returns a new internal implementation of the
// raml.ArrayExample interface.
//
// Generated @ 2020-05-20T21:46:00.638880955-04:00
func NewArrayExample() *ArrayExample {
	return &ArrayExample{
		annotations: NewAnnotationMap(),
		extra:       NewAnyMap(),
	}
}

// ArrayExample is a generated internal implementation of the
// raml.ArrayExample interface.
type ArrayExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       []interface{}
	strict      bool
	extra       raml.AnyMap
}

func (e *ArrayExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *ArrayExample) SetDisplayName(name string) raml.ArrayExample {
	e.displayName = &name
	return e
}

func (e *ArrayExample) UnsetDisplayName() raml.ArrayExample {
	e.displayName = nil
	return e
}

func (e *ArrayExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *ArrayExample) SetDescription(desc string) raml.ArrayExample {
	e.description = &desc
	return e
}

func (e *ArrayExample) UnsetDescription() raml.ArrayExample {
	e.description = nil
	return e
}

func (e *ArrayExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *ArrayExample) SetAnnotations(ann raml.AnnotationMap) raml.ArrayExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *ArrayExample) UnsetAnnotations() raml.ArrayExample {
	e.annotations = NewAnnotationMap()
	return e
}

func (e *ArrayExample) Value() []interface{} {
	return e.value
}

func (e *ArrayExample) SetValue(val []interface{}) raml.ArrayExample {
	e.value = val
	return e
}

func (e *ArrayExample) Strict() bool {
	return e.strict
}

func (e *ArrayExample) SetStrict(b bool) raml.ArrayExample {
	e.strict = b
	return e
}

func (e *ArrayExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *ArrayExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.ForEachMap(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *ArrayExample) MarshalRAML(out raml.AnyMap) (bool, error) {
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

func (e *ArrayExample) assign(key, val *yaml.Node) error {
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

func (e *ArrayExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *ArrayExample) assignVal(val *yaml.Node) error {
	if err := xyml.RequireList(val); err != nil {
		return err
	} else {
		e.value = make([]interface{}, 0, len(val.Content))
		return xyml.ForEachList(val, func(v *yaml.Node) error {
			e.value = append(e.value, v)
			return nil
		})
	}

	return nil
}
