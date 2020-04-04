package rpc

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/txze/visitorpb/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"visitor/client/redis_client"
	"visitor/pkg/visitor_persistence"
)

func (s *Service) AddVisitor(ctx context.Context, req *visitorpb.AddVisitorRequest) (*visitorpb.AddVisitorReply, error) {
	redisCon := redis_client.MasterPool.Get()
	defer redisCon.Close() //其实是放回连接池

	redisP := visitor_persistence.Redis{
		Con: redisCon,
	}
	err := redisP.Add(req.Id, req.Unixnano)
	if err != nil {
		logrus.WithField("req", req).Error("AddVisitor rpc call redis.Add", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &visitorpb.AddVisitorReply{Ok: true}, nil
}

func (s *Service) ListVisitorCount(ctx context.Context, req *visitorpb.ListVisitorCountRequest) (*visitorpb.ListVisitorCountReply, error) {
	redisCon := redis_client.MasterPool.Get()
	defer redisCon.Close() //其实是放回连接池

	redisP := visitor_persistence.Redis{
		Con: redisCon,
	}
	result, err := redisP.List(req.Id, req.Span, req.Unixnano)
	if err != nil {
		logrus.WithField("req", req).Error("ListVisitorCount rpc call redis.List", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &visitorpb.ListVisitorCountReply{
		Counts: result,
	}, nil
}
