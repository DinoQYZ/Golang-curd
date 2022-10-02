package main

import (
	"fmt"
	"strconv"

	uuid "github.com/google/uuid"
	iris "github.com/kataras/iris/v12"
)

func defaultGreet(ctx iris.Context) {
	fmt.Println("<<Greeting message sent>>")
	ctx.JSON(iris.Map{"message": "Welcome customer system"})
}

func firstData() {
	id1 := fmt.Sprintf("%v", uuid.New())
	id2 := fmt.Sprintf("%v", uuid.New())
	data1 := CustomerData{
		id:    id1,
		name:  "Winter",
		age:   21,
		email: "winter@gm.com",
	}

	data2 := CustomerData{
		id:    id2,
		name:  "Karina",
		age:   22,
		email: "karina@gm.com",
	}

	cstmData = append(cstmData, data1)
	cstmData = append(cstmData, data2)
}

func getFormatFullData() map[int]map[string]string {
	dataMap := map[int]map[string]string{}
	for i, ct := range cstmData {
		nestedMap := map[string]string{
			"id":    ct.id,
			"name":  ct.name,
			"age":   strconv.FormatUint(uint64(ct.age), 10),
			"email": ct.email,
		}
		dataMap[i] = nestedMap
	}
	return dataMap
}

func getFormatTargetData(target CustomerData) map[string]string {
	targetData := map[string]string{
		"id":    target.id,
		"name":  target.name,
		"age":   strconv.FormatUint(uint64(target.age), 10),
		"email": target.email,
	}
	return targetData
}

func deleteFromSlice(i int) {
	n := len(cstmData)
	if i != n-1 {
		cstmData[i] = cstmData[n-1]
	}
	cstmData = cstmData[:n-1]
}
