package entities

type Comment struct {
	id      string
	post    Post
	message string
	upvotes []Upvote
}

func NewComment(message string, post Post) Comment {
	return Comment{
		message: message,
		upvotes: []Upvote{},
		post:    post,
	}
}

func LoadComment(id string, post Post, message string, upvotes []Upvote) Comment {
	return Comment{
		id:      id,
		message: message,
		upvotes: upvotes,
		post:    post,
	}
}

func (comment *Comment) GetId() string {
	return comment.id
}

func (comment *Comment) GetMessage() string {
	return comment.message
}

func (comment *Comment) GetUpvotes() []Upvote {
	return comment.upvotes
}

func (comment *Comment) GetPost() Post {
	return comment.post
}

func (comment *Comment) AddUpvote(upvote Upvote) {
	comment.upvotes = append(comment.upvotes, upvote)
}

func (comment *Comment) RemoveUpvote(id string) {
	for i, upvote := range comment.upvotes {
		if upvote.id == id {
			comment.upvotes[i] = comment.upvotes[len(comment.upvotes)-1]
		}
	}
	comment.upvotes = comment.upvotes[:len(comment.upvotes)-1]
}
