package numtrie



import (
    "testing"
)



func TestNew(t *testing.T) {

    nt := New()
    if nil == nt {
        t.Errorf("Received new when calling New(), but should not have.")
    }
}


