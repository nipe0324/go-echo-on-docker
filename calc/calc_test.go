package calc

import "testing"

// TestSumは加算のテストをする
func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1, 2) should be 3, but doesn't match")
	}
}
