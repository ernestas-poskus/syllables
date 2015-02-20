# syllables
A syllable is a unit of organization for a sequence of speech sounds.

[![Build Status](https://travis-ci.org/ernestas-poskus/syllables.svg)](https://travis-ci.org/ernestas-poskus/syllables)
[![GoDoc](http://godoc.org/github.com/ernestas-poskus/syllables?status.svg)](http://godoc.org/github.com/ernestas-poskus/syllables)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)

Programmatically count number of syllables in word, some specific words requires more attention and might vary between +/-1 syllable.


Benchmark
-----------------

'Longest' word has 189,819 letters, takes three hours to pronounce and it takes 1/3 of seconds to process
Read more: http://www.digitalspy.co.uk/fun/news/a444700/longest-word-has-189819-letters-takes-three-hours-to-pronounce.html#ixzz3SITIl2nM 

| Function | operations | ns/op | sec/op |
| -------- | ---------- | ----- | ------ |
| BenchmarkCountSyllables | 20000 | 97600 ns/op | 9.76e-5 sec/op |
| BenchmarkCountSyllables (Longest word 189,819 char)| 5 | 318114133 ns/op | 0.318114133 sec/op |
