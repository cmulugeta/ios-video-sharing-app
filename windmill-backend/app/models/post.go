package models

type Post struct {
	PostId string `json:"postid"`
	UserId string `json:"userid"`
	Verified bool `json:"verified"`
	Username string `json:"username"`
	Caption string `json:"caption"`
	Comments []Comment `json:"comments"`
	NumLikes int `json:"numlikes"`
	Likers []string `json:"likers"`
	Url string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

type Comment struct {
	Username string `json:"username"`
	CommentData string `json:"commentdata"`
}
