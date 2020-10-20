package knife

import (
	"regexp"
	"strings"
	"unicode"
)

// Pluralize
// 'post'.pluralize             => "posts"
// 'octopus'.pluralize          => "octopi"
// 'sheep'.pluralize            => "sheep"
// 'words'.pluralize            => "words"
// 'the blue mailman'.pluralize => "the blue mailmen"
// 'CamelOctopus'.pluralize     => "CamelOctopi"
func Pluralize(str string) string {
	return applyInflections(str, Plurals)
}

//  Singularize
//  'posts'.singularize            => "post"
//  'octopi'.singularize           => "octopus"
//  'sheep'.singularize            => "sheep"
//  'word'.singularize             => "word"
//  'the blue mailmen'.singularize => "the blue mailman"
//  'CamelOctopi'.singularize      => "CamelOctopus"
//  'leyes'.singularize(:es)       => "ley"
func Singularize(str string) string {
	return applyInflections(str, Singulars)
}

// Camelize
//  'active_record'.camelize => "ActiveRecord"
//  'kai_xin'.camelize       => "KaiXin"
func Camelize(item string) string {
	words := strings.Split(item, "_")
	for i := 0; i < len(words); i++ {
		str := words[i]
		for i, v := range str {
			words[i] = string(unicode.ToUpper(v)) + str[i+1:]
		}
	}
	return strings.Join(words, "")
}

// Creates the name of a table like Rails does for models to table names. This method
// uses the +pluralize+ method on the last word in the string.
//
//   'RawScaledScorer'.tableize => "raw_scaled_scorers"
//   'ham_and_egg'.tableize     => "ham_and_eggs"
//   'fancyCategory'.tableize   => "fancy_categories"
func Tableize(tableName string) string {
	return Pluralize(underscore(tableName))
}

// Creates a class name from a plural table name like Rails does for table names to models.
// Note that this returns a string and not a class. (To convert to an actual class
// follow +Classify+ with +constantize+.)
//
//   'ham_and_eggs'.classify  => "HamAndEgg"
//   'posts'.classify        => "Post"
func Classify(tableName string) string {
	reg := regexp.MustCompile(`.*\.`)
	tableName = string(reg.ReplaceAll([]byte(tableName), []byte("")))
	return Camelize(Singularize(tableName))

}
