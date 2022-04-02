package Func

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"text/template"
)

func UpdateBlog(writer http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "SerialN")
	e := req.ParseForm() // populate r.postForm with data from form fields

	HandleErr(e)
	title := req.PostForm.Get("title")
	body := req.PostForm.Get("body")

	dbs.EditBlog(id, body, title)

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
