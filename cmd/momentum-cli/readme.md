# Momentum-cli
Generate RPC functions and Javascript libraries with your Go Code.

# How it works
1. Add  comment `RPC` to the function you wish to convert.
2. Run `$ momentum-cli` within the directory you wish to generate `http.Handlers` of..
3. Add Handlers to your MUX or http handler. Read the [File properties and functions](#file-properties-and-functions) section to view information about the generated code.

## Install

	go install github.com/cheikhshift/momentum/cmd/momentum-cli

## Example
The following function will be parsed by momentum-cli :

	// RPC
	func RPCTest(input int) (result int ) {
		result = input * 400
		return		
	}	

The previous func  will generate JS function : (The variable definitions carry over to javascript.)

	function RPCTest(Input , function callback(ObjectResponse, success) )

Notes : Within callback function parameter `ObjectResponse` (javascript type object) the response key `result` will be available. Following the named variables provided. Callback function parameter `success` (javascript type boolean) is optional and indicates a successful RPC call to the server. If `success` is false, `ObjectResponse` will have one object key, which will be `error`. `error` is a string message of why the request has failed.

# What is generated?
Momentum will generate an additional go file. This file will have the same package name as its neighbors (go files in directory you ran command).

### File properties and functions  

	var FuncFactory =  "/funcfactory.js"

String path to generated javascript RPC library. Use this path with MomentumJS handler function. Include a script tag on your HTML pages with this URI (or other URI as long as it is using the MomentumJS handler) to load the generated javascript functions.

	var MomentumURI = "/momentum/funcs"

String path to momentum RPC functions. Use this path with the Momentum handler function to make your RPC functions accessible.


	func Momentum(w http.ResponseWriter, r *http.Request) 
Momentum http.HandlerFunc. This is responsible for making your RPC functions' functionality accessible.


	func MomentumChain(f http.HandlerFunc) http.HandlerFunc
Use this function to chain multiple Momentum handler with one path,

	func MomentumJS(w http.ResponseWriter, r *http.Request)
Serve RPC generated javascript functions library. 

	func MomentumJSChain(f http.HandlerFunc) http.HandlerFunc
Use this function to chain multiple Momentum javascript libraries.

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
