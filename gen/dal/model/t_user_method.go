package model

import "gorm.io/gen"

type TUserMethod interface {
	// where("name=@name")
	FindByName(name string) (gen.M, error)
}
