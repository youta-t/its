package its_test

import "github.com/youta-t/its"

func ExampleBytesHavingSuffix_ok() {
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("abcde")).OrError(t)
	// Output:
}

func ExampleBytesHavingSuffix_ng() {
	its.BytesHavingSuffix([]byte("cde")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasSuffix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x63, 0x64, 0x65})		--- @ ./bytes_having_suffix_test.go:11
}
