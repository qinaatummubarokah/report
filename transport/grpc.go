package transport

import (
	"context"
	"reflect"
	pb "report/proto"
	"report/repository"

	"github.com/go-kit/kit/log"

	// apmgrpc "git.bluebird.id/lib/apm/grpc"

	"report/endpoint"

	gt "github.com/go-kit/kit/transport/grpc"
)

//GrpcServer ....
type GrpcServer struct {
	getReport gt.Handler
}

// //GRPCServerRun run grpc server
// func GRPCServerRun(
// 	addr string,
// 	recruitmentSrv pb.ServiceServer,
// ) {
// 	errs := make(chan error)
// 	go func() {
// 		c := make(chan os.Signal)
// 		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
// 		errs <- fmt.Errorf("%s", <-c)
// 	}()

// 	grpcListener, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}

// 	var opts []grpc.ServerOption
// 	opts = apmgrpc.GetElasticAPMServerOptions()

// 	go func() {
// 		baseServer := grpc.NewServer(opts...)

// 		pb.RegisterServiceServer(baseServer, recruitmentSrv)

// 		log.Println("🚀 Server recruitment started successfully 🚀 - Running on", addr)
// 		baseServer.Serve(grpcListener)
// 	}()

// 	log.Println("exit", <-errs)
// }

// func NewGRPCServer() *GrpcServer {
// 	// start database conn
// 	repo, err := postgres.NewPostgresConn(repository.DBConfiguration{
// 		// Default Value
// 		DBHost:     config.Get(constant.DBHostKey, "127.0.0.1"),
// 		DBName:     config.Get(constant.DBNameKey, "postgres"),
// 		DBOptions:  config.Get(constant.DBOptionsKey, "?sslmode=disable"),
// 		DBPassword: config.Get(constant.DBPasswordKey, ""),
// 		DBPort:     config.Get(constant.DBPortKey, "5432"),
// 		DBUser:     config.Get(constant.DBUserKey, "postgres"),
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}

// 	svc := service.NewService(repo)

// 	endpoint := endpoint.MakeEndpoints(&svc)

// 	// handlerOpt := []grpctransport.ServerOption{
// 	// 	grpctransport.ServerBefore(jwt.GRPCToContext()),
// 	// }

// 	return &GrpcServer{
// 		getReport: grpctransport.NewServer(
// 			endpoint.GetReport,
// 			decodeGetReport,
// 			encodeGetReport,
// 			// handlerOpt...,
// 		),
// 	}
// }

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoint endpoint.Endpoint, logger log.Logger) pb.ServiceServer {
	return &GrpcServer{
		getReport: gt.NewServer(
			endpoint.GetReport,
			decodeGetReport,
			encodeGetReport,
		),
	}
}

func (g *GrpcServer) GetReport(ctx context.Context, request *pb.Empty) (*pb.Data, error) {
	_, resp, err := g.getReport.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	println("xxxxx")
	return resp.(*pb.Data), nil
}

func decodeGetReport(ctx context.Context, request interface{}) (interface{}, error) {
	// req := request.(pb.Data)
	return &pb.Empty{}, nil
}

// func encodeGetReport(ctx context.Context, resp interface{}) (interface{}, error) {
// 	println("hmmmmmmmmm")
// 	resp = resp.(endpoint.OtherResp)
// 	println("resp: ", resp)
// 	return &pb.Data{Id: resp.(pb.Data).Id, Name: resp.(pb.Data).Name}, nil
// }

func encodeGetReport(ctx context.Context, response interface{}) (interface{}, error) {
	var transforms []*pb.DataDb
	s := reflect.ValueOf(response)

	for i := 0; i < s.Len(); i++ {
		var row pb.DataDb
		curentData := s.Index(i).Interface().(repository.Data)
		row.Id = curentData.ID
		row.Name = curentData.Name
		transforms = append(transforms, &row)
	}
	return &pb.Data{Data: transforms}, nil
}
