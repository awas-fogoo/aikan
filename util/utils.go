package util

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

// RandomString 随机生成10位大小写数字字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func RandomCode(n int) string {
	var letters = []byte("0123456789")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func ReEmail(email string) string {
	return fmt.Sprintf("reg:%s", email)
}

func IsValidURL(urlStr string) bool {
	// Parse the URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	// Check if the scheme is missing or not supported
	if u.Scheme == "" || !(strings.HasPrefix(u.Scheme, "http") || strings.HasPrefix(u.Scheme, "ftp")) {
		return false
	}

	// Check if the host is missing or invalid
	if u.Host == "" || len(u.Host) > 253 {
		return false
	}

	// Check if the path is invalid
	if strings.Contains(u.Path, "\\") {
		return false
	}

	return true
}
