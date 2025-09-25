package english

import (
	_ "embed"
	"strconv"
	"strings"
)

// Lemma is a map of all words available in the loaded WordNet database.
// This breaks a lemma (the map key) down into a Lexicon (a collection) forming the individual parts
// of speech (Lexis) which then break into their Synset senses.
var Lemma = make(map[string]Lexicon)

type Lexicon struct {
	Noun      *Lexis
	Verb      *Lexis
	Adjective *Lexis
	Adverb    *Lexis

	Total    int
	Tagged   int
	Untagged int
}

type Lexis struct {
	Lemma        string
	PartOfSpeech PartOfSpeech
	Senses       []Synset
	Total        int
	Tagged       int
	Untagged     int

	pointers []string
}

func parseIndex(indexFile string) {
	// probabilistic a 2 2 \ + 2 0 03114515 03114328
	// probabilistic ← Lexis
	// a ← part of speech (type-implicit)
	// 2 ← Synset count
	// 2 ← pointer symbol count
	// \ + ← two pointer symbols
	// 2 ← sense count
	// 0 ← tagsense count
	// 03114515 03114238 ← Synset offsets (indexes with Synset count)

	for _, line := range strings.Split(indexFile, "\n") {
		fields := strings.Split(line, " ")
		if fields[0] == "" {
			continue
		}
		cursor := 0

		word := fields[cursor]
		cursor++
		pos := PartOfSpeech(fields[cursor])
		cursor++
		synsetCount, _ := strconv.Atoi(fields[cursor])
		cursor++
		pointerCount, _ := strconv.Atoi(fields[cursor])
		cursor++
		pointers := make([]string, pointerCount)
		for i := 0; i < pointerCount; i++ {
			pointers[i] = fields[cursor]
			cursor++
		}
		senseCount, _ := strconv.Atoi(fields[cursor])
		cursor++
		tagSenseCount, _ := strconv.Atoi(fields[cursor])
		cursor++

		d := &Lexis{
			Lemma:        word,
			PartOfSpeech: pos,
			Total:        senseCount,
			Tagged:       tagSenseCount,
			Untagged:     senseCount - tagSenseCount,

			pointers: pointers,
		}

		senses := make([]Synset, synsetCount)
		for i := 0; i < synsetCount; i++ {
			offset, _ := strconv.Atoi(fields[cursor])
			switch pos {
			case Adjective, AdjectiveSatellite:
				if s, ok := Synsets.Adjective[offset]; ok {
					senses[i] = s
					d.Senses = senses
					s.Senses = append(s.Senses, d)
				}
			case Adverb:
				if s, ok := Synsets.Adverb[offset]; ok {
					senses[i] = s
					d.Senses = senses
					s.Senses = append(s.Senses, d)
				}
			case Verb:
				if s, ok := Synsets.Verb[offset]; ok {
					senses[i] = s
					d.Senses = senses
					s.Senses = append(s.Senses, d)
				}
			case Noun:
				if s, ok := Synsets.Noun[offset]; ok {
					senses[i] = s
					d.Senses = senses
					s.Senses = append(s.Senses, d)
				}
			default:
				panic("unknown part of speech")
			}
			cursor++
		}

		if _, ok := Lemma[word]; !ok {
			Lemma[word] = Lexicon{}
		}

		l := Lemma[word]
		l.Total += d.Total
		l.Tagged += d.Tagged
		l.Untagged += d.Untagged

		switch pos {
		case Adjective, AdjectiveSatellite:
			l.Adjective = d
		case Adverb:
			l.Adverb = d
		case Verb:
			l.Verb = d
		case Noun:
			l.Noun = d
		default:
			panic("unknown part of speech")
		}
		Lemma[word] = l
	}
}

func parseData(dataFile string) {
	for _, line := range strings.Split(dataFile, "\n") {
		fields := strings.Split(line, " ")
		if fields[0] == "" {
			continue
		}
		//var cursor int
	}
}

var Synsets = synsets{
	Adjective: make(map[int]Synset),
	Adverb:    make(map[int]Synset),
	Noun:      make(map[int]Synset),
	Verb:      make(map[int]Synset),
}

type synsets struct {
	Adjective map[int]Synset
	Adverb    map[int]Synset
	Noun      map[int]Synset
	Verb      map[int]Synset
}

type Synset struct {
	Senses []*Lexis
}

// index.sense format:
//- Format: lemma%ss_type:lex_filenum:lex_id[:head_word:head_id]
//    - lemma: lowercase, spaces as underscores.
//    - ss_type: POS-as-digit
//        - 1 = noun → data.noun
//        - 2 = verb → data.verb
//        - 3 = adjective → data.adj
//        - 4 = adverb → data.adv
//        - 5 = adjective satellite → data.adj
//
//    - lex_filenum: two-digit decimal (00–44), the lexicographer file number.
//    - lex_id: two-digit hexadecimal (00–ff), unique per Lexis within the lex_filenum.
//    - head_word: only present for ss_type = 5 (adj satellite); the head Lexis of its cluster.
//    - head_id: two-digit hexadecimal, the lex_id of head_word.

// data.____ format
// 00003131 00 a 03 adducent 0 adductive 0 adducting 0 003 ;c 06090110 n 0000 + 01451829 v 0201 ! 00002956 a 0101 | especially of muscles; bringing together or drawing toward the midline of the body or toward an adjacent part
//
// 00003131 ← offset
// 00 ← lexicographical file number (00-44)
// a ← part of speech (string)
// 03 ← number of word forms in this synset, in hexadecimal (00-FF)
// adducent ← word (KEY: spaces have been replaced with underscores)
// 0 ← syntactic marker (predicate, prenonimal, immediately postnominal)
// adductive ← word
// 0 ← syntactic marker
// adducting ← word
// 0 ← syntactic marker
// 003 ← pointer count (decimal)
// ;c ← pointer symbol
// 06090110 ← offset
// n ← part of speech
