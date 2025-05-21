API key value store, supporting any json payload type, using dynamodb 


PUT http://localhost:8080/key1

`{
    "field": 1
}`

GET http://localhost:8080/key1

`{
    "field": 1
}`

PATCH http://localhost:8080/key1

`{
    "field": 2
}`

...

`{
    "field": 2
}`

PATCH http://localhost:8080/key1

`{
    "anotherField": []
}`

...

`{
    "field": 2
    "anotherField": []
}`

DELETE http://localhost:8080/key1

`{
}`