export interface ValidationError {
  linha: number;
  coluna: string;
  nomeColuna: string;
  tipo: 'VAZIA' | 'NCM' | 'CSOSN' | 'CST_ORIGEM' | 'TIPO_ITEM';
  mensagem: string;
}

export interface ValidationResult {
  nomeArquivo: string;
  processingTime: string;
  totalErros: number;
  errosVazias: number;
  errosNCM: number;
  errosCSTOrigem: number;
  errosCSOSN: number;
  errosTipoItem: number;
  detalhes: ValidationError[];
}

export type ErrorFilter = 'all' | 'VAZIA' | 'NCM' | 'CSOSN' | 'CST_ORIGEM' | 'TIPO_ITEM';
