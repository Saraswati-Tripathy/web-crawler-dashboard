import AddURLForm from "../components/AddURLForm";
import ResultsTable from "../components/tables/ResultsTable";

export default function Dashboard() {
  return (
    <div className="p-4 max-w-6xl mx-auto">
      <h1 className="text-2xl font-semibold mb-4">Web Crawler Dashboard</h1>
      <AddURLForm />
      <ResultsTable />
    </div>
  );
}
