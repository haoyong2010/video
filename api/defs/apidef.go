package defs

//reqeusts
type UserCredebtial struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//data model
type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string //login name
	TTL      int64
}
