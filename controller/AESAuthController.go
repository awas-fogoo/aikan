package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

// 加密函数
func encrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
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

// 解密函数
func decrypt(ciphertext string, key []byte) ([]byte, error) {
	data, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
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

func AESAuthCon() {
	// 定义密钥
	key := []byte("c0351a0cee08e289c73079b01a0d3585")

	// 定义待加密数据
	UserID := 123456
	Expiration := time.Now().Unix() + 86400
	plaintext := []byte(fmt.Sprintf("%d:%d", UserID, Expiration))

	// 加密
	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密结果：%s\n", ciphertext)

	// 解密
	decryptedText, err := decrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密结果：%s\n", decryptedText)

	//ctx := context.Background()
	//// Set the token
	//err = rdb.Set(ctx, "access_token:"+ciphertext, UserID, time.Duration(Expiration)*time.Second).Err()
	//if err != nil {
	//	panic(err)
	//}

}
