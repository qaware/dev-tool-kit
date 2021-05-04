package core

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"regexp"
	"strconv"
	"strings"
)

var diffPattern = regexp.MustCompile(`<(ins|del)`)

func DiffTexts(left string, right string) (string, int) {
	left, right = normalizeLeftRight(left, right)

	patch := diffmatchpatch.New()
	diff := patch.DiffMain(left, right, false)
	return addElementIds(strings.ReplaceAll(patch.DiffPrettyHtml(diff), "&para;", ""))
}

func normalizeLeftRight(left string, right string) (string, string) {
	if IsJson(left) && IsJson(right) {
		leftNormalized, err := FormatJson(left)
		if err == nil {
			left = leftNormalized
		}
		rightNormalized, err := FormatJson(right)
		if err == nil {
			right = rightNormalized
		}
	}

	return left, right
}

func addElementIds(output string) (string, int) {
	count := 0
	return diffPattern.ReplaceAllStringFunc(output, func(expr string) string {
		exprWithId := expr + " id=\"diff-" + strconv.Itoa(count) + "\""
		count++
		return exprWithId
	}), count
}
