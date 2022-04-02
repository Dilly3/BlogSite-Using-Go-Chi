package Func

type ID = string
type Blog struct {
	SerialN ID
	Title   string
	Body    string
	Edit    string
	Delete  string
	Date    string
}

type Blogger struct {
	Blogs []Blog
}
