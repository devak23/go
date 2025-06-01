package trace

import (
	"fmt"
)

// Tracer is an interface that denotes that an object is capable of tracing
// events throughout the code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(params ...interface{}) {
	fmt.Fprint(t.out, params)
	fmt.Println(t.out)
}
