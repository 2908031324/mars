
package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

//https://www.runoob.com/go/go-structures.html 菜鸟教程

var db *sql.DB

func initDB()(err error){
	//dsn := "ruby:123456@tcp(111.230.54.23:3306)/dc_cusn"
	dsn := "root:root@tcp(127.0.0.1:3306)/dc_cusn"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		//dsn配置格式不正确
		return
	}

	err = db.Ping()
	if err != nil{
		//校验连接数据库失败
		return
	}

	db.SetMaxOpenConns(6)//设置最大连接数
	db.SetMaxIdleConns(3)//设置最大空闲数

	return

}

type userInfo struct {
	id int
	username string
	idcard int
}


func queryOne(id int, fields string) {
	err := initDB()
	if err != nil {
		fmt.Printf("数据库连接失败：%v", err)
		return
	}

	var info userInfo
	sqlStr := "select " + fields + " from `dc_user` where id = ?;"
	fmt.Println(sqlStr)
	err = db.QueryRow(sqlStr, id).Scan(&info.id, &info.username, &info.idcard)
	if err != nil {
		fmt.Printf("查询数据失败： %v", err)
		return
	}


	fmt.Printf("success result: %v", info)


}

func main(){

	queryOne(133, "id,username,idcard")

	//err := initDB()
	//if err != nil{
	//	fmt.Printf("数据库连接失败：%s", err)
	//}
	//
	//var u1 userInfo
	//DcUserAll := "select id, username, idcard from `user` where id = 131;"
	//rowObj := db.QueryRow(DcUserAll)
	//
	////for i:=1; i<10; i++{
	////	db.QueryRow(DcUserAll)
	////	fmt.Printf("当前第%d个链接\n", i)
	////}
	//
	//err = rowObj.Scan(&u1.id, &u1.username, &u1.idcard)
	//fmt.Println(err)
	//if err != nil{
	//	fmt.Printf("遍历数据失败了：%v", err)
	//	return
	//}
	//fmt.Printf("%v", u1)



}

//func queryRowDemo() {
//	sqlStr := "select id, name, age from user where id=?"
//	var u user
//	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
//	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
//	if err != nil {
//		fmt.Printf("scan failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
//}

