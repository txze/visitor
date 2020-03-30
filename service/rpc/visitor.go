package rpc

import (
	"context"
	"visitor/visitorpb"
)

func (s *Service) AddVisitor(ctx context.Context, req *visitorpb.AddVisitorRequest) (*visitorpb.AddVisitorReply, error) {
	return &visitorpb.AddVisitorReply{Ok: true}, nil
}
