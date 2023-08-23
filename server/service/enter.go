package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/DailyAlgorithm"
	"github.com/flipped-aurora/gin-vue-admin/server/service/TEST"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	DailyAlgorithmServiceGroup DailyAlgorithm.ServiceGroup
	TESTServiceGroup           TEST.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
