package controllers

import (
	"encoding/json"
	"loveHome/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserControllers struct {
	beego.Controller
}

func (this *UserControllers) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

//登陆
func (this *UserControllers) Login() {
	beego.Info("===========login=========")
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	var loginRequestMap = make(map[string]interface{})

	//1.得到客户端请求的json 数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &loginRequestMap)
	beego.Info("mobile=", loginRequestMap["mobile"])
	beego.Info("password=", loginRequestMap["password"])
	//2.判断数据的合法性
	if loginRequestMap["mobile"] == "" || loginRequestMap["password"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	//3.查询数据库
	o := orm.NewOrm()
	var user models.User
	qs := o.QueryTable("user")
	if err := qs.Filter("mobile", loginRequestMap["mobile"]).One(&user); err != nil {
		//查询失败
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	//4.比对密码
	if user.Password_hash != loginRequestMap["password"].(string) {
		resp["errno"] = models.RECODE_PWDERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
		return
	}
	beego.Info("====登陆成功====", user.Name)

	this.SetSession("name", user.Mobile)
	//this.SetSession("user_id", user.Id)
	this.SetSession("mobile", user.Mobile)
}

//注册
func (this *UserControllers) Reg() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_SESSIONERR
	resp["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
	defer this.RetData(resp)

	//存储前端的信息
	var regRequesMap = make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequesMap)
	//1.得到前端请求的数据
	//2.判断数据的合理性
	//3.将数据存储到MySQL的user表中
	//4.将当前用户的信息存储到session中

	beego.Info("mobile=", regRequesMap["mobile"])
	beego.Info("password=", regRequesMap["password"])
	beego.Info("sms_code=", regRequesMap["sms_code"])
	//this.ServeJSON()

	if regRequesMap["mobile"] == "" || regRequesMap["password"] == "" || regRequesMap["sms_code"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	user := models.User{}
	user.Mobile = regRequesMap["mobile"].(string)
	user.Password_hash = regRequesMap["password"].(string)
	user.Name = regRequesMap["mobile"].(string)

	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Info("注册成  id= ", id)

	this.SetSession("name", user.Mobile)
	this.SetSession("user_id", user.Id)
	this.SetSession("mobile", user.Mobile)
}
