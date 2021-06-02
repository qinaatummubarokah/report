package endpoint

import (
	"context"
	"report/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	GetReport endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoint {
	return Endpoint{
		GetReport: makeGetReportEndpoint(s),
	}
}

func makeGetReportEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, _ := s.GetReport(ctx)
		print("result: ", result)
		return s.GetReport(ctx)
	}
}

// func makeCdaEndpoint(s service.Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 		result, _ := s.Cda(ctx)
// 		print("result: ", result)
// 		return OtherResp{Result: result}, nil
// 	}
// }
