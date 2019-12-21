package nets

import (
	"fmt"
	"testing"
)

func TestReceiveMail(t *testing.T) {
	err := ReceiveMail("pop.qq.com:995", "1029535012@qq.com", "amdepjytfocvbfbc", func(number int, uid, data string, err error) (b bool, e error) {
		fmt.Printf("%d, %s\n", number, uid)
		return false, nil
	})
	if err != nil {
		t.Fatal("Error receive mail:", err)
	}
}
