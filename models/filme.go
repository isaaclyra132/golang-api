package models

type Filme struct {
	ID            string `json:"id"`
	Titulo        string `json:"titulo"`
	Direcao       string `json:"direcao"`
	Producao      string `json:"producao"`
	AnoLancamento int    `json:"ano_lancamento"`
}
