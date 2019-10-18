package serializer

import (
	"blog/models"
	"github.com/russross/blackfriday"
)

// Post response
type Post struct {
	ID uint 
	Title string `json:"title"`
	Content string `json:"content"`
	Time   string `json:"time"`
	Views uint `json:"views"`
}

// PostResponse response for post
type PostResponse struct {
	Response
	Post Post `json:"post"`
}


// PostsResponse posts 
type PostsResponse struct {
	Response
	Posts []Post `json:"posts"`
}


// BuildPosts posts 
func  BuildPosts(posts []models.Post) []Post {
	var resPost []Post
	for _, post:= range posts {

		p := Post{
			ID: post.ID,
			Title: post.Title,
			Time:post.UpdatedAt.Format("2006-01-02"),
			Views: post.View,
		}
	 resPost = append(resPost, p)
	}
	return resPost
}

// BuildPostsResponse response
func BuildPostsResponse(posts []models.Post) PostsResponse {
	return PostsResponse{
		Posts: BuildPosts(posts),
	}
}

// BuildPost post
func BuildPost(post models.Post) Post {
	p := Post{
		ID: post.ID,
		Title: post.Title,
		Content: string(blackfriday.Run([]byte(post.Content))),
		Time:post.UpdatedAt.Format("2006-01-02"),
		Views: post.View,
	}
	return p
}

// BuildPostResponse post response
func BuildPostResponse(post models.Post) PostResponse {
	return PostResponse{
		Post: BuildPost(post),
	}
}