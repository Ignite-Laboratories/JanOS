package english

func init() {
	/**
	Adjective
	*/
	if len(_data_adj) == 0 {
		panic("no adjective Lexis found")
	}
	parseData(_data_adj)
	if len(_index_adj) == 0 {
		panic("no adjective index found")
	}
	parseIndex(_index_adj)

	/**
	Adverb
	*/
	if len(_data_adv) == 0 {
		panic("no adverb Lexis found")
	}
	parseData(_data_adv)
	if len(_index_adv) == 0 {
		panic("no adverb index found")
	}
	parseIndex(_index_adv)

	/**
	Noun
	*/
	if len(_data_noun) == 0 {
		panic("no noun Lexis found")
	}
	parseData(_data_noun)
	if len(_index_noun) == 0 {
		panic("no noun index found")
	}
	parseIndex(_index_noun)

	/**
	Verb
	*/
	if len(_data_verb) == 0 {
		panic("no verb Lexis found")
	}
	parseData(_data_verb)
	if len(_index_verb) == 0 {
		panic("no verb index found")
	}
	parseIndex(_index_verb)

	/**
	Sense
	*/
	//if len(_index_sense) == 0 {
	//	panic("no sense index found")
	//}
	//parseIndex(_index_sense)
}
