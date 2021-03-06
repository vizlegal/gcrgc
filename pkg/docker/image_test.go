package docker

import (
	"regexp"
	"testing"
)

var hasTagTest = []struct {
	img      Image
	test     string
	expected bool
}{
	{Image{}, "some tag", false},
	{Image{Tags: []string{"not matching tag"}}, "some tag", false},
	{Image{Tags: []string{"not matching tag", "matching tag"}}, "matching tag", true},
}

func TestHasTag(t *testing.T) {
	var actual bool

	for _, test := range hasTagTest {
		actual = test.img.HasTag(test.test)
		if actual != test.expected {
			t.Errorf("Expected HasTag to be %v, got %v instead", test.expected, actual)
		}
	}
}

var hasTagRegexpTest = []struct {
	img      Image
	pattern  string
	expected bool
}{
	{Image{}, "^pattern$", false},
	{Image{Tags: []string{"pattern"}}, "^pattern$", true},
	{Image{Tags: []string{"not matching pattern"}}, "^pattern$", false},
}

func TestHasTagRegexp(t *testing.T) {
	var actual bool

	for _, test := range hasTagRegexpTest {
		actual = test.img.HasTagRegexp(regexp.MustCompile(test.pattern))
		if actual != test.expected {
			t.Errorf("Expected HasTagRegexp to be %v, got %v instead", test.expected, actual)
		}
	}
}

var isTaggedTest = []struct {
	img      Image
	expected bool
}{
	{Image{}, false},
	{Image{Tags: []string{}}, false},
	{Image{Tags: []string{"a tag"}}, true},
}

func TestIsTagged(t *testing.T) {
	var actual bool

	for _, test := range isTaggedTest {
		actual = test.img.IsTagged()
		if actual != test.expected {
			t.Errorf("Expected IsTagged to be %v, got %v instead", test.expected, actual)
		}
	}
}
