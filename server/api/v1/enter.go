package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/DailyAlgorithm"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/TEST"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	DailyAlgorithmApiGroup DailyAlgorithm.ApiGroup
	TESTApiGroup           TEST.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
