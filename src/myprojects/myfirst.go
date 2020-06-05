package main

import(

	"fmt"
	"github.com/benbjohnson/clock"
	)
	
type Application struct{
	Clock clock.Clock
}

func main(){
	fmt.Println("hello world")
	var app Application
	app.Clock = clock.New()
	fmt.Println(app.Clock.Now())

}