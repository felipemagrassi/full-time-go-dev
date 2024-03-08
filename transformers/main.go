package main

import (
	"fmt"
	"strings"
)

type StringTransformer func(string) string

func transformStringWithDependency(str string) string {
	return strings.ToUpper(str)
}

func transformString(str string, transformer StringTransformer) string {
	return transformer(str)
}

func Prefixer(prefix string) StringTransformer {
	return func(str string) string {
		return prefix + str
	}
}

func RemoveNonAscii() StringTransformer {
	return func(str string) string {
		var sb strings.Builder
		for _, ch := range str {
			if ch < 128 {
				sb.WriteRune(ch)
			}
		}
		return sb.String()
	}
}

func Truncate(length int) StringTransformer {
	return func(str string) string {
		if len(str) > length {
			return str[:length]
		}
		return str
	}
}

func MultipleTransformer(transformers ...StringTransformer) StringTransformer {
	return func(str string) string {
		for _, transformer := range transformers {
			str = transformer(str)
		}
		return str
	}
}

func main() {
	fmt.Println(transformStringWithDependency("Hello, World!"))
	fmt.Println(transformString("Hello, World!", strings.ToUpper))
	fmt.Println(transformString("Hello, World!", Prefixer("Prefix: ")))

	fmt.Println(transformString("©Hello, World!", MultipleTransformer(
		Prefixer("Prefix: "),
		RemoveNonAscii(),
		Truncate(10),
	)))
}
