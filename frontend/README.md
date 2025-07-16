# Web Crawler Project

A web crawling and analysis application with user authentication, URL management, and crawl results dashboard.

---

## Features

- User registration and login
- Add URLs for crawling and analysis
- Asynchronous crawling with progress status
- View crawl results in a sortable, paginated table
- Detailed views with link statistics and broken links
- Bulk actions to re-run or delete crawls
- CORS support for frontend-backend communication

---

## Tech Stack

- **Backend:** Go (Gin framework), GORM, MySQL/PostgreSQL
- **Frontend:** React, TypeScript, TanStack React Table, Axios
- **API:** REST endpoints with JWT-based authentication

---

## Prerequisites

- Go 1.20+
- Node.js 18+
- MySQL or PostgreSQL database
- `npm` or `yarn`

---

## Setup & Run Backend

1. Clone repository and navigate to backend:

   ```bash
   cd backend

Create .env file with database credentials and JWT secret, for example:

Install Go dependencies:

go mod download
Run database migrations (if applicable):


go run cmd/migrate/main.go
Start backend server:

go run main.go
Backend will listen on http://localhost:8081

Setup & Run Frontend
Navigate to frontend directory:

cd frontend
Install dependencies:

npm install
# or
yarn
Start development server:


npm run dev
# or
yarn dev
Frontend will run on http://localhost:5173

