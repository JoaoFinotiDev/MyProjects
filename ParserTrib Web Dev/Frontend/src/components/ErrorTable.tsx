import { useState, useRef, useEffect, useMemo } from 'react';
import { ChevronDown, ChevronUp, Filter } from 'lucide-react';
import { motion, AnimatePresence } from 'framer-motion';
import type { ValidationError, ErrorFilter } from '../types/validation';

interface ErrorTableProps {
  errors: ValidationError[];
}

const ErrorTable = ({ errors }: ErrorTableProps) => {
  const [itemsToShow, setItemsToShow] = useState(50);
  const [filter, setFilter] = useState<ErrorFilter>('all');
  const [showFilters, setShowFilters] = useState(false);
  const observerTarget = useRef<HTMLDivElement>(null);
  const tableRefs = useRef<{ [key: string]: HTMLTableRowElement | null }>({});

  const ITEMS_PER_LOAD = 50;

  const filterOptions: { value: ErrorFilter; label: string }[] = [
    { value: 'all', label: 'Todos' },
    { value: 'VAZIA', label: 'Células Vazias' },
    { value: 'NCM', label: 'NCM' },
    { value: 'CSOSN', label: 'CSOSN' },
    { value: 'CST_ORIGEM', label: 'CST Origem' },
    { value: 'TIPO_ITEM', label: 'Tipo Item' },
  ];

  const filteredErrors = filter === 'all'
      ? errors
      : errors.filter(e => e.tipo === filter);

  const displayedErrors = filteredErrors.slice(0, itemsToShow);
  const hasMore = itemsToShow < filteredErrors.length;

  // Extrair letras únicas das colunas
  const availableColumns = useMemo(() => {
    const columnsSet = new Set<string>();
    filteredErrors.forEach(error => {
      columnsSet.add(error.coluna.toUpperCase());
    });
    return Array.from(columnsSet).sort();
  }, [filteredErrors]);

  // Navegação para coluna específica
  const scrollToColumn = (coluna: string) => {
    const index = filteredErrors.findIndex(
        error => error.coluna.toUpperCase() === coluna
    );

    if (index === -1) return;

    // Garante que o item esteja renderizado
    if (index >= itemsToShow) {
      setItemsToShow(prev =>
          Math.min(
              Math.ceil((index + 1) / ITEMS_PER_LOAD) * ITEMS_PER_LOAD,
              filteredErrors.length
          )
      );
    }

    // Espera o DOM renderizar
    requestAnimationFrame(() => {
      const error = filteredErrors[index];
      const key = `${error.linha}-${error.coluna}`;
      const element = tableRefs.current[key];

      if (element) {
        element.scrollIntoView({ behavior: 'smooth', block: 'center' });
        element.style.backgroundColor = 'hsl(var(--accent) / 0.2)';
        setTimeout(() => {
          element.style.backgroundColor = '';
        }, 1500);
      }
    });
  };


  // Intersection Observer para scroll infinito
  useEffect(() => {
    const observer = new IntersectionObserver(
        (entries) => {
          if (entries[0].isIntersecting && hasMore) {
            setItemsToShow(prev => Math.min(prev + ITEMS_PER_LOAD, filteredErrors.length));
          }
        },
        { threshold: 0.1 }
    );

    if (observerTarget.current) {
      observer.observe(observerTarget.current);
    }

    return () => observer.disconnect();
  }, [hasMore, filteredErrors.length]);

  // Reset ao mudar filtro
  useEffect(() => {
    setItemsToShow(50);
  }, [filter]);

  const getTypeColor = (tipo: ValidationError['tipo']) => {
    switch (tipo) {
      case 'VAZIA':
        return 'bg-warning/20 text-warning-foreground border-warning/30';
      case 'NCM':
        return 'bg-destructive/20 text-destructive border-destructive/30';
      case 'CSOSN':
        return 'bg-primary/20 text-primary border-primary/30';
      case 'CST_ORIGEM':
        return 'bg-accent/20 text-accent-foreground border-accent/30';
      case 'TIPO_ITEM':
        return 'bg-muted text-muted-foreground border-border';
      default:
        return 'bg-muted text-muted-foreground border-border';
    }
  };

  const getTypeBadge = (tipo: ValidationError['tipo']) => {
    const labels: Record<ValidationError['tipo'], string> = {
      VAZIA: 'Vazia',
      NCM: 'NCM',
      CSOSN: 'CSOSN',
      CST_ORIGEM: 'CST',
      TIPO_ITEM: 'Tipo Item',
    };
    return labels[tipo] || tipo;
  };

  if (errors.length === 0) {
    return (
        <div className="bg-success/10 border border-success/20 rounded-lg p-8 text-center">
          <p className="text-success font-medium text-lg">Nenhum erro encontrado!</p>
          <p className="text-muted-foreground mt-2">Sua planilha está pronta para envio.</p>
        </div>
    );
  }

  return (
      <div className="bg-card rounded-lg border border-border overflow-hidden">
        <div className="p-4 border-b border-border flex flex-col sm:flex-row sm:items-center justify-between gap-3">
        <span className="text-sm text-muted-foreground">
          Mostrando {displayedErrors.length} de {filteredErrors.length} erros
          {filter !== 'all' && ` (filtrado: ${filterOptions.find(f => f.value === filter)?.label})`}
        </span>

          <div className="relative">
            <button
                onClick={() => setShowFilters(!showFilters)}
                className="flex items-center gap-2 px-4 py-2 bg-secondary hover:bg-secondary/80 rounded-lg text-sm font-medium transition-colors"
            >
              <Filter className="w-4 h-4" />
              Filtrar por tipo
              {showFilters ? <ChevronUp className="w-4 h-4" /> : <ChevronDown className="w-4 h-4" />}
            </button>

            <AnimatePresence>
              {showFilters && (
                  <motion.div
                      initial={{ opacity: 0, y: -10 }}
                      animate={{ opacity: 1, y: 0 }}
                      exit={{ opacity: 0, y: -10 }}
                      className="absolute right-0 top-full mt-2 bg-popover border border-border rounded-lg shadow-lg z-20 py-2 min-w-[160px] max-h-[300px] overflow-y-auto"
                  >
                    {filterOptions.map((option) => (
                        <button
                            key={option.value}
                            onClick={() => {
                              setFilter(option.value);
                              setShowFilters(false);
                            }}
                            className={`w-full text-left px-4 py-2 text-sm hover:bg-muted transition-colors whitespace-nowrap ${
                                filter === option.value ? 'bg-primary/10 text-primary font-medium' : 'text-foreground'
                            }`}
                        >
                          {option.label}
                        </button>
                    ))}
                  </motion.div>
              )}
            </AnimatePresence>
          </div>
        </div>

        {filteredErrors.length === 0 ? (
            <div className="p-8 text-center min-h-[230px] flex flex-col items-center justify-center">
              <p className="text-muted-foreground">
                Nenhum erro do tipo <strong>{filterOptions.find(f => f.value === filter)?.label}</strong> encontrado.
              </p>
              <p className="text-sm text-muted-foreground mt-2">
                Tente selecionar outro filtro.
              </p>
            </div>
        ) : (
            <>
              <div className="overflow-x-auto max-h-[600px] overflow-y-auto">
                <table className="w-full">
                  <thead className="bg-muted sticky top-0 z-10">
                  <tr>
                    <th className="px-4 py-3 text-left text-xs font-semibold text-muted-foreground uppercase tracking-wider">
                      Linha
                    </th>
                    <th className="px-4 py-3 text-left text-xs font-semibold text-muted-foreground uppercase tracking-wider">
                      Coluna
                    </th>
                    <th className="px-4 py-3 text-left text-xs font-semibold text-muted-foreground uppercase tracking-wider">
                      Campo
                    </th>
                    <th className="px-4 py-3 text-left text-xs font-semibold text-muted-foreground uppercase tracking-wider">
                      Tipo
                    </th>
                    <th className="px-4 py-3 text-left text-xs font-semibold text-muted-foreground uppercase tracking-wider">
                      Problema
                    </th>
                  </tr>
                  </thead>
                  <tbody className="divide-y divide-border">
                  {displayedErrors.map((error, idx) => {
                    const key = `${error.linha}-${error.coluna}`;
                    return (
                        <motion.tr
                            key={`${key}-${idx}`}
                            ref={(el) => { tableRefs.current[key] = el; }}
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={idx < 20 ? { delay: idx * 0.01 } : {}}
                            className="hover:bg-muted/50 transition-colors"
                        >
                          <td className="px-4 py-3 text-sm font-mono text-foreground">{error.linha}</td>
                          <td className="px-4 py-3 text-sm font-mono text-foreground">{error.coluna}</td>
                          <td className="px-4 py-3 text-sm text-foreground">{error.nomeColuna}</td>
                          <td className="px-4 py-3">
                      <span className={`inline-flex px-2 py-1 text-xs font-medium rounded border ${getTypeColor(error.tipo)}`}>
                        {getTypeBadge(error.tipo)}
                      </span>
                          </td>
                          <td className="px-4 py-3 text-sm text-muted-foreground">{error.mensagem}</td>
                        </motion.tr>
                    );
                  })}
                  </tbody>
                </table>

                {hasMore && (
                    <div ref={observerTarget} className="p-4 text-center text-sm text-muted-foreground">
                      Carregando mais resultados...
                    </div>
                )}

                {!hasMore && filteredErrors.length > 50 && (
                    <div className="p-4 text-center text-sm text-muted-foreground">
                      Todos os {filteredErrors.length} erros foram carregados
                    </div>
                )}
              </div>

              {/* Navegação alfabética por colunas */}
              {availableColumns.length > 0 && (
                  <div className="p-4 border-t border-border bg-muted/30">
                    <div className="flex items-center gap-3 flex-wrap">
                    <span className="text-sm font-medium text-muted-foreground">
                      Navegar por coluna:
                    </span>
                      <div className="flex gap-1 flex-wrap">
                        {availableColumns.map((letter) => (
                            <button
                                key={letter}
                                onClick={() => scrollToColumn(letter)}
                                className="w-8 h-8 rounded-md bg-primary/10 hover:bg-primary hover:text-primary-foreground text-primary font-semibold text-sm transition-colors"
                            >
                              {letter}
                            </button>
                        ))}
                      </div>
                    </div>
                  </div>
              )}
            </>
        )}
      </div>
  );
};

export default ErrorTable;