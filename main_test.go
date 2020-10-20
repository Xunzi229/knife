package knife

import (
    "github.com/alecthomas/assert"
    "testing"
)

func TestPluralize(t *testing.T) {
    assert.Equal(t, Pluralize("post"),"posts")
    assert.Equal(t, Pluralize("octopus"), "octopi")
    assert.Equal(t, Pluralize("sheep"), "sheep")
    assert.Equal(t, Pluralize("words"), "words")
    assert.Equal(t, Pluralize("the blue mailman"), "the blue mailmen")
    assert.Equal(t, Pluralize("CamelOctopus"), "CamelOctopi")
}


func TestSingularize(t *testing.T) {
    assert.Equal(t, Singularize("posts"),"post")
    assert.Equal(t, Singularize("octopi"), "octopus")
    assert.Equal(t, Singularize("sheep"), "sheep")
    assert.Equal(t, Singularize("word"), "word")
    assert.Equal(t, Singularize("the blue mailmen"), "the blue mailman")
    assert.Equal(t, Singularize("CamelOctopi"), "CamelOctopus")
    assert.Equal(t, Singularize("leyes"), "ley")
}

