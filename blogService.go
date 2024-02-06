package di

type BlogService struct {
}

func NewBlogService() BlogService {
	return BlogService{}
}

func (p *BlogService) Print() BlogPost {
	db := NewDataBase()
	return db.GetBlog()
}
