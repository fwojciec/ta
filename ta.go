// A fork of https://github.com/benbjohnson/testing but using go-cmp instead of
// reflect.DeepEqual for equality testing.
// Credit to Ben B. Johnson for the original code.

package ta

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// OK fails the test if an err is not nil.
func OK(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	if diff := cmp.Diff(exp, act); diff != "" {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected diff (-want +got):\n%s", filepath.Base(file), line, diff)
		tb.FailNow()
	}
}
