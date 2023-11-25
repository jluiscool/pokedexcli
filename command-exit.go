package main

import "os"

func commandExit() error {
	//exits the code with a error code of 0
	//which means it successfully exited
	os.Exit(0)
	return nil
}
