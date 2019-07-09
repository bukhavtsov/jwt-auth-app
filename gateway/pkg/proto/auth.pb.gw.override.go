package serverProto

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"log"
	"net/http"
)

func RegisterRestAppHandlerClientOverride(ctx context.Context, mux *runtime.ServeMux, client RestAppClient) error {
	mux.Handle("GET", pattern_RestApp_SignIn_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RestApp_SignIn_0(rctx, inboundMarshaler, client, req, pathParams)
		log.Println(resp)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		log.Println("Register has been recived!")
		forward_RestApp_SignIn_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})
	return nil
}
