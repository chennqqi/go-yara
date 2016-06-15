// Copyright ? 2015 Hilko Bengen <bengen@hilluzination.de>. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

// Package yara provides bindings to the YARA library.
package yara

/*

#cgo CFLAGS:  -I../yara/libyara/include
#include <yara.h>


YR_API int yr_rules_get_standalone_variable(YR_RULES* rules,
	const char* identifier,
	int64_t* pI64,
	double* pD64,
	char** ppSTR)
{
	YR_EXTERNAL_VARIABLE* external = rules->externals_list_head;
	while (!EXTERNAL_VARIABLE_IS_NULL(external))
	{
		if (strcmp(external->identifier, identifier) == 0)
		{
			switch (external->type)
			{
			default:
				return -1;
				break;

				//int key
			case EXTERNAL_VARIABLE_TYPE_INTEGER:
			case EXTERNAL_VARIABLE_TYPE_BOOLEAN:
				*pI64 = external->value.i;
				return 1;

			case EXTERNAL_VARIABLE_TYPE_FLOAT:
				*pD64 = external->value.f;
				return 2;

			case EXTERNAL_VARIABLE_TYPE_STRING:
			case EXTERNAL_VARIABLE_TYPE_MALLOC_STRING:
				*ppSTR = external->value.s;
				//PTRKEY
				return 3;
			}
			return -1;
		}

		external++;
	}
	return 0;
}

*/
import "C"
import (
	"C"
	"errors"
	"unsafe"
)

// DefineVariable defines a named variable for use by the compiler.
// Only return Boolean, int64, float64, and string
func (r *Rules) ReadStandaloneVariable(name string) (value interface{}, err error) {
	cname := C.CString(name)
	C.free(unsafe.Pointer(cname))

	var i64 C.int64_t
	var d64 C.double
	var pstr *C.char /*gcc sizeof(longlong)==sizeof(ptr)*/
	f := C.yr_rules_get_standalone_variable(r.cptr,
		cname, &i64, &d64, &pstr)

	var e interface{}
	switch f {
	case 0:
		return e, errors.New("NOT FOUND")

	case 1:
		return i64, nil
	case 2:
		return d64, nil
	case 3:
		return C.GoString(pstr), nil

	case -1:
		return e, errors.New("FOUND UNSUPPORT DATA TYPE")
	default:
		return e, errors.New("INVALID QUERY TYPE")
	}
	return
	//	return interface{},errors.New("NOT FOUND")
}
