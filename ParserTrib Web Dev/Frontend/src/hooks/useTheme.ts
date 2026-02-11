import { useState, useEffect } from 'react';

export function useTheme() {
  // Inicializa o estado com base no localStorage ou preferência do sistema
  const [isDark, setIsDark] = useState(() => {
    if (typeof window === 'undefined') return false;

    const stored = localStorage.getItem('theme');
    if (stored) {
      return stored === 'dark';
    }

    return window.matchMedia('(prefers-color-scheme: dark)').matches;
  });

  // Aplica o tema ao documento
  useEffect(() => {
    const root = window.document.documentElement;

    if (isDark) {
      root.classList.add('dark');
      localStorage.setItem('theme', 'dark');
    } else {
      root.classList.remove('dark');
      localStorage.setItem('theme', 'light');
    }
  }, [isDark]);

  const toggleTheme = () => {
    setIsDark(prev => !prev);
  };

  return { isDark, toggleTheme };
}