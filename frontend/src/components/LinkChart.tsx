import { PieChart, Pie, Cell, Tooltip } from "recharts";

interface Props {
  internal: number;
  external: number;
}

export default function LinkChart({ internal, external }: Props) {
  const data = [
    { name: "Internal", value: internal },
    { name: "External", value: external },
  ];
  const COLORS = ["#3182CE", "#E53E3E"];

  return (
    <PieChart width={300} height={200}>
      <Pie data={data} dataKey="value" outerRadius={80} label>
        {data.map((_, i) => (
          <Cell key={`cell-${i}`} fill={COLORS[i % COLORS.length]} />
        ))}
      </Pie>
      <Tooltip />
    </PieChart>
  );
}
