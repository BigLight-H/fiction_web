package routers

import (
	"fiction_web/controllers"
	"fiction_web/controllers/users"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //搜索图书
	beego.Router("/home/:name", &controllers.MainController{},"get:Home")
    //获取图书目录列表
	beego.Router("/book/list", &controllers.MainController{},"get:List")
    //获取图书内容详情
	beego.Router("/book/detail", &controllers.MainController{},"get:Detail")

    //登录/注册
	beego.Router("/register", &controllers.LoginController{},"post:Register")
	beego.Router("/login", &controllers.LoginController{},"post:Login")

    //用户登录后
    //获取书架数据
	beego.Router("/user/books", &users.UserController{},"get:GetBookshelf")
    //添加到书架
	beego.Router("/user/add/books", &users.UserController{},"post:AddBookshelf")
    //查看书本是否已经存在书架
	beego.Router("/user/verification/books", &users.UserController{},"get:VerificationBook")
	//更新图书阅读进度
	beego.Router("/user/books/update", &users.UserController{},"post:UpdateBookSchedule")
	//删除图书
	beego.Router("/user/books/del", &users.UserController{},"post:DelBook")
    //退出登录
	beego.Router("/user/Logout", &users.UserController{},"get:Logout")
}
