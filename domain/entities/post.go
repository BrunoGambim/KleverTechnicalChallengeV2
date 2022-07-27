package entities

type Post struct {
	id     string
	userId string
	title  string
	body   string
}

func NewPost(userId string, title string, body string) Post {
	return Post{
		userId: userId,
		title:  title,
		body:   body,
	}
}

func LoadPost(id string, userId string, title string, body string) Post {
	return Post{
		id:     id,
		userId: userId,
		title:  title,
		body:   body,
	}
}

func (post *Post) GetId() string {
	return post.id
}

func (post *Post) GetUserId() string {
	return post.userId
}

func (post *Post) GetTitle() string {
	return post.title
}

func (post *Post) GetBody() string {
	return post.body
}
