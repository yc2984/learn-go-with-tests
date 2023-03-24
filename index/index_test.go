package my_temp_package

import "testing"
import "strings"

func TestIndex(t *testing.T){
	index := strings.Index("apple", "l")
	expected := 3
	if index != expected {
		t.Errorf("expected %d, got %d", expected, index)
	}
}