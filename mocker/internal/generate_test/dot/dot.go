//go:generate go run github.com/youta-t/its/structer
package dot

type DotStruct struct {
	ThisIsDotStruct string
}

func (ds DotStruct) DotMethod(int) string {
	return ""
}

type DotInterface interface {
	DotMethod()
}

type DotMap map[string]string

type DotSlice []string

type DotGene[T any] struct {
	Field T
}
