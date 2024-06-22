package utils

import (
	"context"
	"github.com/DateMine/bot-domain/internal/converter"
	"github.com/DateMine/bot-domain/pkg/models/request"
	"github.com/DateMine/bot-domain/pkg/models/response"
	"github.com/DateMine/grpc-domain/pkg/parser_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type WebClient struct {
	address string
}

func NewWebClient(ctx context.Context) *WebClient {
	return &WebClient{
		address: "localhost:50056",
	}
}

func (c WebClient) Send(request request.Request, ctx context.Context) (*response.HttpClientResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()
	client := parser_v1.NewParseV1Client(conn)
	response, err := client.SendHttp(ctx, converter.ToRequestDesc(request))
	if err != nil {
		return nil, err
	}
	responseDomain := converter.ToResponseDomain(response)

	return responseDomain, nil
}
