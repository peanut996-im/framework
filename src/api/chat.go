package api

type ChatMessage struct {
	From      string `json:"from"`
	To        string `json:"to"`
	MessageID string `json:"message_id"`
	RoomID    string `json:"room_id"`
	TimeStamp int64  `json:"time_stamp"`
}
