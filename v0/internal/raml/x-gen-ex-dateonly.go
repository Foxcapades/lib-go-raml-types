package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rmeta"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

// NewDateOnlyExample returns a new internal implementation of the
// raml.DateOnlyExample interface.
//
// Generated @ 2020-05-25T18:15:33.802521588-04:00
func NewDateOnlyExample() *DateOnlyExample {
	return &DateOnlyExample{
		annotations: raml.NewAnnotationMap(0),
		extra:       raml.NewAnyMap(0),
		strict:      rmeta.ExampleDefaultStrict,
	}
}

// DateOnlyExample is a generated internal implementation of the
// raml.DateOnlyExample interface.
type DateOnlyExample struct {
	displayName *string
	description *string
	annotations raml.AnnotationMap
	value       *string
	strict      bool
	extra       raml.AnyMap
}

func (e *DateOnlyExample) DisplayName() option.String {
	return option.NewMaybeString(e.displayName)
}

func (e *DateOnlyExample) SetDisplayName(name string) raml.DateOnlyExample {
	e.displayName = &name
	return e
}

func (e *DateOnlyExample) UnsetDisplayName() raml.DateOnlyExample {
	e.displayName = nil
	return e
}

func (e *DateOnlyExample) Description() option.String {
	return option.NewMaybeString(e.description)
}

func (e *DateOnlyExample) SetDescription(desc string) raml.DateOnlyExample {
	e.description = &desc
	return e
}

func (e *DateOnlyExample) UnsetDescription() raml.DateOnlyExample {
	e.description = nil
	return e
}

func (e *DateOnlyExample) Annotations() raml.AnnotationMap {
	return e.annotations
}

func (e *DateOnlyExample) SetAnnotations(ann raml.AnnotationMap) raml.DateOnlyExample {
	if ann == nil {
		return e.UnsetAnnotations()
	}
	e.annotations = ann
	return e
}

func (e *DateOnlyExample) UnsetAnnotations() raml.DateOnlyExample {
	e.annotations = raml.NewAnnotationMap(0)
	return e
}

func (e *DateOnlyExample) Value() option.String {
	return option.NewMaybeString(e.value)
}

func (e *DateOnlyExample) SetValue(val string) raml.DateOnlyExample {
	e.value = &val
	return e
}

func (e *DateOnlyExample) UnsetValue() raml.DateOnlyExample {
	e.value = nil
	return e
}

func (e *DateOnlyExample) Strict() bool {
	return e.strict
}

func (e *DateOnlyExample) SetStrict(b bool) raml.DateOnlyExample {
	e.strict = b
	return e
}

func (e *DateOnlyExample) ExtraFacets() raml.AnyMap {
	return e.extra
}

func (e *DateOnlyExample) UnmarshalRAML(value *yaml.Node) error {

	if xyml.IsMap(value) {
		return xyml.MapForEach(value, e.assign)
	}

	return e.assignVal(value)
}

func (e *DateOnlyExample) MarshalRAML(out raml.AnyMap) (bool, error) {
	if e.expand() {
		out.PutIfNotNil(rmeta.KeyDisplayName, e.displayName).
			PutIfNotNil(rmeta.KeyDescription, e.description).
			PutIfNotNil(rmeta.KeyValue, e.value)

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

func (e *DateOnlyExample) assign(key, val *yaml.Node) error {
	if !xyml.IsString(key) {
		if ver, err := xyml.ToScalarValue(key); err != nil {
			return err
		} else {
			e.extra.Put(ver, val)
		}
		return nil
	}

	if key.Value[0] == '(' {
		tmp := NewAnnotation()
		if err := UnmarshalUntypedMapRAML(tmp, val); err != nil {
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

	if ver, err := xyml.ToScalarValue(key); err != nil {
		return err
	} else {
		e.extra.Put(ver, val)
	}

	return nil
}

func (e *DateOnlyExample) expand() bool {
	return e.displayName != nil ||
		e.description != nil ||
		e.annotations.Len() > 0 ||
		e.extra.Len() > 0 ||
		e.strict != rmeta.ExampleDefaultStrict
}

func (e *DateOnlyExample) assignVal(val *yaml.Node) error {
	if err := xyml.RequireString(val); err != nil {
		return err
	}
	e.value = &val.Value

	return nil
}
