package matcher

import (
	"medblogers_base/internal/pkg/errors"
	"os/exec"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

type execErrMatcher struct {
	base types.GomegaMatcher
}

func (e execErrMatcher) Match(actual interface{}) (success bool, err error) {
	return e.base.Match(actual)
}

func (e execErrMatcher) FailureMessage(actual interface{}) (message string) {
	if exitErr, ok := errors.ErrAs[*exec.ExitError](actual.(error)); ok {
		return string(exitErr.Stderr)
	}

	return e.base.FailureMessage(actual)
}

func (e execErrMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	if exitErr, ok := errors.ErrAs[*exec.ExitError](actual.(error)); ok {
		return string(exitErr.Stderr)
	}

	return e.base.NegatedFailureMessage(actual)
}

// HaveOccurred конвертирует ошибку cmd.Exec в формат exec.ExitError, которая содержит подробную информацию об ошибки
func HaveOccurred() gomega.OmegaMatcher {
	return execErrMatcher{
		base: gomega.HaveOccurred(),
	}
}
