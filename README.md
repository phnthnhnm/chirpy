# Chirpy

An HTTP server that mimics the functionality of a Twitter-like social media platform. It allows users to register,
create posts, and follow other users. The server is built using Go and uses PostgreSQL as the database.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You
can then install `chirpy` with:

```bash
go install github.com/phnthnhnm/chirpy@latest
```

## Config

Create a `.env` file in the project's root directory with the following structure:

```ini
DB_URL="postgres://username:@localhost:5432/database?sslmode=disable"
PLATFORM="dev"
JWT_SECRET="256bitsecret"
POLKA_KEY="polkaapikey"
```

- `DB_URL` - The URL to your Postgres database. Make sure to replace `username`, `localhost`, and `database` with your
  own values.
- `PLATFORM` - The platform you're using. `dev` is required for the `reset` API
- `JWT_SECRET` - A secret key used to sign JWT tokens. This should be a long, random string.
- `POLKA_KEY` - A key used to authenticate with the Polka API. This is used for upgrading users to Chirpy Red.

## Usage

To run the server, use the following command:

```bash
chirpy
```

This will start the server on `localhost:8080`.

For details on the API documentation, you can look through [here](docs/API.md). The API is RESTful and follows
standard conventions.

## How to Contribute

1. Fork the repository.
2. Create a new branch for your feature or bug fix:
   ```bash
   git checkout -b feature-name
   ```
3. Make your changes and commit them:
   ```bash
   git commit -m "Description of changes"
   ```
4. Push your changes to your fork:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request on the main repository.