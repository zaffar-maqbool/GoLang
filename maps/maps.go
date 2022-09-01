package main

import "fmt"

func main() {
	// 	wesites := []string{
	// 		"google.com",
	// 		"aws.com",
	// 	}

	websites := map[string]string{
		"Google": "google.com",
		"AWS":    "aws,com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

}
