package _switch

import "fmt"

// Judge ：空接口，代表任意类型
func Judge(t interface{}) {
	// realType也可以命名为t
	switch realType := t.(type) {
	case int:
		fmt.Printf("%d: i am int!\n", realType)
	case string:
		fmt.Printf("%s: i am string!\n", realType)
	case float64:
		fmt.Printf("%f: i am float64!\n", realType)
	default:
		fmt.Printf("%#v: i am other!\n", realType)
	}
}

func Fallthrough(t int) {
	switch t {
	case 0:
		fmt.Println("continue!")
		fallthrough
	case 1, 2:
		fmt.Println("value is ", t)
	default:
		fmt.Println("reached end")
	}
}
