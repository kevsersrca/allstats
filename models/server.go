package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Server struct {
	Id               int64            `orm:"auto""`
	Uuid             string           `orm:"size(255)"`
	Account          *User            `orm:"rel(fk)"`
	Name             string           `orm:"size(255)", form:"name"`
	Ipv4             string           `orm:"size(255)"`
	Ipv6             string           `orm:"size(255)"`
	Mac              string           `orm:"size(255)"`
	BillingCycle     int64            `orm:"int64(11)"`
	Amount           float64          `orm:"type(float)"`
	Systemlimit      string           `orm:"size(255)", form:"loadusage"`
	DiskLimit		 string           `orm:"size(255)", form:"diskusage"`
	RamLimit		 string           `orm:"size(255)", form:"ramusage"`
	Diskoverflow     bool            `orm:"int8(1)"`
	AutoPayment      bool             `orm:"int8(1)"`
	Description      string           `orm:"size(255)"`
	Status           string           `orm:"size(255);default(InProgress)"`
	Created          time.Time        `orm:"auto_now_add;type(datetime)"`
	Updated          time.Time        `orm:"auto_now;type(datetime)"`
}

type ServerBwhistory struct {
	Id      int64     `orm:"auto"`
	Server  *Server   `orm:"rel(fk)"`
	Year    uint16    `orm:"uint16(11)"`
	Month   uint8     `orm:"uint8(11)"`
	Bwusage uint64    `orm:"int64(11)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
type ServerDiskhistory struct {
	Id      int64     `orm:"auto"`
	Server  *Server   `orm:"rel(fk)"`
	Year    uint16    `orm:"uint16(11)"`
	Month   uint8     `orm:"uint8(11)"`
	Diskusage uint64    `orm:"int64(11)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
type ServerRamhistory struct {
	Id      int64     `orm:"auto"`
	Server  *Server   `orm:"rel(fk)"`
	Year    uint16    `orm:"uint16(11)"`
	Month   uint8     `orm:"uint8(11)"`
	Ramusage uint64    `orm:"int64(11)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
type ServerCpuhistory struct {
	Id      int64     `orm:"auto"`
	Server  *Server   `orm:"rel(fk)"`
	Year    uint16    `orm:"uint16(11)"`
	Month   uint8     `orm:"uint8(11)"`
	Cpuusage uint64    `orm:"int64(11)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type Task struct {
	Id      int64     `orm:"auto"`
	Task    string    `orm:"size(255)"`
	Ongoing bool      `orm:"int8(1)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type Api struct {
	Id      int64     `orm:"auto"`
	Key     string    `orm:"size(255)"`
	Active  bool     `orm:"int8(1)"`
	User    *User     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(Server),
		new(ServerBwhistory),
		new(ServerDiskhistory),
		new(ServerCpuhistory),
		new(ServerRamhistory),
		new(Task),
		new(Api),
	)
}