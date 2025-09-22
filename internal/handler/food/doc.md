# food 

## Notice
All routes protected by token...
```
Authorization Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOiJzb21lIGlkIiwidG9rZW4iOiJqd3QgdG9rZW4iLCJleHAiOjE3NTg1NDk1NTEsImlhdCI6MTc1ODU0NTk1MX0.x4vgGw6Zz3W29OHdZWbz4lwKiSv2EfUQt9ErqcaFEc0
``` 

## GET    /food/list 


## POST   /food/rate 
```json
{
    "foods":{
        "1":2,
        "2":5,
        "8":10
    }
}
```

## POST   /food/autosave 
```json
{
    "autosave" : true
}
```
