import { useCallback } from 'react';
import { useDropzone } from 'react-dropzone';
import { Upload, FileSpreadsheet, Loader2, AlertCircle, Info } from 'lucide-react';
import { motion, AnimatePresence } from 'framer-motion';

interface UploadAreaProps {
  onUpload: (file: File) => void;
  loading: boolean;
  erro: string | null;
}

const UploadArea = ({ onUpload, loading, erro }: UploadAreaProps) => {
  const onDrop = useCallback((acceptedFiles: File[]) => {
    if (acceptedFiles.length > 0) {
      onUpload(acceptedFiles[0]);
    }
  }, [onUpload]);

    const { getRootProps, getInputProps, isDragActive, fileRejections } = useDropzone({
        accept: {
            'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet': ['.xlsx']
        },
        maxFiles: 1,
        onDrop,
        disabled: loading
    });

  const hasRejection = fileRejections.length > 0 || !!erro;

  return (
    <section className="max-w-4xl mx-auto px-4 sm:px-6 py-8 sm:py-16">
      <div className="text-center mb-8">
        <motion.h2 
          initial={{ opacity: 0, y: -10 }}
          animate={{ opacity: 1, y: 0 }}
          className="text-2xl sm:text-3xl font-bold text-foreground mb-3"
        >
          Utilize nossa ferramenta para validar sua planilha antes de encaminha-la para a nossa equipe:
        </motion.h2>
          <motion.p
              initial={{ opacity: 0, y: -10 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.1 }}
              className="text-muted-foreground"
          >
              Formato aceito: .xlsx
          </motion.p>
      </div>

      <div
        {...getRootProps()}
        className={`
          relative border-2 border-dashed rounded-xl p-8 sm:p-16 cursor-pointer
          transition-all duration-300 bg-card
          ${isDragActive ? 'border-primary bg-primary/5 scale-[1.02]' : 'border-border hover:border-primary/50'}
          ${loading ? 'pointer-events-none' : ''}
          ${hasRejection ? 'border-destructive bg-destructive/5' : ''}
        `}
      >
        <input {...getInputProps()} />
        
        <div className="text-center">
          <AnimatePresence mode="wait">
            {loading ? (
              <motion.div
                key="loading"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
              >
                <Loader2 className="w-16 h-16 text-primary mx-auto mb-4 animate-spin" />
                <p className="text-lg font-medium text-foreground">Validando planilha...</p>
                <p className="text-sm text-muted-foreground mt-2">Isso pode levar alguns segundos</p>
              </motion.div>
            ) : hasRejection ? (
              <motion.div
                key="error"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
              >
                <AlertCircle className="w-16 h-16 text-destructive mx-auto mb-4" />
                <p className="text-lg font-medium text-destructive mb-2">Arquivo inválido</p>
                  <p className="text-muted-foreground">
                      {erro
                          ? erro
                          : 'Por favor, selecione um arquivo .xlsx válido'
                      }
                  </p>
              </motion.div>
            ) : (
              <motion.div
                key="idle"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
              >
                {isDragActive ? (
                  <FileSpreadsheet className="w-16 h-16 text-primary mx-auto mb-4 animate-pulse-slow" />
                ) : (
                  <Upload className="w-16 h-16 text-muted-foreground mx-auto mb-4" />
                )}
                <p className="text-xl font-medium text-foreground mb-2">
                  {isDragActive ? 'Solte o arquivo aqui' : 'Arraste o arquivo aqui'}
                </p>
                <p className="text-muted-foreground">ou clique para selecionar</p>
              </motion.div>
            )}
          </AnimatePresence>
        </div>
      </div>

      <motion.div 
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.3 }}
        className="mt-6 p-4 bg-primary/5 rounded-lg border border-primary/20 flex items-start gap-3"
      >
        <Info className="w-5 h-5 text-primary flex-shrink-0 mt-0.5" />
        <p className="text-sm text-foreground">
          <strong>Validações realizadas:</strong> NCM (8 dígitos), CSOSN (códigos válidos), 
          CST Origem (0-8), Tipo Item (tabela fiscal), células vazias em campos obrigatórios.
        </p>
      </motion.div>
    </section>
  );
};

export default UploadArea;
