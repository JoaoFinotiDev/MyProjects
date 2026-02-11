import { motion } from 'framer-motion';
import { AlertCircle, FileWarning, Shield, Package, Hash, Tag } from 'lucide-react';
import type { ValidationResult } from '../types/validation';

interface ErrorSummaryCardsProps {
  data: ValidationResult;
}

const ErrorSummaryCards = ({ data }: ErrorSummaryCardsProps) => {
  const cards = [
    {
      label: 'Total de Erros',
      value: data.totalErros,
      icon: AlertCircle,
      colorClass: data.totalErros > 0 ? 'text-destructive' : 'text-success',
      bgClass: data.totalErros > 0 ? 'bg-destructive/10' : 'bg-success/10'
    },
    {
      label: 'Células Vazias',
      value: data.errosVazias,
      icon: FileWarning,
      colorClass: data.errosVazias > 0 ? 'text-warning' : 'text-muted-foreground',
      bgClass: data.errosVazias > 0 ? 'bg-warning/10' : 'bg-muted'
    },
    {
      label: 'Erros NCM',
      value: data.errosNCM,
      icon: Hash,
      colorClass: data.errosNCM > 0 ? 'text-destructive' : 'text-muted-foreground',
      bgClass: data.errosNCM > 0 ? 'bg-destructive/10' : 'bg-muted'
    },
    {
      label: 'Erros CSOSN',
      value: data.errosCSOSN,
      icon: Shield,
      colorClass: data.errosCSOSN > 0 ? 'text-destructive' : 'text-muted-foreground',
      bgClass: data.errosCSOSN > 0 ? 'bg-destructive/10' : 'bg-muted'
    },
    {
      label: 'Erros CST',
      value: data.errosCSTOrigem,
      icon: Package,
      colorClass: data.errosCSTOrigem > 0 ? 'text-destructive' : 'text-muted-foreground',
      bgClass: data.errosCSTOrigem > 0 ? 'bg-destructive/10' : 'bg-muted'
    },
    {
      label: 'Tipo Item',
      value: data.errosTipoItem,
      icon: Tag,
      colorClass: data.errosTipoItem > 0 ? 'text-destructive' : 'text-muted-foreground',
      bgClass: data.errosTipoItem > 0 ? 'bg-destructive/10' : 'bg-muted'
    },
  ];

  return (
    <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-3">
      {cards.map((card, idx) => (
        <motion.div
          key={card.label}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: idx * 0.05 }}
          className={`${card.bgClass} rounded-lg p-4 text-center border border-border`}
        >
          <card.icon className={`w-6 h-6 mx-auto mb-2 ${card.colorClass}`} />
          <div className={`text-2xl font-bold mb-1 ${card.colorClass}`}>
            {card.value}
          </div>
          <div className="text-xs text-muted-foreground font-medium">{card.label}</div>
        </motion.div>
      ))}
    </div>
  );
};

export default ErrorSummaryCards;
