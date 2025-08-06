package spec

import (
	"fmt"

	"github.com/pkg/errors"
)

var specConnector string

func init() {
	SetSpecErrorConnector("или")
}

// SetSpecErrorConnector переопределяет союз, который объединяет ошибки для логической операции OR
func SetSpecErrorConnector(connector string) {
	specConnector = connector
}

type errRightLeft struct {
	left, right error
}

func newErrRightLeft(left, right error) error {
	return errRightLeft{left: left, right: right}
}

func (e errRightLeft) Is(err error) bool {
	return errors.Is(e.left, err) || errors.Is(e.right, err)
}

func (e errRightLeft) Error() string {
	return fmt.Sprintf("%s %s %s", e.left, specConnector, e.right)
}
