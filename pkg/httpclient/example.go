package httpclient

import "fmt"

func Example() {
	profile := New("192.168.1.162:3000")
	p, err := profile.Get("4461915a-d59e-4a0d-9cfc-45fd732eb13a")

	if err != nil {
		panic(err)
	}

	fmt.Println(p.Name)
}

func ExamplePost() {
	profile := New("192.168.1.162:3000")
	p, err := profile.CreateProfile("Alex2", 41)

	if err != nil {
		panic(err)
	}

	fmt.Println(p.ID)
}