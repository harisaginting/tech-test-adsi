package cases

import (
	"context"
	"github.com/harisaginting/tech-test-adsi/pkg/tracer"
)

type Service struct {}

func ProviderService() Service {
	return Service{}
}

func (service *Service) One(ctx context.Context, req *Payload) (res Response, err error) {
	trace := tracer.Span(ctx,"CaseOne:Service")
	defer trace.End()

	

	tracer.SetAttributeInt(trace,"total cases","")
	return
}