# go-test
Generate test and benchmark functions of your Go functions.

## Install

### Install deep.
The following package will be used in your generated code : `go get github.com/go-test/deep`

### Install go-test

	go get github.com/cheikhshift/momentum/cmd/go-test


# How it works
1. Add  comment (on new line) `@test` to the function you wish to test.
3. Run `$ go-test`. (optional) Use flag `workdir` to specify path of directory with go sources.

## CLI flags

	  Usage of go-test:
	  -bench
	    	Run benchmark tests after source files are created.
	  -test $ go test
	    	Run $ go test after source files are created.
	  -workdir string
	    	Path of directory with go sources to convert.


## Define test cases
Add a comment, to your function, with following format `@case <input_variables> @equal <output_variables>` to define a test case. Definitions of `@case` syntax  :
1. `<input_variables>` : comma delimited function parameters (in Go syntax).. For example to test `func Test(fnParam int) int`,  the comment would be `// @case 40 @equal 50` (assuming Test adds 10 to `fnParam`). 
2.  `<output_variables>` : comma delimited function return values (in Go syntax). For example to test `func Test(fnParam int) (int,error)`,  the comment would be `// @case 40 @equal 50,nil` (assuming Test adds 10 to `fnParam`). 

### Interface functions
Following a comment with `@case` you may also define the interface to use with your test, as well as to compare with. The syntax to use is `@obj <init_value> @equal <expected_value>`.Definitions of `@obj` syntax  :

1. `<init_value>` :  Go syntax to create your functions interface. For example to test `func (obj * SomeStruct) Test(fnParam int) int`,  the comment would be `// @obj &SomeStruct{} @equal &SomeStruct{}` .

	2.  `<expected_value>` : comma delimited function return values (in Go syntax). For example to test `func (obj * SomeStruct) Test(fnParam int) (int,error)`,  the comment would be `// @obj &SomeStruct{} @equal &SomeStruct{Var : "100"}` (assuming Test updates the interface field `Var` to `100`). 

For example : 

	// @test
	// @case "string" @equal "value string", nil
	// @obj &Stk{Var : "Newva"} @equal &Stk{Var : "Newval"}
	func (st * Stk) TestFunction (cas string) (res string,err error) {
