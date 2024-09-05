//go:generate go run github.com/youta-t/its/structer
package dot

type DotStruct struct {
	Value string
}

type DotInterface interface {
	ThisIsEmbInterface()
}

type DotFunc func(int) string

type DotName int

type DotG[T any] struct {
	Field T
}

type DotSlice []string

type DotMap map[string]string
