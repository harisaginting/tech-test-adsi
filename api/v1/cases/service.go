package cases

import (
	"context"
	"github.com/harisaginting/tech-test-adsi/pkg/tracer"
	"github.com/harisaginting/tech-test-adsi/pkg/utils/helper"
)

type Service struct {}

func ProviderService() Service {
	return Service{}
}

func (service *Service) One(ctx context.Context, req *Payload, min, max int) (res Response, err error) {
	trace := tracer.Span(ctx,"CaseOne:Service")
	defer trace.End()
	var tmpResult []int
	for i := min; i <= max; i++ {
		if helper.IsOdd(i) && !helper.IntInSlice(i, req.Data) && i > 0{
			tmpResult = append(tmpResult, i)
		}
	} 
	minTmp, _ := helper.MinMaxIntSlice(tmpResult)
	if len(tmpResult) == 0 {
		res.Odd  = 1
	}else{
		res.Odd  = minTmp
	}
	tracer.SetAttributeInt(trace,"total cases",res.Odd)
	return
}