import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Cell } from 'recharts';
import type { ValidationResult } from '../types/validation';

interface ErrorChartProps {
  data: ValidationResult;
}

const ErrorChart = ({ data }: ErrorChartProps) => {
  const chartData = [
    { name: 'Vazias', value: data.errosVazias, color: 'hsl(48, 96%, 47%)' },
    { name: 'NCM', value: data.errosNCM, color: 'hsl(0, 84%, 60%)' },
    { name: 'CSOSN', value: data.errosCSOSN, color: 'hsl(224, 76%, 33%)' },
    { name: 'CST', value: data.errosCSTOrigem, color: 'hsl(25, 95%, 53%)' },
    { name: 'Tipo Item', value: data.errosTipoItem, color: 'hsl(215, 16%, 47%)' },
  ].filter(item => item.value > 0);

  if (chartData.length === 0) {
    return null;
  }

  return (
      <div className="bg-card rounded-lg border border-border p-4 sm:p-6">
        <div className="h-64">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart data={chartData} margin={{ top: 10, right: 10, left: 0, bottom: 0 }}>
              <CartesianGrid strokeDasharray="3 3" stroke="hsl(var(--border))" />
              <XAxis
                  dataKey="name"
                  tick={{ fill: 'hsl(var(--muted-foreground))', fontSize: 12 }}
                  axisLine={{ stroke: 'hsl(var(--border))' }}
              />
              <YAxis
                  tick={{ fill: 'hsl(var(--muted-foreground))', fontSize: 12 }}
                  axisLine={{ stroke: 'hsl(var(--border))' }}
              />
              <Tooltip
                  contentStyle={{
                    backgroundColor: 'hsl(var(--popover))',
                    border: '1px solid hsl(var(--border))',
                    borderRadius: '8px'
                  }}
                  labelStyle={{
                    color: 'hsl(var(--popover-foreground))'
                  }}
                  itemStyle={{
                    color: 'hsl(var(--popover-foreground))'
                  }}
                  formatter={(value: number) => [value, 'Erros']}
              />
              <Bar dataKey="value" radius={[4, 4, 0, 0]}>
                {chartData.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.color} />
                ))}
              </Bar>
            </BarChart>
          </ResponsiveContainer>
        </div>
      </div>
  );
};

export default ErrorChart;