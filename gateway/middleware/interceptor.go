package middleware

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func ServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Skip authorize when GetJWT is requested
	log.Println("hello from middleware")
	if info.FullMethod != "/proto.EventStoreService/GetJWT" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}

	// Calls the handler
	h, err := handler(ctx, req)

	log.Println("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

func authorize(ctx context.Context) error {
/*	log.Println("i'm in method authorize!")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	accessHeader, ok := md["accessToken"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}
	refreshHeader, _ := md["refreshToken"]

	accessToken := accessHeader[0]
	refreshToken := refreshHeader[0]
	updatedAccess, err := jwt.Validate(accessToken, refreshToken)
	if err != nil {
		return err
	}
	if updatedAccess != nil {
		md.Set("accessToken", *updatedAccess)
	}*/
	return nil
}
