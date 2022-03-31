package main

func main() {
	a := "hello"
	var b *string
	var c *string
	for k, v := range a {
		if k < 2 {
			*b = v
		}
		if k >= 2 {
			&c += v
		}
	}
}
