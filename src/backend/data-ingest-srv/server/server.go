package server

import (
	"context"
	"time"

	event "github.com/SandeepMultani/gocommerce/src/backend/data-ingest-srv/protobuf/event"
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
	return &event.IngestResponse{
		RequestId:    req.RequestId,
		IsSuccessful: &wrappers.BoolValue{Value: true},
		Message:      "hello world",
		Timestamp:    &timestamppb.Timestamp{Seconds: time.Now().Unix()},
	}, nil
}
