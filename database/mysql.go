package database

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type PrimaryKey struct {
	Name string
	Type int
}
type ActiveDb struct {
}

type IBaseModel interface {
	Get(primary int64) ActiveDb
	FindOne(sql string) ActiveDb
	GetOrCreate([]interface{})
	FindAll(sql string)
}

type UserInfoModel struct {
	ActiveDb
	Uid      int64
	Name     string
	PostCode string
	Address  string
	Phone    string
	Age      int
}

func (userInfo *UserInfoModel) Get(uid int64) *UserInfoModel {
	db := GetDriver()
	rows, _ := db.Query("select * from user_info where uid=?", uid)
	user := UserInfoModel{}
	if rows.Next() {
		rows.Scan(&user.Uid, &user.Name, &user.PostCode, &user.Address, &user.Phone, &user.Age)
	}
	return &user
}

func GetDriver() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/caiop?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

func Driver() *sql.DB {
	db := GetDriver()
	//query
	rows, _ := db.Query("select * from apis where id=?", 1)
	for rows.Next() {
		var apis = Apis{}
		var url, request, response, demo string
		_ = rows.Scan(&apis.Id, &url, &request, &response, &demo)
		apis.Url = url
		apis.Response = response
		apis.Request = request
		apis.Demo = demo
		fmt.Println(apis)
		fmt.Println(url)
		fmt.Println(request)
		fmt.Println(response)
		fmt.Println(demo)
	}
	//insert
	//_, err = db.Prepare("")
	//if err != nil {
	//	panic(err)
	//}

	return db
}

func GetApisById(id int32) Apis {
	db := GetDriver()
	rows, _ := db.Query("select * from apis where id=?", id)

	for rows.Next() {
		var apis = Apis{}
		var url, request, response, demo string
		_ = rows.Scan(&apis.Id, &url, &request, &response, &demo)
		apis.Url = url
		apis.Response = response
		apis.Demo = demo
		return apis
	}
	return Apis{}
}

func GetApisListByIds(result *[]Apis, ids []int32) {
	db := GetDriver()
	rows, _ := db.Query("select * from apis where id in ( ? )", ids)
	for rows.Next() {

	}

}

func Init() {

	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	_ = orm.RegisterDataBase("default", "mysql", "root:root@/127.0.0.1?charset=utf8", 30)
	// 注册定义的 model
	orm.RegisterModel(new(Apis))
	// 创建 table
	_ = orm.RunSyncdb("default", false, true)

}

type Apis struct {
	Id       int
	Url      string
	Request  string
	Response string
	Demo     string
}

type User struct {
	Id int
}

func main() {
	//Init()
	apis := GetApisById(1)
	_ = orm.RegisterDataBase("default", "mysql", "root:root@/127.0.0.1:3306/caiop?charset=utf8", 30)

	//fmt.Println(apis)
	//Driver()
	orm := orm.NewOrm()

	_, _ = orm.Delete(apis)

}
