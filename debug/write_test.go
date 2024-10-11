package debug

import (
	"errors"
	"fmt"
	"iter"
	"log"
	"strings"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// MockFailWriter is a writer that always returns an error. Used for tests that make use of the
// io.Writer interface.
type MockFailWriter struct{}

// Write implements io.Writer.
//
// If the data is empty, the error "could not write" is returned. Otherwise, 0 is returned.
func (w MockFailWriter) Write(data []byte) (int, error) {
	if len(data) <= 1 {
		return 0, errors.New("could not write")
	}

	return 0, nil
}

// TestSuccessLogSeq tests the LogSeq function when it succeeds.
func TestSuccessLogSeq(t *testing.T) {
	var builder strings.Builder

	type args struct {
		dbg          *log.Logger
		title        string
		seq          iter.Seq[string]
		expected_err string
		expected_out string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			defer builder.Reset()

			err := LogSeq(args.dbg, args.title, args.seq)
			err = test.CheckErr(args.expected_err, err)
			if err != nil {
				t.Error(err)
			}

			msg := builder.String()
			if msg != args.expected_out {
				t.Errorf("want %q, got %q", args.expected_out, msg)
			}
		}
	})

	_ = tests.AddTest("no logger", args{
		dbg:          nil,
		title:        "",
		seq:          nil,
		expected_err: fmt.Sprintf("(BadParameter) parameter (%q) must not be nil", "logger"),
		expected_out: "",
	})

	_ = tests.AddTest("no sequence", args{
		dbg:          log.New(&builder, "", log.Lshortfile),
		title:        "",
		seq:          nil,
		expected_err: "",
		expected_out: "",
	})

	_ = tests.AddTest("with title", args{
		dbg:   log.New(&builder, "", 0),
		title: "my title",
		seq: func(yield func(string) bool) {
			_ = yield("test")
		},
		expected_err: "",
		expected_out: "my title\ntest\n",
	})

	_ = tests.Run(t)
}

// TestFailLogSeq tests the LogSeq function when it fails.
func TestFailLogSeq(t *testing.T) {
	var mfw MockFailWriter

	type args struct {
		dbg          *log.Logger
		title        string
		seq          iter.Seq[string]
		expected_err string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := LogSeq(args.dbg, args.title, args.seq)
			err = test.CheckErr(args.expected_err, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("non-empty sequence", args{
		dbg:   log.New(&mfw, "", 0),
		title: "my title",
		seq: func(yield func(string) bool) {
			_ = yield("test")
		},
		expected_err: "could not print the 1st element: short write",
	})

	_ = tests.AddTest("empty sequence", args{
		dbg:   log.New(&mfw, "", 0),
		title: "",
		seq: func(yield func(string) bool) {
			_ = yield("")
		},
		expected_err: "could not print the 1st element: could not write",
	})

	_ = tests.Run(t)
}
