package server

import (
	"github.com/golang/glog"
	m "github.com/tradingAI/go/db/postgres/model"
)

func (s *Servlet) CheckTokenExisted(token string) (err error) {
	var user m.User
	err = s.DB.Where("token = ?", token).Find(&user).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
