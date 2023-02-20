package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func DivideMessages(message map[string]interface{}) []string {
	// Convert the input map to a JSON string
	jsonString, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling input map to JSON string:", err)
		return nil
	}

	// Remove square brackets from the JSON string
	jsonString = bytes.ReplaceAll(jsonString, []byte("["), []byte(""))
	jsonString = bytes.ReplaceAll(jsonString, []byte("]"), []byte("\n"))
	jsonString = bytes.ReplaceAll(jsonString, []byte("{"), []byte("\n"))
	jsonString = bytes.ReplaceAll(jsonString, []byte("}"), []byte("\n\n"))

	// Split the JSON string by closing curly braces
	jsonSplit := strings.Split(string(jsonString), "}")

	// Trim whitespace and append a closing curly brace to each split string
	for i := range jsonSplit {
		jsonSplit[i] = strings.TrimSpace(jsonSplit[i]) + "}\n"
	}

	return jsonSplit
}
