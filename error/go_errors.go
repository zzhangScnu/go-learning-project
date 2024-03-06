package error

import "github.com/pkg/errors"

// https://bytedance.feishu.cn/docs/doccneaWosTDgx63VHyiWHFC8Zf#

func A() error {

	err := B()
	if err != nil {
		return errors.Wrap(err, "A重复？")
		// return errors.WithMessage(err, "extra_A") //这里直接返回或者附加信息都可以
	}

	return nil
}

func B() error {

	err := C()
	if err != nil {
		return errors.Wrap(err, "B重复？")
		// return errors.WithMessage(err, "extra_B") //这里直直接返回或者附加信息都可以
	}

	return nil
}

func C() error {
	err := D()
	if err != nil {
		return errors.Wrap(err, "C重复？")
		// return err
		//这里直接返回或者附加信息都可以,但是这里如果用Wrap包装的话,找信息就会重复
	}
	return nil
}

func D() error {
	return errors.New("D.err") //这里直接new或者使用Wrap,WithStack包装一个已用的err都可以
}
