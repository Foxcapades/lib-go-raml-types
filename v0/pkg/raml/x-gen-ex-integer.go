package raml

// IntegerExample defines a single example attached to a DataType
// or Property definition.
// Generated @ 2020-05-19T21:52:50.990676962-04:00
type IntegerExample interface {
	Example

	// SetDisplayName sets this example's display name value.
	SetDisplayName(string) IntegerExample

	// UnsetDisplayName removes this example's display name
	// value.
	UnsetDisplayName() IntegerExample

	// SetDescription sets this example's description value.
	SetDescription(string) IntegerExample

	// UnsetDescription removes this example's description
	// value.
	UnsetDescription() IntegerExample

	// SetAnnotations replaces this example's annotation map
	// with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) IntegerExample

	// UnsetAnnotations removes all annotations from this
	// example.
	UnsetAnnotations() IntegerExample

	// Value returns this example's value.
	Value() int64

	// SetValue sets this example's value.
	SetValue(int64) IntegerExample

	// SetStrict sets whether or not this example should be
	// validated against its parent type definition.
	SetStrict(bool) IntegerExample
}
