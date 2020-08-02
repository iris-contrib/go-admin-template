package main

import (
	"go-admin-template/models"
	"go-admin-template/pages"
	"go-admin-template/tables"

	_ "github.com/GoAdminGroup/go-admin/adapter/iris"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                    // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"

	"github.com/kataras/iris/v12"
)

func main() {
	startServer()
}

func startServer() {
	app := iris.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	//cfg := config.Config{
	//	Databases: config.DatabaseList{
	//		"default": {
	//			Host:       "127.0.0.1",
	//			Port:       "3306",
	//			User:       "root",
	//			Pwd:        "root",
	//			Name:       "go-admin",
	//			MaxIdleCon: 50,
	//			MaxOpenCon: 150,
	//			Driver:     db.DriverMysql,
	//		},
	//	},
	//	UrlPrefix: "admin",
	//	IndexUrl:  "/",
	//	Debug:     true,
	//	Language:  language.CN,
	//}

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		AddGenerator("external", tables.GetExternalTable).
		Use(app); err != nil {
		panic(err)
	}

	models.Init(eng.SqliteConnection())

	app.HandleDir("/uploads", iris.Dir("./uploads"), iris.DirOptions{
		Compress: true,
	})

	eng.HTML("GET", "/admin", pages.DashboardPage)
	eng.HTML("GET", "/admin/form", pages.GetFormContent)
	eng.HTML("GET", "/admin/table", pages.GetTableContent)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	app.Listen(":8080")
}
