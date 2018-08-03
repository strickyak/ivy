package value // import "github.com/strickyak/ivy/value"

import (
	"fmt"

	"github.com/strickyak/ivy/config"
)

type Box struct {
	Contents interface{}
}

func (b Box) String() string {
	return fmt.Sprintf("@box<%T>", b.Contents)
}

func (b Box) Sprint(c *config.Config) string {
	return fmt.Sprintf("@box<%T>", b.Contents)
}

func (b Box) Eval(c Context) Value {
	return b
}

func (b Box) Inner() Value {
	return b
}

func (b Box) ProgString() string {
	return "@z@"
}

func (b Box) toType(c *config.Config, vt valueType) Value {
	switch vt {
	case vectorType:
		return NewVector([]Value{b})
	case matrixType:
		return NewMatrix([]Value{one}, []Value{b})
	}
	panic(Error(fmt.Sprintf("Cannot use a Box as %v", vt)))
}
