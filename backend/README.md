To start server run 
bash$ go build ./...
bash$ ./todolist_server

When sending update task request you should specify
next boolean params. Each parameter determines if the field 
has been changed. It is necessary becaus when parameter
is not specified it has default value as empty string.
`json:"update_name"`
`json:"update_description"`
`json:"update_is_open"`