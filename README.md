# Chirpy: Boot.dev Tutorial Golang Server

Welcome to the docs for Chirpy, a guided project from Boot.dev.
This project is sort of like that bird social media platform, but only broadly.
It is an instructional Go server using Chi, JWTs, bcrypt, and godotenv.

## Imported Technologies

Here are the external packages used to create this server, apart from the Go stdlib.

- chi -> github.com/go-chi/chi/v5
  For creating the server.
- JWT -> github.com/golang-jwt/jwt/v5
  For creating and parsing JWTs.
- bcyrpt -> golang.org/x/crypto/bcrypt
  For encrypting and unencrypting passwords.
- godotenv -> github.com/joho/godotenv
  For retrieving environmental variables.

## Getting Started

This is just an instructional project and isn't meant to be cloned or deployed. But if it were, here's how:

1. Clone the repository:
   ```bash
   git clone https://github.com/skovranek/chirpy.git
   ```

2. Install the dependencies:
   ```bash
   go mod download
   ```

3. Configure the server by creating a ".env" file, including these not-so-secret keys from the tutorial:
   ```bash
   POLKA_API_KEY=f271c81ff7084ee5b99a5091b42d486e
   JWT_SECRET=CKpJVLHqOoYtKX/hjkQ6iPtVhqeqmAKYF4uPfqGoQxTVVe8ZMbedqRcjUrhlkiy1keNbSQq3Cn9RnZ2xTKM8GA==
   ```

4. Build and run the server with the debug mode flag, or not. Debug mode deletes the old database.json on start up for testing.
   ```bash
   go build -o out and ./out --debug
   go build -o out and ./out
   ```

## Configure

Here is an example '.env' file with the available environmental variables you can define. At minimum, you must include the provided fictional secret keys:

   ```bash
    ROOT=.
    PORT=8080
    DB_PATH=database.json
    POLKA_API_KEY=f271c81ff7084ee5b99a5091b42d486e
    JWT_SECRET=CKpJVLHqOoYtKX/hjkQ6iPtVhqeqmAKYF4uPfqGoQxTVVe8ZMbedqRcjUrhlkiy1keNbSQq3Cn9RnZ2xTKM8GA==
    ACCESS_JWT_EXP_IN_HOURS=1
    REFRESH_JWT_EXP_IN_HOURS=1440
   ```

## APP

The server includes a basic file server at path '/app', which serves 'index.html' and '/assets/logo.png'.

## ADMIN & METRICS

The server includes an '/admin/metrics' endpoint which responds with the number of hits on the file server path '/app'.

## API

The server includes an API which exposes the following endpoints:

- `GET /api/healthz`: check server readiness.
- `POST  /api/users`: create user.
- `POST /api/login`: login user and create JWTs.
- `POST /api/refresh`: refresh access JWT.
- `POST /api/revoke`: revoke refresh JWT.
- `PUT /api/users`: update user.
- `POST /api/chirps`: create chirp.
- `GET /api/chirps`: retrieve a filtered and sorted list of chirps.
- `GET /api/chirps/{chirp_id}`: retrieve a chirp.
- `DELETE /api/chirps/{chirp_id}`: delete a chirp.
- `POST /api/polka/webhooks`: webhook from a fictional third-party site.

Detailed information about each endpoint, including request and response examples:

1. Check server readiness: `GET /api/healthz`
   
   **Example Request:**

   ```bash
   GET /api/healthz
   ```
   **Example Response:**

   **Header:**
   >Status: 201

   ```
   OK
   ```

2. Create user: `POST /api/users`

   **Example Request:** 

   ```bash
   POST /api/users
   Content-Type: application/json

   {
     "password": "123456",
     "email":    "example@email.com"
   }
   ```

   **Example Response:** 

   **Header:**
   >Status: 201

   ```json
   {
     "id":    1,
     "email": "example@email.com"
   }
   ```

3. Login: `POST /api/login`

   **Example Request:**

   ```bash
   POST /api/login
   Content-Type: application/json

   {
     "password": "123456",
     "email":    "example@email.com"
   }
   ```

   **Example Response:**

   **Header:**
   >Status: 200

   ```json
   {
     "id":            1,
     "email":         "example@email.com",
     "is_chirpy_red": false,
     "token":         "your_valid_ACCESS_token_here",
     "refresh_token": "your_valid_REFRESH_token_here"
   }
   ```

