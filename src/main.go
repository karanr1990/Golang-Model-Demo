package main

import (
	"fmt"

	"github.abc.com/Quest/Golang-Sql-Demo/src/model"
)

func main() {
	//Create object of struct UserModel

	user := model.UserModel{}
	user.SetuserName("John Martin")
	user.GetuserName()
	user.Setemail("john_martin@gmail.com")
	user.Getemail()
	user.Setpassword("johnMartin234")
	user.GetPassword()
	fmt.Println(user)
}
