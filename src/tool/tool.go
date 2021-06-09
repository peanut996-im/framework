// Package tool
// @Title  tool.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 22:30
package tool

import (
	"encoding/json"
	"errors"
	"github.com/bwmarrin/snowflake"
	"net"
	"strconv"
	"time"
)

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

func PrettyPrint(val interface{}) (string, error) {
	s, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func GetNowUnixMilliSecond() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}
func GetNowUnixNanoSecond() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func GetIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!")

}
