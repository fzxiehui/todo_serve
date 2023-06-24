package tools

import (
	"golang.org/x/crypto/bcrypt"
)

// func MD5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

// hash
func Hash(str string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
}

func HashPassword(password string) string {
	hash, _ := Hash(password)
	return string(hash)
}

/*
 * CompareHash
 * @Description: 比较hash
 * @param hash string
 * @param str string
 * @return error
 */
func CompareHash(hash, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
}
