// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package fit

import "reflect"

// Fit fits the given object and any underlying values.
func Fit(in interface{}) interface{} {
	v := reflect.ValueOf(in)

	switch v.Type().Kind() {
	case reflect.Array, reflect.Slice:
		return FitSliceValue(v).Interface()
	case reflect.Map:
		return FitMapValue(v).Interface()
	}

	return in
}
