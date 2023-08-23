// 自动生成模板DailyAlgorithmRecord
package DailyAlgorithm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
	"time"
)

// DailyAlgorithmRecord 结构体
type DailyAlgorithmRecord struct {
	global.GVA_MODEL
	Date      *time.Time     `json:"date" form:"date" gorm:"column:date;comment:日期;"`
	User_name string         `json:"user_name" form:"user_name" gorm:"column:user_name;comment:用户名;"`
	Link      string         `json:"link" form:"link" gorm:"column:link;comment:题目链接;"`
	Code      datatypes.JSON `json:"code" form:"code" gorm:"column:code;comment:代码;"`
}

// TableName DailyAlgorithmRecord 表名
func (DailyAlgorithmRecord) TableName() string {
	return "DailyAlgorithmRecord"
}
