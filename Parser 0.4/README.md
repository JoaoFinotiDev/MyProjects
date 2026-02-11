# Parser de Validação de Planilhas Excel

Aplicação CLI em Go para validação automática de células vazias em arquivos Excel (.xlsx), com geração de logs detalhados.

## 📋 Funcionalidades

- **Detecção automática** de arquivos .xlsx em diretório configurável
- **Validação completa** de células vazias em todas as colunas
- **Geração de logs** com timestamp e detalhamento por linha/coluna
- **Interface interativa** com menu de seleção
- **Ordenação inteligente** de arquivos por data de modificação
- **Performance otimizada** para processamento de grandes planilhas

## 🚀 Começando

### Pré-requisitos

- Go 1.25.4 ou superior
- Arquivos Excel (.xlsx) para validação

### Instalação

```bash
# Clone o repositório
git clone [seu-repositorio]

# Acesse o diretório
cd Parser\ 0.4

# Instale as dependências
go mod download

# Execute
go run main.go
```

## 📁 Estrutura do Projeto

```
Parser 0.4/
├── cmd/
│   └── menu.go                 # Interface do menu principal
├── internal/
│   ├── config/
│   │   └── config.go           # Configurações padrão
│   ├── domain/
│   │   └── models.go           # Modelos de domínio
│   ├── excel/
│   │   ├── reader.go           # Leitura de arquivos Excel
│   │   └── validator.go        # Validação de células
│   ├── filesystem/
│   │   └── scanner.go          # Varredura de diretórios
│   └── formatter/
│       └── formatter.go        # Formatação de erros
├── logger/
│   └── logger.go               # Sistema de logs
└── main.go                     # Ponto de entrada
```

## ⚙️ Configuração

Configurações padrão em `internal/config/config.go`:

```go
CaminhoPadrao: "./xlsxModels"    // Diretório dos arquivos Excel
SheetPadrao:   "Produto"         // Nome da planilha a validar
DiretorioLogs: "./logs"          // Diretório para salvar logs
```

## 💻 Uso

### 1. Preparação

Coloque seus arquivos `.xlsx` na pasta `./xlsxModels` (ou configure outro diretório)

### 2. Execução

```bash
go run main.go
```

### 3. Menu Interativo

```
Quantidade de Arquivos disponíveis: 3
1 - arquivo_recente.xlsx
2 - planilha_teste.xlsx
3 - dados_antigos.xlsx
4 - Atualizar lista de arquivos
5 - Sair

Escolha a opção desejada:
```

### 4. Saída

```
Planilha Carregada!
Colunas: 15 | Linhas: 500

--- ERROS ENCONTRADOS (12) ---
[ERRO] Linha 45, Coluna C (Descrição): CÉLULA VAZIA
[ERRO] Linha 78, Coluna E (Preço): CÉLULA VAZIA
...

Log salvo em: ./logs/log_validacao_arquivo_20241226_143022.txt
Células verificadas: 7500
Células vazias: 12
Tempo total: 234ms
```

## 🏗️ Arquitetura e Estrutura do Código

### Princípios de Design

O projeto segue **Clean Architecture** com separação clara de responsabilidades:

- **Modelos de domínio independentes** de implementação
- **Inversão de dependências** entre camadas
- **Single Responsibility Principle** para cada componente
- **Separação entre lógica de negócio e infraestrutura**

### Camadas Detalhadas

#### 1. **Domain Layer** (`internal/domain/`)

Núcleo da aplicação - define as entidades e contratos de negócio:

```go
// models.go - Entidades puras sem dependências externas

type ArquivoExcel struct {
    Nome            string
    Caminho         string
    DataModificacao time.Time
}

type ErroValidacao struct {
    Linha      int
    Coluna     string
    NomeColuna string
    Mensagem   string
}

// Enums para controle de fluxo
type TipoAcao int
const (
    AcaoArquivo TipoAcao = iota + 1
    AcaoRefresh
    AcaoSair
)
```

