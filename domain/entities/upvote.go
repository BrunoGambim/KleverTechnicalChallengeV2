package entities

import (
	"fmt"
	"time"
)

type Upvote struct {
	id         string
	upvoteType string
	createdAt  int64
}

func NewUpvote(upvoteType string) Upvote {
	createdAt := time.Now().Unix()
	return Upvote{
		upvoteType: upvoteType,
		createdAt:  createdAt,
	}
}

func LoadUpvote(id string, upvoteType string, createdAt int64) Upvote {
	return Upvote{
		id:         id,
		upvoteType: upvoteType,
		createdAt:  createdAt,
	}
}

func (upvote *Upvote) GetId() string {
	return upvote.id
}

func (upvote *Upvote) GetUpvoteType() string {
	return upvote.upvoteType
}

func (upvote *Upvote) GetCreatedAt() int64 {
	return upvote.createdAt
}

func (upvote *Upvote) GetFormatedCreatedAt() string {
	createdAtTime := time.Unix(upvote.createdAt, 0)
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		createdAtTime.Year(), createdAtTime.Month(), createdAtTime.Day(),
		createdAtTime.Hour(), createdAtTime.Minute(), createdAtTime.Second())
}
