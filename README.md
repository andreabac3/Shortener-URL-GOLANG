# Shortener URL in pure Golang and MongoDB

This code allows you to create a web server with the URL shortener function.
is developed purely in Go, uses official MongoDB drivers and Gorilla Mux for the routing.

## Getting Started

### Prerequisites
```
Golang
MongoDB
Setting up your .env file (path: src/.env)
```
#### MongoDB
Database Name = "shorturl" <br>
Collection Name = "shorturldata"
```
{ Hash: 'hashed string', Url: 'string'}
```

### Test

Use it in your favourite browser

#### Record URL
Use this link: http://127.0.0.1:8080/insert/www.yourWebSiteURL.com
#### Redirect from shortner URL
Use this link: http://127.0.0.1:8080/get/www.yourWebSiteURL.com

#### Convert from hashed URL to un-hashed URL 
Use this link: http://127.0.0.1:8080/get_no_redirect/www.yourWebSiteURL.com


## Program execution

Compile it with GO

# Authors

* **Andrea Bacciu**  - [github](https://github.com/andreabac3)

## License
[![](https://img.shields.io/npm/l/unique-names-generator.svg)](https://github.com/andreasonny83/unique-names-generator/blob/master/LICENSE)
This project is licensed under the MIT License - see the MIT license online for details
