package syllables

import "testing"

// Syllable count verified
// http://www.syllablecount.com/syllables/word
var countSyllablesTest = map[string]int{
	"University":                                    5,
	"creaturely":                                    3,
	"animal":                                        3,
	"animallike":                                    4,
	"yank":                                          1,
	"slightly":                                      2,
	"magazine":                                      3,
	"Synonyms":                                      3,
	"understood":                                    3,
	"how":                                           1,
	"tacit":                                         2,
	"comprehended":                                  3,
	"antilogarithm":                                 5,
	"logarithm":                                     3,
	"themselves":                                    2,
	"makes":                                         1,
	"used":                                          1,
	"logorrhoea":                                    4,
	"rule":                                          1,
	"cake":                                          1,
	"open":                                          2,
	"eight":                                         1,
	"eighteen":                                      2,
	"funicular":                                     4,
	"yellow":                                        2,
	"Trihalomethane":                                5,
	"jackdaw":                                       2,
	"less":                                          1,
	"guess":                                         1,
	"guessed":                                       1,
	"conspicuous":                                   4,
	"pronounced":                                    2,
	"pneumonoultramicroscopicsilicovolcanoconiosis": 18,
	"thyroparathyroidectomized":                     9,
	"pseudopseudohypoparathyroidism":                12,
	"deoxyribonucleic":                              6,
	"aminoheptafluorocyclotetraphosphonitrile":      15,
	"The": 1,
	// "Australian": 3,
	"platypus":  3,
	"is":        1,
	"seemingly": 3,
	"a":         1,
	"hybrid":    2,
	"of":        1,
	"mammal":    2,
	"and":       1,
	"reptilian": 4,
	"creature":  2,

	// Difference in 1 syllable
	// "Algorithm":                                     4,
	// "biorhythm":                                     2,
	// "rhythm":                                        2,
	// "theorem":                                       3,
}

func TestCountSyllables(t *testing.T) {
	for word, expectedSyllablesCount := range countSyllablesTest {
		got := CountSyllables([]byte(word))
		if got != expectedSyllablesCount {
			t.Errorf("Expected %s: %d got %d", word, expectedSyllablesCount, got)
		}
	}
}
