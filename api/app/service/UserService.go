package service

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"ekgo/lib/orm"
	"ekgo/lib/response"
	"github.com/jinzhu/gorm"
)

type UserInterface interface {
	Index() *response.Write  //分页
	Store() *response.Write  //添加
	Update() *response.Write //修改
	Delete() *response.Write //删除
	Show() *response.Write   //查询
}

type User struct {
	PageParam response.PageParam
	Model     model.User
	Db        *gorm.DB
}

//分页
func (this *User) Index() *response.Write {
	var list = []model.User{}

	err := this.Db.Scopes(
		orm.WhereQueryBuild(this.PageParam.Filter),
		orm.Order(this.PageParam.Order),
		orm.Paginate(this.PageParam.PageSize, this.PageParam.CurrentPage),
	).Find(&list).Offset(0).Count(&this.PageParam.Total).Error

	if err != nil {
		return response.Fail("获取失败")
	}

	return &response.Write{
		Code: 200,
		Data: response.Page{
			Total: this.PageParam.Total,
			List:  list,
		},
	}
}

//查询单个
func (this *User) Show() *response.Write {
	err := db.Master.First(&this.Model, this.Model.Id).Error
	if err == nil {
		return response.Success("获取成功", this.Model)
	}

	return response.Success("获取失败")
}

//添加
func (this *User) Store() *response.Write {
	err := this.Db.Create(&this.Model).Error

	if err == nil {
		return response.Success("保存成功")
	}

	return response.Success("保存失败", err.Error())
}

//修改
func (this *User) Update() *response.Write {
	err := this.Db.Model(&this.Model).Update(this.Model).Error

	if err == nil {
		return response.Success("修改成功")
	}

	return response.Success("保存失败", err.Error())
}

//删除
func (this *User) Delete() *response.Write {
	err := this.Db.First(&this.Model, this.Model.Id).Error

	if err == nil {
		this.Db.Where("id=?", this.Model.Id).Delete(&this.Model)
		return response.Success("删除成功")
	}

	return response.Success("删除失败", err.Error())
}
