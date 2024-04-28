package dot

type DotStruct struct {
	Value string
}

type DotInterface interface {
	ThisIsEmbInterface()
}

type DotFunc func(int) string

type DotName int
