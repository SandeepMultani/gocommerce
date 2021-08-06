package server

import (
	"context"
	"time"

	event "github.com/SandeepMultani/gocommerce/src/backend/protobuf/event"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ingestService struct {
	event.UnimplementedIngestServiceServer
}

var _ event.IngestServiceServer = &ingestService{}

func NewIngestServer() event.IngestServiceServer {
	return &ingestService{}
}

func (srv *ingestService) Ingest(context context.Context, req *event.IngestRequest) (*event.IngestResponse, error) {
	println("request from rpc client: payload - ", req.Payload)
	println("request from rpc client: ", req)

	return &event.IngestResponse{
		RequestId:    req.RequestId,
		IsSuccessful: &wrappers.BoolValue{Value: true},
		Message:      "hello world",
		Timestamp:    &timestamppb.Timestamp{Seconds: time.Now().Unix()},
	}, nil
}
