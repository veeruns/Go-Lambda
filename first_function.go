package main

import (
	"fmt"
)

//Event Structure
type Event struct {
	Username string
}

func handler(e Event) (string, error) {
	return fmt.Sprintf("<h1> Hello %s from Lambda in Go </h1>", e.Username), nil
}

/*func main() {
	lambda.Start(handler)
}*/
