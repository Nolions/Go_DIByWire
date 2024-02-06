package di

type BlogService struct {
	db Database
}

func NewBlogService(database Database) BlogService {
	return BlogService{
		db: database,
	}
}

func (p *BlogService) Print() BlogPost {
	//db := NewDataBase()
	return p.db.GetBlog()
}
