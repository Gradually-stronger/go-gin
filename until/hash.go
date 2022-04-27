package until

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

type HashType string

const (
	//HashMd5 hash 类型-md5
	HashMd5    HashType = "md5"
	HashSha1   HashType = "sha1"
	HashSha256 HashType = "sha256"
	HashSha512 HashType = "sha512"
)

// MD5Hash MD5哈希值
func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString MD5哈希值
func MD5HashString(s string) string {
	return MD5Hash([]byte(s))
}

// SHA1Hash SHA1哈希值
func SHA1Hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString SHA1哈希值
func SHA1HashString(s string) string {
	return SHA1Hash([]byte(s))
}

// Hash 计算 hash/md5 值
func Hash(text string, hashType HashType, isHex bool) string {
	var hashObj hash.Hash

	switch hashType {
	case HashMd5:
		hashObj = md5.New()
	case HashSha1:
		hashObj = sha1.New()
	case HashSha256:
		hashObj = sha256.New()
	case HashSha512:
		hashObj = sha512.New()
	default:
		hashObj = md5.New()
	}

	if isHex {
		bytes, _ := hex.DecodeString(text)
		hashObj.Write(bytes)
	} else {
		hashObj.Write([]byte(text))
	}

	cipherBytes := hashObj.Sum(nil)
	//等价于 return fmt.Sprintf("%x",cipherBytes)
	return hex.EncodeToString(cipherBytes)
}

// SHA256Double CnPeng 2021/10/15 14:56 对一个字符串做两次 SHA256 处理
func SHA256Double(text string, isHex bool) []byte {
	hashObj := sha256.New()
	if isHex {
		bytes, _ := hex.DecodeString(text)
		hashObj.Write(bytes)
	} else {
		hashObj.Write([]byte(text))
	}
	cipherBytes := hashObj.Sum(nil)

	hashObj.Reset()

	hashObj.Write(cipherBytes)
	cipherBytes = hashObj.Sum(nil)
	return cipherBytes
}
