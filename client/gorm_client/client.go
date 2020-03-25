package gorm_client

import (
	"github.com/jinzhu/gorm"
)

type GormClient struct {
	master *gorm.DB
	slave  *gorm.DB
}

func NewClient() *GormClient {
	return &GormClient{}
}

var Client *GormClient

func Dial(dialect string, args ...interface{}) (*GormClient, error) {
	Client = NewClient()
	return Client.Dial(dialect, args...)
}

func (c *GormClient) Dial(dialect string, args ...interface{}) (*GormClient, error) {
	var err error
	// master dial
	c.master, err = gorm.Open(dialect, args...)
	if err != nil {
		return c, err
	}

	// slave dial
	c.slave, err = gorm.Open(dialect, args...)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c *GormClient) Master() *gorm.DB {
	return c.master
}

func (c *GormClient) Slave() *gorm.DB {
	return c.slave
}
