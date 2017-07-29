// Copyright © 2015-2017 Hilko Bengen <bengen@hilluzination.de>. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

// This file contains functionality available in go1.7 onward

// +build go1.7

package yara

import (
	"runtime"
)

func (r *Rules) keepAlive() {
	runtime.KeepAlive(r)
}
