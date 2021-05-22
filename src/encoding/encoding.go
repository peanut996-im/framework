// Package encoding
// @Title  encoding.go
// @Description  提供需要的简化的编码操作
// @Author  peanut996
// @Update  peanut996  2021/5/22 1:47
package encoding

import (
	"crypto/sha1"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/mitchellh/mapstructure"
)

func NewSnowFlakeID() string {
	node, _ := snowflake.NewNode(1)
	return node.Generate().String()
}

func EncryptBySha1(plain string) string {
	h := sha1.New()
	h.Write([]byte(plain))
	return fmt.Sprintf("%X", h.Sum(nil))
}

func MapToStruct(input, output interface{}) error {
	return mapstructure.Decode(input, output)
}
