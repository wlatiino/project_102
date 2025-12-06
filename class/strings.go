package SO_Class

import "strings"

// Define a struct with a Print method
type sStrings struct{}

func (saya sStrings) ToUpper(k string) string {
	return strings.ToUpper(k)
}

func (saya sStrings) ToLower(k string) string {
	return strings.ToLower(k)
}

func (saya sStrings) TrimSpace(k string) string {
	return strings.TrimSpace(k)
}

func (saya sStrings) TrimRight(k, cutset string) string {
	return strings.TrimRight(k, cutset)
}

func (saya sStrings) Contains(k, substr string) bool {
	return strings.Contains(k, substr)
}

func (saya sStrings) Index(k, substr string) int {
	return strings.Index(k, substr)
}

func (saya sStrings) Replace(k, old, new string, n int) string {
	return strings.Replace(k, old, new, n)
}

func (saya sStrings) Repeat(k string, count int) string {
	return strings.Repeat(k, count)
}

func (saya sStrings) Trim(k, cutset string) string {
	return strings.Trim(k, cutset)
}

func (saya sStrings) Split(k, sep string) []string {
	return strings.Split(k, sep)
}

func (saya sStrings) Count(k, substr string) int {
	return strings.Count(k, substr)
}

func (saya sStrings) ReplaceAll(k, old, new string) string {
	return strings.ReplaceAll(k, old, new)
}

func (saya sStrings) TrimLeft(k, cutset string) string {
	return strings.TrimLeft(k, cutset)
}

func (saya sStrings) Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Exported instance
var Strings sStrings
