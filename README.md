#First start
##REST-API fuctions:
___
###GET
1. **GetAccounts**
Return all accounts. Return Header with state and json ```{{"login":"", "password":""}, ....}```
*Example*
```
curl http://localhost:8081/accounts  
```

###POST
1. **CreateAccount**
Create account. Accepts login and password. Return Header with state and json ```"successfull":bool[true/false]```
*json struct for accept*
```
{"login":""
"password":""}
```
*Example*
```curl -X POST http://localhost:8081/CreateAccount \                    ✔ 
-H "Content-Type: application/json" \
-d '{"login":"Rulsan","password":"123hfdds$"}'
```

2. **EnterAccount**
Logining in account if it is exist. Retun Header with statte and json ```"successfull":bool[true/false]```
*json struct for accept*
```
{"login":""
"password":""}
```
*Example*
```curl -X POST http://localhost:8081/EnterAccount \                    ✔ 
-H "Content-Type: application/json" \
-d '{"login":"Rulsan","password":"123hfdds$"}'
```

___
