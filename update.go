package main

import (
	"fmt"
	"strconv"

	iris "github.com/kataras/iris/v12"
)

// put >> /customer/{id: string}
func updateCustomerData(ctx iris.Context) {
	valueExist := false

	id := ctx.Params().Get("id")
	inputName := ctx.PostValue("name")
	inputAge, _ := strconv.ParseUint(ctx.PostValue("age"), 10, 32)
	inputEmail := ctx.PostValue("email")

	updateTarget := CustomerData{
		id:    id,
		name:  inputName,
		age:   inputAge,
		email: inputEmail,
	}

	var originalTargetName string

	for i, ct := range cstmData {
		if ct.id == id {
			originalTargetName = ct.name
			if updateTarget.name != "" {
				cstmData[i].name = updateTarget.name
			}
			if updateTarget.age != 0 {
				cstmData[i].age = updateTarget.age
			}
			if updateTarget.email != "" {
				cstmData[i].email = updateTarget.email
			}
			valueExist = true
		}
	}

	resp := fmt.Sprintf("Original customer data [%v] has been updated", originalTargetName)

	if valueExist {
		fmt.Printf("<<%v>>\n", resp)
		ctx.JSON(iris.Map{"response": resp})
	} else {
		fmt.Println("<<id not found>>")
		ctx.JSON(iris.Map{"invalid_id": "id not found"})
	}
}
