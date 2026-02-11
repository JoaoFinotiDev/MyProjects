package excel

import (
	"ParserTrib/internal/domain"
	"regexp"
	"strconv"
	"strings"
)

// Validator valida dados da planilha Excel
type Validator struct {
	rows        [][]string
	sheetName   string
	cabecalhos  []string
	mapaIndices map[string]int
}

// NovoValidator cria instância do validador
func NovoValidator(rows [][]string, sheetName string, cabecalhos []string) *Validator {
	mapaIndices := make(map[string]int)
	for i, cab := range cabecalhos {
		mapaIndices[cab] = i
	}

	return &Validator{
		rows:        rows,
		sheetName:   sheetName,
		cabecalhos:  cabecalhos,
		mapaIndices: mapaIndices,
	}
}

// ValidarTudo executa as 4 validações
func (v *Validator) ValidarTudo(totalLinhas int) domain.ResultadoValidacaoCompleto {
	return domain.ResultadoValidacaoCompleto{
		ErrosVazias:    v.validarVazias(),
		ErrosNCM:       v.validarNCM(),
		ErrosCSTOrigem: v.validarCSTOrigem(),
		ErrosCSOSN:     v.validarCSOSN(),
		ErrosTipoItem:  v.validarTipoItem(),
	}
}

func (v *Validator) validarVazias() []domain.ErroValidacao {
	var erros []domain.ErroValidacao

	if len(v.rows) <= 1 {
		return erros
	}

	maxColunas := len(v.cabecalhos)

	for i := 1; i < len(v.rows); i++ {
		linha := v.rows[i]
		numLinha := i + 1

		for j := 0; j < maxColunas; j++ {
			celula := ""
			if j < len(linha) {
				celula = strings.TrimSpace(linha[j])
			}

			if celula == "" {
				colLetra := indiceParaLetra(j)
				nomeColuna := v.cabecalhos[j]

				erros = append(erros, domain.ErroValidacao{
					Linha:      numLinha,
					Coluna:     colLetra,
					NomeColuna: nomeColuna,
					Mensagem:   "CÉLULA VAZIA",
				})
			}
		}
	}

	return erros
}

// validarNCM verifica TODAS as células da coluna NCM com regra fiscal
func (v *Validator) validarNCM() []domain.ErroValidacao {
	var erros []domain.ErroValidacao

	indiceNCM, existe := v.mapaIndices["NCM"]
	if !existe {
		return erros
	}

	regexNCM := regexp.MustCompile(`^\d{8}$`)

	for i := 1; i < len(v.rows); i++ {
		linha := v.rows[i]
		numLinha := i + 1

		valorNCM := ""
		if indiceNCM < len(linha) {
			valorNCM = strings.TrimSpace(linha[indiceNCM])
		}

		if valorNCM == "" {
			// Já reportado por validarVazias — pular
			continue
		} else if !regexNCM.MatchString(valorNCM) {
			erros = append(erros, domain.ErroValidacao{
				Linha:      numLinha,
				Coluna:     indiceParaLetra(indiceNCM),
				NomeColuna: "NCM",
				Mensagem:   "NCM INVÁLIDO - deve conter exatamente 8 dígitos numéricos (atual: '" + valorNCM + "')",
			})
		}
	}

	return erros
}

// validarCSTOrigem verifica TODAS as células da coluna CST Origem com regra fiscal
func (v *Validator) validarCSTOrigem() []domain.ErroValidacao {
	var erros []domain.ErroValidacao

	indiceCST, existe := v.mapaIndices["CST Origem"]
	if !existe {
		return erros
	}

	for i := 1; i < len(v.rows); i++ {
		linha := v.rows[i]
		numLinha := i + 1

		valorCST := ""
		if indiceCST < len(linha) {
			valorCST = strings.TrimSpace(linha[indiceCST])
		}

		if valorCST == "" {
			// Já reportado por validarVazias — pular
			continue
		} else {
			num, err := strconv.Atoi(valorCST)
			if err != nil {
				erros = append(erros, domain.ErroValidacao{
					Linha:      numLinha,
					Coluna:     indiceParaLetra(indiceCST),
					NomeColuna: "CST Origem",
					Mensagem:   "CST ORIGEM INVÁLIDO - deve ser um número entre 0 e 8 (atual: '" + valorCST + "')",
				})
			} else if num < 0 || num > 8 {
				erros = append(erros, domain.ErroValidacao{
					Linha:      numLinha,
					Coluna:     indiceParaLetra(indiceCST),
					NomeColuna: "CST Origem",
					Mensagem:   "CST ORIGEM FORA DO RANGE - deve estar entre 0 e 8 (atual: " + valorCST + ")",
				})
			}
		}
	}

	return erros
}

