# Shortener URL in pure Golang and MongoDB

This code allows you to create a web server with the URL shortener function. <br>
It's developed purely in Go, uses official MongoDB drivers and Gorilla Mux for the routing.

## Getting Started

## Prerequisites
```
Golang
MongoDB
Setting up your .env file (path: src/.env)
```
### MongoDB
Database Name = "shorturl" <br>
Collection Name = "shorturldata"
```
{ Url: 'string', Hash: 'hashed string' }
An example
{ Url: 'www.andreabacciu.com', Hash: 'bd443137fc0662f1e10f7f71ac948766' }
```

### Test

Use it in your favourite browser

#### Record URL
Use this link: http://127.0.0.1:8080/insert/www.yourWebSiteURL.com
#### Redirect from shortner URL
Use this link: http://127.0.0.1:8080/get/bd443137fc0662f1e10f7f71ac948766
#### Convert from hashed URL to un-hashed URL 
Use this link: http://127.0.0.1:8080/get_no_redirect/bd443137fc0662f1e10f7f71ac948766


## Program execution

Compile it with GO

# Known Issues
This is a toy project.
* .env file needs a real path. (Path src/lib/mongodbConnection.go Function: Init() )
* Currently, communications with the DB are not secure.

# Authors

* **Andrea Bacciu**  - [github](https://github.com/andreabac3)

## License
[![](https://img.shields.io/npm/l/unique-names-generator.svg)](https://github.com/andreasonny83/unique-names-generator/blob/master/LICENSE) <br>
This project is licensed under the MIT License - see the MIT license online for details
