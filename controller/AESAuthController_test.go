package controller

import (
	"fmt"
	"testing"
	"time"
)

func TestAESAuthCon(t *testing.T) {
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
