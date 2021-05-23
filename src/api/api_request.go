// Package api
// @Title  api_request.go
// @Description record some defined request struct.
// @Author  peanut996
// @Update  peanut996  2021/5/22 21:57
package api

type AuthRequest struct {
	Token string `json:"token"`
}

type FriendRequest struct {
	FriendA string `json:"friendA"`
	FriendB string `json:"friendB"`
}

type GroupRequest struct {
	UID       string `json:"uid,omitempty"`
	GroupID   string `json:"groupID,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	GroupAdmin string `json:"groupAdmin,omitempty"`
}

type LoadRequest struct {
	UID string `json:"uid"`
}
