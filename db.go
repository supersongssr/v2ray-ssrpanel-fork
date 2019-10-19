package v2ray_ssrpanel_plugin

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserModel struct {
	ID      uint
	VmessID string `gorm:"column:v2ray_uuid"`
	Email   string `gorm:"column:email"`
}

func (*UserModel) TableName() string {
	return "user"
}

type UserTrafficLog struct {
	ID       uint `gorm:"primary_key"`
	UserID   uint
	Uplink   uint64 `gorm:"column:u"`
	Downlink uint64 `gorm:"column:d"`
	NodeID   uint
	Rate     float64
	Traffic  string
	LogTime  int64
}

func (l *UserTrafficLog) BeforeCreate(scope *gorm.Scope) error {
	l.LogTime = time.Now().Unix()
	return nil
}

type NodeOnlineLog struct {
	ID         uint `gorm:"primary_key"`
	NodeID     uint
	OnlineUser int
	LogTime    int64
}

func (*NodeOnlineLog) TableName() string {
	return "ss_node_online_log"
}

func (l *NodeOnlineLog) BeforeCreate(scope *gorm.Scope) error {
	l.LogTime = time.Now().Unix()
	return nil
}

type NodeInfo struct {
	ID      uint `gorm:"primary_key"`
	NodeID  uint
	Uptime  time.Duration
	Load    string
	LogTime int64
}

func (*NodeInfo) TableName() string {
	return "ss_node_info_log"
}

func (l *NodeInfo) BeforeCreate(scope *gorm.Scope) error {
	l.LogTime = time.Now().Unix()
	return nil
}

type Node struct {
	ID          uint `gorm:"primary_key"`
	TrafficRate float64
}

func (*Node) TableName() string {
	return "ss_node"
}

type DB struct {
	DB *gorm.DB
}

func (db *DB) GetAllUsers() ([]UserModel, error) {
	users := make([]UserModel, 0)
	db.DB.Select("id, v2ray_uuid, email").Where("enable = 1 AND u + d < transfer_enable AND plan != 'A'").Find(&users)
	return users, nil
}

func (db *DB) GetNode(id uint) (*Node, error) {
	node := Node{}
	err := db.DB.First(&node, id).Error
	return &node, err
}
