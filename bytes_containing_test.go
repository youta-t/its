package its_test

import "github.com/youta-t/its"

func ExampleBytesContaining_ok() {
	its.BytesContaining([]byte("bcd")).Match([]byte("abcde")).OrError(t)
	// Output:
}

func ExampleBytesContaining_ng() {
	its.BytesContaining([]byte("bcd")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.Contains(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x62, 0x63, 0x64})		--- @ ./bytes_containing_test.go:11
}
