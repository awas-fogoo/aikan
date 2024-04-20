package util

import (
	"testing"
)

func TestCheck(t *testing.T) {
	dir := "./" // 设置一个测试用的目录路径

	checker := NewChecker(dir)
	result, err := checker.Check()
	if err != nil {
		t.Fatalf("Error during checking: %v", err)
	}

	t.Logf("Available Bytes: %d", result.AvailableBytes)
	t.Logf("Total Bytes: %d", result.TotalBytes)

	// 在这里添加适当的断言来验证结果
	if result.AvailableBytes > result.TotalBytes {
		t.Errorf("Available bytes should not be greater than total bytes")
	}
}
