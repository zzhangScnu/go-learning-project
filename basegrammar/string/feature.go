package string

import "fmt"

const (
	str = "hello，世界"
	num = "123"
)

// ConvertToIntByForRange 测试了for-range循环的每一位，都是Unicode位码，只有减去'0'时才能得到数字（int32类型）。循环效果跟for-i以及rune[]是一样的。
func ConvertToIntByForRange() {
	for i, v := range num {
		fmt.Println(i, v, v-'0')
	}
}

func ConvertToIntByForI() {
	l := len(num)
	for i := 0; i < l; i++ {
		fmt.Println(i, num[i], num[i]-'0')
	}
}

func ConvertToIntByRune() {
	r := []rune(num)
	for i, v := range r {
		fmt.Println(i, v, v-'0')
	}
}

// RangeString 每次循环展示【Unicode位码编码后，第一个字节的索引-Unicode位码-位码转换的字符串】
func RangeString() {
	for i, v := range str {
		fmt.Println(i, v, string(v))
	}
}

// RangeStringRune 每次循环展示【Unicode位码索引-Unicode位码-位码转换的字符串】
// Rune是Unicode位码的表示，实际上是int32的等价类型
func RangeStringRune() {
	r := []rune(str)
	for i, v := range r {
		fmt.Println(i, v, string(v))
	}
}

// RangeByte 展示了对于非ASCII码的Unicode位码，编码（用字节数组表示）后，可能会用多个字节表示一个字符。每次循环展示【字节索引-字节-字节解码的字符串（可能映射不到Unicode字符）】
func RangeByte() {
	r := []byte(str)
	for i, v := range r {
		fmt.Println(i, v, string(v))
	}
}

// RangeStringByI 循环的效果和RangeByte一致
func RangeStringByI() {
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i], string(str[i]))
	}
}

func Compare() {
	// 截断操作是前闭后开区间
	eqFlag := str[:5] == "hello"
	fmt.Println("equals compare result: ", eqFlag)
	// 比较操作是逐字节进行的，所以比较结果是按编码后的自然顺序来的
	gtFlag := str[0] <= 'i'
	fmt.Println("greater compare result: ", gtFlag)
}
