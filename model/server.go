package model

import (
	"awesomeProject5/proto"
	"context"
	"go.uber.org/zap"
)

type Server struct {
	proto.UnimplementedUserInfoServer
}

func (s *Server) Userinfo(c context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	log := zap.NewExample()
	log.Info(req.Password)
	return &proto.UserResponse{
		Msg:  "msg 响应码",
		Code: "200",
	}, nil
}
