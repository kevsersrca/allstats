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
	Server          	*Server   `orm:"rel(fk)"`
	Version 			string    `orm:"type(text)"`
	Uptime 				string    `orm:"type(text)"`
	Sessions			string    `orm:"type(text)"`
	Processes			string    `orm:"type(text)"`
	ProcessesArray  	string    `orm:"type(text)"`
	FileHandles  		string    `orm:"type(text)"`
	FileHandlesLimit  	string    `orm:"type(text)"`
	OsKernel  			string    `orm:"type(text)"`
	OsName  			string    `orm:"type(text)"`
	OsArch  			string    `orm:"type(text)"`
	CpuName  			string    `orm:"type(text)"`
	CpuCores  			string    `orm:"type(text)"`
	CpuFreq  			string    `orm:"type(text)"`
	RamTotal  			string    `orm:"type(text)"`
	RamUsage  			string    `orm:"type(text)"`
	SwapTotal  			string    `orm:"type(text)"`
	SwapUsage  			string    `orm:"type(text)"`
	DiskArray  			string    `orm:"type(text)"`
	DiskTotal  			string    `orm:"type(text)"`
	DiskUsage  			string    `orm:"type(text)"`
	Connections  		string    `orm:"type(text)"`
	Nic  				string    `orm:"type(text)"`
	IPV4  				string    `orm:"type(text)"`
	IPV6  				string    `orm:"type(text)"`
	Rx  				string    `orm:"type(text)"`
	Tx  				string    `orm:"type(text)"`
	RxGap  				string    `orm:"type(text)"`
	TxGap  				string    `orm:"type(text)"`
	Load  				string    `orm:"type(text)"`
	LoadCpu  			string    `orm:"type(text)"`
	LoadIO  			string    `orm:"type(text)"`
	PingEu  			string    `orm:"type(text)"`
	PingUs  			string    `orm:"type(text)"`
	PingAs  			string    `orm:"type(text)"`
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