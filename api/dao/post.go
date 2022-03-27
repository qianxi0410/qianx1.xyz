package dao

import "gorm.io/gorm"

type Post struct {
	ID         string `json:"id" gorm:"primaryKey" redis:"id" form:"id"`
	Title      string `json:"title" gorm:"title" redis:"title" form:"title"`
	Tags       string `json:"tags" gorm:"tags" redis:"tags" form:"tags"`
	CreateTime int64  `json:"create_time" gorm:"create_time" redis:"create_time" form:"create_time"`
	UpdateTime int64  `json:"update_time" gorm:"update_time" redis:"update_time" form:"update_time"`
	Content    string `json:"content" gorm:"content" redis:"content" form:"content"`
	Prev       string `json:"prev" gorm:"prev" redis:"prev" form:"prev"`
	Next       string `json:"next" gorm:"next" redis:"next" form:"next"`
	DisplayId  string `json:"display_id" gorm:"display_id" redis:"display_id" form:"display_id"`
	Author     string `json:"author" gorm:"author" redis:"author" form:"author"`
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

func DeletePost(id string) error {
	db, err := Open()
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var cur Post
		tx.First(cur, id)
		tx.Delete(&Post{}, id)

		var next, prev Post

		if cur.Next != "" {
			tx.First(&next, cur.Next)
		}
		if cur.Prev != "" {
			tx.First(&prev, cur.Prev)
		}

		if next.ID != "" {
			prev.Next = next.ID
		}
		if prev.ID != "" {
			next.Prev = prev.ID
		}

		tx.Save(prev)
		tx.Save(next)

		return nil
	})

}

func UpdatePost(p Post) error {
	db, err := Open()
	if err != nil {
		return err
	}

	db.Save(&p)
	return nil
}

func CreatePost(p Post) error {
	db, err := Open()
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&p)

		var count int64
		tx.Model(&Post{}).Count(&count)

		if count != 0 {
			var next Post
			db.Order("create_time desc").First(&next)
			p.Next = next.ID

			next.Prev = p.ID
			db.Save(next)
		}

		return nil
	})

}
