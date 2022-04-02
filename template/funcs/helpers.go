package Func

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"text/template"
	"time"
)

func HandleError1(myC map[string]*template.Template, err1 error) (myc map[string]*template.Template, err error) {
	if err != nil {
		return myC, err1
	}
	return make(map[string]*template.Template), nil
}
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Println(s interface{}) {
	fmt.Println(s)
}

var dbs = Blogger{} // database

func RenderPage(writer http.ResponseWriter, addr string, data interface{}) {

	TemplateFile, err := template.ParseFiles(addr)

	if err != nil {
		Println(err)
	}

	TemplateFile.Execute(writer, data)

}

func RenderHomePage(writer http.ResponseWriter, addr string) {
	parsedTemplate, err := template.ParseFiles(addr)

	if err != nil {
		Println(err)
	}

	parsedTemplate.Execute(writer, addr)

}

func Home(writer http.ResponseWriter, req *http.Request) {
	RenderHomePage(writer, "./tmpl/frontpage.html")
}
func HomePage(writer http.ResponseWriter, req *http.Request) {
	RenderHomePage(writer, "./tmpl/basetemplate.html")
}
func AddBlog(writer http.ResponseWriter, req *http.Request) {
	fmt.Printf("%v", req.Body)
	e := req.ParseForm() // populate r.postForm with data from form fields

	HandleErr(e)
	timeNow := time.Now()
	date := timeNow.Format("Mon, 02 Jan 2006 15:04:05")
	title := req.PostForm.Get("title")
	body := req.PostForm.Get("body")
	if len(body) < 1 || body[0] == 32 || len(title) < 1 || title[0] == 32 {
		RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)
	} else {
		num := strconv.Itoa((217*rand.Intn(250-85) + (rand.Intn(430-127) + 127)) + (rand.Intn(530-327) + 127))
		b := Blog{
			SerialN: num,
			Title:   title,
			Body:    body,
			Delete:  "delete",
			Edit:    "edit",
			Date:    date,
		}
		dbs.Append(b)

		RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)
		// http.Redirect(writer, req, "/blog", http.StatusPermanentRedirect)

	}

}

func ViewBlog(writer http.ResponseWriter, req *http.Request) {
	RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)
}

func About(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "This is my About page")
}
