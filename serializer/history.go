package serializer

type HistorySerializer struct {
	UserID string `json:"user_id" binding:"required"`
}
