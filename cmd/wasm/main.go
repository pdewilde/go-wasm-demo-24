package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper()) // registers function so can be called from js
	select {}                                    // Wait on a channel so go program doesn't exit. Above is pretty much creating a server.
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

// syscall/js provides js.Func type which exports function from go to js.
func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}
