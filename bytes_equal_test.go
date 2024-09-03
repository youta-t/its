package its_test

import "github.com/youta-t/its"

func ExampleBytesEqual_ok() {
	its.BytesEqual([]byte("abc")).Match([]byte("abc")).OrError(t)
	// Output:
}

func ExampleBytesEqual_ng() {
	its.BytesEqual([]byte("abc")).Match([]byte("acb")).OrError(t)
	// Output:
	// ✘ bytes.Equal(/* got */ []byte{0x61, 0x63, 0x62}, /* want */ []byte{0x61, 0x62, 0x63})		--- @ ./bytes_equal_test.go:11
}
