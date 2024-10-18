# Links App

## Backend APIs

Credit for the user API and the GoLang application goes to [Mohit Dubey](https://medium.com/@mohitdubey_83162).

The application is an extension of [this example](https://medium.com/@mohitdubey_83162/mastering-mongodb-integration-with-go-a-step-by-step-developers-guide-3fcbe46a906e).

The GitHub repo is available [here](https://github.com/mohitdb7/full-stack-development/tree/main/mongo-db).

### Run

To run the backend APIs, open a terminal in the root directory of the project.
Change directory into `links-be` with:

```bash
cd links-be
```

and then run:

```
go run main.go
```

### Users

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

### Links

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

## Frontend

The frontend uses [Alpine.js](https://alpinejs.dev/) and was firstly put together using the help of [ChatGPT](https://chatgpt.com/).

To easily serve the frontend, run a Python HTTP server with:

```bash
python -m http.server 3000
```

You'll get the frontend available at http://localhost:3000/. Don't forget to run the APIs in parallel!

Clearly, you can serve the `index.html` file using any web server you like...
