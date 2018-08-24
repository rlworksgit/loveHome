package controllers

import (
	"loveHome/models"

	"github.com/astaxie/beego"
)

type SessionControllers struct {
	beego.Controller
}

func (this *SessionControllers) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *SessionControllers) DeleteSession() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	this.DelSession("name")
	this.DelSession("user_id")
	this.DelSession("mobile")
	defer this.RetData(resp)

}

func (this *SessionControllers) GetSessionInfo() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_SESSIONERR
	resp["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
	defer this.RetData(resp)

	name_map := make(map[string]interface{})
	name := this.GetSession("name")
	if name != nil {
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		name_map["name"] = name.(string)
		resp["data"] = name_map

	}
}
