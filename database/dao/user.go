package dao

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"slices"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PeopleId uint
	Name     string
	Password string
	Solt     string
}

func (u *User) HashPassword() {
	solt := generateSolt(32)
	u.Password = base64.StdEncoding.EncodeToString(hashPassword(u.Password, solt))
	u.Solt = base64.StdEncoding.EncodeToString(solt)
}

func (u *User) ValidatePassword(password string) bool {
	solt, _ := base64.StdEncoding.DecodeString(u.Solt)
	newPassword := hashPassword(password, solt)
	return u.Password == base64.StdEncoding.EncodeToString(newPassword)
}

func hashPassword(password string, solt []byte) []byte {
	return generateSha256(slices.Concat([]byte(password), solt))
}

func generateSolt(size uint) []byte {
	solt := make([]byte, size)
	rand.Read(solt)
	return solt
}

func generateSha256(source []byte) []byte {
	h := sha256.New()
	h.Write(source)
	return h.Sum(nil)
}
