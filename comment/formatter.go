package comment

type CommentFormatter struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	PhotoID int    `json:"photo_id"`
	Message string `json:"message"`
}

func FormatterComment(comment Comment) CommentFormatter {
	formatter := CommentFormatter{
		ID:      comment.ID,
		UserID:  comment.UserId,
		PhotoID: comment.PhotoId,
		Message: comment.Message,
	}
	return formatter
}
