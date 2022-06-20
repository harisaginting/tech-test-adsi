package api

import (
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/tech-test-adsi/pkg/wire"
)

func RestV1(r *gin.RouterGroup) {
	// Dependency injection
	apiCase := wire.ApiCases()

	// group rest
	rest := r.Group("rest")
	{
		// versioning,
		v1 := rest.Group("v1")
		{
			// case group
			apiCaseGroup := v1.Group("case")
			{
				apiCaseGroup.POST("/one", apiCase.One)
			}
		}
	}

	return
}