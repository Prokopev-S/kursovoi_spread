package types

type ElemsKline struct {
	Symbol string     `json:"symbol"`
	List   [][]string `json:"list"`
}

type ExchangeFunding struct {
	Exchange string       `json:"exchange"`
	Elems    []ElemsKline `json:"elems"`
}
