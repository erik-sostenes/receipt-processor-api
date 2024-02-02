# Receipt Processor - Challenge

REST API for the receipt processor

###### Docs

- [API documentation (Swagger)](swagger.yaml)

### how to execute the project?

###### recomendations (optional)

Verify that you do not have any docker image running

```bash
docker stop $(docker ps -q)
```

###### Compile

Follow the instructions below to compile

```bash
# build the image of GO
make build
#up the api stack
make start
```

### how run test?

Follow the instructions below to run the tests

```bash
make test
```

Note: this is another option on how to run the tests

```bash
# Load required environment variables
export $(grep -v ^# .env.example)
# Run tests
go test -v ./...
```

###### Unit tests

#### Structure

```go
package receipt

import (
	"strconv"
	"testing"
)

func TestX(t *testing.T) {
	// Table Driven Testing
	tdt := map[string]struct{}{
	    {},	// Test Case
    	}{}

	// Setup
	x := struct{}{}

	// Cleanup function (Optional)
	t.Cleanup(func() {})

	// Subtests
	for name, v := range tdt {
		t.Run(name, func(t *testing.T){
			t.Log(x, v)
			// Cleanup function for each subtest (Optional)
			t.Cleanup(func() {})
		})
	}
}
```

### Hexagonal Architecture

The hexagonal architecture is based on three principles and techniques:
1.Explicitly separate User-Side, Business Logic, and Server-Side
2.Dependencies are going from User-Side and Server-Side to the Business Logic
3.We isolate the boundaries by using Ports and Adapters

```
├───cmd                         👉🏼 (execute commands)
│   └───bootstrap               👉🏼 (bootstrap package that builds the program with its full set of components)
├───internal
│   ├───backoffice              👉🏼 (core business)
│   │   ├───module              👉🏼 (represents a boundary)
│   │   │   ├───business        👉🏼 (business logic layer)
│   │   │   │   ├───domain      👉🏼 (data transfer objects, business objects, errors, entities, value object)
│   │   │   │   ├───ports       👉🏼 (business contracts)
│   │   │   │   └───services    👉🏼 (logic)
│   │   │   └───infrastructure  👉🏼 (layer infrastructure)
│   │   │       ├───driven      👉🏼 (output adapters)
│   │   │       └───drives      👉 (input adapters)
```

### User-Side

- This is the side through which the user or external programs will interact with the application.

### Business Logic

- This is the part that we want to isolate from both left and right sides. It contains all the code that concerns and implements business logic.

### Server-Side

- This is where we’ll find what your application needs, what it drives to work. It contains essential infrastructure details such as the code that interacts with your database, makes calls to the file system, or code that handles HTTP calls to other applications on which you depend for example.

```

|------------|                   |----------------|             |-------------|
| User side  | =====[port]=====> | Business logic | <==[port]== | server side |
|------------|                   |----------------|             |-------------|

```
