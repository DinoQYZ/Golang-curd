package main

import (
	"fmt"

	iris "github.com/kataras/iris/v12"
)

// delete >> /customer/{id: string}
func deleteCustomerData(ctx iris.Context) {
	id := ctx.Params().Get("id")
	valueExist := false
	target := CustomerData{}
	for i, ct := range cstmData {
		if ct.id == id {
			target = ct
			deleteFromSlice(i)
			valueExist = true
		}
	}

	resp := fmt.Sprintf("Customer data [%v] has been deleted", target.name)

	if valueExist {
		fmt.Printf("<<%v>>", resp)
		ctx.JSON(iris.Map{"response": resp})
	} else {
		fmt.Println("<<id not found>>")
		ctx.JSON(iris.Map{"invalid_id": "id not found"})
	}
}
