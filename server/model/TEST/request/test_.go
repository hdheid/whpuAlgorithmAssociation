package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/TEST"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TESTSearch struct {
	TEST.TEST
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
