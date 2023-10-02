# go-error-handling

This repository presents an approach to handle errors effectively in Go.
The project follows a layered structure, but the approach can actually be applied for other styles. 

Go provides a way to wrap errors (using `fmt.Errorf`) in layers which we often combine with a sentinel error to add extra info on top an existing error.
Usually a sentinel error is constructed with `errors.New`, which returns a simple string error.
```golang
package mypackage

func myFunc() error {
    if err := someFunc(); err != nil {
        return fmt.Errorf("someFunc encounters an error: %w", err)
    }
	
	return nil
}

func someFunc() error {
	return errSentinel
}

var errSentinel = errors.New("sentinel error")
```

While this can be useful to a certain extent, it lacks the ability to communicate with greater details of the error up the stack.
The caller, in this case, does not know if it's ok to share the error message to the outer world or not, because the error message could expose the internal of the application (which database the application is using for example), or indicate a low level error that the client does not need to care about (TCP connection closed for instance).
Furthermore, the caller has no idea in which context the error happens.

It's desirable to design our error handling process that achieves these goals:
- **Wrappable**: the error must be able to be wrapped to create layers of nested error, which can be compatible to `fmt.Errorf`.
- **Communicative**: the error must be able to tell the caller what would be a proper message to share with the outer world, ensuring no internal information is leaked.
- **Contextual**: the error must be able to have context data attached to it. This will be very helpful for our error logging system.