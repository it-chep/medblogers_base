package errors

import "github.com/pkg/errors"

// ErrAs generic version of errors.As()
func ErrAs[TErr any](err error) (TErr, bool) {
	var dstErr TErr
	ok := errors.As(err, &dstErr)
	return dstErr, ok
}
