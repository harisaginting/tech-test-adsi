package user

import "github.com/gin-gonic/gin"
import "github.com/harisaginting/ginting/pkg/http/response"
import "github.com/harisaginting/ginting/pkg/tracer"
import "github.com/harisaginting/ginting/pkg/log"

type Controller struct {
	service Service
}

func ProviderController(s Service) Controller {
	return Controller{
		service: s,
	}
}

func (ctrl *Controller) List(c *gin.Context) {
	ctx  := c.Request.Context()
	span := tracer.Span(ctx, "ListUserController")
	defer span.End()
	log.Info(ctx,"Controller User")

	var resData ResponseList
	ctrl.service.List(ctx, &resData)
	
	// return
	response.Json(c,resData)
	return
}

func (ctrl *Controller) ListGRPC(c *gin.Context) {
	ctx  := c.Request.Context()
	span := tracer.Span(ctx, "ListUserControllerGRPC")
	defer span.End()

	var resData string
	res := ctrl.service.ListGRPC(ctx, resData)
	resData = res
	// return
	response.Json(c,resData)
	return
}