# Todo API

To utilize the API, data will either need to be sent or received in a JSON format that's an array with objects that have fields for id, value, and, checked and some of those fields will be required depending on the API call.

To use this service:

1. Make sure no service is hosting on port 80
2. Git clone this project
3. In the project folder run 'docker-compose up'
4. Wait a minute or two till you see the MySQL message 'mysqld: ready for connections.'

## API Calls

* __Create Todos__ (createTodos) - will create new todo item entries, each entry will need a value field provided

        curl -X POST -d '[{"value": "do 20 pushups"}, {"value": "floss teeth"}]' http://localhost/createTodos

* __Get All Todos__ (getAllTodos) - will retrieve all todo items (value, checked status, id) whether they are checked or not

         curl -X GET http://localhost/getAllTodos

* __Check Off Todos__ (checkTodos) - will check off todo items based on id

        curl -X POST -d '[{"id": 2}, {"id": 3}]' http://localhost/checkTodos

* __Uncheck Todos__ (checkTodos) - will uncheck todo items based on id

        curl -X POST -d '[{"id": 3}]' http://localhost/uncheckTodos
