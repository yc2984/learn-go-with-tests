package iteration

import "testing"
import "fmt"


func TestRepeat(t *testing.T){
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected{
		t.Errorf("expected %q, but got %q", expected, repeated)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat(){
	repeated := Repeat("hey", 3)
	fmt.Println(repeated)
	// Output: heyheyhey
}