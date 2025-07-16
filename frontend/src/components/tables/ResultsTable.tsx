import React, { useEffect, useState } from "react";
import {
  useReactTable,
  getCoreRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  flexRender,
} from "@tanstack/react-table";
import type { ColumnDef, SortingState } from "@tanstack/react-table";
import type { CrawlResult } from "../../types";
import { fetchResults } from "../../utils/api";

export default function ResultsTable() {
  const [data, setData] = useState<CrawlResult[]>([]);
  const [sorting, setSorting] = useState<SortingState>([]); // <---- Explicit type here

  const columns: ColumnDef<CrawlResult>[] = [
    { accessorKey: "title", header: "Title" },
    { accessorKey: "htmlVersion", header: "HTML Version" },
    { accessorKey: "internalLinks", header: "Internal Links" },
    { accessorKey: "externalLinks", header: "External Links" },
    { accessorKey: "status", header: "Status" },
  ];

  const table = useReactTable({
    data,
    columns,
    state: {
      sorting,
    },
    onSortingChange: setSorting,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
  });

  useEffect(() => {
    fetchResults().then((res: any) => setData(res.data));
  }, []);

  return (
    <table className="w-full text-sm border">
      <thead>
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <th key={header.id} className="p-2 border-b cursor-pointer" onClick={header.column.getToggleSortingHandler()}>
                {flexRender(header.column.columnDef.header, header.getContext())}
                {{
                  asc: " 🔼",
                  desc: " 🔽",
                }[header.column.getIsSorted() as string] ?? null}
              </th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody>
        {table.getRowModel().rows.map((row) => (
          <tr key={row.id} className="hover:bg-gray-100">
            {row.getVisibleCells().map((cell) => (
              <td key={cell.id} className="p-2 border-b">
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}
