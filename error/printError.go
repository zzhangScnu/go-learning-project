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

/**
https://bytedance.feishu.cn/docs/doccneaWosTDgx63VHyiWHFC8Zf#
https://blog.dharnitski.com/2019/09/09/go-errors-are-not-pkg-errors/
https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
*/
