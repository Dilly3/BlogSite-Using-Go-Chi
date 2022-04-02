package Func

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

func (b *Blogger) EditBlog(SN string, body string, title string) {
	for ind, blog := range dbs.Blogs {
		if blog.SerialN == SN {
			dbs.Blogs[ind].Body = body
			dbs.Blogs[ind].Title = title
			break
		}
	}
}
