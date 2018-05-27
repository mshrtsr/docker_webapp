package main

import (
	"fmt"
	"time"
	"../webapp/CRUD"
)

func main(){
	tb_name := "test_table"
	var user = CRUD.User{0, "Taro", "test@email.com", time.Now(), time.Now()}
	var tmp CRUD.User
	CRUD.DropTable(tb_name)
	CRUD.CreateTable(tb_name)
	CRUD.CreateData(user.Name, user.Email, tb_name)
	CRUD.CreateData(user.Name, user.Email, tb_name)
	CRUD.CreateData(user.Name, user.Email, tb_name)
	CRUD.ReadDataAll(tb_name)
	tmp = CRUD.ReadData(2,tb_name)
	//loc, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Data row = (%d, %s, %s, %d, %d)\n", tmp.Id, tmp.Name, tmp.Email, tmp.Created_at, tmp.Updated_at)
	fmt.Println(tmp.Updated_at)
	fmt.Println(tmp.Updated_at.Format(time.RFC3339Nano))
	user.Id = 2
	user.Name = "Jiro"
	CRUD.UpdateData(user.Id, user.Name, user.Email, tb_name)
	CRUD.ReadDataAll(tb_name)
	CRUD.DeleteData(3, tb_name)
	CRUD.ReadDataAll(tb_name)
}
