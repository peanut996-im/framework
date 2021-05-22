package model

type ChatMessage struct {
	//From sender user id
	From string `json:"from" bson:"from"`
	//MessageID message id generate by server(logic)
	MessageID string `json:"messageID,omitempty" bson:"messageID"`
	//To receiver id (friend)
	To string `json:"to,omitempty" bson:"to"`
	//RoomID room id
	RoomID string `json:"roomID" bson:"roomID"`
	//Time unix time
	Time int64 `json:"time,omitempty" bson:"time"`
	//Type message type
	Type string `json:"type" bson:"type"`
}
