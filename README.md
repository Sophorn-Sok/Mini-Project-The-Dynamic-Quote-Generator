GenZ Quote Generator
Author: Sophorn Sok

Description
A vibrant web application that generates random GenZ-style quotes. Built with Next.js for the frontend and Golang for the backend, with SQLite for persistent quote storage.

Setup Instructions
1. Clone the repository

git clone https://github.com/Sophorn-Sok/Mini-Project-The-Dynamic-Quote-Generator.git
cd Mini-Project-The-Dynamic-Quote-Generator

2. Install dependencies
Frontend (Next.js)

npm install
Backend (Golang)

go mod init genz-quote-backend
go get github.com/mattn/go-sqlite3


3. Run the backend

# Make sure you have a C compiler installed (TDM-GCC or MinGW-w64 for Windows)
$env:CGO_ENABLED=1
go run app/generate/main.go

4. Run the frontend

npm run dev

5. Open the app
Frontend: http://localhost:3000
Backend API: http://localhost:8080/api/quote

Architecture Explanation
Frontend (Next.js): Displays the UI and fetches quotes from the backend API.
Backend (Golang): Serves random quotes from a SQLite database via a REST API endpoint (/api/quote).
Database (SQLite): Stores all quotes persistently. The backend seeds the database with GenZ quotes if it's empty.
Communication: The frontend sends HTTP requests to the backend API to fetch quotes. The backend queries the database and returns a random quote as JSON.