# An Old Assignment Task for ByFood

## Environment setup

You need to have [Go](https://golang.org/),
[Node.js](https://nodejs.org/),
[Docker](https://www.docker.com/), and
[Docker Compose](https://docs.docker.com/compose/)
(comes pre-installed with Docker on Mac and Windows)
installed on your computer.

Verify the tools by running the following commands:

```sh
go version
npm --version
docker --version
docker-compose --version
```

## Start in development mode

In the project directory run the command (you might
need to prepend it with `sudo` depending on your setup):
```sh
docker-compose -f docker-compose-dev.yml up
```

This starts a local Mongo database on `localhost:27017`.
The database will be populated with test records from
the [init-db.sql](init-db.sql) file.

Nitpick: For unit tests you need to start local Mongo databe with these commands:


```sh
docker pull mongo 
docker run -d --name mongodbtest -p 27017:27017 mongo
```
Navigate to the `server` folder and start the back end:

```sh
cd server
go run server.go
```
The back end will serve on http://localhost:8080.

Navigate to the `webapp` folder, install dependencies,
and start the front end development server by running:

```sh
cd webapp
npm install
npm start
```
The application will be available on http://localhost:3000.
 
## Start in production mode

Perform:
```sh
docker-compose up
```
This will build the application and start it together with
its database. Access the application on http://localhost:8080.

If an application using 8080 or 27017 ports, I recommend to remove docker containers and images in your docker engine.
If you get an problem clean docker then use this commands:
```sh
docker-compose build
docker-compose up
```

