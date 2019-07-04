package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hqhs/types"
)

var MissingRequiredFieldError = errors.New("Some of required fields are missing from provided data")

type Presenter interface {
	IsPresent() bool
}

type ExampleT struct {
	F types.Int64  `json:"f"`
	I types.Int64  `json:"i"`
	E types.String `json:"e"`
	L types.String `json:"l"`
	D types.Int    `json:"d"`
}

func (e *ExampleT) Validate() error {
	requiredF := []Presenter{e.F, e.I, e.E}
	for _, c := range requiredF {
		if !c.IsPresent() {
			return MissingRequiredFieldError
		}
	}
	return nil
}

func main() {
	payloadOk := `
{
  "f": 123,
  "i": 321,
  "e": "hello where!"
}
`
	t1 := &ExampleT{}
	handleErr(json.Unmarshal([]byte(payloadOk), t1))
	handleErr(t1.Validate())
	fmt.Printf("Unmarshaled struct: %+v\n", t1)
	payloadErr := `
{
  "f": 123,
  "i": 321
}
`
	t2 := &ExampleT{}
	handleErr(json.Unmarshal([]byte(payloadErr), t2))
	handleErr(t2.Validate())
	fmt.Printf("Unmarshaled struct: %+v\n", t2)
	/* output:
		Unmarshaled struct: &{F:123 I:321 E:hello where! L: D:0}
		panic: Some of required fields are missing from provided data
	    ...
		*stacktrace*
	*/
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
