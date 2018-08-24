package controllers

import (
	"loveHome/models"
	"time"

	"encoding/json"
	_ "fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
)

type AreaControllers struct {
	beego.Controller
}

func (this *AreaControllers) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *AreaControllers) GetAreaInfo() {
	beego.Info("GetAreaInfo succ .......")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	//思路
	//0.连接redis
	//1.从缓存中读取数据
	//2.如果Redis中有数据，直接返回给前端AreaControllers
	//3.如果Redis中没有数据，从mysql 中查询

	//0连接缓存数据库

	cache_conn, err := cache.NewCache("redis", `{"key":"lovehome","conn":"127.0.0.1:6379", "dbNum":"0"}`) //参数连接缓存类型
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	/*
		//1从缓存中读取数据
		//参数  1 key 2 value  3 数据库能缓存多久
		cache_conn.Put("xixi", "lala", time.Second*300)
		//读
		value := cache_conn.Get("xixi")
		if value != nil {
			beego.Info("读取到缓存：", value)
			fmt.Printf("value=%s\n", value)
		}
	*/
	//从缓存中读数据
	areas_info_value := cache_conn.Get("area_info")
	if areas_info_value != nil {
		beego.Info("=====redis中有数据，数据返回给前端==========")
		var area_info interface{}
		json.Unmarshal(areas_info_value.([]byte), &area_info)
		resp["data"] = area_info
		return
	}
	o := orm.NewOrm()
	var areas []models.Area

	qs := o.QueryTable("area")

	num, err := qs.All(&areas)

	if err != nil {
		//返回错误信息给前端
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
	}

	resp["data"] = areas

	//没有缓存将缓存
	area_str, _ := json.Marshal(areas)
	if err := cache_conn.Put("area_info", area_str, time.Second*300); err != nil {
		//返回错误信息给前端
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	}
	return
}
