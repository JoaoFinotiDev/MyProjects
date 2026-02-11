import { HelpCircle, Moon, Sun, FileSpreadsheet } from 'lucide-react';
import { motion } from 'framer-motion';

interface HeaderProps {
  onHelpClick: () => void;
  isDark: boolean;
  onToggleTheme: () => void;
}

const Header = ({ onHelpClick, isDark, onToggleTheme }: HeaderProps) => {
  return (
      <header className="bg-primary text-primary-foreground h-16 px-4 sm:px-6 flex items-center justify-between shadow-lg sticky top-0 z-40">
        <div className="flex items-center gap-3">
          {/* Logo */}
          <div className="w-64 h-64 flex-shrink-0 flex items-center justify-center -my-16 -mt-20">
            <img
                src="/logo.png"
                alt="Logo"
                className="w-full h-full object-contain"
                onError={(e) => {
                  e.currentTarget.style.display = 'none';
                  e.currentTarget.nextElementSibling?.classList.remove('hidden');
                }}
            />
            <FileSpreadsheet className="w-48 h-48 hidden" />
          </div>

          <div className="flex items-center gap-2">
            <FileSpreadsheet className="w-7 h-7" />
            <h1 className="text-lg sm:text-xl font-bold tracking-tight">
              Validador de Planilhas
            </h1>
          </div>
        </div>

        <div className="flex items-center gap-2">
          <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              onClick={onToggleTheme}
              className="p-2 rounded-lg bg-primary-foreground/10 hover:bg-primary-foreground/20 transition-colors"
              aria-label={isDark ? 'Ativar modo claro' : 'Ativar modo escuro'}
          >
            {isDark ? <Sun className="w-5 h-5" /> : <Moon className="w-5 h-5" />}
          </motion.button>

          <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              onClick={onHelpClick}
              className="flex items-center gap-2 px-3 sm:px-4 py-2 bg-primary-foreground/10 hover:bg-primary-foreground/20 rounded-lg transition-colors"
          >
            <HelpCircle className="w-5 h-5" />
            <span className="hidden sm:inline font-medium">Ajuda</span>
          </motion.button>
        </div>
      </header>
  );
};

export default Header;