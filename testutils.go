package ssgen

import "testing"

func assertEqual(t *testing.T, o, e interface{}) {
	if o != e {
		t.Errorf("expected %d, got %d", e, o)
	}
}

func assertNil(t *testing.T, o interface{}) {
	if o != nil {
		t.Error("expected nil, got", o)
	}
}

func assertNotNil(t *testing.T, o interface{}) {
	if o == nil {
		t.Error("expected NOT nil, got", o)
	}
}
