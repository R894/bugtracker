package entity

type Comment struct {
	Content string
}

func NewComment(content string) *Comment {
	return &Comment{
		Content: content,
	}
}
