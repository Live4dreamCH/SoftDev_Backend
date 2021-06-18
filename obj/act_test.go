package obj

import "testing"

func TestGetAid(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(get_aid())
	}
}
