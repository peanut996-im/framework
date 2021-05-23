// Package api
// @Title  api_request.go
// @Description record some defined request struct.
// @Author  peanut996
// @Update  peanut996  2021/5/22 21:57
package api

type FriendRequest struct {
	FriendA string `json:"friendA"`
	FriendB string `json:"friendB"`
}

type GroupRequest struct {
	UID       string `json:"uid,omitempty"`
	GroupID   string `json:"groupID,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}
