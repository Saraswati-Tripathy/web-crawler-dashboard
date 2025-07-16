import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { fetchResultById } from "../utils/api";
import type { CrawlResult } from "../types";
import LinkChart from "../components/LinkChart";

export default function ResultDetails() {
  const { id } = useParams();
  const [result, setResult] = useState<CrawlResult | null>(null);

  useEffect(() => {
    fetchResultById(Number(id)).then((res) => setResult(res.data));
  }, [id]);

  if (!result) return <p>Loading...</p>;

  return (
    <div className="p-4">
      <h2 className="text-xl font-bold mb-2">{result.title}</h2>
      <LinkChart internal={result.internalLinks} external={result.externalLinks} />
      <p className="mt-4 text-sm text-gray-600">Status: {result.status}</p>
    </div>
  );
}
