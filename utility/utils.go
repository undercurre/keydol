package utility

import "github.com/gogf/gf/v2/crypto/gsha1"

// 密码加密
func EncryptPassword(password, salt string) string {
	return gsha1.Encrypt(gsha1.Encrypt(password) + gsha1.Encrypt(salt))
}
