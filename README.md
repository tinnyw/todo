# Todo API

To utilize the API, data will either need to be sent or received in a JSON format that's an array with objects that have fields for id, value, and, checked and some of those fields will be required depending on the API call.

* _Create API Call_ (createTodos) - will create new todo item entries, each entry will need a value field provided

        curl -X POST -d '[{"value": "do 20 pushups"}, {"value": "floss teeth"}]' http://localhost/createTodos
