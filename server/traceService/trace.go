package trace

import (
	"context"
	tracepb "micro-gogrpc-gateway/server/proto/gen/go"
	"time"
)

//type TraceServiceServer interface {
//	GetTrace(context.Context, *GetTraceRequest) (*GetTraceResponse, error)
//	mustEmbedUnimplementedTraceServiceServer()
//}

// Service implementation.
type Service struct{}

func (*Service) GetTrace(c context.Context, req *tracepb.GetTraceRequest) (*tracepb.GetTraceResponse, error) {
	return &tracepb.GetTraceResponse{
		Id: req.Id,
		Trace: &tracepb.Trace{
			Time:     int64(time.Now().Unix()),
			Location: "shanghai",
			StaySec:  100,
			VieePage: "hompage",
		},
	}, nil
}
