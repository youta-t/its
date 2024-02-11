//go:generate go run github.com/youta-t/its/structer
//go:generate gofmt -w ./gen_structer/example.go
package internal

import "time"

type MyStruct struct {
	Name      string
	Value     []int
	Timestamp time.Time
}
