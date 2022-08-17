package panic

import (
	"log"
	"runtime/debug"
	"time"
)

/*
The return value of recover is nil if any of the following conditions holds:
- panic's argument was nil;
- the goroutine is not panicking;
- recover was not called directly by a deferred function.
*/
func recoverByAnotherGoRoutine() {
	// Recover outside a goroutine
	defer func() {
		log.Println(recover())
		log.Println("recover from panic!")
	}()

	go func() {
		// Panic occurs in a goroutine
		panic("A bad boy stole a server")
	}()
}

func recoverByAnotherFunction() {
	go func() {
		// Call Recover() Tool Function.
		defer func() {
			Recover("Panic in goroutine")
			// Test the panic result.
			log.Println("recover from panic!")
		}()
		// Panic here.
		panic("A bad boy stole the server")
	}()
	time.Sleep(time.Second)

	go func() {
		// Call Recover() Tool Function.
		defer Recover("Panic in goroutine")
		// Panic here.
		panic("A bad boy stole the server again")
	}()
	time.Sleep(time.Second)
}

func Recover(funcName string) {
	if err := recover(); err != nil {
		// If the panic is caught.
		log.Printf("recover from panic! panic para: %v, panic info: %v\n", funcName, err)
	}
}

func recoverByUnnamedFunc() {
	// Call the recover() in a deferred function.
	go func() {
		defer func() { // f1
			recover() // Recover successfully because the caller function f1 is a deferred function
			log.Println("recover from a panic by f1")
		}()
		panic("A bad boy stole a server")
	}()
	time.Sleep(time.Second)

	// Call the recover() directly.
	go func() { // f2
		defer recover() // Recover failed because the caller function f2 is NOT a deferred function
		panic("A bad boy stole a server, again!")
	}()
	time.Sleep(time.Second) // To prevent the main exiting before the f2 goroutine
}

func recoverCorrectly() {
	// Test in goroutine
	go func() {
		protect(func() {
			panicFunc("Run in goroutine")
		})
	}()
	time.Sleep(time.Second)

	// Test in main
	protect(func() {
		panicFunc("Run in main")
	})
}

func protect(objFunc func()) {
	// Handle the recover before the objFunc() running
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover from a runtime panic: %v, error stack: %s\n", err, debug.Stack())
			// Handle the panic below
			// someOperationsHere()
			log.Println("finish completely")
		}
	}()
	// Run the objective function
	objFunc()
}

// Accept a parameter and send it to panic
func panicFunc(x string) {
	errStr := "panic and the para is " + x
	panic(errStr)
}
