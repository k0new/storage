# Go Redis-Like Key-Value Store
This project provides a Redis-like key-value store implemented in Go that supports *Set*, *Get*, and *Delete* operations along with *TTL* functionality.

## Features
 - **Key-Value Storage**: Utilizes a hash map to store key-value pairs for efficient retrieval.
 - **TTL Support**: Implements TTL using a priority queue to manage key expiration.
 - **Concurrency Safety**: Employs mutex locks to ensure thread safety during operations.
 - **Automatic Cleanup**: Removes expired keys from the store automatically.
## Approach
### Data Structures Used
**Hash Map**: Stores key-value pairs for quick access during Get and Set operations.

**TTL Heap**: Manages key expiration times using a priority queue (heap) of entries with their expiration times.
### Operations
***Set***: Adds a key-value pair to the hash map and updates the TTL heap with the key's expiration time.

***Get***: Retrieves the value associated with a key from the hash map. Handles key expiration based on TTL.

***Delete***: Removes a key from both the hash map and the TTL management structures.
### Concurrency
**Mutex Locks**: Ensures thread safety by employing mutex locks during operations that modify the data structures.
### TTL Management
**Background Process**: Implements a routine that periodically checks the TTL heap to remove expired keys, ensuring efficient cleanup.
## Usage
```go
package main

import (
	"time"
	
	"github.com/k0new/storage" // import storage interface 
)

func main() {
	s := storage.New() // Create a new storage object
	
	s.Set("key", "value", 20 * time.Minute) // Set the new value with expiration time
	
	val, err := s.Get("key") // Get value
	if err != nil { 
	    // ...handle error
	}
	
	s.Delete("key") // Delete value
}
```
## Testing

The project includes comprehensive unit tests to ensure the correctness of operations and functionality. Run tests using the following command:
`$ go test`

## CPU Profiling
To profile CPU usage during tests:

`$ go test -run TestKeyValueStore_Performance -cpuprofile=cpu_profile.prof`

Analyze the CPU profile:

`$ go tool pprof cpu_profile.prof`
