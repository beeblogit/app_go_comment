package comment

import (
	"errors"
)


var ErrNameRequired = errors.New("Name is required")
var ErrCommentRequired = errors.New("Comment is required")
var ErrPostIDRequired = errors.New("PostID is required")