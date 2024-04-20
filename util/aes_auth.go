package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// AESFactory 包含密钥和其他属性
type AESFactory struct {
	key []byte
}

// NewAESFactory 创建一个新的AESFactory实例
func NewAESFactory(key string) *AESFactory {
	return &AESFactory{
		key: []byte(key),
	}
}

// Encrypt 使用工厂的密钥进行加密
func (af *AESFactory) Encrypt(data []byte) (string, error) {
	block, err := aes.NewCipher(af.key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 使用工厂的密钥进行解密
func (af *AESFactory) Decrypt(ciphertext string) ([]byte, error) {
	data, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(af.key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}
