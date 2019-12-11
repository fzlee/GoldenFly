package page

type CommentValidator struct {
	Content    string     	`json:"content" binding:"required"`
	Email	   string		`json:"email" binding:"required"`
	Nickname   string       `json:"nickname" binding:"required"`
	Website	   string       `json:"website" binding:"url"`
	CommentID  * int          `json:"comment_id" binding:"omitempty,numeric"`
}


type InPlaceValidator struct {
	URL string `json:"url" binding:"required"`
}


type SavePageValidator struct {
	PageID        *int      `json:"id", "binding:omitempty"`
	AllowComment  bool      `json:"allow_comment", "binding:required"`
	AllowVisit    bool      `json:"allow_visit", "binding:required"`
	IsOriginal    bool      `json:"is_original", "binding:required"`
	NeedKey       bool      `json:"need_key", "binding:required"`
	Content       string    `json:"content", "binding:min:1"`
	Editor        string    `json:"editor", "binding:oneof=markdown html"`
	MetaContent   string    `json:"metacontent", "binding:required"`
	Password      string    `json:"password", "binding:min=1"`
	Tags          string    `json:"tags", "binding:min=1"`
	Title         string    `json:"title", "binding:min=1"`
	URL           string    `json:"url", "binding:min=1"`
}
