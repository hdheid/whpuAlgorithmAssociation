package DailyAlgorithm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DailyAlgorithmRecordRouter struct {
}

// InitDailyAlgorithmRecordRouter 初始化 DailyAlgorithmRecord 路由信息
func (s *DailyAlgorithmRecordRouter) InitDailyAlgorithmRecordRouter(Router *gin.RouterGroup) {
	DARRouter := Router.Group("DAR").Use(middleware.OperationRecord())
	DARRouterWithoutRecord := Router.Group("DAR")
	var DARApi = v1.ApiGroupApp.DailyAlgorithmApiGroup.DailyAlgorithmRecordApi
	{
		DARRouter.POST("createDailyAlgorithmRecord", DARApi.CreateDailyAlgorithmRecord)             // 新建DailyAlgorithmRecord
		DARRouter.DELETE("deleteDailyAlgorithmRecord", DARApi.DeleteDailyAlgorithmRecord)           // 删除DailyAlgorithmRecord
		DARRouter.DELETE("deleteDailyAlgorithmRecordByIds", DARApi.DeleteDailyAlgorithmRecordByIds) // 批量删除DailyAlgorithmRecord
		DARRouter.PUT("updateDailyAlgorithmRecord", DARApi.UpdateDailyAlgorithmRecord)              // 更新DailyAlgorithmRecord
	}
	{
		DARRouterWithoutRecord.GET("findDailyAlgorithmRecord", DARApi.FindDailyAlgorithmRecord)       // 根据ID获取DailyAlgorithmRecord
		DARRouterWithoutRecord.GET("coverDailyAlgorithmRecord", DARApi.CoverDailyAlgorithmRecord)     // 根据date获取DailyAlgorithmRecord
		DARRouterWithoutRecord.GET("getDailyAlgorithmRecordList", DARApi.GetDailyAlgorithmRecordList) // 获取DailyAlgorithmRecord列表
	}
}
