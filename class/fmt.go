package SO_Class

import "fmt"

// Define a struct with a Print method
type fFmt struct{}

func (saya fFmt) Println(flag bool, a ...any) (n int, err error) {
	if flag {
		fmt.Println(a...)
	}
	return n, err
}
func (saya fFmt) Sprint(a ...any) string {
	nilai := fmt.Sprint(a...)
	return nilai
}
func (saya fFmt) Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
func (saya fFmt) Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

// Exported instance
var Fmt fFmt
