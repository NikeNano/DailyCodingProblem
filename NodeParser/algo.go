package NodeParser

import (
	"encoding/json"
)

type node struct {
	value string
	left  *node
	right *node
}

func serialise(input node) (string, error) {
	out, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func deserialise(input string) (node, error) {
	var output node
	err := json.Unmarshal([]byte(input), &output)
	if err != nil {
		return output, err
	}
	return output, nil
}

func serialiseStrict(input node, output []int) string {
	if input.left != nil {
		ouput = append(output, serialiseStrict(*input.left, output))
	} else {
		output = append(output, -1)
	}
	if input.right != nil {
		ouput := append(ouput, serialiseStrict(*input.right, output))
	} else {
		output = append(output, -1)
	}
	return input.value
}
