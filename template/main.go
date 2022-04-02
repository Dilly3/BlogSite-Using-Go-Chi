package main

import (
	H "blogsite/Template/Funcs"
	"time"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	type PortID = string
	const PORT_NUMBER PortID = ":9020" // port number

	chiRouter := chi.NewRouter() // initiate chi router

	chiRouter.Get("/", H.HomePage)
	chiRouter.Get("/view", H.Home)
	chiRouter.Get("/delete/{SerialN}", H.DeleteBlog)
	chiRouter.Get("/blog", H.ViewBlog)
	chiRouter.Post("/blogs", H.AddBlog)
	chiRouter.Post("/update/{SerialN}", H.UpdateBlog)
	chiRouter.Get("/edit/{SerialN}", H.EditBlog)
	chiRouter.Get("/cap/{SerialN}", H.Capitalize)
	chiRouter.Get("/uncap/{SerialN}", H.UnCapitalize)
	chiRouter.Get("/readmore/{SerialN}", H.ReadMore)
	chiRouter.Get("/home/{SerialN}", H.ViewBlog)

	fmt.Printf("Server Running at Port %v \n %v\n", PORT_NUMBER, time.Now())
	//Listening
	e := http.ListenAndServe(PORT_NUMBER, chiRouter) // listening to requests from connections
	H.HandleErr(e)

}
