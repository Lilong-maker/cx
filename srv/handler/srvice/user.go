package srvice

import (
	"context"
	"cx/srv/dasic/config"
	"cx/srv/handler/model"

	__ "cx/srv/dasic/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedOrderServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderAdd(_ context.Context, in *__.OrderAddReq) (*__.OrderAddResp, error) {

	var order model.Order
	err := order.FindOrder(config.DB, in.Name)
	if err != nil {
		return &__.OrderAddResp{
			Msg:  "参数错误",
			Code: 400,
		}, nil
	}
	m := model.Order{
		Name:  in.Name,
		Price: float64(in.Price),
		Num:   int(in.Num),
	}
	err = m.OrderAdd(config.DB)
	if err != nil {
		return &__.OrderAddResp{
			Msg:  "添加失败",
			Code: 400,
		}, nil
	}
	return &__.OrderAddResp{
		Msg:  "添加成功",
		Code: 200,
	}, nil
}
