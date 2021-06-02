package transport

import (
	"context"
	"log"
	"os"
	"report/endpoint"

	"report/proto"
	"report/repository"
	"report/repository/postgres"

	"github.com/go-kit/kit/auth/jwt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

//GrpcServer ....
type GrpcServer struct {
	getReport grpctransport.Handler
}

func NewGRPCServer() *GrpcServer {
	// start database conn
	repo, err := postgres.NewPostgresConn(repository.DBConfiguration{
		// Default Value
		DBHost:     config.Get(constant.DBHostKey, "127.0.0.1"),
		DBName:     config.Get(constant.DBNameKey, "postgres"),
		DBOptions:  config.Get(constant.DBOptionsKey, "?sslmode=disable"),
		DBPassword: config.Get(constant.DBPasswordKey, ""),
		DBPort:     config.Get(constant.DBPortKey, "5432"),
		DBUser:     config.Get(constant.DBUserKey, "postgres"),
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	svc := service.NewRecruitmentService(repo)

	endpoint := endpoint.MakeEndpoints(&svc)

	handlerOpt := []grpctransport.ServerOption{
		grpctransport.ServerBefore(jwt.GRPCToContext()),
	}

	return &GrpcServer{
		getReport: grpctransport.NewServer(
			endpoint.GetReport,
			decodeGetReport,
			encodeGetReport,
			// handlerOpt...,
		),
	}
}

func (g *GrpcServer) GetReport(ctx context.Context, request *proto.Empty) (*proto.Data, error) {
	if _, resp, err := g.createReport.ServeGRPC(ctx, request); err != nil {
		return nil, err
	}
	return resp.(*proto.Data), nil
}

func decodeGetReport(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*proto.GetReportRequest)
	return &proto.Empty{}, nil
}

func encodeGetReport(ctx context.Context, resp interface{}) (interface{}, error) {
	esp := resp.(model.Data)
	return &proto.Data{
		// RegionId:  []byte(resp.RegionID.String()),
		ID:   resp.id,
		Name: resp.name,
	}, nil
}
