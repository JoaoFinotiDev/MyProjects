import { Mail } from 'lucide-react';

const Footer = () => {
  const currentYear = new Date().getFullYear();
  const currentMonth = new Date().toLocaleString('pt-BR', { month: '2-digit' });
  
  return (
    <footer className="bg-card border-t border-border py-4 px-6">
      <div className="max-w-6xl mx-auto flex flex-col sm:flex-row items-center justify-between gap-2 text-sm text-muted-foreground">
        <div className="flex items-center gap-2">
          <Mail className="w-4 h-4" />
          <span>Dúvidas? Entre em contato: </span>
          <a 
            href="mailto:imp@alternativasistemas.com.br"
            className="text-primary hover:underline font-medium"
          >
            imp@alternativasistemas.com.br
          </a>
        </div>
        <div>
          <span>Versão 1.2 • Última atualização: 04/02/2026</span>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
