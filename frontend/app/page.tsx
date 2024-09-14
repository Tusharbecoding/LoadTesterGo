"use client";

import { useState } from "react";
import MetricsTable from "./components/MetricsTable";

interface Metrics {
  totalRequests: number;
  successCount: number;
  failureCount: number;
  avgResponseTime: string;
}

export default function Home() {
  const [targetUrl, setTargetUrl] = useState("");
  const [metrics, setMetrics] = useState<Metrics | null>(null);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setMetrics(null);

    const response = await fetch("/api/load-test", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ target_url: targetUrl }),
    });

    const data = await response.json();
    setMetrics(data);
    setLoading(false);
  };

  return (
    <div>
      <h1 className="text-2xl font-bold mt-8">
        Distributed Load Tester Dashboard
      </h1>

      <form onSubmit={handleSubmit} className="mt-4">
        <input
          type="text"
          value={targetUrl}
          onChange={(e) => setTargetUrl(e.target.value)}
          placeholder="Enter URL to test"
          className="border border-gray-300 px-4 py-2 rounded"
        />
        <button
          type="submit"
          className="bg-blue-500 text-white px-4 py-2 rounded ml-2"
        >
          Run Load Test
        </button>
      </form>

      {loading && <p>Running load test...</p>}

      {metrics && <MetricsTable metrics={metrics} />}
    </div>
  );
}
