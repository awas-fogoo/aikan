package vo

import (
	"fmt"
	"strings"
	"time"
)

type VideoHomeVo struct {
	ID          uint       `json:"id"`
	CreatedAt   CustomTime `json:"created_at"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Url         string     `json:"url"`
	CoverUrl    string     `json:"cover_url"`
	Views       uint       `json:"views"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(data []byte) error {
	layout := "2006-01-02 15:04:05"
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	tt, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// Scan implements the Scanner interface.
func (t *CustomTime) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
	default:
		return fmt.Errorf("invalid value type for CustomTime: %T", value)
	}

	return nil
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))), nil
}

type DanmukuResponseVo struct {
	UserID   uint   `json:"user_id"`
	Content  string `json:"txt"`
	Start    uint64 `json:"start"`
	Duration uint64 `json:"duration"`
	Prior    bool   `json:"prior"`
	Colour   bool   `json:"color"`
	Mode     string `json:"mode"`
	Style    struct {
		Color           string `json:"color"`
		FontSize        string `json:"fontSize"`
		Border          string `json:"border"`
		BorderRadius    string `json:"borderRadius"`
		Padding         string `json:"padding"`
		BackgroundColor string `json:"backgroundColor"`
	} `json:"style"`
}
