package cat

type Cat struct {
	color  string
	height float64
	weight float64
	name   string
}

// NewCat 可以在返回指针的构造函数之上取值，因为指针指向的是分配的真实地址
// 但是不能在返回值的构造函数之上取地址，因为返回的值已经是拷贝的，而不是实例化的原值了
func NewCat(color string, height float64, weight float64) Cat {
	return *NewCatAddr(color, height, weight)
}

func NewCatAddr(color string, height float64, weight float64) *Cat {
	return &Cat{
		color:  color,
		height: height,
		weight: weight,
	}
}

// Color 方法：与函数类似，但是与特定类型关联，体现在c Cat这个接收器参数上
func (c Cat) Color() string {
	return c.color
}

func (c Cat) Height() float64 {
	return c.height
}

func (c Cat) Weight() float64 {
	return c.weight
}

func (c Cat) Name() string {
	return c.name
}

// SetName 方法的接收器参数是指针，这样才能真实地改变实参
func (c *Cat) SetName(name string) {
	c.name = name
}

func (c *Cat) SetColor(color string) {
	c.color = color
}

func (c Cat) SetColorForValue(color string) {
	c.color = color
}
