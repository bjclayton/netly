package main

import (
	"fmt"
)

func main() {
	cfg := config{
		addr: ":5000",
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		fmt.Println(err)
	}
}
