package user

import "context"
import "github.com/harisaginting/ginting/pkg/tracer"
import srv "github.com/harisaginting/ginting/service/samplegrpc"

type Service struct {
	repo Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) List(ctx context.Context, res *ResponseList) {
	trace := tracer.Span(ctx,"ListUser")
	defer trace.End()
	users := service.repo.FindAll()
	res.Items = users
	res.Total = len(users)

	tracer.SetAttributeInt(trace,"total User",res.Total)
	return
}

func (service *Service) ListGRPC(ctx context.Context, param string) (res string) {
	trace := tracer.Span(ctx, "ListUser")
	defer trace.End()

	// CALL GRPC SAMPLE
	res = srv.Sample("TEST SERVICE LIST GRPC")
	// END CALL GRPC SAMPLE

	tracer.SetAttributeInt(trace,"total User",res)
	return
}