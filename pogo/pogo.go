package main

import "flag"

func main() {

	role := flag.String("role", "standalone", "individual developer git commands")
	flag.Parse()

	if *role == "standalone" {
		ProcessStandalone()
	}
}

func ProcessStandalone() {
	inquier := &inquierService{}

	standalone := &pogoStandalone{
		inquier: inquier,
	}

	standalone.process()
}
