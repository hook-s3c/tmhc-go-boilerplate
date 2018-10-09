package main

import ( 
	"testing" 
	"bytes"
	. "github.com/smartystreets/goconvey/convey" 
) 

var InterceptedStdout bytes.Buffer

func TestLaunchingAppAndGreetingUser(t *testing.T) { 

	Convey("Given the application is initialized", t, func() { 
    hello(&InterceptedStdout)

		Convey("When the application has loaded", func() { 

			Convey("Then the user should be prompted with a greeting", func() {
				So(InterceptedStdout.String(), ShouldEqual, "Hello TMHC!\n")
			}) 

		}) 

	}) 

}
