package pkg

import "testing"

func TestCaptcha(t *testing.T) {
	gg := NewCaptcha()
	id, s, err := gg.GenerateCaptcha()
	if err != nil {
		t.Fatalf("产生错误%s", err)
	}
	t.Log(id, s)
}
