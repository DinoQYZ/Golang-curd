package main

import (
	iris "github.com/kataras/iris/v12"
)

var cstmData = make([]CustomerData, 0)

type CustomerData struct {
	id    string
	name  string
	age   uint64
	email string
}

func main() {
	app := iris.New()

	//default data
	firstData()

	setupRoutes(app)

	app.Run(iris.Addr(":4000"))
}

func setupRoutes(app *iris.Application) {
	//default html
	app.RegisterView(iris.HTML("./views", ".html"))

	//default
	app.Get("/", defaultGreet)

	//get customer
	app.Get("/customer", getCustomerData)

	//get single user
	app.Get("/customer/{id: string}", getTargetCustomerData)

	//add customer
	app.Post("/customer", addCustomerData)

	//update customer
	app.Put("/customer/{id: string}", updateCustomerData)

	//delete customer
	app.Delete("/customer/{id: string}", deleteCustomerData)
}
