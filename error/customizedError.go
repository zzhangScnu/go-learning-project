package error

import "fmt"

type WrongAccount struct {
	account string
}

// 实现了error接口
func (w WrongAccount) Error() string {
	return fmt.Sprintf("wrong account is: %s, please correct it.", w.account)
}
