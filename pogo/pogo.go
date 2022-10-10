package main

import "flag"

func main() {

	role := flag.String("role", "basic", "individual developer git commands")
	flag.Parse()

	if *role == "basic" {
		ProcessBasic()
	}

	if *role == "standalone" {
		ProcessStandalone()
	}
}

func ProcessBasic() {
	inquier := &inquierService{}

	basic := &pogoBasic{
		inquier: inquier,
	}

	basic.process()
}

func ProcessStandalone() {
	inquier := &inquierService{}

	standalone := &pogoStandalone{
		inquier: inquier,
	}

	standalone.process()
}
