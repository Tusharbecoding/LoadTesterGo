interface Metrics {
  totalRequests: number;
  successCount: number;
  failureCount: number;
  avgResponseTime: string;
}

interface MetricsTableProps {
  metrics: Metrics;
}

export default function MetricsTable({ metrics }: MetricsTableProps) {
  return (
    <table className="table-auto w-full bg-white shadow-md rounded my-4">
      <thead>
        <tr className="bg-blue-500 text-white">
          <th className="px-4 py-2">Total Requests</th>
          <th className="px-4 py-2">Success</th>
          <th className="px-4 py-2">Failures</th>
          <th className="px-4 py-2">Average Response Time (ms)</th>
        </tr>
      </thead>
      <tbody>
        <tr className="text-center">
          <td className="border px-4 py-2">{metrics.totalRequests}</td>
          <td className="border px-4 py-2">{metrics.successCount}</td>
          <td className="border px-4 py-2">{metrics.failureCount}</td>
          <td className="border px-4 py-2">{metrics.avgResponseTime}</td>
        </tr>
      </tbody>
    </table>
  );
}
