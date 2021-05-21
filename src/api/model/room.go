package model

type Room struct {
	RoomID     string `json:"room_id"`
	RoomName   string `json:"room_name"`
	OneToOne   bool   `json:"one_to_one"`
	CreateTime int64  `json:"create_time,omitempty"`
}
