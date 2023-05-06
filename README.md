# Jester

Store and reuse your CLI tricks.

## API Documentation

When running locally, you can find API documentation at:

```
http://localhost:8080/swagger/index.html
```

## Configuring Jester backend

For log level, set this environment variable (underlying library is `logrus`):

```
LOG_LEVEL=info
```

For backend API you must use `postgres` database, and these are the environment variables you need to set:

```
DB_HOST=127.0.0.1
DB_DRIVER=postgres
DB_USER=user
DB_PASSWORD=password
DB_NAME=jester
DB_PORT=5432
DB_POSTGRES_SSLMODE=disable
```

For JWT authentication you need to set token duration, and API secret. Token duration is expressed in hours:

```
TOKEN_DURATION=1
API_SECRET=c4636aea-8b96-473b-acb0-38913870be13
```

Port and default listen address for the server are set with these variables:

```
LISTEN_ADDRESS=127.0.0.1
PORT=8080
```
