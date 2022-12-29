package models

import (
	"errors"
	"time"
)

const maxLengthComments = 400

//Review present an anon review from some website
type Review struct {
	Id      int64
	Stars   int       //1 - 5
	Comment string    //max 400 chars
	Date    time.Time //created at
}

//CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Stars   int    `json:stars`
	Comment string `json:comment`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("stars must be between 1 and 5")
	}

	if len(cmd.Comment) > maxLengthComments {
		return errors.New("comment must be less than 400 chars")
	}

	return nil
}