// validarCSOSN verifica TODAS as células da coluna CSOSN com regra fiscal
func (v *Validator) validarCSOSN() []domain.ErroValidacao {
	var erros []domain.ErroValidacao

	indiceCSOSN, existe := v.mapaIndices["CSOSN"]
	if !existe {
		return erros
	}

	// CSOSN válidos conforme legislação do Simples Nacional
	csosnValidos := map[string]bool{
		"101": true, // Tributada pelo Simples Nacional com permissão de crédito
		"102": true, // Tributada pelo Simples Nacional sem permissão de crédito
		"103": true, // Isenção do ICMS no Simples Nacional para faixa de receita bruta
		"201": true, // Tributada pelo Simples Nacional com permissão de crédito e com cobrança do ICMS por ST
		"202": true, // Tributada pelo Simples Nacional sem permissão de crédito e com cobrança do ICMS por ST
		"203": true, // Isenção do ICMS no Simples Nacional para faixa de receita bruta e com cobrança do ICMS por ST
		"300": true, // Imune
		"400": true, // Não tributada pelo Simples Nacional
		"500": true, // ICMS cobrado anteriormente por substituição tributária ou por antecipação
		"900": true, // Outros
	}

	for i := 1; i < len(v.rows); i++ {
		linha := v.rows[i]
		numLinha := i + 1

		valorCSOSN := ""
		if indiceCSOSN < len(linha) {
			valorCSOSN = strings.TrimSpace(linha[indiceCSOSN])
		}

		if valorCSOSN == "" {
			// Já reportado por validarVazias — pular
			continue
		} else if !csosnValidos[valorCSOSN] {
			erros = append(erros, domain.ErroValidacao{
				Linha:      numLinha,
				Coluna:     indiceParaLetra(indiceCSOSN),
				NomeColuna: "CSOSN",
				Mensagem:   "CSOSN INVÁLIDO - deve ser um dos códigos válidos: 101, 102, 103, 201, 202, 203, 300, 400, 500, 900 (atual: '" + valorCSOSN + "')",
			})
		}
	}

	return erros
}

// validarTipoItem verifica TODAS as células da coluna "Tipo Item" com regra fiscal
func (v *Validator) validarTipoItem() []domain.ErroValidacao {
	var erros []domain.ErroValidacao

	indiceTipoItem, existe := v.mapaIndices["Tipo Item"]
	if !existe {
		return erros
	}

	// Tipos de item válidos conforme tabela fiscal
	tiposValidos := map[string]bool{
		"00": true, // Mercadoria para Revenda
		"01": true, // Matéria-Prima
		"02": true, // Embalagem
		"03": true, // Produto em Processo
		"04": true, // Produto Acabado
		"05": true, // Subproduto
		"06": true, // Produto Intermediário
		"07": true, // Material de Uso e Consumo
		"08": true, // Ativo Imobilizado
		"09": true, // Serviços
		"10": true, // Outros insumos
		"99": true, // Outras
	}

	for i := 1; i < len(v.rows); i++ {
		linha := v.rows[i]
		numLinha := i + 1

		valorTipoItem := ""
		if indiceTipoItem < len(linha) {
			valorTipoItem = strings.TrimSpace(linha[indiceTipoItem])
		}

		if valorTipoItem == "" {
			// Já reportado por validarVazias — pular
			continue
		} else {
			// Validar se é inteiro
			_, err := strconv.Atoi(valorTipoItem)
			if err != nil {
				erros = append(erros, domain.ErroValidacao{
					Linha:      numLinha,
					Coluna:     indiceParaLetra(indiceTipoItem),
					NomeColuna: "Tipo Item",
					Mensagem:   "TIPO ITEM INVÁLIDO - deve ser um número inteiro (atual: '" + valorTipoItem + "')",
				})
			} else if !tiposValidos[valorTipoItem] {
				// É inteiro mas não está na tabela válida
				erros = append(erros, domain.ErroValidacao{
					Linha:      numLinha,
					Coluna:     indiceParaLetra(indiceTipoItem),
					NomeColuna: "Tipo Item",
					Mensagem:   "TIPO ITEM FORA DA TABELA - deve ser um dos códigos válidos: 00, 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 99 (atual: '" + valorTipoItem + "')",
				})
			}
		}
	}

	return erros
}

// indiceParaLetra converte índice numérico para letra Excel (0=A, 1=B, 26=AA)
func indiceParaLetra(indice int) string {
	letra := ""
	indice++

	for indice > 0 {
		indice--
		letra = string(rune('A'+(indice%26))) + letra
		indice /= 26
	}

	return letra
}
