package cases

import (
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/tech-test-adsi/pkg/http/response"
	"github.com/harisaginting/tech-test-adsi/pkg/utils/helper"
	"github.com/harisaginting/tech-test-adsi/pkg/tracer"
)

type Controller struct {
	service Service
}

func ProviderController(s Service) Controller {
	return Controller{
		service: s,
	}
}

func (ctrl *Controller) One(c *gin.Context) {
	ctx  := c.Request.Context()
	span := tracer.Span(ctx, "CaseOne:Controller")
	defer span.End()


	var req Payload
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.BadRequest(c)
		return
	}
	err = json.Unmarshal(request, &req)
	if err != nil {
		response.BadRequest(c)
		return
	}

	min,max := helper.MinMaxIntSlice(req.Data)
	if &min == nil || &max == nil {
		response.BadRequest(c)
		return
	}
	if max >= 1000 || min <= -1000 {
		response.BadRequest(c)
		return	
	}

	resData, err := ctrl.service.One(ctx, &req, min, max)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Json(c,resData)
	return
}