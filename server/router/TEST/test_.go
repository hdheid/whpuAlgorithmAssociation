package TEST

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TESTRouter struct {
}

// InitTESTRouter 初始化 TEST 路由信息
func (s *TESTRouter) InitTESTRouter(Router *gin.RouterGroup) {
	TEST_SRouter := Router.Group("TEST_S").Use(middleware.OperationRecord())
	TEST_SRouterWithoutRecord := Router.Group("TEST_S")
	var TEST_SApi = v1.ApiGroupApp.TESTApiGroup.TESTApi
	{
		TEST_SRouter.POST("createTEST", TEST_SApi.CreateTEST)             // 新建TEST
		TEST_SRouter.DELETE("deleteTEST", TEST_SApi.DeleteTEST)           // 删除TEST
		TEST_SRouter.DELETE("deleteTESTByIds", TEST_SApi.DeleteTESTByIds) // 批量删除TEST
		TEST_SRouter.PUT("updateTEST", TEST_SApi.UpdateTEST)              // 更新TEST
	}
	{
		TEST_SRouterWithoutRecord.GET("findTEST", TEST_SApi.FindTEST)       // 根据ID获取TEST
		TEST_SRouterWithoutRecord.GET("getTESTList", TEST_SApi.GetTESTList) // 获取TEST列表
	}
}
