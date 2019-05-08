package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Server struct {
	Id               int64            `orm:"auto"`
	Uuid             string           `orm:"size(255)"`
	Account          *User            `orm:"rel(fk)"`
	Name             string           `orm:"size(255)"`
	Ipv4             string           `orm:"size(255)"`
	Ipv6             string           `orm:"size(255)"`
	Mac              string           `orm:"size(255)"`
	OperatingSystem  *OperatingSystem `orm:"rel(fk)"`
	BillingCycle     int64            `orm:"int64(11)"`
	Amount           float64          `orm:"type(float)"`
	RegistrationDate time.Time        `orm:"type(datetime)"`
	NextpaymentDate  time.Time        `orm:"type(datetime)"`
	CancellationDate time.Time        `orm:"type(datetime)"`
	Bwlimit          uint64           `orm:"uint64(64)"`
	DiskLimit		 uint64           `orm:"uint64(64)"`
	CpuLimit		 uint64           `orm:"uint64(64)"`
	RamLimit		 uint64           `orm:"uint64(64)"`
	Diskoverflow       bool             `orm:"int8(1)"`
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
type OperatingSystem struct {
	Id       int64     `orm:"auto"`
	OsGroup  *OsGroup  `orm:"rel(fk)"`
	Name     string    `orm:"size(255)"`
	Version  string    `orm:"size(255)"`
	Visible  bool      `orm:"int8(1)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}
type OsGroup struct {
	Id              int64              `orm:"auto"`
	Name            string             `orm:"size(255)"`
	Logo            string             `orm:"size(255)"`
	Visible         bool               `orm:"int8(1)"`
	Created         time.Time          `orm:"auto_now_add;type(datetime)"`
	Updated         time.Time          `orm:"auto_now;type(datetime)"`
	OperatingSystem []*OperatingSystem `orm:"reverse(many)"`
}
type Task struct {
	Id      int64     `orm:"auto"`
	Task    string    `orm:"size(255)"`
	Ongoing bool      `orm:"int8(1)"`
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
		new(OperatingSystem),
		new(OsGroup),
		new(Task),
	)
}