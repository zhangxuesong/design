package IdGenerator

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//ID生成器
type IdGenerator interface {
	Generate()
}

//日志ID生成器
type LogTraceIdGenerator interface {
}

//ID生成器，用于生成随机ID。
//此类生成的ID并非绝对唯一，
//但是重复的可能性非常低。
type RandomIdGenerator struct {
}

//生成随机ID。这些ID仅在极端情况下可以重复。
//@return 随机ID
//@return error
func (rig RandomIdGenerator) Generate() (string, error) {
	substrOfHostName, err := rig.getLastfieldOfHostName()
	if err != nil {
		return "", err
	}
	currentTimeMillis := time.Now().Unix()
	randomString := rig.GenerateRandomAlphameric(8)
	id := fmt.Sprintf("%s-%d-%s", substrOfHostName, currentTimeMillis, randomString)
	return id, nil
}

func (rig *RandomIdGenerator) getLastfieldOfHostName() (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}
	substrOfHostName := rig.GetLastSubstrSplittedByDot(hostName)
	return substrOfHostName, nil
}

func (rig *RandomIdGenerator) GetLastSubstrSplittedByDot(hostName string) string {
	tokens := strings.Split(hostName, ".")
	return tokens[len(tokens)-1]
}

func (rig *RandomIdGenerator) GenerateRandomAlphameric(length int) string {
	randomChars := make([]byte, length)
	count := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for count < len(randomChars) {
		maxAscii := byte('z')
		randomAscii := r.Intn(int(maxAscii))
		isDigit := randomAscii >= '0' && randomAscii <= '9'
		isUppercase := randomAscii >= 'A' && randomAscii <= 'Z'
		isLowercase := randomAscii >= 'a' && randomAscii <= 'z'
		if isDigit || isUppercase || isLowercase {
			randomChars[count] = byte(randomAscii)
			count++
		}
	}
	//fmt.Printf("this chars is : %v.\n", randomChars)
	return bytes.NewBuffer(randomChars).String()
}
