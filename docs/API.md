# Chirpy API Documentation

## Overview

This document details the RESTful API endpoints provided by the Chirpy backend server. It's intended for developers
interacting with the API or understanding its capabilities.

**Base URL:** All API endpoints described below are prefixed with `/api`. If running locally on the default port 8080,
the full base URL is `http://localhost:8080/api`.

## Authentication

Most endpoints require authentication using a **JSON Web Token (JWT)** provided as a Bearer token in the `Authorization`
header.

1. **Login (`POST /api/login`):** Send email and password to receive an `access token` and a `refresh token`.
2. **Authenticated Requests:** For all endpoints marked as "Authentication: Required", include the obtained access token
   in the request header:
   ```
   Authorization: Bearer <your_access_token>
   ```
3. **Token Refresh (`POST /api/refresh`):** When the access token expires (indicated by a `401 Unauthorized`
   response), call this endpoint. A successful response provides a new access token.
4. **Logout (`POST /api/revoke`):** Call this endpoint to invalidate the current refresh token (effectively
   logging out the session).

## Common Error Responses

The API uses standard HTTP status codes. Error responses typically include a JSON body with an error message:

```json
{
  "error": "A descriptive error message here"
}
```

Common status codes include:

* **`400 Bad Request`**: Invalid request format, missing required fields, validation errors (e.g., invalid date format,
  empty title/content).
* **`401 Unauthorized`**: Missing, invalid, or expired JWT access token (for protected routes), or invalid
  credentials/refresh token for auth routes.
