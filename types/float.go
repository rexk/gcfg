package types

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

func ParseFloat(floatPtr interface{}, val string) error {
	t := reflect.ValueOf(floatPtr).Elem().Type()
	val = strings.TrimSpace(val)
	_, err := fmt.Sscanf(val, "%g", floatPtr)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to parse %q as %v: %v", val, t, err)
	}
	return nil
}
