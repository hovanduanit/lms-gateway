package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hovanduanit/lms-gateway/configs"
	gen "github.com/hovanduanit/lms-proto/generated/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
