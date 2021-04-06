package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/hwkkevin/beego_demo/models"
	"github.com/hwkkevin/beego_demo/util"
	"log"
	"strconv"
)

type IndexController struct {
	BaseController
}

type MoneyForm struct {
	Id int `form:"id"`
	Username string `form:"username"`
	Money float64 `form:"money"`
	Status int8 `form:status`
	Address string `form:"address"`
}

func (b *IndexController) Index()  {
	var (
		page       int
		pagesize   int = 15
		offset     int
		list       []*models.Money
		keyword    string
	)
	keyword = b.GetString("keyword")
	if page,_ = b.GetInt("page") ; page<1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	query := b.o.QueryTable(new(models.Money).TableName())
	if keyword != ""{
		query = query.Filter("username",keyword)
	}   
	count , _ := query.Count()
	if count > 0 {
		query.OrderBy("-id","-created").Limit(pagesize,offset).All(&list)
	}
	b.Data["keyword"] = keyword
	b.Data["count"] = count
	b.Data["list"] = list
	b.Data["pagebar"] = util.NewPager(page, int(count), pagesize,
		fmt.Sprintf("?keyword=%s", keyword), true).ToString()
	b.GetTemplate("index")
}
 


func (b *IndexController) Create() {
	var money MoneyForm
	fmt.Println("对象%v",money)
	id,errId := b.GetInt("id")
	if b.Ctx.Request.Method == "POST" {
		if err := b.ParseForm(&money); err != nil{
			b.History("表单数据为空","")
		}
		
		valid := validation.Validation{}
		valid.Required(money.Username,"username")
		valid.MaxSize(money.Username,20,"usernameMax")
		fmt.Println("对象%v",money)
		if valid.HasErrors() {
			//如果有错误信息，证明验证没通过
			//打印错误信息
			for _,err := range valid.Errors{
				log.Println(err.Key,err.Message)
				b.History(err.Message,"")
			}
		} 
		moneyModel := models.Money{Username:money.Username} 
		if err := b.o.Read(&moneyModel,"Username"); err == nil && id==0 {
			b.History("账号已存在","")
		}
		moneyModel.Username = money.Username
		moneyModel.Money = money.Money
		moneyModel.Status = money.Status
		moneyModel.Address = money.Address
		if id > 0 {
			moneyModel.Id = id
			url := "/create?id="+ strconv.Itoa(id)
			if num, err := b.o.Update(&moneyModel); err == nil {
				b.History("更新成功",url)
				fmt.Println(num)
			}
		}else{
			id, err := b.o.Insert(&moneyModel)
			fmt.Println("id类型:%T",id)
			url := "/create?id=" + strconv.FormatInt(id,0)
			if err == nil {
				b.History("新增成功",url)
				fmt.Println(id)
			}
		}
		b.Data["money"] = moneyModel
	}
	
	if errId == nil{
		if id !=0 {
			moneyModel := models.Money{Id: id} 
			err := b.o.Read(&moneyModel)
			if err != nil {
				b.History("暂无该信息","/")
			}else {
				b.Data["money"] = moneyModel
			}
		}
	}else{
		b.History("暂无该信息","/")
	}
	b.GetTemplate("create") 
}


