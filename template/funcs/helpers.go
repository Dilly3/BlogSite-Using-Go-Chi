package Func

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
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

	TemplateFile, err := template.ParseFiles(fmt.Sprintf(addr))

	if err != nil {
		Println(err)
	}

	TemplateFile.Execute(writer, data)

}

func RenderHomePage(writer http.ResponseWriter, addr string) {
	parsedTemplate, err := template.ParseFiles(fmt.Sprintf(addr))

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
func UpdateBlog(writer http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "SerialN")
	e := req.ParseForm() // populate r.postForm with data from form fields

	HandleErr(e)
	title := req.PostForm.Get("title")
	body := req.PostForm.Get("body")

	for ind, blog := range dbs.Blogs {
		if blog.SerialN == id {
			dbs.Blogs[ind].Title = title
			dbs.Blogs[ind].Body = body
			break
		}
	}
	// http.Redirect(writer, req, "/blog", http.StatusPermanentRedirect)
	RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)
}

func DeleteBlog(writer http.ResponseWriter, req *http.Request) {
	fmt.Printf("%v", req)
	SN := chi.URLParam(req, "SerialN")
	fmt.Println(SN)

	for _, blog := range dbs.Blogs {
		if blog.SerialN == SN {
			dbs.Delete(blog.SerialN)

		}
	}
	RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)

}
func Capitalize(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")

	for ind, blog := range dbs.Blogs {
		if blog.SerialN == SN {

			strH := blog.Title

			strH2 := strings.ToUpper(strH)

			dbs.Blogs[ind].Title = strH2
		}
	}

	RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)

}

func ReadMore(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")
	//template.parse and execute

	parsedTemplate, err := template.ParseFiles("./tmpl/readmore.html")

	if err != nil {
		Println(err)
	}

	for _, blog := range dbs.Blogs {
		if SN == blog.SerialN {
			parsedTemplate.Execute(writer, blog)
		}

	}

}
func UnCapitalize(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")

	for ind, blog := range dbs.Blogs {
		if blog.SerialN == SN {
			str := blog.Body
			strH := blog.Title
			str2 := strings.ToLower(str)
			strH2 := strings.ToLower(strH)
			dbs.Blogs[ind].Body = str2
			dbs.Blogs[ind].Title = strH2
		}
	}

	RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)

}

func EditBlog(writer http.ResponseWriter, req *http.Request) {
	SN := chi.URLParam(req, "SerialN")

	for _, blog := range dbs.Blogs {
		if blog.SerialN == SN {

			//This points to the html location
			TemplateFile, err1 := template.ParseFiles("./tmpl/index.html")
			HandleErr(err1)

			err2 := TemplateFile.Execute(writer, blog)
			HandleErr(err2)

		}
	}
}
func ViewBlog(writer http.ResponseWriter, req *http.Request) {
	RenderPage(writer, "./tmpl/basetemplate.html", dbs.Blogs)
}

func About(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "This is my About page")
}
