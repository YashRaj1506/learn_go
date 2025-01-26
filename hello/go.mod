module example/hello

go 1.23.4

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

//code imports packages contained in other modules, 
//you manage those dependencies through your code's own module. 
//That module is defined by a go.mod file that tracks the modules 
//that provide those packages.

// "go mod init example/hello"