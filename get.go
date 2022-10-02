package main

import (
	"fmt"

	iris "github.com/kataras/iris/v12"
)

// get >> /customer/
func getCustomerData(ctx iris.Context) {
	fmt.Println("<<Customer data sent>>")
	dataMap := getFormatFullData()
	ctx.JSON(iris.Map{"response": dataMap})
}

// get >> /customer/{id: string}
func getTargetCustomerData(ctx iris.Context) {
	id := ctx.Params().Get("id")
	valueExist := false
	target := CustomerData{}
	for _, ct := range cstmData {
		if ct.id == id {
			target = ct
			valueExist = true
		}
	}

	targetData := getFormatTargetData(target)

	if valueExist {
		ctx.JSON(iris.Map{"response": targetData})
	} else {
		fmt.Println("<<id not found>>")
		ctx.JSON(iris.Map{"invalid_id": "id not found"})
	}
}
