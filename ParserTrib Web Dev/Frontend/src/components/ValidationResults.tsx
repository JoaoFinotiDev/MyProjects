import { Download, RotateCcw, AlertTriangle, CheckCircle, Clock, FileText } from 'lucide-react';
import { motion } from 'framer-motion';
import ErrorSummaryCards from './ErrorSummaryCards';
import ErrorTable from './ErrorTable';
import ErrorChart from './ErrorChart';
import type { ValidationResult } from '../types/validation';

interface ValidationResultsProps {
  data: ValidationResult;
  onNewValidation: () => void;
}

const ValidationResults = ({ data, onNewValidation }: ValidationResultsProps) => {
  const hasErrors = data.totalErros > 0;

  const downloadLog = () => {
    const lines = [
      `Relatório de Validação - ${data.nomeArquivo}`,
      `Gerado em: ${new Date().toLocaleString('pt-BR')}`,
      `Tempo de processamento: ${data.processingTime}`,
      '',
      '='.repeat(60),
      'RESUMO',
      '='.repeat(60),
      '',
      `Total de Erros: ${data.totalErros}`,
      `Células Vazias: ${data.errosVazias}`,
      `Erros NCM: ${data.errosNCM}`,
      `Erros CSOSN: ${data.errosCSOSN}`,
      `Erros CST Origem: ${data.errosCSTOrigem}`,
      `Erros Tipo Item: ${data.errosTipoItem}`,
      '',
      '='.repeat(60),
      'DETALHES DOS ERROS',
      '='.repeat(60),
      '',
      'Linha\tColuna\tCampo\t\tTipo\t\tMensagem',
      '-'.repeat(80),
    ];

    data.detalhes.forEach((erro) => {
      lines.push(
        `${erro.linha}\t${erro.coluna}\t${erro.nomeColuna}\t\t${erro.tipo}\t\t${erro.mensagem}`
      );
    });

    lines.push('', '-'.repeat(60), 'Fim do relatório');

    const content = lines.join('\n');
    const blob = new Blob([content], { type: 'text/plain;charset=utf-8' });
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = `relatorio_validacao_${data.nomeArquivo.replace('.xlsx', '')}_${Date.now()}.txt`;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    URL.revokeObjectURL(url);
  };

  return (
    <section className="max-w-6xl mx-auto px-4 sm:px-6 py-8 sm:py-12">
      {/* Header com status */}
      <motion.div 
        initial={{ opacity: 0, y: -10 }}
        animate={{ opacity: 1, y: 0 }}
        className={`p-4 sm:p-6 rounded-lg mb-6 ${hasErrors ? 'bg-destructive/10 border border-destructive/20' : 'bg-success/10 border border-success/20'}`}
      >
        <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
          <div>
            <div className="flex items-center gap-3 mb-2">
              {hasErrors ? (
                <AlertTriangle className="w-6 h-6 text-destructive" />
              ) : (
                <CheckCircle className="w-6 h-6 text-success" />
              )}
              <h2 className="text-xl sm:text-2xl font-bold text-foreground">
                {hasErrors ? 'Validação Concluída! Erros de preenchimento encontrados!' : 'Validação Concluída com Sucesso'}
              </h2>
            </div>
            <div className="flex flex-wrap items-center gap-4 text-sm text-muted-foreground">
              <span className="flex items-center gap-1">
                <FileText className="w-4 h-4" />
                {data.nomeArquivo}
              </span>
              <span className="flex items-center gap-1">
                <Clock className="w-4 h-4" />
                Processado em {data.processingTime}
              </span>
            </div>
          </div>
        </div>
      </motion.div>

      {hasErrors && (
        <motion.div 
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 0.1 }}
          className="mb-6 p-4 bg-warning/10 border border-warning/20 rounded-lg"
        >
          <p className="text-foreground font-medium flex items-center gap-2">
            <AlertTriangle className="w-5 h-5 text-warning" />
            Atenção: Foram encontrados <strong className="text-destructive">{data.totalErros} erros</strong> que precisam ser corrigidos antes do envio da planilha para a nossa equipe.
          </p>
        </motion.div>
      )}

      {/* Cards de resumo */}
      <motion.div
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.15 }}
      >
        <ErrorSummaryCards data={data} />
      </motion.div>

      {/* Tabela de erros */}
      <motion.div 
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.2 }}
        className="mt-8"
      >
        <h3 className="text-lg sm:text-xl font-bold text-foreground mb-4">Detalhes dos Erros</h3>
        <ErrorTable errors={data.detalhes} />
      </motion.div>

      {/* Gráfico de distribuição */}
      {hasErrors && (
        <motion.div 
          initial={{ opacity: 0, y: 10 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.25 }}
          className="mt-8"
        >
          <h3 className="text-lg sm:text-xl font-bold text-foreground mb-4">Distribuição de Erros</h3>
          <ErrorChart data={data} />
        </motion.div>
      )}

      {/* Botões de ação */}
      <motion.div 
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.3 }}
        className="mt-10 flex flex-col sm:flex-row gap-4 justify-center"
      >
        {hasErrors && (
          <button
            onClick={downloadLog}
            className="btn-accent flex items-center justify-center gap-2 px-8 py-4"
          >
            <Download className="w-5 h-5" />
            Baixar Relatório Completo (.txt)
          </button>
        )}
        <button
          onClick={onNewValidation}
          className="btn-secondary flex items-center justify-center gap-2 px-8 py-4"
        >
          <RotateCcw className="w-5 h-5" />
          Validar Outro Arquivo
        </button>
      </motion.div>
    </section>
  );
};

export default ValidationResults;
