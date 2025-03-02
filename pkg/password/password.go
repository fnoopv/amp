package password

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/bcrypt"
	"goyave.dev/goyave/v5/util/errors"
)

// defaultPasswordLength 默认密码长度
const defaultPasswordLength = 12

// GeneratePassword 生成随机密码
func GeneratePassword(length int, includeNumber bool, includeSpecial bool) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	password := make([]byte, length)
	var charsetSource string

	if includeNumber {
		charsetSource += "0123456789"
	}
	if includeSpecial {
		charsetSource += "!@#$%^&*()_+=-"
	}
	charsetSource += charset

	charsetLength := big.NewInt(int64(len(charsetSource)))

	for i := range length {
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", errors.New(fmt.Errorf("generateing random index: %v", err))
		}
		password[i] = charsetSource[index.Int64()]
	}
	return string(password), nil
}

// GeneratePasswordAndHash 生成随机密码并加密
func GeneratePasswordAndHash() (string, string, error) {
	password, err := GeneratePassword(defaultPasswordLength, true, true)
	if err != nil {
		return "", "", errors.New(err)
	}

	hashedPassword, err := HashPassword(password)

	return password, hashedPassword, errors.New(err)
}

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New(err)
	}
	// INFO: 不要使用base64.StdEncoding.EncodeToString, 会导致增加字节
	return string(bytes), nil
}

// CheckPassword 验证密码
func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
