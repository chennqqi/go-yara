// Copyright © 2015 Hilko Bengen <bengen@hilluzination.de>. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.
// +build windows
package yara

//#cgo LDFLAGS: -Lyara/windows/libyara/Release -L../yara/windows/libyara/Release -llibyara32
import "C"
