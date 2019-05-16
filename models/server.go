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

type ServerStatics struct {
	Id      			int64     `orm:"auto"`
	Version 			[]byte    `orm:"type(text)"`
	Uptime 				[]byte    `orm:"type(text)"`
	Sessions			[]byte    `orm:"type(text)"`
	Processes			[]byte    `orm:"type(text)"`
	ProcessesArray  	[]byte    `orm:"type(text)"`
	FileHandles  		[]byte    `orm:"type(text)"`
	FileHandlesLimit  	[]byte    `orm:"type(text)"`
	OsKernel  			[]byte    `orm:"type(text)"`
	OsName  			[]byte    `orm:"type(text)"`
	OsArch  			[]byte    `orm:"type(text)"`
	CpuName  			[]byte    `orm:"type(text)"`
	CpuCores  			[]byte    `orm:"type(text)"`
	CpuFreq  			[]byte    `orm:"type(text)"`
	RamTotal  			[]byte    `orm:"type(text)"`
	RamUsage  			[]byte    `orm:"type(text)"`
	SwapTotal  			[]byte    `orm:"type(text)"`
	SwapUsage  			[]byte    `orm:"type(text)"`
	DiskArray  			[]byte    `orm:"type(text)"`
	DiskTotal  			[]byte    `orm:"type(text)"`
	DiskUsage  			[]byte    `orm:"type(text)"`
	Connections  		[]byte    `orm:"type(text)"`
	Nic  				[]byte    `orm:"type(text)"`
	IPV4  				[]byte    `orm:"type(text)"`
	IPV6  				[]byte    `orm:"type(text)"`
	Rx  				[]byte    `orm:"type(text)"`
	Tx  				[]byte    `orm:"type(text)"`
	RxGap  				[]byte    `orm:"type(text)"`
	TxGap  				[]byte    `orm:"type(text)"`
	Load  			[]byte    `orm:"type(text)"`
	LoadCpu  			[]byte    `orm:"type(text)"`
	LoadIO  			[]byte    `orm:"type(text)"`
	PingEu  			[]byte    `orm:"type(text)"`
	PingUs  			[]byte    `orm:"type(text)"`
	PingAs  			[]byte    `orm:"type(text)"`
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
	Active  bool      `orm:"int8(1)"`
	User    *User     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(Server),
		new(ServerStatics),
		new(Task),
		new(Api),
	)
}