# Links App

Credit for the user API and the GoLang application goes to [Mohit Dubey](https://medium.com/@mohitdubey_83162).

The application is an extension of [this example](https://medium.com/@mohitdubey_83162/mastering-mongodb-integration-with-go-a-step-by-step-developers-guide-3fcbe46a906e).

## Users

Create a new user:

```
curl --location 'http://localhost:8080/create_user' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Moh",
    "age": 2,
    "gender": false
}'
```

Get all users:

```
curl --location 'http://localhost:8080/users'
```

Get user with ID `<ID>`:

```
curl --location 'http://localhost:8080/users/<ID>'
```

Delete user with ID `<ID>`:

```
curl --location 'http://localhost:8080/users/delete/<ID>'
```

## Links

Create a new link:

```
curl --location 'http://localhost:8080/create_link' \
--header 'Content-Type: application/json' \
--data '{
    "link": "https://medium.com/@mohitdubey_83162"
}'
```

Get all links:

```
curl --location 'http://localhost:8080/links'
```

Get link with ID `<ID>`:

```
curl --location 'http://localhost:8080/links/<ID>'
```

Delete link with ID `<ID>`:

```
curl --location 'http://localhost:8080/links/delete/<ID>'
```