4. Refresh access JWT: `POST /api/refresh`

   **Example Request:**

   **Header:**
   >Authorization: Bearer your_valid_REFRESH_token_here

   ```bash
   POST /api/refresh
   ```

   **Example Response:**

   **Header:**
   >Status: 200

   ```json
   {
     "token": "your_valid_ACCESS_token_here"
   }
   ```

5. Revoke refresh JWT: `METHOD /api/revoke`

   **Example Request:**

   **Header:**
   >Authorization: Bearer your_valid_REFRESH_token_here

   ```bash
   POST /api/revoke
   ```

   **Example Response:**

   **Header:**
   >Status: 200

   ```json
   {}
   ```

6. Update user: `PUT /api/users`

   **Example Request:**

   **Header:**
   >Authorization: Bearer your_valid_ACCESS_token_here

   ```bash
   PUT /api/users
   Content-Type: application/json

   {
     "password": "123456",
     "email":    "example@email.com"
   }
   ```

   **Example Response:**

   **Header:**
   >Status: 200

   ```json
   {
     "id":    1,
     "email": "example@email.com"
   }
   ```

7. Create chirp: `POST /api/chirps`

   **Example Request:**

   **Header:**
   >Authorization: Bearer your_valid_ACCESS_token_here

   ```bash
   POST /api/chirps
   Content-Type: application/json

   {
     "body": "messsage"
   }
   ```

   **Example Response:**

   **Header:**
   >Status: 201

   ```json
   {
     "id":    1,
     "author_id": 1,
     "body": "message"
   }
   ```

8. Retrieve a filtered and sorted list of chirps: `GET /api/chirps`

   **Example Request:**

   **Header:**
   >Authorization: Bearer your_valid_ACCESS_token_here

   **Optional Query Parameters:**

   Get chirps by author:
   >author_id={id}

   Sort chirps by ascending (default) or descending order:
   >sort={asc or desc}

   ```bash
   
   GET /api/chirps
   GET /api/chirps?sort=asc
   GET /api/chirps?sort=desc
   GET /api/chirps?sort=asc&author_id=1

   ```

   **Example Response:**

   **Header:**
   >Status: 200

   ```json
   [
     {
       "id": 1,
       "author_id": 1,
       "body": "message 1"
     },
     {
       "id": 2,
       "author_id": 1,
       "body": "message 2"
     }
   ]
   ```
9. Retrieve a chirp: `GET /api/chirps/{chirp_id}`

   **Example Request:**

   **Required Query Parameters:**

   Get chirp by id:
   >/api/chirps/{chirp_id}

   ```bash
   GET /api/chirps/1
   ```

   **Example Response:**

   **Header:**
   >Status: 200

   ```json
   {
     "id": 1,
     "author_id": 1,
     "body": "message"
   }
   ```

10. Delete a chirp: `DELETE /api/chirps/{chirp_id}`

    **Example Request:**

    **Header:**
    >Authorization: Bearer your_valid_ACCESS_token_here

    **Required Query Parameters:**

    Delete chirp by id:
    >/api/chirps/{chirp_id}

    ```bash
    DELETE /api/chirps/1
    ```

    **Example Response:**

    **Header:**
    >Status: 200

    ```json
    {}
    ```

11. Webhook from a fictional third-party named "polka": `POST /api/polka/webhooks`

    **Example Request:**

    **Header:** *(use this fictional secret key)*
    >Authorization: ApiKey f271c81ff7084ee5b99a5091b42d486e

    ```bash
    POST /api/polka/webhooks
    Content-Type: application/json

    {
      "event": "user.upgraded",
      "data":  {
        "user_id": 1
      }
    }
    ```

    **Example Response:**

    **Header:**
    >Status: 200

    ```json
    {}
    ```

## Error Handling

If an error occurs while processing a request, an appropriate error message and status code will be returned in the response. Following the instructions for the appropriate endpoints above: First, create a user account. Then login. Then you may "chirp" (post chirps messages).

## Testing

Testing was done on <Boot.dev> through the browser with the tutorial interface. The encoded tests in this project are minimal. They do not and are not intended to validate required behavior, automate QA or anything else. To run the tests regardless, here's how:

```shell
go test -v
```

## Contribute

"No project is ever finished, but it can be done." --Gene Kranz

This project is done and is no longer maintained. To create your own version, go to <Boot.dev> and complete the course "Learn Web Servers". 

## License

Do not deploy this server. It was created for instructional purposes once upon a time. To create your own version, go to <Boot.dev> and complete the course "Learn Web Servers".

## Contact

Please do not contact me in regards to the maintenance of this project. Thank you for your interest.