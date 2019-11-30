package service

import (
	"context"

	"github.com/ybalexdp/sample-grpc/pb"
)

type SampleService struct {
}

func (s *SampleService) GetSample(ctx context.Context, message *pb.GetSampleRequest) (*pb.SampleResponse, error) {
	return &pb.SampleResponse{
		Message: "Hello :" + message.Name,
	}, nil
}
