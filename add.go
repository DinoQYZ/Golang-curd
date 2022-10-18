package main

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// post >> /customer/
func addCustomerData(ctx iris.Context) {
	inputName := ctx.PostValue("name")
	inputAge, _ := strconv.ParseUint(ctx.PostValue("age"), 10, 32)
	inputEmail := ctx.PostValue("email")
	id := fmt.Sprintf("%v", uuid.New())

	newData := CustomerData{
		id:    id,
		name:  inputName,
		age:   inputAge,
		email: inputEmail,
	}

	cstmData = append(cstmData, newData)

	fmt.Println("<<New data added>>")
	fmt.Printf("<<New customer data [%v] sent>>\n", newData.name)
	ctx.JSON(iris.Map{"response": getFormatTargetData(newData)})
}
