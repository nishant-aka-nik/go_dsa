mutex is locking and unlocking mechanism

mutex is used to protect the data resources from concurrent access

a mutex guards the shared resource

go routine who has put the mutex lock on variable should be the one to unlock it


•	Blocking Behavior: When a goroutine tries to acquire a lock on a mutex that is already held by another goroutine, it will wait (block) until the mutex becomes available.
•	Deadlocks: If not managed properly, this can lead to deadlocks, where two or more goroutines are waiting on each other to release locks.
•	No Timeout: The basic sync.Mutex in Go does not support timeouts for acquiring a lock. Once a goroutine calls Lock(), it will wait indefinitely until the lock is available.