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
