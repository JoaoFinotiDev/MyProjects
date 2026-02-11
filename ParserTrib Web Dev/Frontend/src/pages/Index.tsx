import { useState } from 'react';
import { useMutation } from '@tanstack/react-query';
import Header from '../components/Header';
import UploadArea from '../components/UploadArea';
import ValidationResults from '../components/ValidationResults';
import Footer from '../components/Footer';
import HelpModal from '../components/HelpModal';
import { useTheme } from '../hooks/useTheme';
import type { ValidationResult } from '../types/validation';

// Função que faz a chamada POST /api/validar
async function validarArquivo(file: File): Promise<ValidationResult> {
  const formData = new FormData();
  formData.append('file', file);

  const response = await fetch('/api/validar', {
    method: 'POST',
    body: formData,
  });

  if (!response.ok) {
    const erro = await response.json().catch(() => ({ erro: 'Erro desconhecido' }));
    throw new Error(erro.erro || `Erro HTTP ${response.status}`);
  }

  return response.json();
}

const Index = () => {
  const [results, setResults] = useState<ValidationResult | null>(null);
  const [showHelp, setShowHelp] = useState(false);
  const { isDark, toggleTheme } = useTheme();

  // react-query mutation para o upload
  const mutation = useMutation({
    mutationFn: validarArquivo,
    onSuccess: (data) => {
      setResults(data);
    },
    onError: (error) => {
      console.error('Erro na validação:', error);
    },
  });

  const handleFileUpload = (file: File) => {
    mutation.mutate(file);
  };

  const handleReset = () => {
    setResults(null);
    mutation.reset();
  };

  return (
    <div className="min-h-screen bg-background flex flex-col">
      <Header 
        onHelpClick={() => setShowHelp(true)} 
        isDark={isDark}
        onToggleTheme={toggleTheme}
      />
      
      <main className="flex-1">
        {!results ? (
          <UploadArea 
            onUpload={handleFileUpload}
            loading={mutation.isPending}
            erro={mutation.error?.message || null}
          />
        ) : (
          <ValidationResults 
            data={results}
            onNewValidation={handleReset}
          />
        )}
      </main>
      
      <Footer />
      
      {showHelp && (
        <HelpModal onClose={() => setShowHelp(false)} />
      )}
    </div>
  );
};

export default Index;
