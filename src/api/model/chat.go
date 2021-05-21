package model

type ChatMessage struct {
	//From sender user id
	From string `json:"from"`
	//MessageID message id generate by server(logic)
	MessageID string `json:"message_id,omitempty"`
	//To receiver id (friend)
	To string `json:"to,omitempty"`
	//RoomID room id
	RoomID string `json:"room_id"`
	//Time unix time
	Time int64 `json:"time,omitempty"`
	//Type message type
	Type string `json:"type"`
}
