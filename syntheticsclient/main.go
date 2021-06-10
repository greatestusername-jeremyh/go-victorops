package main

import (
	"fmt"
	"os"

)

func main() {

	token := os.Getenv("API_ACCESS_TOKEN")

	c := NewClient(token)

	//HTTP CHECK Endpoints
	//res, _, err := c.GetCheck(175389) //175389

	o := GetChecksOptions{
		Page: 0,
		PerPage: 0,
	}
	res, _, err := c.GetChecks(&o)


	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}
}