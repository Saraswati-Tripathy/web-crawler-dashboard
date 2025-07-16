export interface CrawlResult {
  id: number;
  url: string;
  title: string;
  htmlVersion: string;
  h1Count: number;
  h2Count: number;
  internalLinks: number;
  externalLinks: number;
  inaccessibleLinks: number;
  hasLoginForm: boolean;
  status: 'queued' | 'running' | 'done' | 'error';
  createdAt: string;
  userId: number;
}
