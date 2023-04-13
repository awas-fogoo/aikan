package server

import (
	"fmt"
	"testing"
	"time"
)

func TestGetHotVideoServer(t *testing.T) {
	tt := time.Date(2023, 4, 5, 10, 20, 44, 22, time.UTC)
	ct := calculateVideoWeight(1, 0, 0, 0, 10, tt)
	fmt.Printf("%2f\n", ct)
}
