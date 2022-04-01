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

func (b *Blogger) Append(blog Blog) {
	b.Blogs = append(b.Blogs, blog)

}
func (b *Blogger) Delete(sn string) {
	for ind, blog := range b.Blogs {
		if blog.SerialN == sn {
			b.Blogs = append(b.Blogs[:ind], b.Blogs[ind+1:]...)
			break
		}
	}
}

func (b *Blogger) EditBlog(sn string, strBody string, strTitle string) {
	for _, blog := range b.Blogs {
		if blog.SerialN == sn {
			blog.Body = strBody
			blog.Title = strTitle
			break
		}
	}
}

//RandInt := rand.Intn(13-2) + 2
