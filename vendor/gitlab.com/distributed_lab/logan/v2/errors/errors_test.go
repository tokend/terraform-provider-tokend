package errors_test

import (
	goerrors "errors"
	"fmt"
	"io"
	"testing"

	"gitlab.com/distributed_lab/logan/v2/errors"
	"reflect"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  error
	}{
		{"empty", "", goerrors.New("")},
		{"simple", "foo", goerrors.New("foo")},
		{"with format", "%v", goerrors.New("%v")},
	}

	for _, tc := range cases {
		got := errors.New(tc.in)
		if got.Error() != tc.out.Error() {
			t.Errorf("%s: expected %s got %s",
				tc.name, tc.out.Error(), got.Error())
		}
	}
}

func TestPrint(t *testing.T) {
	err := errors.New("test print")
	fmt.Println(err)
	// Output: test print
}

func TestEqual(t *testing.T) {
	foo := "foo"
	if errors.New(foo) == errors.New(foo) {
		t.Error("different allocations should not be equal")
	}
	err := errors.New(foo)
	if err != err {
		t.Error("same allocation should be equal")
	}
}

func TestWithFields_Error(t *testing.T) {
	err := errors.New("foo")
	if err.Error() != "foo" {
		t.Error("expected to be equal")
	}
}

func TestWrapNil(t *testing.T) {
	if err := errors.Wrap(nil, "foo"); err != nil {
		t.Errorf("expected nil got %s", err)
	}
}

func TestWrap(t *testing.T) {
	cases := []struct {
		name string
		in   error
		msg  string
		out  string
	}{
		{"simple", io.EOF, "foo", "foo: EOF"},
		{"wrapped", errors.Wrap(io.EOF, "bar"), "foo", "foo: bar: EOF"},
	}

	for _, tc := range cases {
		got := errors.Wrap(tc.in, tc.msg)
		if got.Error() != tc.out {
			t.Errorf("%s: expected %s got %s", tc.name, tc.out, got.Error())
		}
	}
}

type nilError struct{}

func (nilError) Error() string { return "nil error" }

func TestWithFields_Cause(t *testing.T) {
	x := errors.New("foo")
	tests := []struct {
		name string
		err  error
		out  error
	}{{
		name: "nil error is nil",
		err:  nil,
		out:  nil,
	}, {
		name: "explicit nil error is nil",
		err:  (error)(nil),
		out:  nil,
	}, {
		name: "typed nil is nil",
		err:  (*nilError)(nil),
		out:  (*nilError)(nil),
	}, {
		name: "uncaused error is unaffected",
		err:  io.EOF,
		out:  io.EOF,
	}, {
		name: "caused error returns cause",
		err:  errors.Wrap(io.EOF, "bar"),
		out:  io.EOF,
	}, {
		name: "return from errors.New",
		err:  x,
		out:  x,
	}, {
		name: "wrapped errors.New",
		err:  errors.Wrap(x, "bar"),
		out:  x,
	},
	}

	for _, tc := range tests {
		got := errors.Cause(tc.err)
		if got != tc.out {
			t.Errorf("%s: got %#v, want %#v", tc.name, got, tc.out)
		}
	}
}

func TestWithFields_GetFields(t *testing.T) {
	f1 := errors.WithField("key1", "value1").Add("key2", "value2")
	f2 := errors.WithField("key3", "value3").Add("key4", "value4")

	tests := []struct {
		name string
		err  error
		out  errors.F
	}{{
		name: "fields.ToError",
		err:  f1.ToError("foo"),
		out:  f1,
	}, {
		name: "fields.Wrap",
		err:  f1.Wrap(errors.New("foo"), "bar"),
		out:  f1,
	}, {
		name: "getting transitive fields",
		err:  f2.Wrap(f1.ToError("foo"), "bar"),
		out:  f1.AddFields(f2),
	}, {
		name: "getting transitive fields with non-fielded mediator",
		err:  f2.Wrap(errors.Wrap(f1.ToError("foo"), "bar"), "bar2"),
		out:  f1.AddFields(f2),
	},
	}

	for _, tc := range tests {
		got := errors.GetFields(tc.err)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%s: got %#v, want %#v", tc.name, got, tc.out)
		}
	}
}

func Test(t *testing.T) {
	ErrInternal := errors.New("internal")

	failure := func() error {
		return ErrInternal
	}

	success := func() error {
		return nil
	}

	middle := func(f func() error) error {
		return errors.Wrap(f(), "middle")
	}

	outer := func(f func() error) error {
		return errors.Wrap(middle(f), "outer")
	}

	if err := outer(success); err != nil {
		t.Errorf("expected nil got %#v", err)
	}

	if err := errors.Cause(outer(failure)); err != ErrInternal {
		t.Errorf("expected %#v got %#v", ErrInternal, err)
	}
}
