package main

import (
	"fmt"
)

func main() {
	/* pp := p.New(p.Box7, 100)
	c := pp.Run()
	for {
		select {
		case <-p.Race(c):
			fmt.Print(
				p.PrintMultiText(pp.Print(" HAHAHA")),
			)
		}
	} */
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//ctx := context.WithValue(context.Background(), "arash", "midos")

	//	defer cancel()
	fmt.Println("start",
		ctx.Value("arash"),
		ctx.Value("arash1"),
		ctx.Value("arash2"),
	)

	fmt.Println("fin")
}
