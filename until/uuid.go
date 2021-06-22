package until

import (
	"github.com/google/uuid"
	"unicode"
)

// NewUUID 创建uuid
func NewUUID() (string, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return v.String(), err
}

func MustUUID() string {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return v
}

func CheckUUID(s string) bool {
	if len(s) == 36 {
		for _, v := range s {
			if unicode.Is(unicode.Han, v) {
				return false
			}
		}
		return true
	}
	return false
}
