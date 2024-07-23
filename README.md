# Crowdfunding - Startup

Crowdfunding adalah sebuah service yang dibangun untuk membuat API aplikasi crowdfunding - startup.

## Daftar Isi

1. [Prasyarat](#prasyarat)
2. [Teknologi yang Digunakan](#teknologi-yang-digunakan)
3. [Fitur-fitur](#fitur---fitur)
4. [Pemasangan](#pemasangan)

## Prasyarat

- [GIT](https://www.git-scm.com/downloads)
- [MySQL 8.0](https://dev.mysql.com/downloads/installer/)
- [Go](https://go.dev/)

## Teknologi yang Digunakan

- Go
- Gin
- GORM
- Midtrans
- JWT

## Fitur - fitur

1. **Autentikasi Pengguna:**

- Register, login, dan check email.

2. **Profile:**

- Menampilkan data profile.
- Upload image.

3. **Manajemen Campaign:**

- Menampilkan, membuat, dan merubah data.
- Upload images.
- Menampilkan data transactions yang dimiliki campaign.

4. **Manajemen Transaction:**

- Menampilkan daftar transactions.
- Membuat transaction baru.

5. **Webhook Midtrans:**

- Handle notification dari midtrans.

6. **Role:**

- Menampilkan dan membuat role.

7. **Manajemen File:**

- Upload file.

## Pemasangan

Langkah-langkah untuk menginstall proyek ini.

Clone proyek

```bash
git clone https://github.com/DimasPondra/startup.git
```

Masuk ke dalam folder proyek

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

Dengan mengikuti langkah-langkah di atas, Anda akan dapat menjalankan service yang dibangun untuk membuat API aplikasi crowdfunding - startup.
