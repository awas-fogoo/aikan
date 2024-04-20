package util

import (
	"fmt"
	"testing"
	"time"
)

func TestAESFactory_EncryptDecrypt(t *testing.T) {
	// 定义密钥
	key := "c0351a0cee08e289c73079b01a0d3585"

	// 创建工厂
	af := NewAESFactory(key)

	// 定义待加密数据
	userID := 123456
	expiration := time.Now().Unix() + 86400
	plaintext := []byte(fmt.Sprintf("%d:%d", userID, expiration))

	// 加密
	ciphertext, err := af.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	fmt.Println(ciphertext)

	// 解密
	decryptedText, err := af.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	fmt.Println(string(decryptedText))

	// 验证解密后的数据与原始数据是否一致
	if string(decryptedText) != string(plaintext) {
		t.Fatalf("Decrypted data does not match plaintext")
	}
	// 在这里可以继续执行其他操作，使用密文或解密后的数据
	//ctx := context.Background()
	//// Set the token
	//err = rdb.Set(ctx, "access_token:"+ciphertext, UserID, time.Duration(Expiration)*time.Second).Err()
	//if err != nil {
	//	panic(err)
	//}
}
