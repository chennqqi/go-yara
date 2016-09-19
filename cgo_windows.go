// Copyright © 2015 Hilko Bengen <bengen@hilluzination.de>. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.
// +build windows
package yara

//#cgo LDFLAGS: -lyara

/*
size_t strnlen(const char* pStr, size_t maxcount)
{
	size_t idx = 0;
	for (idx=0; idx<maxcount; idx++)
	{
		if (pStr[idx]=='\0')
		{
			return idx;
		}
	}
	return maxcount;
}
*/

import "C"
