# 📧 OopsMail - Temporary Email Service

OopsMail is a temporary email service built with Go. It allows users to create temporary email addresses and receive emails at those addresses. The service automatically deletes mailboxes and their contents after a specified period.

## ✨ Features

- 📨 Create temporary email addresses
- 📥 Receive emails at temporary addresses
- 🔍 View received emails via REST API
- 🧹 Automatic cleanup of expired mailboxes
- 🗄️ Redis-based storage for scalability

## 📋 Prerequisites

- 🐹 Go 1.21 or later
- 🔴 Redis server (for email storage)

## 🚀 Installation

1. Clone the repository:
```bash
git clone https://github.com/kohaku1907/oopsmail.git
cd oopsmail
```

2. Install dependencies:
```bash
go mod download
```

3. Start Redis server:
```bash
# Using Docker
docker run -d -p 6379:6379 redis
```

## ⚙️ Running the Service

1. Start the service:
```bash
go run cmd/main.go
```

The service will start:
- 📧 SMTP server on port 1025
- 🌐 HTTP API server on port 8080

## 🔌 API Endpoints

### 📝 Create Mailbox
```
POST /mailbox
```
Response:
```json
{
    "id": "random_id",
    "email": "random_id@oopsmail.com",
    "expires_at": "2024-03-20T12:00:00Z"
}
```

### 📬 Get Emails
```
GET /mailbox/:id
```
Response:
```json
[
    {
        "id": "email_id",
        "from": "sender@example.com",
        "to": "random_id@oopsmail.com",
        "subject": "Email Subject",
        "body": "Email content",
        "created_at": "2024-03-20T11:00:00Z"
    }
]
```

## 📖 Usage Example

1. Create a temporary mailbox:
```bash
curl -X POST http://localhost:8080/mailbox
```

2. Send an email to the generated address (e.g., random_id@oopsmail.com)

3. Retrieve emails:
```bash
curl http://localhost:8080/mailbox/random_id
```

## 🛠️ Development

The project is structured as follows:
- `cmd/` - Main application entry point
- `internal/` - Core application code
  - `api/` - HTTP API server
  - `mailbox/` - Mailbox service
  - `smtp/` - SMTP server
  - `storage/` - Storage interface and implementations

## 📄 License

MIT 