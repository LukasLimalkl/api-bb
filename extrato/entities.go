package extrato

type Oauth struct {
	AccessToken string `json:"access_token"`
	Expire      int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type ExtratoBB struct {
	NumeroPaginaAtual             int       `json:"numeroPaginaAtual"`
	QuantidadeRegistroPaginaAtual int       `json:"quantidadeRegistroPaginaAtual"`
	NumeroPaginaAnterior          int       `json:"numeroPaginaAnterior"`
	NumeroPaginaProximo           int       `json:"numeroPaginaProximo"`
	QuantidadeTotalPagina         int       `json:"quantidadeTotalPagina"`
	QuantidadeTotalRegistro       int       `json:"quantidadeTotalRegistro"`
	ListaLancamento               []Release `json:"listaLancamento"`
}

type Release struct {
	IndicadorTipoLancamento          string  `json:"indicadorTipoLancamento"`
	DataLancamento                   int     `json:"dataLancamento"`
	DataMovimento                    int     `json:"dataMovimento"`
	CodigoAgenciaOrigem              int     `json:"codigoAgenciaOrigem"`
	NumeroLote                       int     `json:"numeroLote"`
	NumeroDocumento                  int     `json:"numeroDocumento"`
	CodigoHistorico                  int     `json:"codigoHistorico"`
	TextoDescricaoHistorico          string  `json:"textoDescricaoHistorico"`
	ValorLancamento                  float64 `json:"valorLancamento"`
	IndicadorSinalLancamento         string  `json:"indicadorSinalLancamento"`
	TextoInformacaoComplementar      string  `json:"textoInformacaoComplementar"`
	NumeroCpfCnpjContrapartida       int     `json:"numeroCpfCnpjContrapartida"`
	IndicadorTipoPessoaContrapartida string  `json:"indicadorTipoPessoaContrapartida"`
	CodigoBancoContrapartida         int     `json:"codigoBancoContrapartida"`
	CodigoAgenciaContrapartida       int     `json:"codigoAgenciaContrapartida"`
	NumeroContaContrapartida         string  `json:"numeroContaContrapartida"`
	TextoDvContaContrapartida        string  `json:"textoDvContaContrapartida"`
}
