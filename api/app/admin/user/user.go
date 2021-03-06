package user

import (
	"ekgo/api/app/model"
	"ekgo/api/app/service"
	"ekgo/api/boot/db"
	"ekgo/api/lib/response"
	"github.com/gin-gonic/gin"
)

//接口服务
var Interface service.UserInterface

//分页
func Index(this *gin.Context) {
	var param = response.PageParam{CurrentPage: 1, PageSize: 10}
	this.ShouldBindQuery(&param)
	param.Filter = this.QueryArray("filter[]")
	Interface = &service.User{PageParam: param, Db: db.Master}

	this.SecureJSON(200, Interface.Index())

}

//查询
func Read(this *gin.Context) {
	var data = model.User{}
	this.ShouldBindUri(&data)
	Interface = &service.User{Model: data, Db: db.Master}

	this.SecureJSON(200, Interface.Show())
}

//创建
func Store(this *gin.Context) {
	var data = model.User{}
	this.ShouldBind(&data)
	Interface = &service.User{Model: data, Db: db.Master}

	this.SecureJSON(200, Interface.Store())
}

//修改
func Update(this *gin.Context) {
	var data = model.User{}
	this.ShouldBind(&data)
	Interface = &service.User{Model: data, Db: db.Master}

	this.SecureJSON(200, Interface.Update())
}

//删除
func Delete(this *gin.Context) {
	var data = model.User{}
	this.ShouldBindUri(&data)
	Interface = &service.User{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Delete())
}
