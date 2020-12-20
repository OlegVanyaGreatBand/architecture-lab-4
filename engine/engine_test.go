package engine

import (
	"fmt"
	"time"
)

func printCommand(msg string) Command {
	return CommandFunc(func(h Handler) {
		println(msg)
	})
}

func lenCommand(str string) Command {
	return CommandFunc(func(h Handler) {
		h.Post(printCommand(fmt.Sprintf("%s len = %d", str, len(str))))
	})
}

func ExampleEventLoop()  {
	eventLoop := new(EventLoop)
	eventLoop.Start()

	eventLoop.Post(printCommand("Test0"))
	eventLoop.Post(lenCommand("Test0"))

	eventLoop.AwaitFinish()

	eventLoop.Post(printCommand("Test1")) // muted due to event loop finished
	eventLoop.Resume()
	eventLoop.Post(printCommand("Test2")) // posted successfully
	eventLoop.AwaitFinish()

	eventLoop.Resume()
	eventLoop.Post(printCommand("Test3")) // posted successfully
	c := CommandFunc(func(h Handler) {
		go func() {
			h.Post(printCommand("Started waiting")) // printed immediately
			time.Sleep(5 * time.Second)
			h.Post(printCommand("Waited 5 sec")) // printed in 5 sec
		}()
	})
	eventLoop.Post(c)

	eventLoop.AwaitFinish()
	time.Sleep(7 * time.Second) // needed because main routine exits and don't wait for children goroutines
}
