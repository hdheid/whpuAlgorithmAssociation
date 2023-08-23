package DailyAlgorithm

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DailyAlgorithmRecordApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
