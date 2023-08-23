// 自动生成模板TEST
package TEST

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TEST 结构体
type TEST struct {
	global.GVA_MODEL
	Test string `json:"test" form:"test" gorm:"column:test;comment:;"`
}

// TableName TEST 表名
func (TEST) TableName() string {
	return "TEST"
}
