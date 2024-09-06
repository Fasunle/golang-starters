# Golang Stater Kit

This is a starter kit for golang projects. The aim of this project is to provide a starting point for golang projects. This increase the development speed as it reduces the setup time.

It should be noted that this project is not a framework. It is a starter kit which utilize [Fiber](https://docs.gofiber.io/)  as a router and middleware. It setup http router by default on the `fiber` branch.

For installations:

```bash
    git clone -b fiber https://github.com/Fasunle/golang-starters.git && rm -rf .git

```

## Folder Structure

`cmd/api/main.go` - This is the entry point of the application. It is responsible for starting the server and listening for requests.

`cmd/handlers/` - This is the folder where all the route handlers are stored.

`cmd/routes/` - This is the folder where all the routes are stored.

`cmd/api` - folder is specifically reserved for business logic.

## Installing dependencies

After cloning the project, install it dependencies by  running the following command:

```bash
    go mod tidy
```

## Running the application on your local machine

Run the following command on your terminal:

```bash
    go run ./cmd/api
```

You application should be open on port http://localhost:80

Check the route:

http://localhost:80/api/v1