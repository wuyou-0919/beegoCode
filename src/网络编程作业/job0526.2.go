package main

import (
	"database/sql"
	"fmt"
	_"网络编程作业/go-sql-driver/mysql"
)
func main() {
	connStr := "root:lch123456@tcp(127.0.0.1:3306)/big?charset=utf8"
	sql, err :=sql.Open("mysql",connStr)
	fmt.Printf("%T",sql)
	if err != nil{
		fmt.Println("数据库连接失败：",err.Error())
		return
	}
	defer sql.Close()
	stmt,err :=sql.Prepare(" update student set Name = ? , Sex = ?,Id = ? where Id = 11")
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("执行失败")
		return
	}
	result,err :=stmt.Exec("wang","ma",12)
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("执行失败")
		return
	}
	fmt.Println("执行成功")
	fmt.Println(result.RowsAffected())
	stmt2, err2 := sql.Prepare("insert into student(Id,Name,Sex,Old,Brithday,Score) values(11, 'long', 'we',12,'2001-02-14',81.5)")
	if err2 != nil{
		fmt.Println(err.Error())
		fmt.Println("数据插入执行失败")
		return
	}
	result,err =stmt2.Exec()
	fmt.Println("执行成功")
	fmt.Println(result.RowsAffected())
	sel := sql.QueryRow(" select * from student where id = 11")
	var id int
	var name, sex string
	var birthday string
	var old int
	var score float64
	err1 := sel.Scan(&id,&name,&sex,&old,&birthday,&score)
	fmt.Println(id,name,sex,old,birthday,score)
	if err1 != nil {
		fmt.Println("没有查询到数据")
		return
	}
	list, err :=sql.Query("select * from student")
	if err != nil {
		fmt.Println("数据库查询失败：",err.Error())
		return
	}
	columns, err := list.Columns()
	if err != nil {
		fmt.Println("获取失败：",err.Error())
		return
	}
	fmt.Println(columns)
	stmt3, err3 := sql.Prepare("delete from student where Id = 12")
	if err3 != nil{
		fmt.Println(err3.Error())
		fmt.Println("数据删除执行失败")
		return
	}
	result,err =stmt3.Exec()
	fmt.Println("执行成功")
	fmt.Println(result.RowsAffected())

}