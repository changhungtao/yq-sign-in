package main

import (
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	_ "github.com/mattn/go-sqlite3"
)

// Employee ...
type Employee struct {
	ID             int64
	EmployeeID     int64  `xorm:"varchar(200)"`
	Name           string `xorm:"varchar(200)"`
	Mobile         string `xorm:"varchar(200)"`
	Password       string `xorm:"varchar(200)"`
	UserToken      string `xorm:"varchar(200)"`
	LastSignInTime int64
}

var page = struct {
	Title string
}{"Welcome"}

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./build", ".html").Binary(Asset, AssetNames))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Page", page)
		ctx.View("index.html")
	})

	assetHandler := iris.StaticEmbeddedHandler("./build", Asset, AssetNames, false)
	app.SPA(assetHandler).AddIndexName("index.html")

	orm, err := xorm.NewEngine("sqlite3", "./database.db")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}

	iris.RegisterOnInterrupt(func() {
		err := orm.Close()
		if err != nil {
			app.Logger().Fatalf("orm failed to close: %v", err)
		}
	})

	err = orm.Sync2(new(Employee))

	if err != nil {
		app.Logger().Fatalf("orm failed to initialized Employee table: %v", err)
	}

	apiRoutes := app.Party("/api")
	// Todo 注册employee
	apiRoutes.Post("/employee", func(ctx iris.Context) {
		type LoginData struct {
			Mobile   string `json:"mobile"`
			Password string `json:"password"`
		}
		var loginData LoginData
		if err := ctx.ReadJSON(&loginData); err != nil {
			ctx.JSON(iris.Map{"code": "1", "message": "登录参数错误", "err": err})
		}
		// 判断employee是否存在
		has, err := orm.Exist(&Employee{Mobile: loginData.Mobile})
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "message": "判断employee是否存在失败！", "err": err})
		}
		if has {
			ctx.JSON(iris.Map{"code": "1", "message": "创建员工已存在！"})
		} else {
			// 进行登录
			data, err := Login(loginData.Mobile, loginData.Password)
			if err != nil {
				ctx.JSON(iris.Map{"code": "1", "message": "登录失败！", "err": err})
			}
			// 登录验证错误
			if data["code"] != "0" {
				ctx.JSON(iris.Map{"code": data["code"], "message": data["message"]})
			} else {
				resultMap := data["resultMap"].(map[string]interface{})
				userInfo := resultMap["userInfo"].(map[string]interface{})
				var employee Employee
				employee.EmployeeID, _ = strconv.ParseInt(userInfo["id"].(string), 10, 64)
				employee.Name = userInfo["name"].(string)
				employee.Mobile = loginData.Mobile
				employee.Password = loginData.Password
				employee.UserToken = userInfo["usertoken"].(string)
				_, err2 := orm.Insert(employee)
				if err2 != nil {
					ctx.JSON(iris.Map{"code": "1", "message": "创建失败！", "err": err})
				}
				ctx.JSON(iris.Map{"code": "0", "message": "创建成功！"})
			}
		}
	})

	// Delete
	apiRoutes.Delete("/employee/{id}", func(ctx iris.Context) {
		id, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		employee := &Employee{ID: id}
		has, err := orm.Get(employee)
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "message": "删除失败!", "err": err})
		} else {
			if !has {
				ctx.JSON(iris.Map{"code": "1", "message": "删除的员工不存在！"})
			} else {
				_, err := orm.Delete(employee)
				if err != nil {
					ctx.JSON(iris.Map{"code": "1", "message": "删除失败！", "err": err})
				} else {
					ctx.JSON(iris.Map{"code": "0", "employee": employee, "message": "删除成功！"})
				}
			}
		}
	})

	// Modify
	apiRoutes.Put("/employee/{id}", func(ctx iris.Context) {
		id, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		employee := &Employee{ID: id}
		has, err := orm.Get(employee)
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "message": "查询错误！", "err": err})
		} else {
			if !has {
				ctx.JSON(iris.Map{"code": "1", "message": "修改的员工不存在！"})
			} else {
				type LoginData struct {
					Mobile   string `json:"mobile"`
					Password string `json:"password"`
				}
				var loginData LoginData
				if err := ctx.ReadJSON(&loginData); err != nil {
					ctx.JSON(iris.Map{"code": "1", "message": "登录参数错误", "err": err})
				}
				data, err := Login(loginData.Mobile, loginData.Password)
				if err != nil {
					ctx.JSON(iris.Map{"code": "1", "message": "更新失败！", "err": err})
				}
				// 登录验证错误
				if data["code"] != "0" {
					ctx.JSON(iris.Map{"code": data["code"], "message": data["message"]})
				} else {
					resultMap := data["resultMap"].(map[string]interface{})
					userInfo := resultMap["userInfo"].(map[string]interface{})
					employee.EmployeeID, _ = strconv.ParseInt(userInfo["id"].(string), 10, 64)
					employee.Name = userInfo["name"].(string)
					employee.Mobile = loginData.Mobile
					employee.Password = loginData.Password
					employee.UserToken = userInfo["usertoken"].(string)
					_, err2 := orm.Id(id).Update(employee)
					if err2 != nil {
						ctx.JSON(iris.Map{"code": "1", "message": "更新失败！", "err": err})
					}
					ctx.JSON(iris.Map{"code": "0", "message": "更新成功！"})
				}
			}
		}
	})

	// Detail
	apiRoutes.Get("/employee/{id}", func(ctx iris.Context) {
		id, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		employee := &Employee{ID: id}
		has, err := orm.Get(employee)
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "message": "查询错误！", "err": err})
		} else {
			if !has {
				ctx.JSON(iris.Map{"code": "1", "message": "查询的员工不存在！"})
			} else {
				ctx.JSON(iris.Map{"code": "0", "employee": employee, "message": "查询成功！"})
			}
		}
	})

	apiRoutes.Get("/employee", func(ctx iris.Context) {
		employeeList := make([]Employee, 0)
		err := orm.Desc("i_d").Find(&employeeList)
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "message": "查询员工失败！", "err": err})
		} else {
			ctx.JSON(iris.Map{"code": "0", "count": len(employeeList), "employeeList": employeeList, "message": "查询成功！"})
		}
	})

	apiRoutes.Post("/employee/{id}/sign", func(ctx iris.Context) {
		id, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		employee := &Employee{ID: id}
		has, err := orm.Get(employee)
		if err != nil {
			ctx.JSON(iris.Map{"code": "1", "message": "查询错误！", "err": err})
		} else {
			if !has {
				ctx.JSON(iris.Map{"code": "1", "message": "签到的员工不存在！"})
			} else {
				data, err := SignIn(id, employee)
				if err != nil {
					ctx.JSON(iris.Map{"code": "1", "message": "更新失败！", "err": err})
				}
				// 登录验证错误
				if data["code"] != "0" {
					ctx.JSON(iris.Map{"code": data["code"], "message": data["message"]})
				} else {
					// 更新最后签
					employee.LastSignInTime = time.Now().Unix()
					_, err2 := orm.Id(id).Update(employee)
					if err2 != nil {
						ctx.JSON(iris.Map{"code": "1", "message": "更新失败！", "err": err})
					}
					resultMap := data["resultMap"].(map[string]interface{})
					ctx.JSON(iris.Map{"code": data["code"], "message": data["message"], "resultMap": resultMap})
				}
			}
		}
	})

	app.Run(iris.Addr(":8080"))
}
