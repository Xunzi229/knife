package knife

import (
    "regexp"
    "strings"
)

type Inflection struct {
	Reg     *regexp.Regexp
	Replace string
	Sites   []int
}

var (
	Plurals = make([]Inflection, 35)
	Singulars = make([]Inflection, 41)
	Uncountables = make([]Inflection, 11)
	MaxSeg = 10
)

func init() {
	var reg *regexp.Regexp
 
	reg = regexp.MustCompile(`(?i)(h)ouse_corpuses$`)
	Plurals[0] = Inflection{Reg: reg, Replace: `${1}ouse_corpuses`}
	reg = regexp.MustCompile(`(?i)(h)ouse_corpus$`)
	Plurals[1] = Inflection{Reg: reg, Replace: `${1}ouse_corpuses`}
	reg = regexp.MustCompile(`(?i)(z)ombies$`)
	Plurals[2] = Inflection{Reg: reg, Replace: `${1}ombies`}
	reg = regexp.MustCompile(`(?i)(z)ombie$`)
	Plurals[3] = Inflection{Reg: reg, Replace: `${1}ombies`}
	reg = regexp.MustCompile(`(?i)(m)oves$`)
	Plurals[4] = Inflection{Reg: reg, Replace: `${1}oves`}
	reg = regexp.MustCompile(`(?i)(m)ove$`)
	Plurals[5] = Inflection{Reg: reg, Replace: `${1}oves`}
	reg = regexp.MustCompile(`(?i)(s)exes$`)
	Plurals[6] = Inflection{Reg: reg, Replace: `${1}exes`}
	reg = regexp.MustCompile(`(?i)(s)ex$`)
	Plurals[7] = Inflection{Reg: reg, Replace: `${1}exes`}
	reg = regexp.MustCompile(`(?i)(c)hildren$`)
	Plurals[8] = Inflection{Reg: reg, Replace: `${1}hildren`}
	reg = regexp.MustCompile(`(?i)(c)hild$`)
	Plurals[9] = Inflection{Reg: reg, Replace: `${1}hildren`}
	reg = regexp.MustCompile(`(?i)(m)en$`)
	Plurals[10] = Inflection{Reg: reg, Replace: `${1}en`}
	reg = regexp.MustCompile(`(?i)(m)an$`)
	Plurals[11] = Inflection{Reg: reg, Replace: `${1}en`}
	reg = regexp.MustCompile(`(?i)(p)eople$`)
	Plurals[12] = Inflection{Reg: reg, Replace: `${1}eople`}
	reg = regexp.MustCompile(`(?i)(p)erson$`)
	Plurals[13] = Inflection{Reg: reg, Replace: `${1}eople`}
	reg = regexp.MustCompile(`(?i)(quiz)$`)
	Plurals[14] = Inflection{Reg: reg, Replace: `${1}zes`}
	reg = regexp.MustCompile(`(?i)^(oxen)$`)
	Plurals[15] = Inflection{Reg: reg, Replace: `${1}`}
	reg = regexp.MustCompile(`(?i)^(ox)$`)
	Plurals[16] = Inflection{Reg: reg, Replace: `${1}en`}
	reg = regexp.MustCompile(`(?i)^(m|l)ice$`)
	Plurals[17] = Inflection{Reg: reg, Replace: `${1}ice`}
	reg = regexp.MustCompile(`(?i)^(m|l)ouse$`)
	Plurals[18] = Inflection{Reg: reg, Replace: `${1}ice`}
	reg = regexp.MustCompile(`(?i)(matr|vert|ind)(?:ix|ex)$`)
	Plurals[19] = Inflection{Reg: reg, Replace: `${1}ices`}
	reg = regexp.MustCompile(`(?i)(x|ch|ss|sh)$`)
	Plurals[20] = Inflection{Reg: reg, Replace: `${1}es`}
	reg = regexp.MustCompile(`(?i)([^aeiouy]|qu)y$`)
	Plurals[21] = Inflection{Reg: reg, Replace: `${1}ies`}
	reg = regexp.MustCompile(`(?i)(hive)$`)
	Plurals[22] = Inflection{Reg: reg, Replace: `${1}s`}
	reg = regexp.MustCompile(`(?i)(?:([^f])fe|([lr])f)$`)
	Plurals[23] = Inflection{Reg: reg, Replace: `${1}${2}ves`, Sites: []int{2}}
	reg = regexp.MustCompile(`(?i)sis$`)
	Plurals[24] = Inflection{Reg: reg, Replace: `${1}ses`, Sites: []int{}}
	reg = regexp.MustCompile(`(?i)([ti])a$`)
	Plurals[25] = Inflection{Reg: reg, Replace: `${1}a`}
	reg = regexp.MustCompile(`(?i)([ti])um$`)
	Plurals[26] = Inflection{Reg: reg, Replace: `${1}a`}
	reg = regexp.MustCompile(`(?i)(buffal|tomat)o$`)
	Plurals[27] = Inflection{Reg: reg, Replace: `${1}oes`}
	reg = regexp.MustCompile(`(?i)(bu)s$`)
	Plurals[28] = Inflection{Reg: reg, Replace: `${1}ses`}
	reg = regexp.MustCompile(`(?i)(alias|status)$`)
	Plurals[29] = Inflection{Reg: reg, Replace: `${1}es`}
	reg = regexp.MustCompile(`(?i)(octop|vir)i$`)
	Plurals[30] = Inflection{Reg: reg, Replace: `${1}i`}
	reg = regexp.MustCompile(`(?i)(octop|vir)us$`)
	Plurals[31] = Inflection{Reg: reg, Replace: `${1}i`}
	reg = regexp.MustCompile(`(?i)^(ax|test)is$`)
	Plurals[32] = Inflection{Reg: reg, Replace: `${1}es`}
	reg = regexp.MustCompile(`(?i)s$`)
	Plurals[33] = Inflection{Reg: reg, Replace: "s", Sites: []int{}}
	reg = regexp.MustCompile(`$`)
	Plurals[34] = Inflection{Reg: reg, Replace: "s", Sites: []int{}}
 
	reg = regexp.MustCompile(`(?i)(h)ouse_corpuses$`)
    Singulars[0] = Inflection{Reg: reg, Replace: "${1}ouse_corpus"}
    reg = regexp.MustCompile(`(?i)(h)ouse_corpus$`)
    Singulars[1] = Inflection{Reg: reg, Replace: "${1}ouse_corpus"}
    reg = regexp.MustCompile(`(?i)(z)ombies$`)
    Singulars[2] = Inflection{Reg: reg, Replace: "${1}ombie"}
    reg = regexp.MustCompile(`(?i)(z)ombie$`)
    Singulars[3] = Inflection{Reg: reg, Replace: "${1}ombie"}
    reg = regexp.MustCompile(`(?i)(m)oves$`)
    Singulars[4] = Inflection{Reg: reg, Replace: "${1}ove"}
    reg = regexp.MustCompile(`(?i)(m)ove$`)
    Singulars[5] = Inflection{Reg: reg, Replace: "${1}ove"}
    reg = regexp.MustCompile(`(?i)(s)exes$`)
    Singulars[6] = Inflection{Reg: reg, Replace: "${1}ex"}
    reg = regexp.MustCompile(`(?i)(s)ex$`)
    Singulars[7] = Inflection{Reg: reg, Replace: "${1}ex"}
    reg = regexp.MustCompile(`(?i)(c)hildren$`)
    Singulars[8] = Inflection{Reg: reg, Replace: "${1}hild"}
    reg = regexp.MustCompile(`(?i)(c)hild$`)
    Singulars[9] = Inflection{Reg: reg, Replace: "${1}hild"}
    reg = regexp.MustCompile(`(?i)(m)en$`)
    Singulars[10] = Inflection{Reg: reg, Replace: "${1}an"}
    reg = regexp.MustCompile(`(?i)(m)an$`)
    Singulars[11] = Inflection{Reg: reg, Replace: "${1}an"}
    reg = regexp.MustCompile(`(?i)(p)eople$`)
    Singulars[12] = Inflection{Reg: reg, Replace: "${1}erson"}
    reg = regexp.MustCompile(`(?i)(p)erson$`)
    Singulars[13] = Inflection{Reg: reg, Replace: "${1}erson"}
    reg = regexp.MustCompile(`(?i)(database)s$`)
    Singulars[14] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)(quiz)zes$`)
    Singulars[15] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)(matr)ices$`)
    Singulars[16] = Inflection{Reg: reg, Replace: "${1}ix"}
    reg = regexp.MustCompile(`(?i)(vert|ind)ices$`)
    Singulars[17] = Inflection{Reg: reg, Replace: "${1}ex"}
    reg = regexp.MustCompile(`(?i)^(ox)en`)
    Singulars[18] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)(alias|status)(es)?$`)
    Singulars[19] = Inflection{Reg: reg, Replace: ""}
    reg = regexp.MustCompile(`(?i)(octop|vir)(us|i)$`)
    Singulars[20] = Inflection{Reg: reg, Replace: "${1}us"}
    reg = regexp.MustCompile(`(?i)^(a)x[ie]s$`)
    Singulars[21] = Inflection{Reg: reg, Replace: "${1}xis"}
    reg = regexp.MustCompile(`(?i)(cris|test)(is|es)$`)
    Singulars[22] = Inflection{Reg: reg, Replace: "${1}is"}
    reg = regexp.MustCompile(`(?i)(shoe)s$`)
    Singulars[23] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)(o)es$`)
    Singulars[24] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)(bus)(es)?$`)
    Singulars[25] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)^(m|l)ice$`)
    Singulars[26] = Inflection{Reg: reg, Replace: "${1}ouse"}
    reg = regexp.MustCompile(`(?i)(x|ch|ss|sh)es$`)
    Singulars[27] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)(m)ovies$`)
    Singulars[28] = Inflection{Reg: reg, Replace: "${1}ovie"}
    reg = regexp.MustCompile(`(?i)(s)eries$`)
    Singulars[29] = Inflection{Reg: reg, Replace: "${1}eries"}
    reg = regexp.MustCompile(`(?i)([^aeiouy]|qu)ies$`)
    Singulars[30] = Inflection{Reg: reg, Replace: "${1}y"}
    reg = regexp.MustCompile(`(?i)([lr])ves$`)
    Singulars[31] = Inflection{Reg: reg, Replace: "${1}f"}
    reg = regexp.MustCompile(`(?i)(tive)s$`)
    Singulars[32] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)(hive)s$`)
    Singulars[33] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)([^f])ves$`)
    Singulars[34] = Inflection{Reg: reg, Replace: "${1}fe"}
    reg = regexp.MustCompile(`(?i)(^analy)(sis|ses)$`)
    Singulars[35] = Inflection{Reg: reg, Replace: "${1}sis"}
    reg = regexp.MustCompile(`(?i)((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$`)
    Singulars[36] = Inflection{Reg: reg, Replace: "${1}sis"}
    reg = regexp.MustCompile(`(?i)([ti])a$`)
    Singulars[37] = Inflection{Reg: reg, Replace: "${1}um"}
    reg = regexp.MustCompile(`(?i)(n)ews$`)
    Singulars[38] = Inflection{Reg: reg, Replace: "${1}ews"}
    reg = regexp.MustCompile(`(?i)(ss)$`)
    Singulars[39] = Inflection{Reg: reg, Replace: "${1}"}
    reg = regexp.MustCompile(`(?i)s$`)
    Singulars[40] = Inflection{Reg: reg, Replace: "", Sites: []int{}}
	
    
    reg = regexp.MustCompile(`equipment`)
    Uncountables[0] = Inflection{Reg: reg, Replace: "equipment"}
    reg = regexp.MustCompile(`information`)
    Uncountables[1] = Inflection{Reg: reg, Replace: "information"}
    reg = regexp.MustCompile(`rice`)
    Uncountables[2] = Inflection{Reg: reg, Replace: "rice"}
    reg = regexp.MustCompile(`money`)
    Uncountables[3] = Inflection{Reg: reg, Replace: "money"}
    reg = regexp.MustCompile(`species`)
    Uncountables[4] = Inflection{Reg: reg, Replace: "species"}
    reg = regexp.MustCompile(`series`)
    Uncountables[5] = Inflection{Reg: reg, Replace: "series"}
    reg = regexp.MustCompile(`fish`)
    Uncountables[6] = Inflection{Reg: reg, Replace: "fish"}
    reg = regexp.MustCompile(`sheep`)
    Uncountables[7] = Inflection{Reg: reg, Replace: "sheep"}
    reg = regexp.MustCompile(`jeans`)
    Uncountables[9] = Inflection{Reg: reg, Replace: "jeans"}
    reg = regexp.MustCompile(`police`)
    Uncountables[10] = Inflection{Reg: reg, Replace: "police"}
}

func uncountable(str string) bool {
    for i := 0; i < len(Uncountables); i++ {
        if Uncountables[i].Reg.Match([]byte(str)) {
            return true
        }
    }
    return false
}

// Makes an underscored, lowercase form from the expression in the string.
//
// Changes '::' to '/' to convert namespaces to paths.
//
//   underscore('ActiveModel')         # => "active_model"
//   underscore('ActiveModel::Errors') # => "active_model/errors"
//
// As a rule of thumb you can think of +underscore+ as the inverse of
// #camelize, though there are cases where that does not hold:
//
//  Camelize(underscore('SSLError'))  # => "SslError"
func underscore(word string) string {
    reg := regexp.MustCompile(`[A-Z-]|::`)
    if !reg.Match([]byte(word)) {
        return word
    }
    reg = regexp.MustCompile(`::`)
    word = string(reg.ReplaceAll([]byte(word), []byte("/")))
    
    //word.gsub!(/(?:(?<=([A-Za-z\d]))|\b)(#{inflections.acronym_regex})(?=\b|[^a-z])/) { "#{$1 && '_'.freeze }#{$2.downcase}" }
    
    reg = regexp.MustCompile(`([A-Z\d]+)([A-Z][a-z])`)
    word = reg.ReplaceAllString(word, "${1}_${2}")
    
    reg = regexp.MustCompile(`([a-z\d])([A-Z])`)
    word = reg.ReplaceAllString(word, "${1}_${2}")
    
    reg = regexp.MustCompile(`-`)
    word = string(reg.ReplaceAll([]byte(word), []byte("_")))
    
    word = strings.ToLower(word)
    return word
}

func applyInflections(word string, rules []Inflection) string {
    if len(word) == 0 && uncountable(word) {
        return word
    }
    
    for i := 0; i < len(rules); i++ {
		sp := rules[i]
		reg := sp.Reg
		if sb := reg.FindSubmatch([]byte(word)); len(sb) > 0 {
			word = reg.ReplaceAllString(word, sp.Replace)
			return word
		}
	}
    return word
}