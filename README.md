# Common services for the [REST service](https://github.com/rus-sharafiev/go-rest)

The module includes the following packages:
- db
  > PostgreSQL instance, instance creator and custom methods (based on the [pgx driver](https://github.com/jackc/pgx))
- exception
  > a set of functions that write JSON encoded errors to `http.ResponseWriter`
- formdata
  > Middleware that saves files and "restore" JSON (requires [custom FormData converter](https://github.com/rus-sharafiev/fetch-api) on frontend)
- jwt
  > Generates and validates JWT (based on [jwt-go](https://github.com/golang-jwt/jwt))
- mail
  > Some basic functions to send a mail using `net/smtp`
- spa
  > `http.Handler` for serving SPA and static files
- uploads
  > `http.Handler` for serving uploads protected by access token.
