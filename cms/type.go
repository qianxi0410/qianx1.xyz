package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type R struct {
	Data string `json:"data"`
	Err  string `json:"err"`
}

type Post struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Prev       string `json:"prev"`
	Next       string `json:"next"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Author     string `json:"author"`
	Tags       string `json:"tags"`
}

type Label struct {
	ID          string `json:"id"`
	NodeId      string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     string `json:"default"`
	Description string `json:"description"`
}
