package raml

// UnionExample defines a single example attached to a DataType
// or Property definition.
// Generated @ 2020-05-19T21:52:50.990676962-04:00
type UnionExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) UnionExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() UnionExample

	// SetDescription sets this example's description value.
	SetDescription(string) UnionExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() UnionExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) UnionExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() UnionExample

	// Value returns this example's value.
	Value() interface{}

	// SetValue sets this example's value.
	SetValue(interface{}) UnionExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) UnionExample
}