# Crowdfunding - Startup
[Go](https://go.dev/) for Portfolio.

## Run locally

Clone the project

```bash
  git clone https://github.com/DimasPondra/startup.git
```

Go to the project directory

```bash
  cd startup
```

Create configuration file

```bash
  cp .env.example .env
```

Modify `.env` to Configure the following variables

- `DB_HOST` - The hostname or IP address of MySQL server.
- `DB_DATABASE` - The database created for the application.
- `DB_USERNAME` - Username to access the database.
- `DB_PASSWORD` - Password to access the database.

- `JWT_SECRET_KEY` - Key for JWT secret.
  
- `MIDTRANS_SERVER_KEY` - Key for Midtrans server.
- `MIDTRANS_ENVIRONMENT` - SANDBOX or PRODUCTION.

### Main Server

Start the server

```bash
  go run main.go
```

Using Air

[Air for live reload golang project](https://github.com/cosmtrek/air) 

```bash
  air
```
