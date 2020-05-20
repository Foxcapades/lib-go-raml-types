package raml

import (
	

	"github.com/Foxcapades/goop/v1/pkg/option"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
)

// NewUnionType returns a new internal implementation
// of the raml.UnionType interface.
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
func NewUnionType(log *logrus.Entry) *UnionType {
	log = xlog.WithType(log, "internal.UnionType")

	out := &UnionType{
		examples: NewUnionExampleMap(log),
	}
	
	out.ExtendedDataType = NewExtendedDataType(rmeta.TypeUnion, log, out)

	return out
}

// UnionType is a default generated implementation of
// the raml.UnionType interface
//
// Generated @ 2020-05-20T00:33:46.349824232-04:00
type UnionType struct {
	*ExtendedDataType

	def      *interface{}
	example  raml.UnionExample
	examples raml.UnionExampleMap
	enum     []interface{}
	
}

func (o *UnionType) SetType(s string) raml.UnionType {
	o.schema = s
	return o
}

func (o *UnionType) Default() option.Untyped {
	return option.NewMaybeUntyped(o.def)
}

func (o *UnionType) SetDefault(i interface{}) raml.UnionType {
	o.def = &i
	return o
}

func (o *UnionType) UnsetDefault() raml.UnionType {
	o.def = nil
	return o
}

func (o *UnionType) Example() raml.UnionExample {
	return o.example
}

func (o *UnionType) SetExample(ex raml.UnionExample) raml.UnionType {
	o.example = ex
	return o
}

func (o *UnionType) UnsetExample() raml.UnionType {
	o.example = nil
	return o
}

func (o *UnionType) Examples() raml.UnionExampleMap {
	return o.examples
}

func (o *UnionType) SetExamples(examples raml.UnionExampleMap) raml.UnionType {
	if examples == nil {
		return o.UnsetExamples()
	}
	o.examples = examples
	return o
}

func (o *UnionType) UnsetExamples() raml.UnionType {
	o.examples = NewUnionExampleMap(o.DataType.log)
	return o
}

func (o *UnionType) SetDisplayName(s string) raml.UnionType {
	o.displayName = &s
	return o
}

func (o *UnionType) UnsetDisplayName() raml.UnionType {
	o.displayName = nil
	return o
}

func (o *UnionType) SetDescription(s string) raml.UnionType {
	o.description = &s
	return o
}

func (o *UnionType) UnsetDescription() raml.UnionType {
	o.description = nil
	return o
}

func (o *UnionType) SetAnnotations(annotations raml.AnnotationMap) raml.UnionType {
	if annotations == nil {
		return o.UnsetAnnotations()
	}
	o.hasAnnotations.mp = annotations
	return o
}

func (o *UnionType) UnsetAnnotations() raml.UnionType {
	o.hasAnnotations.mp = NewAnnotationMap(o.DataType.log)
	return o
}

func (o *UnionType) SetFacetDefinitions(facets raml.FacetMap) raml.UnionType {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}
	o.facets = facets
	return o
}

func (o *UnionType) UnsetFacetDefinitions() raml.UnionType {
	o.facets = NewFacetMap(o.DataType.log)
	return o
}

func (o *UnionType) SetXml(x raml.Xml) raml.UnionType {
	o.xml = x
	return o
}

func (o *UnionType) UnsetXml() raml.UnionType {
	o.xml = nil
	return o
}

func (o *UnionType) Enum() []interface{} {
	return o.enum
}

func (o *UnionType) SetEnum(i []interface{}) raml.UnionType {
	o.enum = i
	return o
}

func (o *UnionType) UnsetEnum() raml.UnionType {
	o.enum = nil
	return o
}

func (o *UnionType) SetExtraFacets(facets raml.AnyMap) raml.UnionType {
	if facets == nil {
		return o.UnsetExtraFacets()
	}
	o.hasExtra.mp = facets
	return o
}

func (o *UnionType) UnsetExtraFacets() raml.UnionType {
	o.hasExtra.mp = NewAnyMap(o.DataType.log)
	return o
}

func (o *UnionType) SetRequired(b bool) raml.UnionType {
	o.required = b
	return o
}

func (o *UnionType) marshal(out raml.AnyMap) error {
	o.DataType.log.Trace("internal.UnionType.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	
	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *UnionType) assign(key, val interface{}, log *logrus.Entry) error {
	log.Trace("internal.UnionType.assign")
	switch key {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind, log); err != nil {
			return xlog.Error(log, err)
		} else if err := ex.UnmarshalRAML(val, log); err != nil {
			return err
		} else {
			o.example = ex.(raml.UnionExample)
		}
		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val, log)
	case rmeta.KeyEnum:
		arr, err := assign.AsAnyList(val, log)
		if err != nil {
			return xlog.Error(log, "the enum facet must be an array. " + err.Error())
		}
		for i := range arr {
			
			o.enum = append(o.enum, arr[i])
			
		}
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required, log)
	}
	
	return o.ExtendedDataType.assign(key, val, log)
}
