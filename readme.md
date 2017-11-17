# Momentum

A GopherSauce package to convert template and func tags into accessible Javascript functions. The goal of this project is to be able to: 1) generate your templates on your server  2) Ease access to server side functionality.

## Requirements
- [GopherSauce](http://gophersauce.com)

## How it works

### Install
Add this import tag to the root of your `gos.gxml` file. Within the `<gos>` tag. This will activate momentum within your project. Remember to include the funcfactory within the pages you wish to use these functions. `<script src="/funcfactory.js"></script>`

	<import src="github.com/cheikhshift/momentum/gos.gxml">

### Templates
Each template created will have a JS function equivalent with the same name. The format of the generated functions are as follows :	

	function TemplateName(ObjectWithInterfaceFields, function callback(StringOfRenderedTemplate) {} )

Notes :  Replace TemplateName with the string specified as the name attribute of your template tag. `ObjectWithInterfaceFields` is a JS object that will be converted to the interface specified within your template tag's `struct` attribute. On error, the reason why will be returned instead of the template HTML.


### Funcs
Each `func` tag within your `<methods>` section will generate a JS equivalent function. To better visualize this follow the example below.
Notes : Channels are not supported as a valid type within `var` and `return` attributes of your func tag. Using the term `args ...interface` will not work as well. This is why the `<func>` tag was introduced, to provide users with a tag to explicitly declare variable types. With this in mind `<method>` tags will not work with JS function factory because it relies `args ..interface` to pass variables.

A sample `<func>` tag is declared within the methods section of a `gos.gxml` file. The return type specifies names returned values.

	 <func name="TestAdd" var="varx string,numv int" return="(test string, err error)">
		err = errors.New("error test on GOOOO")	
			test  = "Test" + varx
			return 
	</func>

The previous func tag will generate JS function : (The variable definitions carry over to javascript.)

	function TestAdd(Varx,Numv, function callback(ObjectResponse, success) )


Notes :  ObjectResponse variable is an object with your function's returned values. With this `<func>` tag,  the ObjectResponse will have keys `err` and `test` following the tag's return attribute. Success is an indication of successful method invocation. On error, Object response will have one key : `error`, which is a string explanation of why the request failed.

**More notes : You must specify the names of the return types as well, AKA `Named returned values`. These names will be used as key names to your response oject.

*** Angular notes : If you plan on using Angular JS, use this $scope function to update your data after setting it within your callback `$scope.$apply();`.

#### Api format
You can access each of your functions via server URI `/momentum/funcs`.

##### Get Parameters
- name : specifies the name of the function to invoke.

##### Post Body
The body is used to pass a json of your function's variables/parameters. Each json key should match a correponding function variable name, with that variable's data specified as the key's value.
