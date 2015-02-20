package syllables

import (
	"bytes"
	"regexp"
)

// Attept to guess/capture all syllables in single word
var (
	emptyByte = []byte("")
	ending    = regexp.MustCompile(`(es|ed|ing)$`) // Ignored when counting syllables
	alpha     = regexp.MustCompile(`[^[:alpha:]]`)

	addSylMap = []*regexp.Regexp{
		regexp.MustCompile(`i[aiou]`),       // alias, science, phobia
		regexp.MustCompile(`[dls]ien`),      // salient, gradient, transient
		regexp.MustCompile(`[aeiouym]ble$`), // -Vble, plus -mble
		regexp.MustCompile(`[aeiou]{3}`),    // agreeable
		regexp.MustCompile(`^mc`),           // mcwhatever
		regexp.MustCompile(`ism$`),          // sexism, racism

		// regexp.MustCompile(`(?:([^aeiouy])\1|ck|mp|ng)le$`), // bubble, cattle, cackle, sample, angle // Golang does not support regexp backreference
		regexp.MustCompile(`(?:re([^aeiouy])|ck|mp|ng)le$`), // bubble, cattle, cackle, sample, angle
		regexp.MustCompile(`dnt$`),                          // couldn/t
		regexp.MustCompile(`[aeiou]y[aeiou]`),               // annoying, layer
	}

	subSylMap = []*regexp.Regexp{
		regexp.MustCompile(`[^aeiou]e$`), // give, love, bone, done, ride ...

		// regexp.MustCompile(`[aeiou](?:([cfghklmnprsvwz])\1?|ck|sh|[rt]ch)e[ds]$`), // Golang does not support regexp backreference
		regexp.MustCompile(`[aeiou](?:([cfghklmnprsvwz])?|ck|sh|[rt]ch)e[ds]$`),
		// (passive) past participles and 3rd person sing present verbs:
		// bared, liked, called, tricked, bashed, matched

		regexp.MustCompile(`.e(?:ly|less(?:ly)?|ness?|ful(?:ly)?|ments?)$`),
		// nominal, adjectival and adverbial derivatives from -e$ roots:
		// absolutely, nicely, likeness, basement, hopeless
		// hopeful, tastefully, wasteful

		regexp.MustCompile(`ion`),        // action, diction, fiction
		regexp.MustCompile(`[ct]ia[nl]`), // special(ly), initial, physician, christian
		regexp.MustCompile(`[^cx]iou`),   // illustrious, NOT spacious, gracious, anxious, noxious
		regexp.MustCompile(`sia$`),       // amnesia, polynesia
		regexp.MustCompile(`.gue$`),      // dialogue, intrigue, colleague
	}
)

// CountSyllables -
func CountSyllables(word []byte) (matches int) {
	word = alpha.ReplaceAll(bytes.ToLower(word), emptyByte) // Check only :alpha: / letters
	word = ending.ReplaceAll(word, emptyByte)               // Replace ending symbols

	// Words under equal or less than 3 characters are considered having 1 syllable
	if len(word) <= 3 {
		return 1
	}

	matches = len(regexp.MustCompile(`[aeiouy]+`).FindAll(word, -1))

	// For special cases this adds 1 syllable if match occurs
	for _, reg := range addSylMap {
		if result := len(reg.FindAll(word, -1)); result > 0 {
			matches++
		}
	}

	// For special cases this subtracts 1 syllable if match occurs
	for _, reg := range subSylMap {
		if result := len(reg.FindAll(word, -1)); result > 0 {
			matches--
		}
	}
	return matches
}

// CountSyllablesInText - count syllables in text
func CountSyllablesInText(text [][]byte) (count int) {
	for _, word := range text {
		count += CountSyllables(word)
	}
	return count
}
