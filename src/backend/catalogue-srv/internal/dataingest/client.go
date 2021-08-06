package dataingest

import (
	"context"
	"time"

	event "github.com/SandeepMultani/gocommerce/src/backend/protobuf/event"
	"google.golang.org/grpc"
)

func Ingest() {
	println("start to dial rpc")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial("localhost:5501", opts...)
	if err != nil {
		println("fail to dial: %v", err)
	}
	defer conn.Close()
	client := event.NewIngestServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.Ingest(ctx, &event.IngestRequest{
		Payload: "test-payload",
	})

	if err != nil {
		panic(err)
	}

	println("response from rpc server: ", res.Message)
}
