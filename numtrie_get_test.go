package numtrie



import (
	"fmt"
	"testing"
)



func TestGetPathNothing(t *testing.T) {

	nt := New()
	if nil == nt {
		t.Errorf("Received new when calling New(), but should not have.")
	}


	r := nt.GetPath( []byte("a"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	r = nt.GetPath( []byte("ab"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}
}

func TestGetPathDepthOne(t *testing.T) {

	nt := New()
	if nil == nt {
		t.Errorf("Received new when calling New(), but should not have.")
	}


	r := nt.GetPath( []byte("a"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	nt.SetPath( []byte("a"), 7 )

	r = nt.GetPath( []byte("a"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = nt.GetPath( []byte("b"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	nt.SetPath( []byte("b"), 22 )

	r = nt.GetPath( []byte("b"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = nt.GetPath( []byte("a"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = nt.GetPath( []byte("c"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}
}

func TestGetPathDepthTwo(t *testing.T) {

	nt := New()
	if nil == nt {
		t.Errorf("Received new when calling New(), but should not have.")
	}


	r := nt.GetPath( []byte("ab"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	nt.SetPath( []byte("ab"), 7 )

	r = nt.GetPath( []byte("a"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = nt.GetPath( []byte("ab"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = nt.GetPath( []byte("b"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	r = nt.GetPath( []byte("cd"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}
}


