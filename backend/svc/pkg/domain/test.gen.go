// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package domain

const TableNameTest = "test"

// Test mapped from table <test>
type Test struct {
	ID   int32   `gorm:"column:id;type:integer;primaryKey" json:"id"`
	Name *string `gorm:"column:name;type:text" json:"name"`
	Test *string `gorm:"column:test;type:text" json:"test"`
}

// TableName Test's table name
func (*Test) TableName() string {
	return TableNameTest
}
