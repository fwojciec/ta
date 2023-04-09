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

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if diff := cmp.Diff(exp, act); diff != "" {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected diff (-want +got):\n%s", filepath.Base(file), line, diff)
		tb.FailNow()
	}
}
