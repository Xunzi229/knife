package knife

import (
	"regexp"
	"strings"
)

// Pluralize
// Pluralize("post")              => "posts"
// Pluralize("octopus")           => "octopi"
// Pluralize("sheep")             => "sheep"
// Pluralize("words")             => "words"
// Pluralize("the blue mailman")  => "the blue mailmen"
// Pluralize("CamelOctopus")      => "CamelOctopi"
func Pluralize(str string) string {
	return applyInflections(str, Plurals)
}

//  Singularize
//  Singularize("posts)            => "post"
//  Singularize("octopi)           => "octopus"
//  Singularize("sheep)            => "sheep"
//  Singularize("word)             => "word"
//  Singularize("the blue mailman") => "the blue mailman"
//  Singularize("CamelOctopi)      => "CamelOctopus"
//  Singularize("leyes)            => "leye"
func Singularize(str string) string {
	return applyInflections(str, Singulars)
}

//  Camelize
//  Camelize("active_record") => "ActiveRecord"
//  Camelize("kai_xin")       => "KaiXin"
func Camelize(item string) string {
	words := strings.Split(item, "_")

	for i := 0; i < len(words); i++ {
		str := words[i]

		if len(str) > 0 {
			r := str[0]
			if 'a' <= r && r <= 'z' {
				r -= 'a' - 'A'
			}
			words[i] = string(r) + str[1:]
		}
	}
	return strings.Join(words, "")
}

//  Tableize
//  Tableize("RawScaledScorer") => "raw_scaled_scorers"
//  Tableize("ham_and_egg")     => "ham_and_eggs"
//  Tableize("fancyCategory")   => "fancy_categories"
func Tableize(tableName string) string {
	return Pluralize(underscore(tableName))
}

//  Classify
//  Classify("ham_and_eggs")  => "HamAndEgg"
//  Classify("posts")        => "Post"
func Classify(tableName string) string {
	reg := regexp.MustCompile(`.*\.`)
	tableName = Sub(tableName, reg, "")

	return Camelize(Singularize(tableName))
}
