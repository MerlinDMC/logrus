// Based on ssh/terminal:
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build solaris

package logrus

// IsTerminal returns true if the given file descriptor is a terminal.
// @TODO: implement in a sane way as soon as the syscalls are available
func IsTerminal() bool {
	return false
}
