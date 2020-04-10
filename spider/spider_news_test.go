package spider

import "testing"

func TestGetNewsShort(t *testing.T) {
	url := "https://www.guancha.cn/"
	err := GetNewsShort(url)
	if err != nil {
		t.Fatal("Error get news short:", err)
	}
}