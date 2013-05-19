package numtrie



import (
	"fmt"
    "testing"
)



func TestGetPathNothing(t *testing.T) {

	bt := New()
	if nil == bt {
		t.Errorf("Received new when calling New(), but should not have.")
	}


	r := bt.GetPath( []byte("a"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	r = bt.GetPath( []byte("ab"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}
}

func TestGetPathDepthOne(t *testing.T) {

	bt := New()
	if nil == bt {
		t.Errorf("Received new when calling New(), but should not have.")
	}


	r := bt.GetPath( []byte("a"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	bt.SetPath( []byte("a"), 7 )

	r = bt.GetPath( []byte("a"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = bt.GetPath( []byte("b"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	bt.SetPath( []byte("b"), 22 )

	r = bt.GetPath( []byte("b"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = bt.GetPath( []byte("a"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = bt.GetPath( []byte("c"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}
}

func TestGetPathDepthTwo(t *testing.T) {

	bt := New()
	if nil == bt {
		t.Errorf("Received new when calling New(), but should not have.")
	}


	r := bt.GetPath( []byte("ab"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	bt.SetPath( []byte("ab"), 7 )

	r = bt.GetPath( []byte("a"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = bt.GetPath( []byte("ab"))
	if nil == r {
        t.Errorf("Did not received result, but should not have.")
	}

	r = bt.GetPath( []byte("b"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}

	r = bt.GetPath( []byte("cd"))
	if nil != r {
        t.Errorf(  fmt.Sprintf("Received result, but should not have. Received: %v", r)  )
	}
}


