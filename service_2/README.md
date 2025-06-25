### 📘 Service 2 – Python Flask API

This is a simple Flask-based HTTP service. It includes health check and greeting endpoints, designed to run independently or behind a reverse proxy (like Nginx).

---

### 🔧 Endpoints

| Method | Path       | Description        | Response Example                          |
|--------|------------|--------------------|-------------------------------------------|
| GET    | `/ping`    | Health check       | `{"status": "ok", "service": "2"}`        |
| GET    | `/hello`   | Greeting message   | `{"message": "Hello from Service 2"}`     |
| GET    | `/health`  | Docker healthcheck | `OK`                                      |

---

### 🚀 Running Locally

#### Prerequisites

- Python 3.9+
- pip

#### Steps

```bash
cd service_2
pip install -r requirements.txt
python app.py
