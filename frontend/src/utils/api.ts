import axios from "axios";
const token = localStorage.getItem('token');
const API = axios.create({
  baseURL: "http://localhost:8082/api",
  headers: {
    Authorization: `Bearer ${token}`, // Make sure token is valid and present
  },
  withCredentials: true,
});

export const fetchResults = () => API.get("/results");
export const crawlURL = (url: string) => API.post("/crawl", { url });
export const deleteURLs = (ids: number[]) => API.post("/delete", { ids });
export const reanalyzeURLs = (ids: number[]) => API.post("/reprocess", { ids });
export const fetchResultById = (id: number) => API.get(`/results/${id}`);

export default API;
