package english

import "core/enum/partOfSpeech"

type Synset struct {
	Lexeme

	PartOfSpeech partOfSpeech.PartOfSpeech
	Gloss        string

	Synonyms []Synset
	Antonyms []Synset

	Hypernyms []Synset
	Hyponyms  []Synset

	Holonyms []Synset
	Meronyms []Synset

	Pertainyms []Synset
}

func (s Synset) String() string {
	return s.Lemma
}
