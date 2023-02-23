package exporters

import "fmt"

type ErrUnexpectedTransitionType struct {
	Expected string
	Actual   string
}

func (e *ErrUnexpectedTransitionType) Error() string {
	return fmt.Sprintf("unexpected transition type, expected %v, actual %v", e.Expected, e.Actual)
}
