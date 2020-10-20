package knife

import "regexp"

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
  //  'active_record'.camelize                => "ActiveRecord"
  //  'active_record'.camelize(:lower)        => "activeRecord"
  //  'active_record/errors'.camelize         => "ActiveRecord::Errors"
  //  'active_record/errors'.camelize(:lower) => "activeRecord::Errors"
  func Camelize(firstLetter string) string {
     return ""
  }
  func Camelcase(firstLetter string) string {
    return Camelize(firstLetter)
  }

  //
  // +Titleize+ is also aliased as +Titlecase+.
  //
  //   'man from the boondocks'.titleize => "Man From The Boondocks"
  //   'x-men: the last stand'.titleize  => "X Men: The Last Stand"
  func Titleize(str string){
    //ActiveSupport::Inflector.titleize(self)
  }
  //alias_method :Titlecase, :Titleize
  func Titlecase(str string)  {
    Titleize(str)
  }
  
  // Creates the name of a table like Rails does for models to table names. This method
  // uses the +pluralize+ method on the last word in the string.
  //
  //   'RawScaledScorer'.tableize => "raw_scaled_scorers"
  //   'ham_and_egg'.tableize     => "ham_and_eggs"
  //   'fancyCategory'.tableize   => "fancy_categories"
  func Tableize(tableName string) string{
    return Pluralize(underscore(tableName))
  }

  // Creates a class name from a plural table name like Rails does for table names to models.
  // Note that this returns a string and not a class. (To convert to an actual class
  // follow +Classify+ with +constantize+.)
  //
  //   'ham_and_eggs'.classify  => "HamAndEgg"
  //   'posts'.classify        => "Post"
  func Classify(tableName string) string{
      reg := regexp.MustCompile(`.*\.`)
      tableName = string(reg.ReplaceAll([]byte(tableName), []byte("")))
      return Camelize(Singularize(tableName))
  
  }
