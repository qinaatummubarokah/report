package transport

import (
	"context"
	"reflect"
	pb "report/proto"
	"report/repository"

	"github.com/go-kit/kit/log"
	"github.com/webx-top/echo/logger/zerolog"
	"google.golang.org/grpc"

	// apmgrpc "git.bluebird.id/lib/apm/grpc"

	"report/endpoint"

	gt "github.com/go-kit/kit/transport/grpc"
)

//GrpcServer ....
type GrpcServer struct {
	getReport gt.Handler
}

type ClientGRPC struct {
	logger    zerolog.Logger
	conn      *grpc.ClientConn
	client    pb.GuploadServiceClient
	chunkSize int
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

func encodeGetReport(ctx context.Context, response interface{}) (interface{}, error) {
	var transforms []*pb.DataDb
	s := reflect.ValueOf(response)

	for i := 0; i < s.Len(); i++ {
		var row pb.DataDb
		curentData := s.Index(i).Interface().(repository.Data)
		// row.Id = curentData.ID
		// row.Fare = curentData.Name
		// row.Extra = curentData.AccountCode

		row.Id = *curentData.ID
		row.Fare = *curentData.Fare
		row.Extra = *curentData.Extra
		row.DiscountAmt = *curentData.DiscountAmt
		row.PaidAmount = *curentData.PaidAmount
		row.PromoCode = *curentData.PromoCode
		row.PaymentToken = *curentData.PaymentToken
		row.TransactionTime = *curentData.TransactionTime
		row.Identifier = *curentData.Identifier
		row.PaymentType = *curentData.PaymentType
		row.VehicleId = *curentData.VehicleId
		row.VehicleName = *curentData.VehicleName
		row.ServiceType = *curentData.ServiceType
		row.DriverId = *curentData.DriverID
		row.PickUpSuburb = *curentData.PickUpSuburb
		row.PickUpArea = *curentData.PickUpArea
		row.DestinationArea = *curentData.DestinationArea
		row.DestinationSuburb = *curentData.DSestinationSuburb
		// row.PickUpLatitude = curentData.PickUpLatitude
		row.DestinationLat = *curentData.DestinationLat
		row.PickUpLng = *curentData.PickUpLng
		row.PaymentProfileId = *curentData.PaymentProfileID
		row.State = *curentData.State
		row.ReleasedAt = *curentData.ReleasedAt
		row.CompletedAt = *curentData.CompletedAt
		row.CreatedAt = *curentData.CreatedAt
		row.UpdatedAt = *curentData.Updated_at
		row.CcIdentifier = *curentData.CcIdentifier
		row.AccountId = *curentData.AccountID
		row.SapSentAt = *curentData.SapSentAt
		row.SapState = *curentData.SapState
		row.MsakuState = *curentData.MsakuState
		row.CvNumber = *curentData.CvNumber
		row.ValidityPeriod = *curentData.ValidityPeriod
		row.ItopId = *curentData.ItopID
		row.OrderId = *curentData.OrderID
		row.PickedUpAt = *curentData.PickedUp
		row.TripPurpose = *curentData.TripPurpose
		row.MsakuTransactionId = *curentData.MsakuTransactionID
		row.ExternalOrderId = *curentData.ExternalOrderID
		row.RouteImage = *curentData.RouteImage
		row.DepartmentName = *curentData.DepartmentName
		row.AccountCode = *curentData.AccountCode
		row.UserName = *curentData.UserName
		row.InvoiceNumber = *curentData.InvoiceNumber
		row.PostingDate = *curentData.PostingDate
		row.Distance = *curentData.Distance
		row.OtherInformation = *curentData.OtherInformation
		row.PickUpLat = *curentData.PickUpLat
		row.DestinationLng = *curentData.DestinationLng
		row.MsakuResponse = *curentData.MsakuResponse
		row.PickupAddress = *curentData.PickupAddress
		row.DropoffAddress = *curentData.DropoffAddress
		row.Tips = *curentData.Tips
		row.DriverName = *curentData.DriverName

		// row.Id = curentData.ID
		// row.Fare = curentData.Fare
		// row.Extra = curentData.Extra
		// row.DiscountAmt = curentData.DiscountAmt
		// row.PaidAmount = curentData.PaidAmount
		// row.PromoCode = curentData.PromoCode
		// row.PaymentToken = curentData.PaymentToken
		// row.TransactionTime = curentData.TransactionTime
		// row.Identifier = curentData.Identifier
		// row.PaymentType = curentData.PaymentType
		// row.VehicleId = curentData.VehicleId
		// row.VehicleName = curentData.VehicleName
		// row.ServiceType = curentData.ServiceType
		// row.DriverId = curentData.DriverID
		// row.PickUpSuburb = curentData.PickUpSuburb
		// row.PickUpArea = curentData.PickUpArea
		// row.DestinationArea = curentData.DestinationArea
		// row.DestinationSuburb = curentData.DSestinationSuburb
		// // row.PickUpLatitude = curentData.PickUpLatitude
		// row.DestinationLat = curentData.DestinationLat
		// row.PickUpLng = curentData.PickUpLng
		// row.PaymentProfileId = curentData.PaymentProfileID
		// row.State = curentData.State
		// row.ReleasedAt = curentData.ReleasedAt
		// row.CompletedAt = curentData.CompletedAt
		// row.CreatedAt = curentData.CreatedAt
		// row.UpdatedAt = curentData.Updated_at
		// row.CcIdentifier = curentData.CcIdentifier
		// row.AccountId = curentData.AccountID
		// row.SapSentAt = curentData.SapSentAt
		// row.SapState = curentData.SapState
		// row.MsakuState = curentData.MsakuState
		// row.CvNumber = curentData.CvNumber
		// row.ValidityPeriod = curentData.ValidityPeriod
		// row.ItopId = curentData.ItopID
		// row.OrderId = curentData.OrderID
		// row.PickedUpAt = curentData.PickedUp
		// row.TripPurpose = curentData.TripPurpose
		// row.MsakuTransactionId = curentData.MsakuTransactionID
		// row.ExternalOrderId = curentData.ExternalOrderID
		// row.RouteImage = curentData.RouteImage
		// row.DepartmentName = curentData.DepartmentName
		// row.AccountCode = curentData.AccountCode
		// row.UserName = curentData.UserName
		// row.InvoiceNumber = curentData.InvoiceNumber
		// row.PostingDate = curentData.PostingDate
		// row.Distance = curentData.Distance
		// row.OtherInformation = curentData.OtherInformation
		// row.PickUpLat = curentData.PickUpLat
		// row.DestinationLng = curentData.DestinationLng
		// row.MsakuResponse = curentData.MsakuResponse
		// row.PickupAddress = curentData.PickupAddress
		// row.DropoffAddress = curentData.DropoffAddress
		// row.Tips = curentData.Tips
		// row.DriverName = curentData.DriverName
		transforms = append(transforms, &row)
	}
	return &pb.Data{Data: transforms}, nil
}

// func (c *ClientGRPC) UploadFile(ctx context.Context, f string) (stats Stats, err error) {

// 	// Get a file handle for the file we
// 	// want to upload
// 	file, err = os.Open(f)

// 	// Open a stream-based connection with the
// 	// gRPC server
// 	stream, err := c.client.Upload(ctx)

// 	// Start timing the execution
// 	stats.StartedAt = time.Now()

// 	// Allocate a buffer with `chunkSize` as the capacity
// 	// and length (making a 0 array of the size of `chunkSize`)
// 	buf = make([]byte, c.chunkSize)
// 	for writing {
// 		// put as many bytes as `chunkSize` into the
// 		// buf array.
// 		n, err = file.Read(buf)

// 		// ... if `eof` --> `writing=false`...

// 		stream.Send(&messaging.Chunk{
// 			// because we might've read less than
// 			// `chunkSize` we want to only send up to
// 			// `n` (amount of bytes read).
// 			// note: slicing (`:n`) won't copy the
// 			// underlying data, so this as fast as taking
// 			// a "pointer" to the underlying storage.
// 			Content: buf[:n],
// 		})
// 	}

// 	// keep track of the end time so that we can take the elapsed
// 	// time later
// 	stats.FinishedAt = time.Now()

// 	// close
// 	status, err = stream.CloseAndRecv()
// }
