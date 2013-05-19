package numtrie



import (
    "testing"
)



func TestNew(t *testing.T) {

    bt := New()
    if nil == bt {
        t.Errorf("Received new when calling New(), but should not have.")
    }
}


