package cat

type Cat struct {
	color  string
	height float64
	weight float64
	name   string
}

func NewCat(color string, height float64, weight float64) *Cat {
	return &Cat{
		color:  color,
		height: height,
		weight: weight,
	}
}

func (c Cat) Color() string {
	return c.color
}

func (c Cat) Height() float64 {
	return c.height
}

func (c Cat) Weight() float64 {
	return c.weight
}

func (c *Cat) SetName(name string) {
	c.name = name
}
