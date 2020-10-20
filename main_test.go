package knife

import (
	"github.com/alecthomas/assert"
	"regexp"
	"testing"
)

func TestPluralize(t *testing.T) {
	assert.Equal(t, Pluralize("post"), "posts")
	assert.Equal(t, Pluralize("octopus"), "octopi")
	assert.Equal(t, Pluralize("sheep"), "sheep")
	assert.Equal(t, Pluralize("words"), "words")
	assert.Equal(t, Pluralize("the blue mailman"), "the blue mailmen")
	assert.Equal(t, Pluralize("CamelOctopus"), "CamelOctopi")
}

func TestSingularize(t *testing.T) {
	assert.Equal(t, Singularize("posts"), "post")
	assert.Equal(t, Singularize("octopi"), "octopus")
	assert.Equal(t, Singularize("sheep"), "sheep")
	assert.Equal(t, Singularize("word"), "word")
	assert.Equal(t, Singularize("the blue mailmen"), "the blue mailman")
	assert.Equal(t, Singularize("CamelOctopi"), "CamelOctopus")
	assert.Equal(t, Singularize("leyes"), "leye")
}

func TestTableize(t *testing.T) {
	assert.Equal(t, Tableize("RawScaledScorer"), "raw_scaled_scorers")
	assert.Equal(t, Tableize("ham_and_egg"), "ham_and_eggs")
	assert.Equal(t, Tableize("fancyCategory"), "fancy_categories")
}

func TestSub(t *testing.T) {
	reg := regexp.MustCompile("os")
	res := Sub("post", reg, "-")
	assert.Equal(t, res, "p-t")
}

func TestGsub(t *testing.T) {
	reg := regexp.MustCompile("os")
	res := Gsub("posost", reg, "-")
	assert.Equal(t, res, "p--t")
}

func TestClassify(t *testing.T) {
	assert.Equal(t, Classify("ham_and_eggs"), "HamAndEgg")
	assert.Equal(t, Classify("posts"), "Post")
}
