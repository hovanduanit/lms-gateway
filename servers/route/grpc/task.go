package grpc

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gen "github.com/hovanduanit/lms-proto/generated/task"
)

func RegisterTaskGrpcHandler(mux *runtime.ServeMux) {
	err := gen.RegisterTaskServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		fmt.Sprintf("%s:%d", configs.GlobalConfig.GrpcTaskHost, configs.GlobalConfig.GrpcTaskPort),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)

	if err != nil {
		log.Fatalf("Can't register Task GRPC Handler: %v", err)
	}
}
