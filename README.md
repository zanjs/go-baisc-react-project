# React + Go 

This repository is only for educational purposes. You will find a basic implementation of an application using [React](https://facebook.github.io/react/) and the [Go Programming Language](https://golang.org/)

If you are going to create an production application, you would need somewhere to store and retrieve data. However, that is beyond the scope of this repository, so I simply created a very crude (and not thread safe) mock database.

This project will not cover Authentication and other security factores, is only for demonstrate the basic communication of the Frontend with the Backend without so many complexity.



### Running and Building

Requires [Go](https://golang.org/doc/install) installed.

Requires [Node.js](https://nodejs.org/) v4+ to build the frontend.

Create the build folder with the optimized files running the command:
```sh
\YOUR-PATH\go-baisc-react-project\client
$ npm run build
```

Running the application:
```sh
\YOUR-PATH\go-baisc-react-project\server\main
$ go run main.go
```

The project will run in http://localhost:8000

![Go](http://nordicapis.com/wp-content/uploads/golang-hemmingway-with-a-martini-02-243x300.png)