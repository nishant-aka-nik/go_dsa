channel theory
---
A channel is a communication mechanism that lets one goroutine send values to another goroutine

A channel has two principal operations, send and receive

A send statement transmits a value from one goroutine, through the channel, to another goroutine executing a corresponding receive expression.

ch <- x  // a send statement

x = <-ch // a receive expression in an assignment statement

### unbuffered channel
it does not have any capacity
A send operation on an unbuffered channel blocks the sending goroutine
until another goroutine executes a corresponding receive on the same channel, at which
point the value is transmitted and both goroutines may continue.


### buffered channel
A buffered channel has a queue of elements.
The queue’s maximum size is determined when it is created, by the
capacity argument to make.

A send operation on a buffered channel inserts an element at the
back of the queue, and a receive operation removes an element from the front.

If the channel is full, the send operation blocks its goroutine until
space is made available by another goroutine’s receive.

Conversely, if the channel is empty, a receive operation blocks until
a value is sent by another goroutine.

### Cancellation
Sometimes we need to instruct a goroutine to stop what it is doing
we can use select statement to create a cancellation pattern 

### interface
interface is a type in golang which defines a set of method signatures but does not provide implementations for those methods

Interfaces helps implement polymorphism 
Polymorphism means when a program behaves differently depending on the data it operates on
Multiple types can satisfy the same interface, enabling polymorphic behavior.

Interfaces decouple the definition of behaviors from their implementations, promoting modularity.
They make it easier to mock dependencies during testing.

---
Remember “P-D-T-D-S” (Polymorphism, Decoupling, Testability, Dependency Injection, Supporting Multiple Implementations) as key purposes.

---

Polymorphism example
```
type Speaker interface {
    Speak() string
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

func MakeSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    p := Person{Name: "Alice"}
    d := Dog{}
    MakeSpeak(p) // Outputs: Hello, my name is Alice
    MakeSpeak(d) // Outputs: Woof!
}
```

Decoupling and Abstraction
```
type Database interface {
    Connect() error
    Query(query string) (Results, error)
}

type MySQLDB struct{}
func (db MySQLDB) Connect() error { /* ... */ }
func (db MySQLDB) Query(query string) (Results, error) { /* ... */ }

type MongoDB struct{}
func (db MongoDB) Connect() error { /* ... */ }
func (db MongoDB) Query(query string) (Results, error) { /* ... */ }

func FetchData(db Database) {
    db.Connect()
    // Perform queries...
}
```

### REST vs gRPC
REST (Representational State Transfer)

	•	Definition: An architectural style for designing networked applications using standard HTTP protocols.
	•	Data Format: Typically uses JSON or XML.
	•	Communication: Stateless client-server communication.
	•	Usage: Widely used for public APIs and web services.

gRPC (gRPC Remote Procedure Calls)

	•	Definition: An open-source, high-performance RPC framework developed by Google, built on HTTP/2.
	•	Data Format: Uses Protocol Buffers (protobuf) by default.
	•	Communication: Supports both synchronous and asynchronous communication with streaming capabilities.
	•	Usage: Commonly used for internal microservices communication and real-time systems.


### 
https://medium.com/@code-geass/top-golang-interview-questions-related-to-goroutines-part-1-dff76c66b086
https://www.educative.io/blog/50-golang-interview-questions