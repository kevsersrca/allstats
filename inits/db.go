package inits

import (
	_ "allstats/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	dbname := "default"
	runmode := beego.AppConfig.String("runmode")
	datasource := beego.AppConfig.String("datasource")

	switch runmode {
	case "prod":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase(dbname, "mysql", datasource, 30)
		orm.SetMaxIdleConns(dbname, 100)
		orm.SetMaxOpenConns(dbname, 100)
	case "dev":
		orm.Debug = true
		fallthrough
	default:
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase(dbname, "mysql", datasource)
	}

	force, verbose := false, true
	err := orm.RunSyncdb(dbname, force, verbose)
	if err != nil {
		panic(err)
	}

	orm.RunCommand()
}
