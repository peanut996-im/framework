// Package api
// @Title  api_request.go
// @Description record some defined request struct.
// @Author  peanut996
// @Update  peanut996  2021/5/22 21:57
package api

type ChatRequest struct {
	//From sender user id
	From    string      `json:"from" bson:"from"`
	To      string      `json:"to,omitempty" bson:"to"`
	RoomID  string      `json:"roomID" bson:"roomID"`
	Time    int64       `json:"time,omitempty" bson:"time"`
	Type    string      `json:"type" bson:"type"`
	Content interface{} `json:"content" bson:"content"`
}

type AuthRequest struct {
	Token string `json:"token"`
}

type FriendRequest struct {
	FriendA string `json:"friendA"`
	FriendB string `json:"friendB"`
}

type GroupRequest struct {
	UID        string `json:"uid,omitempty"`
	GroupID    string `json:"groupID,omitempty"`
	GroupName  string `json:"groupName,omitempty"`
	GroupAdmin string `json:"groupAdmin,omitempty"`
}

type LoadRequest struct {
	UID string `json:"uid"`
}
