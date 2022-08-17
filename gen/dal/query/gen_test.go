// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Input struct {
	Args []interface{}
}

type Expectation struct {
	Ret []interface{}
}

type TestCase struct {
	Input
	Expectation
}

const dbName = "gen_test.db"

var db *gorm.DB
var once sync.Once

func init() {
	InitializeDB()
	db.AutoMigrate(&_another{})
}

func InitializeDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
	})
}

func assert(t *testing.T, methodName string, res, exp interface{}) {
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("%v() gotResult = %v, want %v", methodName, res, exp)
	}
}

type _another struct {
	ID uint64 `gorm:"primaryKey"`
}

func (*_another) TableName() string { return "another_for_unit_test" }

func Test_Available(t *testing.T) {
	if !Use(db).Available() {
		t.Errorf("query.Available() == false")
	}
}

func Test_WithContext(t *testing.T) {
	query := Use(db)
	if !query.Available() {
		t.Errorf("query Use(db) fail: query.Available() == false")
	}

	type Content string
	var key, value Content = "gen_tag", "unit_test"
	qCtx := query.WithContext(context.WithValue(context.Background(), key, value))

	for _, ctx := range []context.Context{
		qCtx.TUser.UnderlyingDB().Statement.Context,
	} {
		if v := ctx.Value(key); v != value {
			t.Errorf("get value from context fail, expect %q, got %q", value, v)
		}
	}
}

func Test_Transaction(t *testing.T) {
	query := Use(db)
	if !query.Available() {
		t.Errorf("query Use(db) fail: query.Available() == false")
	}

	err := query.Transaction(func(tx *Query) error { return nil })
	if err != nil {
		t.Errorf("query.Transaction execute fail: %s", err)
	}

	tx := query.Begin()

	err = tx.SavePoint("point")
	if err != nil {
		t.Errorf("query tx SavePoint fail: %s", err)
	}
	err = tx.RollbackTo("point")
	if err != nil {
		t.Errorf("query tx RollbackTo fail: %s", err)
	}
	err = tx.Commit()
	if err != nil {
		t.Errorf("query tx Commit fail: %s", err)
	}

	err = query.Begin().Rollback()
	if err != nil {
		t.Errorf("query tx Rollback fail: %s", err)
	}
}