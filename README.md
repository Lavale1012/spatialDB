# 🧠 Semantic Vector Database in Go

![build](https://img.shields.io/badge/build-passing-brightgreen)
![license](https://img.shields.io/badge/license-MIT-blue)
![postgres](https://img.shields.io/badge/database-PostgreSQL-blue)
![language](https://img.shields.io/badge/language-Go-informational)

> 🔍 Real-time semantic search. Fast, smart, persistent.

---

## 🚀 Live Demo (GIF)

<img src="https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExbmdsMXhxOGE2aHFtcTdqbHNod3c5cDVtaTlmN2Fya3E1cXo5bnVsaiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/MF9ItaT1f9bZziwRf9/giphy.gif" width="600"/>

> _(Replace this GIF with one of your terminal, Postman, or browser hitting the API)_

---

## ✨ Features

- 🔠 Accepts raw text and stores vector embeddings
- 🤝 Embeds text using Python + sentence-transformers
- 🗃️ Stores vectors in PostgreSQL via GORM
- 📐 Searches using cosine similarity
- 🧱 Modular and designed to be hackable

---

## 📦 Tech Stack

### 🛠 Backend (Go)

- ⚡️ Gin
- 🧬 GORM + PostgreSQL
- 📏 Cosine similarity search

### 🤖 Embedding Service (Python)

- 🚀 FastAPI
- 🧠 SentenceTransformers (`all-MiniLM-L6-v2`)

---

## 🗂 Project Structure

```
vector-db/
├── go-server/
│   ├── api/          # HTTP routes
│   ├── core/         # Search engine logic
│   ├── models/       # GORM data models
│   └── cmd/          # Entrypoint
├── python-embedder/
│   └── app/          # FastAPI embedding service
```

---

## 📥 Setup Instructions

### 1️⃣ Start Python Embedding API

```bash
cd python-embedder
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
uvicorn app.main:app --reload --port 8000
```

### 2️⃣ Set up PostgreSQL

> Use Supabase, Docker, or local Postgres.

`.env` example:

```
POSTGRES_DSN="host=localhost user=lavale password=secret dbname=vector port=5432 sslmode=disable"
```

### 3️⃣ Start the Go Server

```bash
cd go-server
go run cmd/main.go
```

---

## 🔎 Example Requests

### ➕ Insert a Vector

```bash
curl -X POST http://localhost:8080/api/embed \
  -H "Content-Type: application/json" \
  -d '{"id": "vec1", "text": "Golang is amazing!"}'
```

### 🔍 Search for Similar

```bash
curl -X POST http://localhost:8080/api/search \
  -H "Content-Type: application/json" \
  -d '{"text": "I love programming", "k": 3}'
```

---

## 📈 Future Features

- 📌 Metadata filtering (e.g. "category = tech")
- ⚡ HNSW for faster search
- 🌍 REST + gRPC support
- 🔒 API auth and rate limiting
- 🚢 Deployment via Docker + Fly.io

---

## 🧑‍💻 Author

Built with 💡 by **Lavale Butterfield**  
[GitHub Profile → @Lavale1012](https://github.com/Lavale1012)

---

## 🧪 Want to Help?

Star the repo ⭐ | Fork it 🔀 | Build your own feature 🚀
# spatialDB
