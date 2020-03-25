package model

import (
	"time"
	"visitor/client/gorm_client"
	"visitor/pkg/common"
	"visitor/pkg/ierr"
)

type User struct {
	Id     int    `json:"id" gorm:"id;primary_key;AUTO_INCREMENT"`                   // 自增id
	OpenId string `json:"open_id" gorm:"open_id;type:varchar(32);index:idx_open_id"` // 用户微信openId
	Token  string `json:"token" gorm:"token;type:varchar(32);index:idx_token"`       // 系统用户身份token

	CreatedAt time.Time `json:"created_at" gorm:"created_at;NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0)"`
	Status    bool      `json:"status" gorm:"status;NOT NULL;default:0"`
}

func (u User) TableName() string {
	return "user"
}

func (u *User) Create() (code int, err error) {
	if u.OpenId == "" {
		code = ierr.InvalidOpenId
		err = ierr.IErr(code)
		return
	}

	err = gorm_client.Client.Master().Where("open_id = ? and status = ?", u.OpenId, common.StatusNormal).FirstOrCreate(&u).Error
	if err != nil {
		code = ierr.CreateDataFail
		return
	}
	return
}

func (u *User) LoadByToken() (code int, err error) {
	err = gorm_client.Client.Master().Table(u.TableName()).Where("token = ? and status = ?", u.Token, common.StatusNormal).
		Find(&u).Error
	if err != nil {
		code = ierr.QueryDataFail
		return
	}
	return
}
