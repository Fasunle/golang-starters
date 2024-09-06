# Golang Stater Kit

This is a starter kit for golang projects. The aim of this project is to provide a starting point for golang projects. This increase the development speed as it reduces the setup time.

It should be noted that this project is not a framework. It is a starter kit which utilize go-chi as a router and middleware. It setup http router by default on th e main branch but other variants can be access in the branches.

For default setup, use the main branch as follows:

```bash
    git clone https://github.com/Fasunle/golang-starter-kit.git && rm -rf .git

```

For installations, other than the default, you just need to specify the target branch other than main:

```bash
    git clone -b TARGET-BRANCH https://github.com/Fasunle/golang-starter-kit.git && rm -rf .git

```

## Folder Structure

`cmd/api/main.go` - This is the entry point of the application. It is responsible for starting the server and listening for requests.

`cmd/handlers/` - This is the folder where all the route handlers are stored.

`cmd/routes/` - This is the folder where all the routes are stored.

`cmd/api` - folder is specifically reserved for business logic.
