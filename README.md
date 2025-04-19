# 📧 OopsMail - Temporary Email Service

OopsMail is a temporary email service built with Go. It allows users to create temporary email addresses and receive emails at those addresses. The service automatically deletes mailboxes and their contents after a specified period.

## ✨ Features

- 📨 Create temporary email addresses
- 📥 Receive emails at temporary addresses
- 🔍 View received emails via web interface or REST API
- 🧹 Automatic cleanup of expired mailboxes
- 🗄️ Redis-based storage for scalability
- 🌐 Modern web interface

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
- 🌐 Web interface on port 8080

## 🌐 Web Interface

Visit http://localhost:8080 in your browser to access the web interface.

### Home Page
- Overview of the service features
- Quick access to create new mailboxes
- Information about automatic cleanup

### Create Mailbox
- Generate a new temporary email address
- Copy email address to clipboard
- View expiration time

### View Emails
- Enter mailbox ID or full email address
- View all received emails
- See email details including subject, sender, and content

## 🔌 API Endpoints

### 📝 Create Mailbox
```
POST /api/mailbox
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
GET /api/mailbox/:id
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
   - Visit http://localhost:8080/create
   - Click "Generate Temporary Email"
   - Copy the generated email address

2. Send an email to the generated address

3. View your emails:
   - Visit http://localhost:8080/view
   - Enter your mailbox ID or full email address
   - View all received emails

## 🛠️ Development

The project is structured as follows:
- `cmd/` - Main application entry point
- `internal/` - Core application code
  - `api/` - HTTP API server
  - `mailbox/` - Mailbox service
  - `smtp/` - SMTP server
  - `storage/` - Storage interface and implementations
  - `web/` - Web interface
    - `templates/` - HTML templates
    - `static/` - CSS and JavaScript files

## 📄 License

MIT 