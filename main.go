package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wisphes/goAndMongoDB/controllers"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()
	us := controllers.NewUserController(getSession())
	r.GET("/user/:id", us.GetUser)
	r.POST("/user", us.CreateUser)
	r.DELETE("/user/:id", us.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		panic(err)
	}
	return s
}
