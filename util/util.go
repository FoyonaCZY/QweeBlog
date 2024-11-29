package util

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateRandomString 生成随机字符串
func GenerateRandomString(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(res)
}

// HashPassword 生成密码哈希
func HashPassword(str string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(res), err
}

// ComparePassword 比较密码哈希
func ComparePassword(hashed string, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

// IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	// 邮箱正则表达式
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
