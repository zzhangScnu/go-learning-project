package error

import (
	"fmt"
	"github.com/pkg/errors"
)

func fmtSomeError() error {
	err := WrongAccount{"xxx"}
	return fmt.Errorf("something failed: %v", err)
}

func wrapSomeError() error {
	err := WrongAccount{"xxx"}
	wrapped := errors.Wrap(err, "something failed")
	fmt.Println(errors.Cause(wrapped).(WrongAccount).account)
	return wrapped
}
