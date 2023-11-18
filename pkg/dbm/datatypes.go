package dbm

import "fmt"

type TypeInteger struct {
}

func Integer() *TypeInteger {
	return &TypeInteger{}
}

func (t *TypeInteger) SQLType() string {
	return "INTEGER"
}

type TypeVarchar struct {
	size int
}

func Varchar(size int) *TypeVarchar {
	return &TypeVarchar{size: size}
}

func VarcharDefault() *TypeVarchar {
	return &TypeVarchar{}
}

func (t *TypeVarchar) SQLType() string {
	if t.size == 0 {
		return "VARCHAR"
	}
	return fmt.Sprintf("VARCHAR(%d)", t.size)
}

type TypeSmallint struct{}

func SmallInteger() *TypeSmallint {
	return &TypeSmallint{}
}

func (t *TypeSmallint) SQLType() string {
	return "SMALLINT"
}
