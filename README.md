# Mortgage project

Some tests around a simple exercise from a course on Udemy (Go Programming by Example - Kam Hojati) to try to set a better code organization with more flexibility.<br/>
This is in NO WAY intended to be used in production, it's just a playground to test some ideas and see how things work in Go. <br/><br/>

## Some adds from the original exercise:
- General app config from json file (cf. internals/config/config.env.json)
- DB config from json files. They must be included under internals/db/ with the following names:
  -  db-dev.env.json
  -  db-prod.env.json

  The DB is automatically set according to the selected mode in config.env.json => dev vs prod

  expected format:
  ```json
  {
  "engine": "mysql",
  "server": "127.0.0.1",
  "port": "3306",
  "user": "user_name",
  "password": "super_pwd",
  "database": "db_name"
  }
  ```
- Make use of the DAO pattern to easily switch the db type (MySQL and PostgreSQL only)
- Different error messages according to the selected mode in config.env.json => dev vs prod
- Use of `runtime.Caller()` to get filename and line where the error occured<br/>
  Format of the errors:
  ```
  2022/05/12 02:00:02  Query problem: 
        SELECTd first_name, last_name, credit_score, salary, downpayment, house_id 
        FROM customers
        WHERE id=$1
         pq: syntax error at or near "SELECTd" 
         [/Users/jonathan/informatique/languages/go/udemy/go-programming-by-example/mortgage-project/customer/daoPostgres.go:29]
  ```
- Use of enums like in C++ and Java but since go doesn't have this feature I've created the "enums" package to enforce a restricted choice for: 
  - Decision => `Decision.Approved` and `Decision.Restricted`
  - Mode &nbsp; &nbsp; &nbsp;=> `Mode.Dev` and `Mode.Prod`
- The following architecture for house and customer: <br/>
  &nbsp; &nbsp; &nbsp; &nbsp; **controller => service => repository => dao**<br/> 
  allows us to have more flexibility. Modifying the program to control it via an HTTP server would be pretty straightforward, we'd just have to add the server and a specialized controller to call the expected services and nothing else.<br/>
  If we need to add a new DB we just have to add a new DAO implementation and in internals/db/ we add a case in getConnection() with the corresponding function