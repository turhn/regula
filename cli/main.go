package main

import (
	"encoding/json"
	"fmt"
	"github.com/turhn/regula/interpreter/program"
)

func main() {
	source := `
Metadata:
Name = "Demo"
Version = "1.0"

Rules:
When
  1 is 1 then yes
    result = "ok"
#  yes is true
#    result = "double ok"
`
	fact := `
{
}
`

	prg := program.New(source)

	result := prg.Run(fact)

	bytes, _ := json.Marshal(result)

	fmt.Println(string(bytes))
}
