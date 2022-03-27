package dao

type Post struct {
	ID         string `json:"id" gorm:"primaryKey" redis:"id"`
	Title      string `json:"title" gorm:"title" redis:"title"`
	Tags       string `json:"tags" gorm:"tags" redis:"tags"`
	CreateTime string `json:"create_time" gorm:"create_time" redis:"create_time"`
	UpdateTime string `json:"update_time" gorm:"update_time" redis:"update_time"`
	Content    string `json:"content" gorm:"content" redis:"content"`
	Prev       string `json:"prev" gorm:"prev" redis:"prev"`
	Next       string `json:"next" gorm:"next" redis:"next"`
	DisplayId  string `json:"display_id" gorm:"display_id" redis:"display_id"`
	Author     string `json:"author" gorm:"author" redis:"author"`
}

// GetPost get a post from db by ids
func GetPost(id string) (Post, error) {
	db, err := Open()
	if err != nil {
		return Post{}, err
	}

	var post Post
	db.First(&post, id)

	return post, nil
}

func GetPostWithDisplayId(displayId string) (Post, error) {
	db, err := Open()
	if err != nil {
		return Post{}, err
	}

	var post Post
	db.First(&post, "display_id = ?", displayId)

	return post, nil
}

func GetPosts(offset, size int) ([]Post, error) {
	db, err := Open()
	if err != nil {
		return nil, err
	}

	var posts []Post
	db.Select([]string{"id", "title", "tags", "create_time"}).Order("create_time desc").Offset(offset).Limit(size).Find(&posts)

	return posts, nil
}

func GetPostsCount() (int64, error) {
	db, err := Open()
	if err != nil {
		return 0, err
	}

	var count int64
	db.Model(&Post{}).Count(&count)

	return count, nil
}
