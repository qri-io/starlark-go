// Copyright 2020 The Bazel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

// Run 'go generate' in this directory to regenerate the example protobuf Go code.
// Prerequisites:
// $ sudo apt-get install protobuf-compiler  			# install protoc
// $ go install google.golang.org/protobuf/cmd/protoc-gen-go	# install protoc-gen-go
//
//go:generate protoc --go_out=. --go_opt=paths=source_relative testdata/bugtracker/bugtracker.proto testdata/test/test.proto

import (
	"fmt"
	"path/filepath"
	"testing"

	starlarkproto "go.starlark.net/lib/proto"
	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarktest"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"

	// linked-in descriptors for tests:
	"go.starlark.net/lib/proto/testdata/bugtracker"
	_ "go.starlark.net/lib/proto/testdata/test"
)

func init() {
	resolve.AllowLambda = true
}

var predeclared = starlark.StringDict{
	"proto": starlarkproto.Module,
}

func TestProto(t *testing.T) {
	thread := &starlark.Thread{Load: load}

	starlarkproto.SetPool(thread, protoregistry.GlobalFiles)
	starlarktest.SetReporter(thread, t)

	thisdir := starlarktest.DataFile("lib/proto", ".")
	for _, f := range []string{
		"testdata/example.star",
		"testdata/proto.star",
	} {
		filename := filepath.Join(thisdir, f)
		if _, err := starlark.ExecFile(thread, filename, nil, predeclared); err != nil {
			reportEvalError(t, err)
		}
	}
}

func TestMarshal(t *testing.T) {
	thread := &starlark.Thread{Load: load}

	starlarkproto.SetPool(thread, protoregistry.GlobalFiles)
	starlarktest.SetReporter(thread, t)

	protoInStarlark := `
bugtracker = proto.file("testdata/bugtracker/bugtracker.proto")

issue = bugtracker.Issue(
	id = 12345,
	status = bugtracker.Status.ASSIGNED,
	reporter = "adonovan",
	note = ["A note", "Another note"],
	metadata = [bugtracker.KeyValuePair(key="greeting", value="hello")],
)
proto.set_field(issue, bugtracker.Outer.ext, 1)
proto.set_field(issue, bugtracker.ext, 2)
`
	globals, err := starlark.ExecFile(thread, "", protoInStarlark, predeclared)
	if err != nil {
		reportEvalError(t, err)
	}
	issueVal := globals["issue"].(*starlarkproto.Message)
	data, err := proto.Marshal(issueVal.Message())
	if err != nil {
		t.Fatalf("starlarkproto.Marshal failed: %+v", err)
	}
	issue := &bugtracker.Issue{}
	if err := proto.Unmarshal(data, issue); err != nil {
		t.Fatalf("unmarshal failed: %+v", err)
	}
	if issue.GetId() != 12345 {
		t.Fatalf("resulting proto unexpected. got: %+v", issue)
	}
}

func reportEvalError(tb testing.TB, err error) {
	if err, ok := err.(*starlark.EvalError); ok {
		tb.Fatal(err.Backtrace())
	}
	tb.Fatal(err)
}

// load implements the 'load' operation as used in the tests.
// It does not memoize repeated calls.
func load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	// "standard" library
	switch module {
	case "assert.star":
		return starlarktest.LoadAssertModule()
	}

	return nil, fmt.Errorf("no such module: %s", module)
}
