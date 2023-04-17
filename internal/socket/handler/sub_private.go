package handler

import (
	"github.com/mofei1/gpush/internal/socket/manager"
	"github.com/mofei1/gpush/internal/socket/model/request"
	"github.com/mofei1/gpush/internal/socket/model/response"
)

type SubPrivate struct {
}

func (SubPrivate) Handle(r request.Message, conn *manager.Connection) {
	//todo 认证
	if !conn.IsLogin() {
		conn.Send(response.SubWithoutLogin)
		return
	}
}
