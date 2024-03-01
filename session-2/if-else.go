package main

import "fmt"

func ifElse() {
	var score int = 79

	if score >= 90 && score <= 100 {
		fmt.Println("Nilai kamu dapet A")
	} else if score >= 80 && score <= 90 {
		fmt.Println("Nilai kamu dapet B")
	} else {
		fmt.Println("Nilai kamu dapet C")
	}

	var statusCode int = 204
	switch {
	case statusCode == 200 || statusCode == 201 || statusCode == 204:
		if statusCode == 200 {
			fmt.Println("STATUS_OK")
		} else if statusCode == 201 {
			fmt.Println("STATUS_CREATED")
		} else {
			fmt.Println("STATUS_NO_CONTENT")
		}
	case statusCode < 500 && statusCode >= 400:
		fmt.Println("USER_ERROR")
	default:
		fmt.Println("INTERNAL_SERVER_ERROR")
	}

	switch statusCode {
	case 400:
		fmt.Println("ERROR_BAD_REQUEST")
	case 500:
		fmt.Println("ERROR_INTERNAL_SERVER_ERROR")
	case 409:
		fmt.Println("ERROR_CONFLICT")
	default:
		{
			fmt.Println("UNDEFINED")
		}
	}

	// if isError := err; isError != nil {
	// 	fmt.Println(isError.Error())
	// } else {
	// 	fmt.Println("success")
	// }

}
