// monkey/object/object.go

package object

import "fmt"

// type ObjectType string
// type Object interface {
//     Type() ObjectType
//     Inspect() string
// }

const (
	IntType = "INTEGER_TYPE"
)

type Object interface {
	Type() string
	String() string
}

type Integer struct {
	Value int64
}

// @Impl
func (this *Integer) Type() string {
	return IntType
}

// @Impl
func (this *Integer) String() string {
	return fmt.Sprintf("%d", this.Value)
}
