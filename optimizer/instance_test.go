package optimizer

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

// TestCopyList is a test for copyList.
func TestCopyList(t *testing.T) {
	in := make([]string, 10)
	for i := 0; i < 10; i++ {
		in[i] = fmt.Sprintf("test%d", i)
	}
	out := CopyList(in)
	if len(in) != len(out) {
		t.Errorf("copyList(%v) = %v, want %v", in, out, in)
	}
}

// BenchmarkCopyList is a benchmark for copyList.
func BenchmarkCopyList(b *testing.B) {
	// b.ReportAllocs()
	go func() {
		log.Println(http.ListenAndServe(":6123", nil))
	}()
	in := make([]string, 10)
	for i := 0; i < 10; i++ {
		in[i] = fmt.Sprintf("test%d", i)
	}

	for i := 0; i < b.N; i++ {
		CopyList(in)
	}
}
