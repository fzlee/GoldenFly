package page

type CommentValidator struct {
	Content    string     	`json:"content" binding:"required"`
	Email	   string		`json:"email" binding:"required"`
	Nickname   string       `json:"nickname" binding:"required"`
	Website	   string       `json:"website" binding:"url"`
	CommentID  * int          `json:"comment_id" binding:"omitempty,numeric"`
}

