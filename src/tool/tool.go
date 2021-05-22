// Package tool
// @Title  tool.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 22:30
package tool

import "github.com/bwmarrin/snowflake"

func RemoveDuplicateString(origin []string) []string {
	result := make([]string, 0, len(origin))
	temp := map[string]struct{}{}
	for _, item := range origin {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func NewSnowFlakeID() string {
	node, _ := snowflake.NewNode(1)
	return node.Generate().String()
}
