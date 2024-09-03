package its_test

import "github.com/youta-t/its"

func ExampleText_multibytes() {
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
	// ✘ (+ = got, - = want)		--- @ ./text_test.go:7
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

func ExampleText_ascii() {
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
	// ✘ (+ = got, - = want)		--- @ ./text_test.go:46
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

func ExampleText() {
	its.Text(`
Lorem Ipsum:

    Lorem ipsum dolor sit amet,
    consectetur adipiscing elit,
    sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
    nisi ut aliquip ex ea commodo consequat.
    Duis aute irure dolor in reprehenderit in voluptate velit
    esse cillum dolore eu fugiat nulla pariatur.

    Excepteur sint occaecat cupidatat non proident,
    sunt in culpa qui officia deserunt mollit anim id est laborum.
`).Match(`
Lorem Ipsum:

    Lorem ipsum dolor sit amet,
    consectetur adipiscing elit,
    sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

    nisi ut aliquip ex ea commodo consequat.
    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
    Duis aute irure dolor in reprehenderit in voluptate velit
    esse cillum dolore eu fugiat nulla pariatur.

    sunt in culpa qui officia deserunt mollit anim id est laborum.
    Excepteur sint occaecat cupidatat non proident,
`).OrError(t)
	// Output:
	// ✘ (+ = got, - = want)		--- @ ./text_test.go:72
	//       |
	//       | Lorem Ipsum:
	//       |
	//       |     Lorem ipsum dolor sit amet,
	//       |     consectetur adipiscing elit,
	//       |     sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
	//       |
	//     - |     Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
	//       |     nisi ut aliquip ex ea commodo consequat.
	//     + |     Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
	//       |     Duis aute irure dolor in reprehenderit in voluptate velit
	//       |     esse cillum dolore eu fugiat nulla pariatur.
	//       |
	//     - |     Excepteur sint occaecat cupidatat non proident,
	//       |     sunt in culpa qui officia deserunt mollit anim id est laborum.
	//     + |     Excepteur sint occaecat cupidatat non proident,
	//       |
}
