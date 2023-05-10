
## Getting started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Presiquites
#### Golang
You need to have Go v1.19 installed on your machine. Follow the official installation guide to install Go. Or, follow managing installations guide to have multiple Go versions on your machine.

### MySQL
This service has dependency with MySQL. For development environment, you need to have a MySQL server running on your machine.

## installation

1. clone this repository :  

git clone https://github.com/alwisisva/twitter-app.git 

2. Build binaries using go.build :

go build

## Running

Execute binary to start the service :

cmd : online-learning.exe

## Note 

Run auto migration first to migrate table and set up RBAC

cmd : online-learning.exe -migrate up


## API Documentation

Use Postman API Documentation :

https://www.postman.com/solar-water-440274/workspace/test/collection/20757492-1e249602-d084-42f5-91b0-a02e727510a9?action=share&creator=20757492