**Responsabilidades:**
- Definir estruturas de dados do negócio
- Métodos de formatação (`String()`)
- Tipos enumerados para estados da aplicação

#### 2. **Infrastructure Layer**

##### **FileSystem** (`internal/filesystem/`)

Abstrai operações de I/O do sistema operacional:

```go
type Scanner struct {
    Diretorio string
}

// ListarArquivos() implementa:
// 1. os.ReadDir() - leitura do diretório
// 2. Filtragem por extensão .xlsx
// 3. Extração de metadados (data de modificação)
// 4. Ordenação por recência
```

**Algoritmo de ordenação:**
```go
sort.Slice(arquivos, func(i, j int) bool {
    return arquivos[i].DataModificacao.After(arquivos[j].DataModificacao)
})
// Complexidade: O(n log n)
```

##### **Excel** (`internal/excel/`)

Encapsula a biblioteca `excelize` para operações em planilhas:

**Reader:**
```go
type Reader struct {
    arquivo   *excelize.File  // Arquivo aberto em memória
    sheetName string           // Sheet ativa
}

// ObterMetadados() retorna:
// - Cabeçalhos (linha 1)
// - Total de linhas (len(rows) - 1)
// - Nome da planilha
```

**Validator:**
```go
type Validator struct {
    arquivo    *excelize.File
    sheetName  string
    cabecalhos []string
}

// ValidarTudo() - Algoritmo de varredura:
// Para cada linha (2 até totalLinhas+1):
//     Para cada coluna (1 até len(cabecalhos)):
//         Converter número de coluna para letra (A, B, C...)
//         Ler célula (ex: "B5")
//         Se vazia → armazenar erro
```

**Complexidade:** O(n × m) onde n = linhas, m = colunas

##### **Logger** (`logger/`)

Sistema de persistência de relatórios:

```go
// Funções puras para composição:

gerarCaminhoLog()  // Cria nome: log_validacao_[arquivo]_[timestamp].txt
escreverLog()      // Grava erros linha a linha
SalvarLog()        // Orquestra: gerar caminho + criar dir + escrever
```

**Padrão de nomenclatura:**
```
log_validacao_planilha_produtos_20241226_143022.txt
                 └─ arquivo base ─┘ └─── timestamp ──┘
```

#### 3. **Application Layer**

##### **Formatter** (`internal/formatter/`)

Responsável por transformação de dados para apresentação:

```go
type Formatter struct{}

// FormatarErros() - Converte domain.ErroValidacao → string
// OrdenarErros() - Algoritmo de ordenação customizado:
//   1. Prioridade: ordem alfabética de coluna (A < B < C)
//   2. Desempate: ordem numérica de linha
```

**Exemplo de ordenação:**
```
Antes:           Depois:
Linha 5, Col C   Linha 3, Col A
Linha 3, Col A   Linha 5, Col A
Linha 5, Col A   Linha 5, Col C
```

##### **Config** (`internal/config/`)

Centraliza configurações da aplicação:

```go
type Config struct {
    CaminhoPadrao string  // Onde buscar .xlsx
    SheetPadrao   string  // Qual planilha validar
    DiretorioLogs string  // Onde salvar logs
}

// Pattern: Factory Method
func Nova() *Config {
    return &Config{
        CaminhoPadrao: "./xlsxModels",
        SheetPadrao:   "Produto",
        DiretorioLogs: "./logs",
    }
}
```

#### 4. **Presentation Layer** (`cmd/`)

Interface com usuário via terminal:

```go
type Menu struct{}

func (m *Menu) Executar(arquivos []ArquivoExcel) ResultadoMenu {
    // 1. Exibe lista numerada
    // 2. Destaca mais recente (bold)
    // 3. Adiciona opções de sistema
    // 4. Loop até entrada válida
    // 5. Retorna ação + contexto
}
```

**Fluxo de controle:**
```
Menu → ResultadoMenu{Acao, Arquivo} → main() → switch/case
```

### 5. **Orquestração** (`main.go`)

Ponto de entrada - coordena todos os componentes:

