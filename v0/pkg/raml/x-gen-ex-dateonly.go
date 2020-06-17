package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

// DateOnlyExample defines a single example attached to a DataType
// or Property definition.
//
// Generated @ 2020-05-25T18:15:33.802521588-04:00
type DateOnlyExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) DateOnlyExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() DateOnlyExample

	// SetDescription sets this example's description value.
	SetDescription(string) DateOnlyExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() DateOnlyExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) DateOnlyExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() DateOnlyExample

	// Value returns this example's value.
	Value() option.String

	// SetValue sets this example's value.
	SetValue(v string) DateOnlyExample

	// UnsetValue removes this example's value.
	UnsetValue() DateOnlyExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) DateOnlyExample
}
