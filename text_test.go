package its_test

import "github.com/youta-t/its"

func ExampleText_multibyte_ok() {
	// "絶句" (杜甫)
	its.Text(`
江碧鳥逾白
山青花欲燃
今春看又過
何日是歸年
`).Match(`
江碧鳥逾白
山青花欲燃
今春看又過
何日是歸年
`).OrError(t)
	// Output:
}

func ExampleText_multibytes_ng() {
	// "風景" (純銀もざいく; 山村暮鳥): https://www.aozora.gr.jp/cards/000136/files/52348_42039.html
	its.Text(`
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
かすかなるむぎぶえ
いちめんのなのはな
`).Match(`
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
いちめんのなのはな
ひばりのおしやべり
いちめんのなのはな
`).OrError(t)
	// Output:
	// ✘ (+ = got, - = want)		--- @ ./text_test.go:23
	//       |
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//       | いちめんのなのはな
	//     - | かすかなるむぎぶえ
	//     + | ひばりのおしやべり
	//       | いちめんのなのはな
	//       |
}

func ExampleText_ascii_ok() {
	its.Text(`
Lorem Ipsum:

    Lorem ipsum dolor sit amet,
    consectetur adipiscing elit,
    sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
`).Match(`
Lorem Ipsum:

    Lorem ipsum dolor sit amet,
    consectetur adipiscing elit,
    sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
`).OrError(t)
	// Output:
}

func ExampleText_ascii_ng() {
	// London Bridge Is Falling Down
	its.Text(`
Build it up with bricks and mortar,
Bricks and mortar, bricks and mortar,
Build it up with bricks and mortar,
My fair lady.
`).Match(`
Build it up with iron and steel,
Iron and steel, iron and steel,
Build it up with iron and steel,
My fair lady.
`).OrError(t)
	// Output:
	//
	// ✘ (+ = got, - = want)		--- @ ./text_test.go:79
	//       |
	//     - | Build it up with bricks and mortar,
	//     - | Bricks and mortar, bricks and mortar,
	//     - | Build it up with bricks and mortar,
	//     + | Build it up with iron and steel,
	//     + | Iron and steel, iron and steel,
	//     + | Build it up with iron and steel,
	//       | My fair lady.
	//       |
}