```go
func main() {
    cfg := config.Nova()              // 1. Configuração
    scanner := filesystem.NovoScanner() // 2. Scanner
    arquivos := scanner.ListarArquivos() // 3. Descoberta
    
    menu := cmd.NovoMenu()
    for {                              // 4. Loop principal
        resultado := menu.Executar(arquivos)
        switch resultado.Acao {
            case AcaoArquivo:
                processar()            // 5. Pipeline de validação
            case AcaoRefresh:
                arquivos = scanner.ListarArquivos()
            case AcaoSair:
                return
        }
    }
}

func processar(caminho, cfg) {
    reader := excel.NovoReader()       // 1. Leitura
    planilha := reader.ObterMetadados() // 2. Metadados
    validator := excel.NovoValidator()  // 3. Validação
    erros := validator.ValidarTudo()   // 4. Execução
    formatador := formatter.Novo()     // 5. Formatação
    formatador.OrdenarErros()
    logger.SalvarLog()                 // 6. Persistência
}
```

### Padrões de Projeto Implementados

#### Factory Pattern
```go
// Construtores padronizados
config.Nova()
filesystem.NovoScanner()
excel.NovoReader()
formatter.Novo()
```

#### Strategy Pattern (implícito)
```go
// Diferentes validadores podem implementar mesma interface
type Validador interface {
    ValidarTudo(totalLinhas int) []ErroValidacao
}
```

#### Dependency Injection
```go
// Validator recebe dependências via construtor
func NovoValidator(arquivo *excelize.File, sheetName string, cabecalhos []string)
```

### Fluxo de Dados Completo

```
[Sistema de Arquivos]
        ↓
    Scanner.ListarArquivos()
        ↓
    []ArquivoExcel (ordenados)
        ↓
    Menu.Executar()
        ↓
    ResultadoMenu{Acao: AcaoArquivo, Arquivo}
        ↓
    Reader.NovoReader(caminho)
        ↓
    Reader.ObterMetadados()
        ↓
    Planilha{Cabecalhos, TotalLinhas}
        ↓
    Validator.ValidarTudo()
        ↓
    []ErroValidacao (bruto)
        ↓
    Formatter.OrdenarErros()
        ↓
    Formatter.FormatarErros()
        ↓
    []string (mensagens)
        ↓
    Logger.SalvarLog()
        ↓
    [Arquivo de log no disco]
```

### Tratamento de Erros

Estratégia: **Propagação com contexto**

```go
// Padrão usado em toda aplicação
if err != nil {
    return nil, fmt.Errorf("contexto específico: %w", err)
}
```

**Exemplo real:**
```go
// filesystem/scanner.go
if err != nil {
    return nil, fmt.Errorf("erro ao obter info do arquivo '%s': %w", 
                          entrada.Name(), err)
}
```

### Gerenciamento de Recursos

```go
// Pattern: defer para cleanup garantido
reader, err := excel.NovoReader(caminho, sheet)
if err != nil {
    return
}
defer reader.Close()  // ← Garante fechamento do arquivo Excel
```


## 📦 Dependências

```
github.com/xuri/excelize/v2 v2.10.0
```

## 🎯 Casos de Uso

- **Validação pré-importação** de dados em sistemas ERP
- **Auditoria de qualidade** de planilhas comerciais
- **Automação de checagem** antes de processamento em lote
- **Detecção de inconsistências** em dados exportados

## 🛠️ Melhorias Futuras

- [ ] Validação de tipos de dados (numérico, data, etc)
- [ ] Regras customizáveis por coluna
- [ ] Exportação de relatórios em Excel
- [ ] Interface em web
- [ ] Suporte a múltiplas sheets
- [ ] Processamento paralelo com goroutines
- [ ] Cache de arquivos já validados

## 📝 Licença

Este projeto é de uso interno da Alternativa Sistemas e seus colaboradores.

---

**Desenvolvido por:** João Vitor  
**Empresa:** Alternativa Sistemas  
**Versão:** 0.4
