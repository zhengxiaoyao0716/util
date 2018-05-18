package easyjson

import "fmt"

// TypeNotMatchError .
type TypeNotMatchError struct {
	value  interface{}
	expect string
}

func (err *TypeNotMatchError) Error() string {
	return fmt.Sprintf("type not match: %#v is not %s", err.value, err.expect)
}

type refContainer struct {
	container
}

// IsRef .
func (ref refContainer) IsRef(json container) bool {
	return ref.container == json
}

// ValueNotFoundError .
type ValueNotFoundError struct {
	refContainer
	at interface{}
}

func (err *ValueNotFoundError) Error() string {
	return fmt.Sprintf("value not found: `%v` of %v", err.at, err.container)
}

// ValueTypeNotMatchError .
type ValueTypeNotMatchError struct {
	refContainer
	at interface{}
	*TypeNotMatchError
}

func (err *ValueTypeNotMatchError) Error() string {
	return fmt.Sprintf(
		"value type not match: `%v` of %v,  %#v is not %s",
		err.at, err.container, err.value, err.expect)
}
