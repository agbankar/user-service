## User-service

This is sample go rest apis database, main purpose is how to build rest apis with go
and architect the project to create server serving apis and handle validations and error.
This sample also demonstrates how are junit tests are done in golang .

##### System Requirement:

* Go should be installed on your local machine
* pull the mysql docker image and run the mysql container

 ````
docker run --name goapis-mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql
````

Run the below DDL

  ```
  CREATE TABLE users (
  id bigint(20) NOT NULL,
  last_name varchar(255) ,
  first_name varchar(255),
  address varchar(255),
  city varchar(255),
  PRIMARY KEY (`id`)
) ;

CREATE TABLE orders (
  id bigint(20) NOT NULL,
  order_id int(20),
  PRIMARY KEY (`id`)
) ;

  ```

#### How to run the app

Clone this project git@github.com:agbankar/user-service.git on your local at any location
c:/users/{yourid}/ and cd c:/users/{yourid}/user-service . Run g**o build ./...** build should be ready.
once build is ready run **go run ./cmd/main.go**
server will be up and running on your local at http://localhost:8080/

#### APIs

#### 1.CreateUser --> success

* Request

```bash
curl --location --request POST 'http://localhost:8080/CreateUser' \
--header 'Content-Type: application/json' \
--data-raw '{
"Name": "Ashish",
"Id": 1000000
}'
```

* Response

```json
{
  "Name": "Ashish",
  "Id": 1000000
}
```

#### 2.CreateUser --> Validations Error

* Request

```bash
curl --location 'http://localhost:8080/CreateUser' \
--header 'Content-Type: application/json' \
--data '{
"id": {{$randomBankAccount}},
"firstName":"Ashish",
"lastName":"Bankar",
"address":"AMS",
"city":"AMS"

}'
```

* Response

```json
[
  {
    "errorCode": "CODE102",
    "errorDescription": "Id  is required. Numeric only Max 10 characters."
  }
]
```

#### 3.GetBonus --> Success

* Request

```bash
curl --location --request GET 'http://localhost:8080/getBonus/1'
```

* Response

```json
{
  "userID": "1",
  "bonus": 100
}
```

#### Next Steps

Add loggers


