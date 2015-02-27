package main

import (
	"encoding/json"
	"github.com/DavidHuie/quartz/go/quartz"
)

type Foos struct{}

type Foo struct {
	Name string
}
type FooArgs struct {
	Num int
}

type Response struct {
	Output string
}

func (t *Foos) Test(args FooArgs, response *Response) error {
	foos := []Foo{}
	for i := 0; i < 1000; i++ {
		foos = append(foos, Foo{"Test"})
	}
	j, _ := json.Marshal(&foos)
	*response = Response{string(j)}
	return nil
}

func main() {
	foos := &Foos{}
	quartz.RegisterName("foos", foos)
	quartz.Start()
}
