package authentication

import (
	"crypto/md5"
	"strings"
	"time"
)

type AuthToken struct {
	//token操作类
	token               string
	createTime          int64
	expiredTimeInterval int64
}

func New(expiredTimeInterval int64) (*AuthToken, error) {
	at := &AuthToken{expiredTimeInterval: expiredTimeInterval}
	if at.expiredTimeInterval == 0 {
		at.expiredTimeInterval = int64(time.Nanosecond * 60 * 1000)
	}
	return at, nil
}

func (at *AuthToken) Create(baseUrl string, createTime int64, params map[string]string) (string, error) {
	str := strings.Builder{}
	str.WriteString(baseUrl)
	str.WriteRune(rune(createTime))
	m := md5.New()
	token := m.Sum([]byte(str.String()))
	return string(token), nil
}
