package english

import _ "embed"

//go:embed dict/index.adj
var _index_adj string

//go:embed dict/data.adj
var _data_adj string

//go:embed dict/index.adv
var _index_adv string

//go:embed dict/data.adv
var _data_adv string

//go:embed dict/index.noun
var _index_noun string

//go:embed dict/data.noun
var _data_noun string

//go:embed dict/index.verb
var _index_verb string

//go:embed dict/data.verb
var _data_verb string

//go:embed dict/index.sense
var _index_sense string
