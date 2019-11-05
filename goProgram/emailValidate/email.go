package emailValidate

import (
	"blog/logger"
	"crypto/md5"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func init() {
	tokens = make(map[string]string)
}

var tokens map[string]string

func SendEmail(email, token string) {
	CleanToken()
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "邮箱", "博客验证码管理")
	m.SetHeader("To", m.FormatAddress(email, email))
	m.SetHeader("Subject", "登录验证码")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	m.SetBody("text/html",
		fmt.Sprintf("<h3 style='color:blue;'>您的管理员登录验证码是:</h3><h1 style='color:red;'>%s</h1>", code))
	d := gomail.NewDialer("smtp.mxhichina.com", 465,
		"账号", "密码")
	if err := d.DialAndSend(m); err != nil {
		logger.ErrLog.Println(err)
		return
	}
	tokens[token] = fmt.Sprintf("%s", code)
}

func GetToken() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	str := fmt.Sprintf("%s", code)
	token := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", token[:])
}
func GetTokenValue(token string) string {
	return tokens[token]
}

func CleanToken() {
	i := 1
	for k, _ := range tokens {
		if i <= 10 {
			delete(tokens, k)
		} else {
			break
		}
	}
}
