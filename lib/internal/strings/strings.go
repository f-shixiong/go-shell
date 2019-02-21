package strings

import (
	"strings"
	"unicode"
)

type Replacer strings.Replacer

type Builder strings.Builder

type Reader strings.Reader

var StuMap = map[string]interface{}{
	"Replacer": strings.Replacer{},
	"Builder":  strings.Builder{},
	"Reader":   strings.Reader{},
}

func SplitAfter(s, sep string) []string {
	r_0 := strings.SplitAfter(s, sep)
	return r_0
}

func TrimLeft(s string, cutset string) string {
	r_0 := strings.TrimLeft(s, cutset)
	return r_0
}

func TrimSuffix(s, suffix string) string {
	r_0 := strings.TrimSuffix(s, suffix)
	return r_0
}

func IndexRune(s string, r rune) int {
	r_0 := strings.IndexRune(s, r)
	return r_0
}

func ToUpperSpecial(c unicode.SpecialCase, s string) string {
	r_0 := strings.ToUpperSpecial(c, s)
	return r_0
}

func IndexFunc(s string, f func(rune) bool) int {
	r_0 := strings.IndexFunc(s, f)
	return r_0
}

func Split(s, sep string) []string {
	r_0 := strings.Split(s, sep)
	return r_0
}

func FieldsFunc(s string, f func(rune) bool) []string {
	r_0 := strings.FieldsFunc(s, f)
	return r_0
}

func HasSuffix(s, suffix string) bool {
	r_0 := strings.HasSuffix(s, suffix)
	return r_0
}

func Contains(s, substr string) bool {
	r_0 := strings.Contains(s, substr)
	return r_0
}

func Index(s, substr string) int {
	r_0 := strings.Index(s, substr)
	return r_0
}

func Compare(a, b string) int {
	r_0 := strings.Compare(a, b)
	return r_0
}

func LastIndexAny(s, chars string) int {
	r_0 := strings.LastIndexAny(s, chars)
	return r_0
}

func Join(a []string, sep string) string {
	r_0 := strings.Join(a, sep)
	return r_0
}

func IndexByte(s string, c byte) int {
	r_0 := strings.IndexByte(s, c)
	return r_0
}

func ToTitleSpecial(c unicode.SpecialCase, s string) string {
	r_0 := strings.ToTitleSpecial(c, s)
	return r_0
}

func TrimSpace(s string) string {
	r_0 := strings.TrimSpace(s)
	return r_0
}

func ToTitle(s string) string {
	r_0 := strings.ToTitle(s)
	return r_0
}

func ContainsRune(s string, r rune) bool {
	r_0 := strings.ContainsRune(s, r)
	return r_0
}

func Repeat(s string, count int) string {
	r_0 := strings.Repeat(s, count)
	return r_0
}

func Count(s, substr string) int {
	r_0 := strings.Count(s, substr)
	return r_0
}

func HasPrefix(s, prefix string) bool {
	r_0 := strings.HasPrefix(s, prefix)
	return r_0
}

func ToLower(s string) string {
	r_0 := strings.ToLower(s)
	return r_0
}

func EqualFold(s, t string) bool {
	r_0 := strings.EqualFold(s, t)
	return r_0
}

func ToUpper(s string) string {
	r_0 := strings.ToUpper(s)
	return r_0
}

func Replace(s, old, new string, n int) string {
	r_0 := strings.Replace(s, old, new, n)
	return r_0
}

func ContainsAny(s, chars string) bool {
	r_0 := strings.ContainsAny(s, chars)
	return r_0
}

func LastIndex(s, substr string) int {
	r_0 := strings.LastIndex(s, substr)
	return r_0
}

func IndexAny(s, chars string) int {
	r_0 := strings.IndexAny(s, chars)
	return r_0
}

func TrimLeftFunc(s string, f func(rune) bool) string {
	r_0 := strings.TrimLeftFunc(s, f)
	return r_0
}

func SplitAfterN(s, sep string, n int) []string {
	r_0 := strings.SplitAfterN(s, sep, n)
	return r_0
}

func ToLowerSpecial(c unicode.SpecialCase, s string) string {
	r_0 := strings.ToLowerSpecial(c, s)
	return r_0
}

func Title(s string) string {
	r_0 := strings.Title(s)
	return r_0
}

func LastIndexFunc(s string, f func(rune) bool) int {
	r_0 := strings.LastIndexFunc(s, f)
	return r_0
}

func Trim(s string, cutset string) string {
	r_0 := strings.Trim(s, cutset)
	return r_0
}

func SplitN(s, sep string, n int) []string {
	r_0 := strings.SplitN(s, sep, n)
	return r_0
}

func TrimPrefix(s, prefix string) string {
	r_0 := strings.TrimPrefix(s, prefix)
	return r_0
}

func NewReplacer(oldnew ...string) *strings.Replacer {
	r_0 := strings.NewReplacer(oldnew...)
	return r_0
}

func Map(mapping func(rune) rune, s string) string {
	r_0 := strings.Map(mapping, s)
	return r_0
}

func NewReader(s string) *strings.Reader {
	r_0 := strings.NewReader(s)
	return r_0
}

func TrimRight(s string, cutset string) string {
	r_0 := strings.TrimRight(s, cutset)
	return r_0
}

func LastIndexByte(s string, c byte) int {
	r_0 := strings.LastIndexByte(s, c)
	return r_0
}

func Fields(s string) []string {
	r_0 := strings.Fields(s)
	return r_0
}

func TrimRightFunc(s string, f func(rune) bool) string {
	r_0 := strings.TrimRightFunc(s, f)
	return r_0
}

func TrimFunc(s string, f func(rune) bool) string {
	r_0 := strings.TrimFunc(s, f)
	return r_0
}
