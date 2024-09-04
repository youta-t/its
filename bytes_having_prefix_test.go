package its_test

import "github.com/youta-t/its"

func ExampleBytesHavingPrefix_ok() {
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("abcde")).OrError(t)
	// Output:
}

func ExampleBytesHavingPrefix_ng() {
	its.BytesHavingPrefix([]byte("abc")).Match([]byte("adcbe")).OrError(t)
	// Output:
	// ✘ bytes.HasPrefix(/* got */ []byte{0x61, 0x64, 0x63, 0x62, 0x65}, /* want */ []byte{0x61, 0x62, 0x63})		--- @ ./bytes_having_prefix_test.go:11
}
