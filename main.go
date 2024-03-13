package main

import (
	_ "ParkMe/routers"
	"ParkMe/utils/rateLimiter"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func dbConnection() {
	sqlCon, err := beego.AppConfig.String("sql_conn")
	if err != nil {
		fmt.Println(err)
	}
	_ = orm.RegisterDataBase("default", "mysql", sqlCon)
	logs.Info("Database Connection Established")
}

func init() {

	logs.SetLogger(logs.AdapterFile, `{"filename":"error.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],"perm":"777", "maxlines":0, "maxsize":0, "daily":true, "maxdays":10, "color":true}`)

	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	dbConnection()

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "QueryType", "Skip", "Limit", "query_type", "skip", "limit", "cookie"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	requestTime, _ := beego.AppConfig.Int64("requestTime")
	totalRequest, _ := beego.AppConfig.Int64("totalRequest")

	limiter := rateLimiter.NewLimiter(rateLimiter.WithRate(time.Duration(requestTime)*time.Millisecond), rateLimiter.WithCapacity(uint(totalRequest)), rateLimiter.WithSessionKey(rateLimiter.RemoteIPSessionKey)) // Register limiter to specify route
	err := beego.InsertFilter("*", web.BeforeRouter, limiter)

	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
