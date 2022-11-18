package model

import (
	"awesomeProject5/global"
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
	err := global.GVA_DB.Create(&UserInfo{
		Username: req.Username,
		Password: req.Password,
	}).Error
	if err != nil {
		return &proto.UserResponse{
			Msg:  "创建失败了",
			Code: "400",
		}, nil
	}
	return &proto.UserResponse{
		Msg:  "创建成功了",
		Code: "200",
	}, nil
}
