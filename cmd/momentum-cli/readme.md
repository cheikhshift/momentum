# Momentum-cli
Generate JSON-RPC go handlers and Javascript libraries with your Go Code.

# How it works
1. Add  comment `RPC` to the function you wish to convert.
2. Run `$ momentum-cli` within the directory you wish to generate `http.Handlers` of..
3. Add Handlers to your MUX or http handler. Read the [File properties and functions](#file-properties-and-functions) section to view information about the generated code.

## Install

	go get github.com/cheikhshift/momentum/cmd/momentum-cli

## Example 1
The following function will be parsed by momentum-cli :

	// RPC
	func RPCTest(input int) (result int ) {
		result = input * 400
		return		
	}	

The previous func  will generate JS function : (The variable definitions carry over to javascript.)

	function RPCTest(Input , function callback(ObjectResponse, success) )

Notes : Within callback function parameter `ObjectResponse` (javascript type object) the object key `result` will have your function's result, following the named variables provided. Callback function parameter `success` (javascript type boolean) is optional and indicates a successful HTTP request to the server. If `success` is false, `ObjectResponse` will have one object key, which will be `error`. `error` is a string message of why the request has failed.

## Example 2
An interface function will be used with this example. The following function will be parsed by momentum-cli :

	type Stk struct {
		Var string
	}
	
	// RPC
	func (st * Stk) TestFunction () (res string) {
		log.Println("hello")
		log.Println(st.Var)
		st.Var = "Newval"
		res = "value string"
		return 
	}
	
The previous func  will generate the following JS code : (The interface initializer will be available in JS)

	/**
	* Stk
	* @namespace main
	* @param Obj - Object with Go interface fields.
	* @return Stk - Object with specified interface.
	*/
	var Stk = function(Obj) {
		Object.assign(this, Obj)
	}											
	Stk.prototype.TestFunction = function(  cb){
		var t = {};
		
		var payload = {Obj : this, Params : t};
		t.Working = true;
		this.cb = cb;
		jsrequestmomentum("/momentum/funcs?name=*Stk.TestFunction", payload, "POSTJSON", ObjectCallb.bind(this))
	}
 
Notes : definition of `function ObjectCallb`

	function ObjectCallb(jsonobj,success) {
		this.Working = false;
		Object.assign(this, jsonobj.Obj)
		this.cb(jsonobj.Result, success)
	}



# What is generated?
Momentum will generate an additional go file. This file will have the same package name as its neighbors (go files in directory you ran command).

### File properties and functions  

	var FuncFactory =  "/funcfactory.js"

String path to generated javascript RPC library. Use this path with MomentumJS handler function. Include a script tag on your HTML pages with this URI as the tag's `src` attribute (or other URI as long as it is using the MomentumJS handler) to load the generated javascript functions.

	var MomentumURI = "/momentum/funcs"

String path to momentum RPC functions. Use this path with the Momentum handler function to make your RPC functions accessible.


	func Momentum(w http.ResponseWriter, r *http.Request) 
Momentum http.HandlerFunc. This is responsible for making your RPC functions' functionality accessible.


	func MomentumChain(f http.HandlerFunc) http.HandlerFunc
Use this function to chain multiple Momentum handler with one path,

	func MomentumJS(w http.ResponseWriter, r *http.Request)
Serve RPC generated javascript functions library. 

	func MomentumJSChain(f http.HandlerFunc) http.HandlerFunc
Use this function to chain multiple generated Momentum javascript libraries.

## Sample code :

	package main

	import (
		"fmt"
		"github.com/gorilla/mux"
		"net/http"
		"log"
	)
	
	// Some comments maybe
	// mmm
	// RPC
	func RPCTest(input string) (result string) {
		result = fmt.Sprintf("Mod : %s", input)
		return
	}

	type Stk struct {
	Var string
}

	// RPC
	func (st * Stk) TestFunction () (res string) {
		log.Println("hello")
		log.Println(st.Var)
		st.Var = "Newval"
		res = "value string"
		return 
	}


	func main(){
		r := mux.NewRouter()
		// Serve generated library
		r.HandleFunc(FuncFactory, MomentumJS )
		// Make RPC functions' functionality
		// accessible
		r.HandleFunc(MomentumURI, Momentum)
	
		http.Handle("/", r)
		log.Fatal(http.ListenAndServe(":8000", r))
	}

## Api format

You can access each of your functions via server URI /momentum/funcs.

## Get Parameters

name : specifies the name of the function to invoke.

## Post Body

The body is used to pass a JSON of your function's variables/parameters. Each JSON key should match a corresponding function parameter name, with that parameter's data specified as the key's value. Even if your function has no parameters, pass at least an empty JSON (`{}`). If no data is passed momentum will return an EOF JSON error.

### Response format 
With interface methods the response format differs. Here is the schema

	{
		"Obj" : {} ,// interface with updates from function
		"Result" : {} , // result of function. Null if function returns nothing.
	}