# Mortgage project

Little tests around a simple exercise to try to set a better code organization with more flexibility.<br/>
This is in NO WAY intended to be used in production, it's just a playground to test some ideas and see how things happen in Go. <br/><br/>

## Some adds from the original exercise:
- General app config from json file (cf. internals/config/config.env.json)
- DB config from json files. They must be included under internals/db/ with the following names:
  -  db-dev.env.json
  -  db-prod.env.json

  The DB is automatically selected according to the selected mode in config.env.json => dev vs prod

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
- Use of `runtime.Caller()` to get filename and line where the error occured
- Use of enums like in C++ and Java but since go doesn't have this feature I've created the "enums" package to enforce a restricted choice for: 
  - Decision => `Decision.Approved` and `Decision.Restricted`
  - Mode &nbsp; &nbsp; &nbsp;=> `Mode.Dev`and `Mode.Prod`