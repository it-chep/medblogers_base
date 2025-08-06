package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errParentStub struct {
	child error
}

func (err *errParentStub) Error() string {
	return err.child.Error()
}

func (err *errParentStub) Unwrap() error {
	return err.child
}

func newErrParentStub(err error) error {
	return &errParentStub{child: err}
}

type errChildStub struct {
	child error
}

func (err *errChildStub) Error() string {
	return err.child.Error()
}

func (err *errChildStub) Unwrap() error {
	return err.child
}

func newErrChildStub(err error) error {
	return &errChildStub{child: err}
}

func TestErrAs(t *testing.T) {
	t.Parallel()

	t.Run("return true errors.As with generics types", func(t *testing.T) {
		t.Parallel()

		err := fmt.Errorf("new error")
		childErr := newErrChildStub(err)
		parentErr := newErrParentStub(childErr)

		exChildErr, ok := ErrAs[*errChildStub](parentErr)
		assert.True(t, ok)
		assert.Equal(t, exChildErr, childErr)
	})
}