* **`403 Forbidden`**: Authenticated user lacks permission for the requested action (e.g., non-admin trying admin tasks,
  user trying to modify another user's chirp).
* **`404 Not Found`**: The requested resource (e.g., a specific chirp ID) does not exist.
* **`500 Internal Server Error`**: An unexpected error occurred on the server (database issue, unhandled code error).

---

## Endpoints

### Authentication

---

#### `POST /api/login`

Authenticates a user and provides access/refresh tokens.

* **Authentication:** None
* **Request Body:**

```json
{
  "email": "user@example.com",
  "password": "user_password"
}
```

* **Success Response (`200 OK`):**
    * Sets refresh token to DB.
    * Body:

```json
{
  "id": "uuid-string-user-id",
  "created_at": "account_creation_timestamp",
  "updated_at": "account_update_timestamp",
  "email": "user@example.com",
  "is_chirpy_red": "chirpy_red_status",
  "token": "your_access_token_jwt_string",
  "refresh_token": "refresh_token_string"
}
```

* **Errors:** 401, 500

---

#### `POST /api/refresh`

Issues a new access token using a valid refresh token.

* **Authentication:** A refresh token in the headers, in the same `Authorization: Bearer <token>` format
* **Request Body:** None
* **Success Response (`200 OK`):**
    * Body:

```json
{
  "token": "new_access_token_jwt_string"
}
```

* **Errors:** 400, 401

---

#### `POST /api/revoke`

Revokes the refresh token associated with the request's token.

* **Authentication:** A refresh token in the headers, in the same `Authorization: Bearer <token>` format
* **Request Body:** None
* **Success Response (`204 No Content`):** No response body.
* **Errors:** 400, 500

---

### Users

---

#### `POST /api/users`

Creates a new user.

* **Authentication:** None
* **Request Body:**

```json
{
  "email": "user@example.com",
  "password": "user_password"
}
```

* **Success Response (`201 Created`):**
    * Body: Returns the created user object (excluding password).

```json
{
  "id": "uuid-string-user-id",
  "created_at": "account_creation_timestamp",
  "updated_at": "account_update_timestamp",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

* **Errors:** 500

---

#### `PUT /api/users`

Updates the currently authenticated user's email/password.

* **Authentication:** Required (Any authenticated user).
* **Request Body:**

```json
{
  "email": "new_user@example.com",
  "password": "new_user_password"
}
```

* **Success Response (`200 OK`):**
    * Body:

```json
{
  "id": "uuid-string-user-id",
  "created_at": "account_creation_timestamp",
  "updated_at": "account_update_timestamp",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

* **Errors:** 401, 500

---

### Chirps

---

#### `GET /api/chirps`

Retrieves a list of all chirps.

* **Authentication:** None
* **Query Parameters:**
    * `author_id` (UUID): Return this author's chirps only (optional).
    * `sort` (`asc` or `desc`): How to sort the returned chirps (default: `asc`).
* **Request Body:** None
* **Success Response (`200 OK`):**
    * Body: Returns a JSON array of `Chirp` objects.

```json
[
  {
    "id": "32ec8ab5-3109-402b-b0a9-3fcc7bed6987",
    "created_at": "2025-04-16T10:05:30.181161Z",
    "updated_at": "2025-04-16T10:05:30.181161Z",
    "user_id": "565f26f2-b8fb-45e6-b01b-064e0ad23e5e",
    "body": "I'm the one who knocks!"
  },
  {
    "id": "4ca5d7ce-5ed1-4177-b690-f36800836716",
    "created_at": "2025-04-16T10:05:30.183239Z",
    "updated_at": "2025-04-16T10:05:30.183239Z",
    "user_id": "565f26f2-b8fb-45e6-b01b-064e0ad23e5e",
    "body": "Gale!"
  },
  {
    "id": "ddae2739-f463-4284-98f7-e7a880be6625",
    "created_at": "2025-04-16T10:05:30.185212Z",
    "updated_at": "2025-04-16T10:05:30.185212Z",
    "user_id": "565f26f2-b8fb-45e6-b01b-064e0ad23e5e",
    "body": "Cmon Pinkman"
  },
  {
    "id": "fe0106fc-9747-4cac-9a1e-bbfbf357cd5d",
    "created_at": "2025-04-16T10:05:30.186844Z",
    "updated_at": "2025-04-16T10:05:30.186844Z",
    "user_id": "565f26f2-b8fb-45e6-b01b-064e0ad23e5e",
    "body": "Darn that fly, I just wanna cook"
  }
]
```

* **Errors:** 400, 500

---

#### `GET /api/chirp/{chirpID}`

Return a single chirp by its ID.

* **Authentication:** None
* **Request Body:** None
* **Path Parameters:**
    * `{chirpID}` (UUID): The ID of the chirp.
* **Success Response (`200 OK`):**
    * Body: Returns the requested `Chirp` object.
* **Errors:** 400, 404

---

#### `POST /api/chirps`

Creates a new chirp and associates specified targets.

* **Authentication:** Required
* **Constraints:**
    * `body` must be between 1 and 140 characters.
    * Any bad words in the `body` will be censored with `****`.
* **Request Body:**

```json
{
  "body": "Test"
}
```

* **Success Response (`201 Created`):**
    * Body: Returns the a `Chirp` object with the requested body.

```json
{
  "id": "806dcbcf-4a28-49c8-8204-fbebfffec0a3",
  "created_at": "2025-04-16T11:05:30.660638Z",
  "updated_at": "2025-04-16T11:05:30.660638Z",
  "user_id": "0f703a63-7548-49d4-a538-9fa5e270fac8",
  "body": "Test"
}
```

* **Errors:** 400, 401, 500

---

#### `DELETE /api/chirps/{chirpID}`

Deletes a specific chirp.

* **Authentication:** Required.
* **Path Parameters:**
    * `{chirpID}` (UUID): The ID of the chirp.
* **Request Body:** None
* **Success Response (`204 No Content`):** No response body.
* **Errors:** 400, 401, 403, 404, 500

---

### Webhooks

---

#### `POST /api/polka/webhooks`

Webhook endpoint for receiving events from the Polka service.

This will be used to upgrade a user's account to Chirpy
Red

* **Authentication:** Required (API Key)
* **Constraints:**
    * `event` must be `user.upgraded` or request will be ignored.
* **Request Body:**

```json
{
  "data": {
    "user_id": "${userID}"
  },
  "event": "user.upgraded"
}
```

* **Success Response (`204 No Content`):** No response body.
* **Errors:** 401, 404, 500

---

### Server Health

---

#### `GET /api/healthz`

Checks the health of the server.

* **Authentication:** None
* **Request Body:** None
* **Success Response (`200 OK`):** No response body.
* **Errors:** None

---

### Admin

---

#### `GET /admin/metrics`

Returns server metrics.

* **Authentication:** None
* **Request Body:** None
* **Success Response (`200 OK`):**
    * Body: Returns an HTML page with server metrics.

```html

<html>
<body>
<h1>Welcome, Chirpy Admin</h1>
<p>Chirpy has been visited 0 times!</p>
</body>
</html>
```

* **Errors:** None

---

#### `POST /admin/reset`

Resets the server metrics to its initial state.

* **Authentication:** Needs to be in `dev` environment
* **Request Body:** None
* **Success Response (`200 OK`):**
    * Body: Returns a plain text message.

```text
Hits reset to 0 and database reset to initial state.
```

* **Errors:** 403