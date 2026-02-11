import { X, FileUp, Search, AlertTriangle, Download, RotateCcw } from 'lucide-react';
import { motion, AnimatePresence } from 'framer-motion';

interface HelpModalProps {
  onClose: () => void;
}

const HelpModal = ({ onClose }: HelpModalProps) => {
  const steps = [
    {
      icon: FileUp,
      title: '1. Prepare sua planilha',
      description: 'A planilha deve ser a mesma fornecida pela nossa equipe, agora devidamente preenchida e com as colunas fiscais obrigatórias.'
    },
    {
      icon: Search,
      title: '2. Faça o upload',
      description: 'Arraste o arquivo para a área indicada ou clique para selecionar manualmente.'
    },
    {
      icon: AlertTriangle,
      title: '3. Aguarde a validação',
      description: 'O sistema verificará: NCM (8 dígitos), CSOSN (códigos válidos), CST Origem (0-8), Tipo Item e células vazias.'
    },
    {
      icon: Download,
      title: '4. Baixe o relatório',
      description: 'Baixe o relatório completo com todos os erros encontrados para corrigir na planilha original.'
    },
    {
      icon: RotateCcw,
      title: '5. Valide novamente',
      description: 'Após corrigir os erros, valide novamente até que não haja mais pendências.'
    }
  ];

  return (
    <AnimatePresence>
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        exit={{ opacity: 0 }}
        className="fixed inset-0 bg-foreground/50 backdrop-blur-sm flex items-center justify-center p-4 z-50"
        onClick={onClose}
      >
        <motion.div
          initial={{ scale: 0.95, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          exit={{ scale: 0.95, opacity: 0 }}
          transition={{ type: 'spring', damping: 25, stiffness: 300 }}
          className="bg-card rounded-xl max-w-2xl w-full max-h-[85vh] overflow-y-auto shadow-2xl"
          onClick={(e) => e.stopPropagation()}
        >
          <div className="sticky top-0 bg-card border-b border-border px-6 py-4 flex justify-between items-center">
            <h2 className="text-2xl font-bold text-foreground">Como Usar</h2>
            <button 
              onClick={onClose} 
              className="p-2 rounded-lg text-muted-foreground hover:text-foreground hover:bg-muted transition-colors"
            >
              <X className="w-6 h-6" />
            </button>
          </div>

          <div className="p-6 space-y-6">
            {steps.map((step, index) => (
              <motion.div
                key={step.title}
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: index * 0.1 }}
                className="flex gap-4"
              >
                <div className="flex-shrink-0 w-12 h-12 bg-primary/10 rounded-lg flex items-center justify-center">
                  <step.icon className="w-6 h-6 text-primary" />
                </div>
                <div>
                  <h3 className="font-semibold text-lg text-foreground mb-1">{step.title}</h3>
                  <p className="text-muted-foreground">{step.description}</p>
                </div>
              </motion.div>
            ))}

            <div className="mt-8 p-4 bg-primary/5 rounded-lg border border-primary/20">
              <h4 className="font-semibold text-foreground mb-2">Validações Realizadas</h4>
              <ul className="text-sm text-muted-foreground space-y-1">
                <li>• <strong>NCM:</strong> Deve conter exatamente 8 dígitos numéricos, sem pontos ou caracteres especiais</li>
                <li>• <strong>CSOSN:</strong> Códigos válidos do Simples Nacional</li>
                <li>• <strong>CST Origem:</strong> Valores de 0 a 8</li>
                <li>• <strong>Tipo Item:</strong> Conforme tabela fiscal vigente</li>
                <li>• <strong>Células Vazias:</strong> Campos obrigatórios preenchidos</li>
              </ul>
            </div>

            <div className="p-4 bg-accent/10 rounded-lg border border-accent/20">
              <p className="text-sm text-foreground">
                <strong>Precisa de ajuda?</strong> Entre em contato pelo email{' '}
                <a href="mailto:imp@alternativasistemas.com.br" className="text-accent hover:underline font-medium">
                  imp@alternativasistemas.com.br
                </a>
              </p>
            </div>
          </div>
        </motion.div>
      </motion.div>
    </AnimatePresence>
  );
};

export default HelpModal;
