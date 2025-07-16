import { useState } from "react";
import { crawlURL } from "../utils/api";

export default function AddURLForm() {
  const [url, setUrl] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!url) return;
    await crawlURL(url);
    setUrl("");
  };

  return (
    <form onSubmit={handleSubmit} className="flex gap-2 p-2">
      <input
        type="url"
        className="border p-2 rounded w-full"
        placeholder="Enter URL"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded">
        Analyze
      </button>
    </form>
  );
}
