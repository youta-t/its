//go:generate go run github.com/youta-t/its/structer
//go:generate gofmt -w ./gen_structer/sub.go
package sub1

type Sub1 struct {
	StringField string
}
