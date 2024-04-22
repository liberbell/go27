package main

import "fmt"

func main() {
	website := map[string]string{
		"Google":             "https://google.com",
		"Amazon Web Service": "https://aws.com",
	}
	fmt.Println(website)
	fmt.Println(website["Amazon Web Service"])

	website["Linkedin"] = "https://www.linkedin.com"
	fmt.Println(website)

	delete(website, "Google")
	fmt.Println(website)
}
