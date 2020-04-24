package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ConnectionConf struct {
	Host string
	User string
	Password string
	Port	string
	Database string
}

func (c *ConnectionConf) Connect() (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port,c.Database))
}
