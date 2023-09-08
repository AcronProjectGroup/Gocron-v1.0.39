package features

type Features struct {
	EveryWordsInGolang string
}


var OriginalFeatures Features = Features{
	EveryWordsInGolang: `

in Command Line Interface Go:
1- go mod init YOUR_NAME
2- go work init YOUR_WORK_DIRECTORY
3- go run main.go  OR  go run projectName.go


Using the Go Command
1.build
    The go build command compiles the source code in the current directory 
    and generates an executable file.

2.clean
    The go clean command removes the output produced by the go build command, 
    including the executable and any temporary files that were created during the build.

3.doc
    The go doc command generates documentation from source code.

4.fmt
    The go fmt command ensures consistent indentation and alignment in source code files.

5.get
    The go get command downloads and installs external packages.
    flag
    flag package:
    Command-line flags are a common way to specify options for command-line programs. 
    For example, in wc -l the -l is a command-line flag.
    Go provides a flag package supporting basic command-line flag parsing.
    We'll use this package to implement our example command-line program.

6.install
    The go install command downloads packages and is usually used to install tool packages.

7.help
    The go help command displays help information for other Go features. The command go
    help build, for example, displays information about the build argument.

8.mod
    The go mod command is used to create and manage a Go module.

9.run
    The go run command builds and executes the source code in a specified folder without
    creating an executable output

10.test
    The go test command executes unit tests

11.version
    The go version command writes out the Go version number.

12.vet
    The go vet command detects common problems in Go code

    

Useful Debugger State Commands
13.print <expr>
    This command evaluates an expression and displays the result. It can
    be used to display a value (print i) or perform a more complex test
    (print i > 0).

14.set <variable> = <value>
    This command changes the value of the specified variable.
    این دستور مقدار متغیر مشخص شده را تغییر می دهد.

15.locals
    This command prints the value of all local variables.

16.whatis <expr>
    This command prints the type of the specified expression such as whatis


Useful Debugger Commands for Controlling Execution
17.continue
    This command resumes execution of the application.

18.next
    This command moves to the next statement.

19.step
    This command steps into the current statement.

20.stepout
    This command steps out of the current statement.

21.restart
    This command restarts the process. Use the continue command to begin execution.
    
22.exit
    This command exits the debugger.



Understanding the Basic Data Types
23.int = integers
    This type represents a whole number, which can be positive or negative. The
    int type size is platform-dependent and will be either 32 or 64 bits. There are
    also integer types that have a specific size, such as int8, int16, int32, and
    int64, but the int type should be used unless you need a specific size.
    
    Literal Value Examples:
        20, -20. Values can also be expressed in hex (0x14), octal (0o24), and binary notation
        (0b0010100).

24.unit
    uint8,  uint16,  uint32,  uint64,int8,  int16,  int32
    This type represents a positive whole number. The uint type size is platform-
    dependent and will be either 32 or 64 bits. There are also unsigned integer
    types that have a specific size, such as uint8, uint16, uint32, and uint64, but
    the uint type should be used unless you need a specific size.

    Literal Value Examples:
        There are no uint literals. All literal whole numbers are treated as int values.

25.byte = uint8
    This type is an alias for uint8 and is typically used to represent a byte of data.
    byte = 8 bits, 
    1024bytes = 1 kilobyte
    1024 kilobytes = 1 megabyte

    Literal Value Examples:
        There are no byte literals. Bytes are typically expressed as integer literals (such as 101) or run
        literals ('e') since the byte type is an alias for the uint8 type.

26.float32, float64
    Floating Point Numbers. e.g 3.14, 123.541, 100.4020401
    These types represent numbers with a fraction. These types allocate 32 or 64
    bits to store the value.

    Literal Value Examples:
        20.2, -20.2, 1.2e10, 1.2e-10. Values can also be expressed in hex notation (0x2p10), although
        the exponent is expressed in decimal digits.

27.complex64, complex128
    These types represent numbers that have real and imaginary components.
    These types allocate 64 or 128 bits to store the value.

28.bool 
    bolean
    This type represents a Boolean truth with the values true and false.

    Literal Value Examples:
        true, false.

29.string
    This type represents a sequence of characters.

    Literal Value Examples:
        "Hello". Character sequences escaped with a backslash are interpreted if the value is enclosed
        in double quotes ("Hello\n"). Escape sequences are not interpreted if the value is enclosed in
        backquotes ('Hello\n').

30.rune = int32
    This type represents a single Unicode code point. Unicode is complicated,
    but—loosely—this is the representation of a single character. The rune type is
    an alias for int32.
   
    Literal Value Examples:
        'A', '\n', '\u00A5', '¥'. Characters, glyphs, and escape sequences are enclosed in single
        quotes (the ' character).

31.var = variable 
    Variables are defined using the var keyword, and, unlike constants, the value assigned to a variable can be
    changed

32.const = constant
    Constants are basically variables whose values cannot be changed later.


The Zero Values for the Basic Data Types
Type    Zero Value
------  -----------
int         0
unit        0
byte        0
float64     0
bool        false
string      “” (the empty string)
rune        0





Operations and Conversions
33.+, -, *, /, %
    These operators are used to perform arithmetic using numeric values.

34.==, !=, <, <=, >, >=
    These operators compare two values.

35.||, &&, !
    These are the logical operators, which are applied to bool values and return a bool value.

36.=, :=
    These are the assignment operators. The standard assignment operator (=) is used to set
    the initial value when a constant or variable is defined, or to change the value assigned to
    a previously defined variable. The shorthand operator (:=) is used to define a variable and
    assign a value.

37.-=, +=, ++, --
    These operators increment and decrement numeric values.

38.&, |, ^, &^, <<, >>
    These are the bitwise operators, which can be applied to integer values. These operators are
    not often required in mainstream development
    where the | operator is used to configure the Go logging features.

The Arithmetic Operators
Operator    Description
-------     ---------------------
   +        This operator returns the sum of two operands.
   -        This operator returns the difference between two operands.
   *        This operator returns the product of two operands.
   /        This product returns the quotient of two operators.
   %        This product returns the remainder, which is similar to the modulo operator provided by
            other programming languages but can return negative values, as described in the “Using the
            Remainder Operator” section.

The Comparison Operators
Operator    Description
-------     ---------------------
   ==       This operator returns true if the operands are equal.
   !=       This operator returns true if the operands are not equal.
   <        This operator returns true if the first operand is less than the second operand.
   >        This operator returns true if the first operand is greater than the second operand.
   <=       This operator returns true if the first operand is less than or equal to the second operand.
   >=       This operator returns true if the first operand is greater than or equal to the second
            operand.

The Logical Operators
Operator    Description
-------     ---------------------
   ||       This operator returns true if either operand is true. 
            If the first operand is true, then the second
            operand will not be evaluated.

   &&       This operator returns true if both operands are true. 
            If the first operand is false, then the
            second operand will not be evaluated.

   !        This operator is used with a single operand. 
            It returns true if the operand is false and false if
            the operand is true.

            
Converting Floating-Point Values to Integers
Functions in the math Package for Converting Numeric Types
39.Ceil(value)
    This function returns the smallest integer that is greater than the specified floating-
    point value. The smallest integer that is greater than 27.1, for example, is 28.

40.Floor(value)
    This function returns the largest integer that is less than the specified floating-point
    value. The largest integer that is less than 27.1, for example, is 28.

41.Round(value)
    This function rounds the specified floating-point value to the nearest integer.

42.RoundToEven(value)
    This function rounds the specified floating-point value to the nearest even integer.



Parsing from Strings
Functions for Parsing Strings into Other Data Types
43.ParseBool(str)
    This function parses a string into a bool value. Recognized string values are "true",
    "false", "TRUE", "FALSE", "True", "False", "T", "F", "0", and "1".

44.ParseFloat(str,size)
    This function parses a string into a floating-point value with the specified size, as
    described in the “Parsing Floating-Point Numbers” section.

45.ParseInt(str,base, size)
    This function parses a string into an int64 with the specified base and size. Acceptable
    base values are 2 for binary, 8 for octal, 16 for hex, and 10.

46.ParseUint(str,base, size)
    This function parses a string into an unsigned integer value with the specified base and size.

47.Atoi(str)
    This function parses a string into a base 10 int and is equivalent to calling
    ParseInt(str, 10, 0)


Formatting Values as Strings
The strconv Functions for Converting Values into Strings


48.FormatBool(val)
    This function returns the string true or false based on the value of the
    specified bool.

49.FormatInt(val, base)
    This function returns a string representation of the specified int64 value,
    expressed in the specified base.

50.FormatUint(val, base)
    This function returns a string representation of the specified uint64 value,
    expressed in the specified base.

51.FormatFloat(val, format, precision, size) 
    This function returns a string representation of the specified float64 value,
    expressed using the specified format, precision, and size.

52.Itoa(val)
    This function returns a string representation of the specified int value,
    expressed using base 10.


Commonly Used Format Options for Floating-Point String Formatting
Function
---------
f
    The floating-point value will be expressed in the form ±ddd.ddd without an exponent, such as
    49.95.

e, E
    The floating-point value will be expressed in the form ±ddd.ddde±dd, such as 4.995e+01 or
    4.995E+01. The case of the letter denoting the exponent is determined by the case of the rune
    used as the formatting argument.

g, G
    The floating-point value will be expressed using format e/E for large exponents or format f for
    smaller values.


Understanding Flow Control
53.if
    The if keyword is followed by the expression and then the group of statements to be executed,
    surrounded by braces

54.else
    The else keyword can be used to create additional clauses in an if statement
    
55.else if
    The else/if combination can be repeated to create a sequence of clauses

56.for
    Go allows loops only inside of functions.
    The for keyword is used to create loops that repeatedly execute statements. The most basic for loops will
    repeat indefinitely unless interrupted by the break keyword

    Incorporating the Condition into the Loop
    
    example: 
    |    for (counter <= 3) {}
    |___________________________________________________________

    Enumerating Sequences:
    |    for index, character := range product {
    |                fmt.Println("Index:", index, "Character:", string(character))
    |            }
    |___________________________________________________________


57.switch
    A switch statement provides an alternative way to control execution flow, based on matching the result of an
    expression to a specific value, as opposed to evaluating a true or false result
    example:
    |    switch (character) {
    |        case 'K':
    |            fmt.Println("K at position", index)
    |        case 'y':
    |            fmt.Println("y at position", index)
    |    }
    |___________________________________________________________

    Matching Multiple Values:
    example:
    |    switch (character) {
    |        case 'K', 'k':
    |            fmt.Println("K or k at position", index)
    |        case 'y':
    |            fmt.Println("y at position", index)
    |    }
    |___________________________________________________________

    Terminate case Statement Execution:
    example:
    |    switch (character) {
    |        case 'K', 'k':
    |    if (character == 'k') {
    |        fmt.Println("Lowercase k at position", index)
    |        break
    |    }
    |    fmt.Println("Uppercase K at position", index)
    |        case 'y':
    |    fmt.Println("y at position", index)
    |    }
    |__________________________________________________________


58.Falling Through
    Go switch statements don't automatically fall through, but this behavior can be enabled using the
    fallthrough keyword

    example:
    |    switch (character) {
    |        case 'K':
    |            fmt.Println("Uppercase character")
    |            fallthrough
    |        case 'k':
    |            fmt.Println("k at position", index)
    |        case 'y':
    |            fmt.Println("y at position", index)
    |        }
    |_________________________________________________________


59.default
    The default keyword is used to define a clause that will be executed when none of the case statements
    matches the switch statement's value
    
    example:
    |    switch (character) {
    |            case 'K', 'k':
    |            if (character == 'k') {
    |                fmt.Println("Lowercase k at position", index)
    |                break
    |            }
    |            fmt.Println("Uppercase K at position", index)
    |            case 'y':
    |            fmt.Println("y at position", index)
    |            default:
    |            fmt.Println("Character", string(character), "at position", index)
    |        }
    |___________________________________________________________

60.goto label 
    Label statements allow execution to jump to a different point, 
    giving greater flexibility than other flow control features
    example:
    |    counter := 0
    |    target: fmt.Println("Counter", counter)
    |    counter++
    |    if (counter < 5) {
    |            goto target
    |    }
    |__________________________________________________________


61.Array
    Go arrays are a fixed length and contain elements of a single type, which are accessed by index
    var names [3]string

    Array Literal Syntax:
        names := [3]string { "Kayak", "Lifejacket", "Paddle" }

    The number of elements specified with the literal syntax can be less than the capacity of the array. 
    Any position in the array for which a value is not provided will be assigned the zero value for the array type.

    Enumerating Arrays:
    example:
    |    names := [3]string { "Kayak", "Lifejacket", "Paddle" }
    |    for index, value := range names {
    |        fmt.Println("Index:", index, "Value:", value)
    |    }
    |___________________________________________________________


62.Slices
    The slice type in this example is []string
    The best way to think of slices is as a variable-length array because they are useful when you don't know how
    many values you need to store or when the number changes over time. One way to define a slice is to use the
    built-in make function

    example:
    |    names := make([]string, 3)
    |    names[0] = "Kayak"
    |    names[1] = "Lifejacket"
    |    names[2] = "Paddle"
    |____________________________________________________________

    Literal Syntax:
    example:
    |    names := []string {"Kayak", "Lifejacket", "Paddle"}
    |____________________________________________________________


63.append
    If you define a slice variable but don't initialize it, then the result is a slice that has a length of zero
    and a capacity of zero, and this will cause an error when an element is appended to it.
    One of the key advantages of slices is that they can be expanded to accommodate additional elements
    
    Appending Elements to a Slice:
    example:
    |    names := []string {"Kayak", "Lifejacket", "Paddle"}
    |    names = append(names, "Hat", "Gloves")
    |_____________________________________________________________

    Appending Items to a Slice:
    example:
    |    names := []string {"Kayak", "Lifejacket", "Paddle"}
    |    appendedNames := append(names, "Hat", "Gloves")
    |    names[0] = "Canoe"
    |_____________________________________________________________

    Allocating Additional Slice Capacity:
    example:
    |    names := make([]string, 3, 6)
    |    names[0] = "Kayak"
    |    names[1] = "Lifejacket"
    |    names[2] = "Paddle"
    |_____________________________________________________________

    Adding Elements to a Slice:
    example:
    |    names := make([]string, 3, 6)
    |    names[0] = "Kayak"
    |    names[1] = "Lifejacket"
    |    names[2] = "Paddle"
    |    appendedNames := append(names, "Hat", "Gloves")
    |    names[0] = "Canoe"
    |    fmt.Println("names:",names)
    |    fmt.Println("appendedNames:", appendedNames)
    |______________________________________________________________

    Appending One Slice to Another:
    example:
    |    names := make([]string, 3, 6)
    |    names[0] = "Kayak"
    |    names[1] = "Lifejacket"
    |    names[2] = "Paddle"
    |    moreNames := []string { "Hat Gloves"}
    |    appendedNames := append(names, moreNames...)
    |    fmt.Println("appendedNames:", appendedNames)
    |______________________________________________________________

    Creating Slices from Existing Arrays:
    example:
    |    products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    someNames := products[1:3]
    |    allNames := products[:]
    |______________________________________________________________

    Specifying Capacity When Creating a Slice from an Array:
    example:
    |    products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    someNames := products[1:3:3]
    |    allNames := products[:]
    |    someNames = append(someNames, "Gloves")
    |______________________________________________________________

    Creating Slices from Other Slices:
    example:
    |    products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    allNames := products[1:]
    |    someNames := allNames[1:3]
    |    allNames = append(allNames, "Gloves")
    |    allNames[1] = "Canoe"
    |______________________________________________________________

    Comparing Slices:
    example:
    |    Go restricts the use of the comparison operator so that slices can be compared only to the nil value.
    |        p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |        p2 := p1
    |        fmt.Println("Equal:", p1 == p2)
    |_______________________________________________________________

    Output:
    |    .\main.go:13:30: invalid operation: p1 == p2 (slice can only be compared to nil)
    |_______________________________________________________________

64.copy
    The copy function is used to copy elements between slices. This function can be used to ensure that slices
    have separate arrays and to create slices that combine elements from different sources.

         Destination   Source
         ---------     --------
    copy(someNames,    allNames)


    Duplicating a Slice:
    The copy function accepts two arguments:
    example:
    |    products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    allNames := products[1:]
    |    someNames := make([]string, 2)
    |    copy(someNames, allNames)
    |_________________________________________

    Output:
    |    allNames [Lifejacket Paddle Hat]
    |    someNames: [Lifejacket Paddle]
    |____________________________________________

    the Uninitialized Slice Pitfall:
    the copy function doesn't resize the destination slice. A common
    pitfall is to try to copy elements into a slice that has not been initialized
    example:
    |    products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    allNames := products[1:]
    |    var someNames []string
    |    copy(someNames, allNames)
    |___________________________________________

    Output:
    |    someNames: []
    |    allNames [Lifejacket Paddle Hat]
    |____________________________________________

    Specifying Ranges When Copying Slices:
    Fine-grained control over the elements that are copied can be achieved using ranges
    example:
    |    products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    allNames := products[1:]
    |    someNames := []string { "Boots", "Canoe"}
    |    copy(someNames[1:], allNames[2:3])
    |______________________________________________

    Output:
    |    allNames [Lifejacket Paddle Hat]
    |    someNames: [Boots Hat]
    |__________________________________________

    Copying Slices with Different Sizes:
    |    products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    replacementProducts := []string { "Canoe", "Boots"}
    |    copy(products, replacementProducts)
    |_______________________________________________

    Output:
    |    products: [Canoe Boots Paddle Hat]
    |______________________________________________   

    Copying a Larger Source Slice
    example:
    |    products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |        replacementProducts := []string { "Canoe", "Boots"}
    |        copy(products[0:1], replacementProducts)
    |        fmt.Println("products:", products)
    |__________________________________________________

    Output:
    |    products: [Canoe Lifejacket Paddle Hat]
    |__________________________________________________

    Deleting Slice Elements:
        products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat"}
        deleted := append(products[:2], products[3:]...)
        fmt.Println("Deleted:", deleted)

    
    Enumerating Slices:
    example:
    |    products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    for index, value := range products[2:] {
    |        fmt.Println("Index:", index, "Value:", value)
    |    }
    |_______________________________________

    Output:
    |    Index: 0 Value: Paddle
    |    Index: 1 Value: Hat
    |________________________________


65.sort
    There is no built-in support for sorting slices, but the standard library includes the sort package, which
    defines functions for sorting different types of slice.

    example:
    |    products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |    sort.Strings(products)
    |    for index, value := range products {
    |        fmt.Println("Index:", index, "Value:", value)
    |    }
    |_____________________________________________________


    Output:
    |    Index: 0 Value: Hat
    |    Index: 1 Value: Kayak
    |    Index: 2 Value: Lifejacket
    |    Index: 3 Value: Paddle
    |_________________________________________

66.DeepEqual
    The DeepEqual function can be used to compare a wider range of data types than the
    equality operator, including slices.

    example:
    |    package main
    |    import (
    |        "fmt"
    |        "reflect"
    |    )
    |    func main() {
    |        p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
    |        p2 := p1
    |        fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
    |    }
    |_____________________________________________________________
    
    Output:
    |    Equal: true
    |___________________________________________________


67.map
    Maps are a built-in data structure that associates data values with keys. 
    Unlike arrays, where values are associated with sequential integer locations, 
    maps can use other data types as keys

    example:
    |    products := make(map[string]float64, 10)
    |    products["Kayak"] = 279
    |    products["Lifejacket"] = 48.95
    |____________________________________________________

    Map Literal Syntax:
        products := map[string]float64 {
            "Kayak" : 279,
            "Lifejacket": 48.95,
        }

    Checking for Items in a Map:
    example:
    |    products := map[string]float64 {
    |                "Kayak" : 279,
    |                "Lifejacket": 48.95,
    |                "Hat": 0,
    |            }
    |            value, ok := products["Hat"]
    |            if (ok) {
    |                fmt.Println("Stored value:", value)
    |            } else {
    |                fmt.Println("No stored value")
    |            }
    |_________________________________________________________

    Removing Items from a Map:
    example:
    |    products := map[string]float64 {
    |        "Kayak" : 279,
    |        "Lifejacket": 48.95,
    |        "Hat": 0,
    |    }
    |    delete(products, "Hat")            
    |___________________________________________________


    Enumerating the Contents of a Map:
    example:
    |    products := map[string]float64 {
    |            "Kayak" : 279,
    |            "Lifejacket": 48.95,
    |            "Hat": 0,
    |    }
    |    for key, value := range products {
    |        fmt.Println("Key:", key, "Value:", value)
    |    }
    |__________________________________________________________

    Enumerating a Map in Order:
    example:
    |    import (
    |            "fmt"
    |            "sort"
    |        )
    |    func main() {
    |        products := map[string]float64 {
    |            "Kayak" : 279,
    |            "Lifejacket": 48.95,
    |            "Hat": 0,
    |        }
    |        keys := make([]string, 0, len(products))
    |        for key, _ := range products {
    |            keys = append(keys, key)
    |        }
    |        sort.Strings(keys)
    |        for _, key := range keys {
    |            fmt.Println("Key:", key, "Value:", products[key])
    |        }
    |    }
    |______________________________________________________
        

Understanding the Dual Nature of Strings
68.strconv
    Indexing and Slicing a String:
        var price string = "$48.95"
        var currency byte = price[0]
        var amountString string = price[1:]
        amount, parseErr := strconv.ParseFloat(amountString, 64)
        fmt.Println("Currency:", currency)
        if (parseErr == nil) {
                fmt.Println("Amount:", amount)
        } else {
                fmt.Println("Parse Error:", parseErr)
        }
    Output:
        Currency: 36
        Amount: 48.95

    Converting the Result:
    |    var price string = "$48.95"
    |    var currency string = string(price[0])
    |    var amountString string = price[1:]
    |    amount, parseErr  := strconv.ParseFloat(amountString, 64)
    |    fmt.Println("Currency:", currency)
    |    if (parseErr == nil) {
    |        fmt.Println("Amount:", amount)
    |    } else {
    |        fmt.Println("Parse Error:", parseErr)
    |    }
    |
    Output:
        Currency: $
        Amount: 48.95

    Changing the Currency Symbol:
    example:
    |   var price string = "€48.95"
    |   var currency string = string(price[0])
    |   var amountString string = price[1:]
    |   amount, parseErr  := strconv.ParseFloat(amountString, 64)
    |   fmt.Println("Currency:", currency)
    |   if (parseErr == nil) {
    |       fmt.Println("Amount:", amount)
    |   } else {
    |       fmt.Println("Parse Error:", parseErr)
    |   }
    |______________________________________________

    Output:
    |    Currency: â
    |    Parse Error: strconv.ParseFloat: parsing "\x82\xac48.95": invalid syntax
    |______________________________________________
    
    Obtaining the Length:
        fmt.Println("Length:", len(price))
    Output:
        Length: 8

Converting a String to Runes
69.rune
    The rune type represents a Unicode code point, 
    which is essentially a single character. 
    To avoid slicing strings in the middle of characters, 
    an explicit conversion to a rune slice can be performed
    Unicode is incredibly complex, as you would expect 
    from any standard that aims to describe multiple
    writing systems evolved over thousands of years.

    Converting to Runes:
        var price []rune = []rune("€48.95")
        var amountString string = string(price[1:])



    Enumerating Strings        
    example:
    |    var price = "€48.95"
    |    for index, char := range price {
    |        fmt.Println(index, char, string(char))
    |    }
    |_______________________________________________

    Output:
    |    0 8364 €
    |    3 52 4
    |    4 56 8
    |    5 46 .
    |    6 57 9
    |    7 53 5
    |_________________________________________________

    Enumerating the Bytes
    example:
    |    var price = "€48.95"
    |    for index, char := range []byte(price) {
    |        fmt.Println(index, char)
    |    }
    |_________________________________________________

    Output:
    |    0 226
    |    1 130
    |    2 172
    |    3 52
    |    4 56
    |    5 46
    |    6 57
    |    7 53
    |____________________________

70.func
    Functions are groups of code statements that are executed only when the function is
    invoked during the flow of execution.

71.Function Parameters
    Parameters allow a function to receive data values when it is called, 
    allowing its behavior to be altered.
    Values for parameters are supplied as arguments when invoking the function, meaning that different
    values can be provided each time the function is called. Arguments are provided between the parentheses
    that follow the function name, separated by commas and in the same order in which the parameters have
    been defined

    example:
    |    func printPrice(product string, price float64, taxRate float64) {
    |        taxAmount := price * taxRate
    |        fmt.Println(product, "price:", price, "Tax:", taxAmount)
    |    }
    |___________________________________________

72.Defining Variadic Parameters
    example:
    |    func printSuppliers(product string, suppliers []string ) {
    |        for _, supplier := range suppliers {
    |            fmt.Println("Product:", product, "Supplier:", supplier)
    |        }
    |    }
    |______________________________________________

    example:
    |    func printSuppliers(product string, suppliers ...string ) {
    |        for _, supplier := range suppliers {
    |            fmt.Println("Product:", product, "Supplier:", supplier)
    |        }
    |    }
    |_______________________________________________

73.Dealing with No Arguments for a Variadic Parameter
    example:
    |    func printSuppliers(product string, suppliers ...string ) {
    |        for _, supplier := range suppliers {
    |            fmt.Println("Product:", product, "Supplier:", supplier)
    |        }
    |    }
    |    func main() {
    |        printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
    |        printSuppliers("Lifejacket", "Sail Safe Co")
    |        printSuppliers("Soccer Ball")
    |    }
    |______________________________________________

    Output:
    |    Product: Kayak Supplier: Acme Kayaks
    |    Product: Kayak Supplier: Bob's Boats
    |    Product: Kayak Supplier: Crazy Canoes
    |    Product: Lifejacket Supplier: Sail Safe Co
    |________________________________________

74.return Function Results
    example:
    |    func calcTax(price float64) float64 {
    |        return price + (price * 0.2)
    |    }
    |_________________________________________

75.Returning Multiple Function Results
    example:
    |    func swapValues(first, second int) (int, int) {
    |        return second, first
    |    }
    |    func main() {
    |        val1, val2 := 10, 20
    |        fmt.Println("Before calling function", val1, val2)
    |        val1, val2 = swapValues(val1, val2)
    |        fmt.Println("After calling function", val1, val2)
    |    }
    |______________________________________

76.Using Named Results
    example:
    |    func calcTax(price float64) (float64, bool) {
    |        if (price > 100) {
    |            return price * 0.2, true
    |        }
    |        return 0, false
    |    }
    |    func calcTotalPrice(products map[string]float64,
    |            minSpend float64) (total, tax float64)  {
    |        total = minSpend
    |        for _, price := range products {
    |            if taxAmount, due := calcTax(price); due {
    |                total += taxAmount;
    |                tax += taxAmount
    |            } else {
    |                total += price
    |            }
    |        }
    |        return
    |    }
    |_____________________________________________________

77.defer
    The defer keyword is used to schedule a function call that will be performed immediately before the current
    function returns
    The defer keyword lets you group the statements that create, use, and
    release the resource together.
    The defer keyword can be used with any function call
    a single function can use the defer keyword multiple times.
    Immediately before the function returns, Go will perform the
    calls scheduled with the defer keyword in the order in which they were defined.

78.Function Types
    Functions in Go have a data type, which describes the combination of parameters the
    function consumes and the results the function produces. This type can be specified
    explicitly or inferred from a function defined using a literal syntax.
    Function types are defined using the func keyword, followed by a signature that
    describes the parameters and results. No function body is specified.

    Go does not support arrow functions, where functions are expressed more concisely using the =>
    operator, without the func keyword and a code block surrounded by braces. In Go, functions must always be
    defined with the keyword and a body.

79.Function Comparisons and the Zero Type
    example:
    |    func calcWithTax(price float64) float64 {
    |        return price + (price * 0.2)
    |    }
    |    func calcWithoutTax(price float64) float64 {
    |        return price
    |    }
    |    func main() {
    |        products := map[string]float64 {
    |            "Kayak" : 275,
    |            "Lifejacket": 48.95,
    |        }
    |        for product, price := range products {
    |            var calcFunc func(float64) float64
    |            fmt.Println("Function assigned:", calcFunc == nil)
    |            if (price > 100) {
    |                calcFunc = calcWithTax
    |            } else {
    |                calcFunc = calcWithoutTax
    |            }
    |            fmt.Println("Function assigned:", calcFunc == nil)
    |            totalPrice := calcFunc(price)
    |            fmt.Println("Product:", product, "Price:", totalPrice)
    |        }
    |    }
    |_________________________________________________

80.Functions as Arguments
    example:
    |    func calcWithTax(price float64) float64 {
    |        return price + (price * 0.2)
    |    }
    |    func calcWithoutTax(price float64) float64 {
    |        return price
    |    }
    |    func printPrice(product string, price float64, calculator func(float64) float64 ) {
    |        fmt.Println("Product:", product, "Price:", calculator(price))
    |    }
    |    func main() {
    |        products := map[string]float64 {
    |            "Kayak" : 275,
    |            "Lifejacket": 48.95,
    |        }
    |        for product, price := range products {
    |            if (price > 100) {
    |                printPrice(product, price, calcWithTax)
    |            } else {
    |                printPrice(product, price, calcWithoutTax)
    |            }
    |        }
    |    }
    |____________________________________________________

81.Functions as Results
    example:
    |    func calcWithTax(price float64) float64 {
    |        return price + (price * 0.2)
    |    }
    |    func calcWithoutTax(price float64) float64 {
    |        return price
    |    }
    |    func printPrice(product string, price float64, calculator func(float64) float64 ) {
    |        fmt.Println("Product:", product, "Price:", calculator(price))
    |    }
    |    func selectCalculator(price float64) func(float64) float64 {
    |        if (price > 100) {
    |            return calcWithTax
    |        }
    |        return calcWithoutTax
    |    }
    |    func main() {
    |        products := map[string]float64 {
    |            "Kayak" : 275,
    |            "Lifejacket": 48.95,
    |        }
    |        for product, price := range products {
    |            printPrice(product, price, selectCalculator(price))
    |        }
    |    }
    |_________________________________________________

82.Function Type Aliases
    Go supports type aliases, which can be used to assign a name to
    a function signature so that the parameter and result types are not specified every time the function type is
    used.
    example:
    |    type calcFunc func(float64) float64
    |    func calcWithTax(price float64) float64 {
    |        return price + (price * 0.2)
    |    }
    |    func calcWithoutTax(price float64) float64 {
    |        return price
    |    }
    |    func printPrice(product string, price float64, calculator calcFunc) {
    |        fmt.Println("Product:", product, "Price:", calculator(price))
    |    }
    |    func selectCalculator(price float64) calcFunc {
    |        if (price > 100) {
    |            return calcWithTax
    |        }
    |        return calcWithoutTax
    |    }
    |    func main() {
    |        products := map[string]float64 {
    |            "Kayak" : 275,
    |            "Lifejacket": 48.95,
    |        }
    |        for product, price := range products {
    |            printPrice(product, price, selectCalculator(price))
    |        }
    |    }
    |___________________________________________________

83.the Literal Function Syntax
    example:
    |    func selectCalculator(price float64) calcFunc {
    |        if (price > 100) {
    |            var withTax calcFunc = func (price float64) float64 {
    |                return price + (price * 0.2)
    |            }
    |            return withTax
    |        }
    |________________________________________________

84.Function Variable Scope
    example:
    |    func selectCalculator(price float64) calcFunc {
    |        if (price > 100) {
    |            var withTax calcFunc = func (price float64) float64 {
    |                return price + (price * 0.2)
    |            }
    |            return withTax
    |        } else if (price < 10) {
    |            return withTax
    |        }
    |         withoutTax := func (price float64) float64 {
    |                return price
    |            }
    |        return withoutTax
    |    }
    |_____________________________________________________

85.Functions Values Directly
    example:
    |    func selectCalculator(price float64) calcFunc {
    |        if (price > 100) {
    |            return func (price float64) float64 {
    |                return price + (price * 0.2)
    |            }
    |        }
    |         return func (price float64) float64 {
    |            return price
    |        }
    |    }
    |____________________________________________________



86.Literal Function Argument
    example:
    |    func main() {
    |        products := map[string]float64 {
    |            "Kayak" : 275,
    |            "Lifejacket": 48.95,
    |        }
    |        for product, price := range products {
    |            printPrice(product, price, func (price float64) float64 {
    |                return price + (price * 0.2)
    |            })
    |        }
    |    }
    |____________________________________________________

87.Function Closure
    Functions defined using the literal syntax can reference variables from the surrounding code, a feature
    known as closure.

    example:
    |    type calcFunc func(float64) float64
    |    func printPrice(product string, price float64, calculator calcFunc) {
    |        fmt.Println("Product:", product, "Price:", calculator(price))
    |    }
    |    func main() {
    |        watersportsProducts := map[string]float64 {
    |            "Kayak" : 275,
    |            "Lifejacket": 48.95,
    |        }
    |        soccerProducts := map[string] float64 {
    |            "Soccer Ball": 19.50,
    |            "Stadium": 79500,
    |        }
    |        calc := func(price float64) float64 {
    |            if (price > 100) {
    |                return price + (price * 0.2)
    |            }
    |            return price;
    |        }
    |        for product, price := range watersportsProducts {
    |            printPrice(product, price, calc)
    |        }
    |        calc = func(price float64) float64 {
    |            if (price > 50) {
    |                return price + (price * 0.1)
    |            }
    |            return price
    |        }
    |        for product, price := range soccerProducts {
    |            printPrice(product, price, calc)
    |        }
    |    }
    |_____________________________________________________________

        
88.struct
    What are they?
    Structs are data types, comprised of fields.
    
    Why are they useful?
    Structs allow custom data types to be defined.
    
    How are they used?
    The type and struct keywords are used to define a type, 
    allowing field names and types to be specified.
    
    A struct can mix regular and embedded field types.

    example:
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |        }
    |        kayak := Product {
    |            name: "Kayak",
    |            category: "Watersports",
    |            price: 275,
    |        }
    |        fmt.Println(kayak.name, kayak.category, kayak.price)
    |        kayak.price = 300
    |        fmt.Println("Changed price:", kayak.price)
    |    }
    |_________________________________________________________________

    Go doesn't differentiate between structs and classes, in the way that other languages do. All custom
    data types are defined as structs, and the decision to pass them by reference or by value is made
    depending on whether a pointer is used.

    Go doesn't allow structs to be used with the const keyword, and the compiler will report an error if
    you try to define a constant struct.

89.struct tag
    The struct type can be defined with tags, which provide additional information about how a field should
    be processed. Struct tags are just strings that are interpreted by the code that processes struct values,
    using the features provided by the reflect package.

90.struct values
    Values do not have to be provided for all fields when creating a struct value
    
    example:
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |        }
    |        kayak := Product {
    |                name: "Kayak",
    |                category: "Watersports",
    |        }
    |    }
    |_________________________________________________________

91.new 
    the new function to create struct values
        var lifejacket = new(Product)

    The result is a pointer to a struct value whose fields are 
    initialized with their type's zero value. 
    This is equivalent to this statement:
        var lifejacket = &Product{}

    These approaches are interchangeable, and choosing between them is a matter of preference.

92.Field Positions to Create Struct Values
    example:
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |        }
    |        var kayak = Product { "Kayak", "Watersports", 275.00 }
    |        fmt.Println("Name:", kayak.name)
    |        fmt.Println("Category:", kayak.category)
    |        fmt.Println("Price:", kayak.price)
    |    }
    |_____________________________________________________________

93.Defining Embedded Fields
    If a field is defined without a name, 
    it is known as an embedded field, and it is accessed using the name of its type.

    example:
    |    func main() {
    |            type Product struct {
    |                name, category string
    |                price float64
    |            }
    |            type StockLevel struct {
    |                Product
    |                count int
    |            }
    |            stockItem := StockLevel {
    |                Product: Product { "Kayak", "Watersports", 275.00 },
    |                count: 100,
    |            }
    |            fmt.Println("Name:", stockItem.Product.name)
    |            fmt.Println("Count:", stockItem.count)
    |    }
    |____________________________________________________________

94.Defining an Additional Field
    example:
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |        }
    |        type StockLevel struct {
    |            Product
    |            Alternate Product
    |            count int
    |        }
    |        stockItem := StockLevel {
    |            Product: Product { "Kayak", "Watersports", 275.00 },
    |            Alternate: Product{"Lifejacket", "Watersports", 48.95 },
    |            count: 100,
    |        }
    |        fmt.Println("Name:", stockItem.Product.name)
    |        fmt.Println("Alt Name:", stockItem.Alternate.name)
    |    }
    |__________________________________________________________________________

95.Comparing Struct Values
    example:
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |        }
    |        p1 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
    |        p2 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
    |        p3 := Product { name: "Kayak", category: "Boats", price: 275.00 }
    |        fmt.Println("p1 == p2:", p1 == p2)
    |        fmt.Println("p1 == p3:", p1 == p3)
    |    }
    |__________________________________________________________________________


96.Anonymous Struct Types
    Anonymous struct types are defined without using a name
    example:
    |    package main
    |    import "fmt"
    |    func writeName(val struct {
    |            name, category string
    |            price float64}) {
    |        fmt.Println("Name:", val.name)
    |    }
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |            //otherNames []string
    |        }
    |        type Item struct {
    |            name string
    |            category string
    |            price float64
    |        }
    |        prod := Product { name: "Kayak", category: "Watersports", price: 275.00 }
    |        item := Item { name: "Stadium", category: "Soccer", price: 75000 }
    |        writeName(prod)
    |        writeName(item)
    |    }
    |_________________________________________________________________________

97.Creating Arrays, Slices, and Maps Containing Struct Values
    Omitting the Struct Type
    example:
    |    package main
    |    import "fmt"
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |            //otherNames []string
    |        }
    |        type StockLevel struct {
    |            Product
    |            Alternate Product
    |            count int
    |        }
    |        array := [1]StockLevel {
    |            {
    |                Product: Product { "Kayak", "Watersports", 275.00 },
    |                Alternate: Product{"Lifejacket", "Watersports", 48.95 },
    |                count: 100,
    |            },
    |        }
    |        fmt.Println("Array:", array[0].Product.name)
    |        slice := []StockLevel {
    |            {
    |                Product: Product { "Kayak", "Watersports", 275.00 },
    |                Alternate: Product{"Lifejacket", "Watersports", 48.95 },
    |                count: 100,
    |            },
    |        }
    |        fmt.Println("Slice:", slice[0].Product.name)
    |        kvp := map[string]StockLevel {
    |            "kayak": {
    |                Product: Product { "Kayak", "Watersports", 275.00 },
    |                Alternate: Product{"Lifejacket", "Watersports", 48.95 },
    |                count: 100,
    |            },
    |        }
    |        fmt.Println("Map:", kvp["kayak"].Product.name)
    |    }
    |_________________________________________________________________________________________

98.Structs and Pointers
    Assigning a struct to a new variable or using a struct as a function parameter 
    creates a new value that copies the field values.

99.Copying a Struct Value
    example:
    |    package main
    |    import "fmt"
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |        }
    |        p1 := Product {
    |            name: "Kayak",
    |            category: "Watersports",
    |            price: 275,
    |        }
    |        p2 := p1
    |        p1.name = "Original Kayak"
    |        fmt.Println("P1:", p1.name)
    |        fmt.Println("P2:", p2.name)
    |    }
    |_______________________________________________


100.Using a Pointer to a Struct
    example:
    |    package main
    |    import "fmt"
    |    func main() {
    |        type Product struct {
    |            name, category string
    |            price float64
    |        }
    |        p1 := Product {
    |            name: "Kayak",
    |            category: "Watersports",
    |            price: 275,
    |        }
    |        p2 := &p1
    |        p1.name = "Original Kayak"
    |        fmt.Println("P1:", p1.name)
    |        fmt.Println("P2:", (*p2).name)
    |    }
    |_______________________________________________

101.the Struct Pointer Convenience Syntax
    example:
    |    package main
    |    import "fmt"
    |    
    |    type Product struct {
    |        name, category string
    |        price float64
    |    }
    |    
    |    func calcTax(product *Product) {
    |        if ((*product).price > 100) {
    |            (*product).price += (*product).price * 0.2
    |        }
    |    }
    |    
    |    func main() {
    |        kayak := Product {
    |            name: "Kayak",
    |            category: "Watersports",
    |            price: 275,
    |        }
    |        calcTax(&kayak)
    |        fmt.Println("Name:", kayak.name, "Category:",
    |        kayak.category, "Price", kayak.price)
    |    }
    |_______________________________________________________________


102.Struct Constructor Functions
    A constructor function is responsible for creating struct values using values received through parameters
    Constructor functions are used to create struct values consistently. Constructor functions are usually
    named new or New followed by the struct type so that the constructor function for creating Product values
    is named newProduct.

    example:
    |    package main
    |    import "fmt"
    |    type Product struct {
    |        name, category string
    |        price float64
    |    }
    |    
    |    func newProduct(name, category string, price float64) *Product {
    |        return &Product{name, category, price}
    |    }
    |    
    |    func main() {
    |        products := [2]*Product {
    |            newProduct("Kayak", "Watersports", 275),
    |            newProduct("Hat", "Skiing", 42.50),
    |        }
    |        for _, p := range products {
    |            fmt.Println("Name:", p.name, "Category:",  p.category, "Price", p.price)
    |        }
    |    }
    |______________________________________________________________________________


103.Modifying a Constructor
    The benefit of using constructor functions is consistency, 
    ensuring that changes to the construction
    process are reflected in all the struct values created by the function.

    example:
    |    func newProduct(name, category string, price float64) *Product {
    |        return &Product{name, category, price - 10}
    |    }
    |____________________________________________________________________________

104.Pointer Types for Struct Fields
    example:
    |    package main
    |    import "fmt"
    |    type Product struct {
    |        name, category string
    |        price float64
    |        *Supplier
    |    }
    |    type Supplier struct {
    |        name, city string
    |    }
    |    func newProduct(name, category string, price float64, supplier *Supplier) *Product {
    |        return &Product{name, category, price -10, supplier}
    |    }
    |    func main() {
    |        acme := &Supplier { "Acme Co", "New York"}
    |        products := [2]*Product {
    |            newProduct("Kayak", "Watersports", 275, acme),
    |            newProduct("Hat", "Skiing", 42.50, acme),
    |        }
    |        for _, p := range products {
    |            fmt.Println("Name:", p.name, "Supplier:",
    |                p.Supplier.name, p.Supplier.city)
    |        }
    |    }
    |________________________________________________________________________________________


105.Pointer Field Copying
    Care must be taken when copying structs to consider the effect on pointer fields
    
    example:
    |    package main
    |    import "fmt"
    |    type Product struct {
    |        name, category string
    |        price float64
    |        *Supplier
    |    }
    |    type Supplier struct {
    |        name, city string
    |    }
    |    func newProduct(name, category string, price float64, supplier *Supplier) *Product {
    |        return &Product{name, category, price -10, supplier}
    |    }
    |    func main() {
    |        acme := &Supplier { "Acme Co", "New York"}
    |        p1 := newProduct("Kayak", "Watersports", 275, acme)
    |        p2 := *p1
    |        p1.name = "Original Kayak"
    |        p1.Supplier.name = "BoatCo"
    |        for _, p := range []Product { *p1, p2 } {
    |            fmt.Println("Name:", p.name, "Supplier:",
    |                p.Supplier.name, p.Supplier.city)
    |        }
    |    }
    |______________________________________________________________________________________________


106.Method
    What are they?
    Methods are functions that are invoked on a struct and have access to all of the
    fields defined by the value's type. Interfaces define sets of methods, which can be
    implemented by struct types.

    Why are they useful?
    These features allow types to be mixed and used through their common
    characteristics.

    How are they used?
    Methods are defined using the func keyword, but with the addition of a receiver.
    Interfaces are defined using the type and interface keywords.

    Are there any pitfalls or limitations?
    Careful use of pointers is important when creating methods, and care must be taken
    when using interfaces to avoid problems with the underlying dynamic types.

107.Defining and Using Method
    example:
    |    package main
    |    import "fmt"
    |    type Product struct {
    |        name, category string
    |        price float64
    |    }
    |    func newProduct(name, category string, price float64) *Product {
    |            return &Product{ name, category, price }
    |        }
    |        func (product *Product) printDetails() {
    |            fmt.Println("Name:", product.name, "Category:", product.category,
    |                "Price", product.price)
    |        }
    |    func main() {
    |        products := []*Product {
    |            newProduct("Kayak", "Watersports", 275),
    |            newProduct("Lifejacket", "Watersports", 48.95),
    |            newProduct("Soccer Ball", "Soccer", 19.50),
    |        }
    |        for _, p := range products {
    |            p.printDetails()
    |        }
    |    }
    |____________________________________________________________________________


108.Defining Method Parameters and Results
    example:
    |    package main
    |    import "fmt"
    |    type Product struct {
    |        name, category string
    |        price float64
    |    }
    |    func newProduct(name, category string, price float64) *Product {
    |            return &Product{ name, category, price }
    |        }
    |        func (product *Product) printDetails() {
    |            fmt.Println("Name:", product.name, "Category:", product.category,
    |                "Price", product.price)
    |        }
    |    func main() {
    |        products := []*Product {
    |            newProduct("Kayak", "Watersports", 275),
    |            newProduct("Lifejacket", "Watersports", 48.95),
    |            newProduct("Soccer Ball", "Soccer", 19.50),
    |        }
    |        for _, p := range products {
    |            p.calcTax(0.2, 100) //<-------------------------------here
    |        }
    |    }
    |_____________________________________________________________________________________

109.Defining and Using Interfaces
    One interface can enclose another, with the effect that types must implement all the methods defined
    by the enclosing and enclosed interfaces. Interfaces are simpler than structs, and there are no fields or
    method to promote. The result of composing interfaces is a union of the method defined by the enclosing
    and enclosed types.
    
    example:
    |    package main
    |    import "fmt"
    |    type Expense interface {
    |        getName() string
    |        getCost(annual bool) float64
    |    }
    |    func main() {
    |        expenses := []Expense {
    |            Product { "Kayak", "Watersports", 275 },
    |            Service {"Boat Cover", 12, 89.50 },
    |        }
    |        for _, expense := range expenses {
    |            fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
    |        }
    |    }
    |______________________________________________________________________

110.an Interface in a Function
    example:
    |    package main
    |    import "fmt"
    |    type Expense interface {
    |        getName() string
    |        getCost(annual bool) float64
    |    }
    |    func calcTotal(expenses []Expense) (total float64) {
    |        for _, item := range expenses {
    |            total += item.getCost(true)
    |        }
    |        return
    |    }
    |    func main() {
    |        expenses := []Expense {
    |            Product { "Kayak", "Watersports", 275 },
    |            Service {"Boat Cover", 12, 89.50 },
    |        }
    |        for _, expense := range expenses {
    |            fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
    |        }
    |        fmt.Println("Total:", calcTotal(expenses))
    |    }
    |___________________________________________________________________

111.an Interface for Struct Fields
    example:
    |    package main
    |    import "fmt"
    |    type Expense interface {
    |        getName() string
    |        getCost(annual bool) float64
    |    }
    |    func calcTotal(expenses []Expense) (total float64) {
    |        for _, item := range expenses {
    |            total += item.getCost(true)
    |        }
    |        return
    |    }
    |    type Account struct {
    |        accountNumber int
    |        expenses []Expense
    |    }
    |    func main() {
    |        account := Account {
    |            accountNumber: 12345,
    |            expenses: []Expense {
    |                Product { "Kayak", "Watersports", 275 },
    |                Service {"Boat Cover", 12, 89.50 },
    |            },
    |        }
    |        for _, expense := range account.expenses {
    |            fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
    |        }
    |        fmt.Println("Total:", calcTotal(account.expenses))
    |    }
    |___________________________________________________________________________

112.Comparing Interface Values
    Care must be taken when comparing interface values, and inevitably, some knowledge of the dynamic
    types is required.
    The first two Expense values are not equal. 
    That's because the dynamic type for these values is a pointer
    type, and pointers are equal only if they point to the same memory location
    
    example:
    |    package main
    |    import "fmt"
    |    type Expense interface {
    |        getName() string
    |        getCost(annual bool) float64
    |    }
    |    func main() {
    |        var e1 Expense = &Product { name: "Kayak" }
    |        var e2 Expense = &Product { name: "Kayak" }
    |        var e3 Expense = Service { description: "Boat Cover" }
    |        var e4 Expense = Service { description: "Boat Cover" }
    |        fmt.Println("e1 == e2", e1 == e2)
    |        fmt.Println("e3 == e4", e3 == e4)
    |    }
    |_____________________________________________________________________

    Output:
    |    e1 == e2 false
    |    e3 == e4 true
    |_________________________________________________


113.Empty Interface
    Go allows the user of the empty interface—which means an interface that defines no methods—to represent
    any type, which can be a useful way to group disparate types that share no common features
    The empty interface represents all types, including the built-in types and any structs and interfaces
    that have been defined.

114.Empty Interface for Function Parameters
    The empty interface can be used as the type for a function parameter, allowing a function to be called with
    any value
    
    example:
    |    package main
    |    import "fmt"
    |    type Expense interface {
    |        getName() string
    |        getCost(annual bool) float64
    |    }
    |    type Person struct {
    |        name, city string
    |    }
    |    func processItem(item interface{}) {
    |            switch value := item.(type) {
    |                case Product:
    |                    fmt.Println("Product:", value.name, "Price:", value.price)
    |                case *Product:
    |                    fmt.Println("Product Pointer:", value.name, "Price:", value.price)
    |                case Service:
    |                    fmt.Println("Service:", value.description, "Price:",
    |                        value.monthlyFee * float64(value.durationMonths))
    |                case Person:
    |                    fmt.Println("Person:", value.name, "City:", value.city)
    |                case *Person:
    |                    fmt.Println("Person Pointer:", value.name, "City:", value.city)
    |                case string, bool, int:
    |                    fmt.Println("Built-in type:", value)
    |                default:
    |                    fmt.Println("Default:", value)
    |            }
    |        }
    |    func main() {
    |        var expense Expense = &Product { "Kayak", "Watersports", 275 }
    |        data := []interface{} {
    |            expense,
    |            Product { "Lifejacket", "Watersports", 48.95 },
    |            Service {"Boat Cover", 12, 89.50, []string{} },
    |            Person { "Alice", "London"},
    |            &Person { "Bob", "New York"},
    |            "This is a string",
    |            100,
    |            true,
    |        }
    |        for _, item := range data {
    |            processItem(item)
    |        }
    |    }
    |_________________________________________________________________________________

115.Package
    Packages are the Go feature that allows projects to be structured so that related functionality can be grouped
    together, without the need to put all the code into a single file or folder.
    
    What are they?
    Packages allow projects to be structured so that related features can be developed together.

    Why are they useful?
    Packages are how Go implements access controls so that the implementation of
    a feature can be hidden from the code that consumes it.
    
    How are they used?
    Packages are defined by creating code files in folders and using the package
    keyword to denote which package they belong to.

116.the Module File
    This name is important because it is used to import features from other packages created within
    the same project and third-party packages
    The go statement specifies the version of Go that is used.

117.Package Access Control
    Go has an unusual approach to access control. Instead of relying on dedicated keywords, like public
    and private, Go examines the first letter of the names given to the features in a code file, such as types,
    functions, and methods. If the first letter is lowercase, then the feature can be used only within the package
    that defines it. Features are exported for use outside of the package by giving them an uppercase first letter.

    The access control rules do not apply to individual function or method parameters, which means that
    the NewProduct function has to have an uppercase first character to be exported, but the parameter names
    can be lowercase.


118.Adding Code Files to Packages
    Packages can contain multiple code files, and to simplify development, access control rules and package
    prefixes do not apply when accessing features defined in the same package.

119.func init()
    Each code file can contain an initialization function that is executed only when all packages have been
    loaded and all other initialization—such as defining constants and variables—has been done. The most
    common use for initialization functions is to perform calculations that are difficult to perform or that require
    duplication to perform

    The initialization function is called init, and it is defined without parameters and a result. The init
    function is called automatically and provides an opportunity to prepare the package for use.

    The init function is not a regular Go function and cannot be invoked directly. And, unlike regular
    functions, a single file can define multiple init functions, all of which will be executed.

    AVOIDING THE MULTIPLE INITIALIZATION FUNCTION PITFALL
    Each code file can have its own initialization function. When using the standard Go compiler, the
    initialization functions are executed based on the alphabetic order of the filenames, so the function in
    the a.go file will be executed before the function in the b.go file, and so on.
    But this order is not part of the Go language specification and should not be relied on. Your initialization
    functions should be self-contained and not rely on other init functions having been invoked previously.

120.Creating Nested Packages
    Packages can be defined within other packages, making it easy to break up complex features into as many
    units as possible.
    The package statement is used just as with any other package, without the need to include the name
    of the parent or enclosing package. And dependency on custom packages must include the full package
    path. 

    example:
    |    Create the packages/store/cart folder 
    |    and add to it a file named cart.go with the contents:
    |    #1
    |    contents:
    |        package cart
    |        import "packages/store"
    |        type Cart struct {
    |            CustomerName string
    |            Products []store.Product
    |        }
    |        func (cart *Cart) GetTotal() (total float64) {
    |            for _, p := range cart.Products {
    |                total += p.Price()
    |            }
    |            return
    |        }
    |_______________________________________________________________________________________________

    The features defined by the nested package are accessed using the package name, just like any other
    package. When importing a nested package, the package path starts with the module name and lists the
    sequence of packages.
    example:
        package main
        import (
            "fmt"
            "packages/store"
            . "packages/fmt"
            "packages/store/cart"
        )
        func main() {
            product := store.NewProduct("Kayak", "Watersports", 279)
            cart := cart.Cart {
                CustomerName: "Alice",
                Products: []store.Product{ *product },
            }
            fmt.Println("Name:", cart.CustomerName)
            fmt.Println("Total:",  ToCurrency(cart.GetTotal()))
        }


121.Initialization Function
    example:
        func init() {
                for category, price := range categoryMaxPrices {
                    categoryMaxPrices[category] = price + (price * defaultTaxRate)
                }
        }

122.Importing a Package Only for Initialization Effects
    Go prevents packages from being imported but not used, 
    which can be a problem if you rely on the effect of
    an initialization function but don't need to use any of the features the package exports.
    If I need the effect of the initialization function, 
    but I don't need to use the GetData function the package exports.

    example:
        package main
        import (
            "fmt"
            "packages/store"
            . "packages/fmt"
            "packages/store/cart"
            _ "packages/data"
        )

123.Finding Go Packages
    #1 https://pkg.go.dev
    #2 https://github.com/golang/go/wiki/Projects

    Many Go modules are written by individual developers
    to solve a problem and then published for anyone else to use. 
    This creates a rich module ecosystem,
    but it does mean that maintenance and support can be inconsistent.

124.indirect
    The indirect comment at the end of the statements is added automatically because
    the packages are not used by the code in the project. A file named go.sum is created when the module is
    obtained and contains checksums used to validate the packages.

    example:
        module packages
        go 1.17
        require (
            github.com/fatih/color v1.10.0 // indirect
            github.com/mattn/go-colorable v0.1.8 // indirect
            github.com/mattn/go-isatty v0.0.12 // indirect
            golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
        )

    Note:
        You can also use the go.mod file to create dependencies on projects you have created locally

    The first time you run the go get command, you will see a list of the modules that are
    downloaded, which illustrated that modules have their own dependencies and that these are resolved
    automatically:
        go: downloading github.com/fatih/color v1.10.0
        go: downloading github.com/mattn/go-isatty v0.0.12
        go: downloading github.com/mattn/go-colorable v0.1.8
        go: downloading golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae


125.Managing External Packages
    Removing a Package
    To update the go.mod file to reflect the change, run the command:
        go mod tidy


126.Putting Type and Interface Composition in Context
    What is it?
    Composition is the process by which new types are created by combining
    structs and interfaces.

    Why is it useful?
    Composition allows types to be defined based on existing types.

    How is it used?
    Existing types are embedded in new types.
    
    Are there any pitfalls or limitations?
    Composition doesn't work in the same way as inheritance, and care must be
    taken to achieve the desired outcome.

    Are there any alternatives? 
    Composition is optional, and you can create entirely independent types.


    Go doesn't support classes or inheritance and focuses on composition instead. But, despite the
    differences, composition can be used to create hierarchies of types, just in a different way.

    Steps:
        1-Defining the Base Type -> struct - method 
        2-Defining a Constructor -> for create struct
        example:
        // struct:
        type Product struct {
                Name, Category string
                price float64
        }
        // method:
        func (p *Product) Price(taxRate float64) float64 {
            return p.price + (p.price * taxRate)
        }
        // Constructor:
        func NewProduct(name, category string, price float64) *Product {
            return &Product{ name, category, price }
        }
        
127.Creating Struct Values in Packages
    example:
    package main
    import (
        "fmt"
        "composition/store"
    )
    func main() {
        
        // use Constructor:
        kayak := store.NewProduct("Kayak", "Watersports", 275)

        // use the literal syntax:
        lifejacket := &store.Product{ Name: "Lifejacket", Category:  "Watersports"}
        for _, p := range []*store.Product { kayak, lifejacket} {
            fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
        }
    }
    

128.Steps of composition
    
    1-Defining the Base Type -> struct - method 
    2-Defining a Constructor -> for create struct
    3-Composing Types
    4-Using the Boat Struct in the main.go


    A struct can mix regular and embedded field types, but the embedded fields are an important part of
    the composition feature

    Go gives special treatment to struct types that have fields whose type is another struct type, in the way
    that the Boat type has a *Product field in the example project. You can see this special treatment in the
    statement in the for loop, which is responsible for writing out details of each Boat.
    Go allows the fields of the nested type to be accessed in two ways. The first is the conventional approach
    of navigating the hierarchy of types to reach the value that is required. The *Product field is embedded, which
    means that its name its its type.

    The Boat type doesn't define a Name field, but it can be treated as though it did because of the direct
    access feature. This is known as field promotion, and Go essentially flattens the types so that the Boat
    type behaves as though it defines the fields that are provided by the nested Product type

    example:
    product.go:
    |    // struct:
    |    type Product struct {
    |            Name, Category string
    |            price float64
    |    }
    |    // method:
    |    func (p *Product) Price(taxRate float64) float64 {
    |        return p.price + (p.price * taxRate)
    |    }
    |    // Constructor:
    |    func NewProduct(name, category string, price float64) *Product {
    |        return &Product{ name, category, price }
    |    }
    |________________________________

    boat.go:
    |   package store
    |
    |   The Boat struct type defines an embedded *Product field
    |   type Boat struct {
    |       *Product       -------------------> Embedded Type
    |       Capacity int   -----------------> Reguler Fields
    |       Motorized bool ________________|
    |   }
    |
    |   func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
    |           return &Boat {
    |               NewProduct(name, "Watersports", price), capacity, motorized,
    |           }
    |       }
    |_________________________________

    main.go:
    |    package main
    |    import (
    |        "fmt"
    |        "composition/store"
    |    )
    |
    |    func main() {
    |       boats := []*store.Boat {
    |           store.NewBoat("Kayak", 275, 1, false),
    |           store.NewBoat("Canoe", 400, 3, false),
    |           store.NewBoat("Tender", 650.25, 2, true),
    |       }
    |       for _, b := range boats {
    |            fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name)
    |        }
    |    }
    |
    |_____________________________________________________________________
    
    Output:
    |    Conventional: Kayak Direct: Kayak
    |    Conventional: Canoe Direct: Canoe
    |    Conventional: Tender Direct: Tender
    |___________________________________________


    If the field type is a value, such as Product, then any methods defined with Product or *Product
    receivers will be promoted. If the field type is a pointer, such as *Product, then only methods with *Product
    receivers will be prompted.
    There is no Price method defined for the *Boat type, but Go promotes the method defined with a
    *Product receiver.

    Calling a Method in the main.go:
    |    package main
    |    import (
    |        "fmt"
    |        "composition/store"
    |    )
    |    func main() {
    |        boats := []*store.Boat {
    |            store.NewBoat("Kayak", 275, 1, false),
    |            store.NewBoat("Canoe", 400, 3, false),
    |            store.NewBoat("Tender", 650.25, 2, true),
    |        }
    |        for _, b := range boats {
    |            fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
    |        }
    |    }
    |____________________________________________________________________

    Output:
    |    Boat: Kayak Price: 330
    |    Boat: Canoe Price: 480
    |    Boat: Tender Price: 780.3
    |___________________________________________



129.Creating a Chain of Nested Types
    The composition feature can be used to create complex chains of nested types, 
    whose fields and methods are promoted to the top-level enclosing type.

    Go performs promotion so that the fields defined by all three types in the chain can be
    accessed directly.

    Go promotes fields from the nested Boat and 
    Product types so they can be accessed through the top-
    level RentalBoat type, which allows the Name field to be read.
    Methods are also promoted to the top-level type, 
    which is why I can use the Price method, even though it is defined on the *Product type,
    which is at the end of the chain.

    example:
    rentalboats.go:
    |    package store
    |
    |    type RentalBoat struct {
    |       *Boat
    |       IncludeCrew bool
    |    }
    |
    |    func NewRentalBoat(name string, price float64, capacity int,
    |          motorized, crewed bool) *RentalBoat {
    |       return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed}
    |    }
    |__________________________________________________________

    
    Accessing Nested Fields Directly in the main.go:
    |    package main
    |    import (
    |        "fmt"
    |        "composition/store"
    |    )
    |    func main() {
    |        rentals := []*store.RentalBoat {
    |            store.NewRentalBoat("Rubber Ring", 10, 1, false, false),
    |            store.NewRentalBoat("Yacht", 50000, 5, true, true),
    |            store.NewRentalBoat("Super Yacht", 100000, 15, true, true),
    |        }
    |        for _, r := range rentals {
    |            fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2))
    |        }
    |    }
    |__________________________________________________________________
        
    Output:
    |    Rental Boat: Rubber Ring Rental Price: 12
    |    Rental Boat: Yacht Rental Price: 60000
    |    Rental Boat: Super Yacht Rental Price: 120000
    |_______________________________________________________________

130.Multiple Nested Types in the Same Struct
    example:
    rentalboats.go:
    |    package store
    |    type Crew struct {
    |        Captain, FirstOfficer string
    |    }
    |    type RentalBoat struct {
    |        *Boat
    |        IncludeCrew bool
    |        *Crew
    |    }
    |    func NewRentalBoat(name string, price float64, capacity int,
    |            motorized, crewed bool, captain, firstOfficer string) *RentalBoat {
    |        return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed,
    |            &Crew{captain, firstOfficer}}
    |    }
    |________________________________________________________________

    Using Promoted Fields in the main.go:
    |    package main
    |    import (
    |       "fmt"
    |       "composition/store"
    |    )
    |    func main() {
    |       rentals := []*store.RentalBoat {
    |          store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
    |          store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
    |          store.NewRentalBoat("Super Yacht", 100000, 15, true, true,
    |             "Dora", "Charlie"),
    |       }
    |       for _, r := range rentals {
    |          fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2),
    |             "Captain:", r.Captain)
    |       }
    |    }
    |___________________________________________________________________

    Output:
    |    Rental Boat: Rubber Ring Rental Price: 12 Captain: N/A
    |    Rental Boat: Yacht Rental Price: 60000 Captain: Bob
    |    Rental Boat: Super Yacht Rental Price: 120000 Captain: Dora
    |___________________________________________________________________

131.When Promotion Cannot Be Performed
    example:
    specialdeal.go File in the store Folder:
    |    package store
    |    type SpecialDeal struct {
    |       Name string
    |       *Product
    |       price float64
    |    }
    |    func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
    |       return &SpecialDeal{ name, p, p.price - discount }
    |    }
    |    func (deal *SpecialDeal ) GetDetails() (string, float64, float64) {
    |       return deal.Name, deal.price, deal.Price(0)
    |    }
    |____________________________________________________________________

    Using a New Type in the main.go:
    |    package main
    |    import (
    |        "fmt"
    |        "composition/store"
    |    )
    |    func main() {
    |        product := store.NewProduct("Kayak", "Watersports", 279)
    |        deal := store.NewSpecialDeal("Weekend Special", product, 50)
    |        Name, price, Price := deal.GetDetails()
    |        fmt.Println("Name:", Name)
    |        fmt.Println("Price field:", price)
    |        fmt.Println("Price method:", Price)
    |    }
    |_____________________________________________________________________

    Output:
    |    Name: Weekend Special
    |    Price field: 229
    |    Price method: 279
    |______________________________________________

    is a more concise way of expressing this statement:
        return deal.Name, deal.price, deal.Product.Price(0)

    The new Price method stops Go from promoting the Product method 
    and produces the following result
    Defining a Method in the specialdeal.go:
    |    package store
    |    type SpecialDeal struct {
    |       Name string
    |       *Product
    |       price float64
    |    }
    |    func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
    |       return &SpecialDeal{ name, p, p.price - discount }
    |    }
    |    func (deal *SpecialDeal ) GetDetails() (string, float64, float64) {
    |       return deal.Name, deal.price, deal.Price(0)
    |    }
    |    func (deal *SpecialDeal) Price(taxRate float64) float64 {
    |       return deal.price
    |    }
    |_____________________________________________________
    
    Output:
    |    Name: Weekend Special
    |    Price field: 229
    |    Price method: 229
    |_________________________________

132.Promotion Ambiguity
    A related issue arises when two embedded fields use the same field or method names
    example:
    An Ambiguous Method in the main.go:
    |    package main
    |    import (
    |       "fmt"
    |       "composition/store"
    |    )
    |    func main() {
    |       kayak := store.NewProduct("Kayak", "Watersports", 279)
    |       type OfferBundle struct {
    |          *store.SpecialDeal
    |          *store.Product
    |       }
    |       bundle := OfferBundle {
    |          store.NewSpecialDeal("Weekend Special", kayak, 50),
    |          store.NewProduct("Lifrejacket", "Watersports", 48.95),
    |       }
    |       fmt.Println("Price:", bundle.Price(0))
    |    }
    |___________________________________________________

    Output:
    |    .\main.go:22:33: ambiguous selector bundle.Price    
    |________________________________________________________       

133.Composition and Interfaces
    example:
    Mixing Types in the main.go:
    |    package main
    |    import (
    |        "fmt"
    |        "composition/store"
    |    )
    |    func main() {
    |        products := map[string]*store.Product {
    |            "Kayak": store.NewBoat("Kayak", 279, 1, false),
    |            "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
    |        }
    |        for _, p := range products {
    |            fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
    |        }
    |    }
    |__________________________________________

    Output:
    |    .\main.go:11:9: cannot use store.NewBoat("Kayak", 279, 1, false) (type *store.Boat) as
    |    type *store.Product in map value
    |____________________________________________

134.Composition to Implement Interfaces
    Go takes promoted methods into account when determining whether a type conforms to an interface,
    which avoids the need to duplicate methods that are already present through an embedded field.

    add a file named forsale.go to the store folder:
    The Contents of the forsale.go:
    |    package store
    |    type ItemForSale interface {
    |        Price(taxRate float64) float64
    |    }
    |__________________________________________        

    Using an Interface in the main.go:
    |    package main
    |    import (
    |       "fmt"
    |       "composition/store"
    |    )
    |    func main() {
    |       products := map[string]store.ItemForSale {
    |          "Kayak": store.NewBoat("Kayak", 279, 1, false),
    |          "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
    |       }
    |       for key, p := range products {
    |          fmt.Println("Key:", key, "Price:", p.Price(0.2))
    |       }
    |    }
    |____________________________________________

    Output:
    |    Key: Kayak Price: 334.8
    |    Key: Ball Price: 23.4
    |____________________________________

135.the Type Switch Limitation
    
    example:
    main.go:
    |    package main
    |    import (
    |        "fmt"
    |        "composition/store"
    |    )
    |    func main() {
    |        products := map[string]store.ItemForSale {
    |            "Kayak": store.NewBoat("Kayak", 279, 1, false),
    |            "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
    |        }
    |        for key, p := range products {
    |            switch item := p.(type) {
    |                case *store.Product, *store.Boat:
    |                    fmt.Println("Name:", item.Name, "Category:", item.Category,
    |                        "Price:", item.Price(0.2))
    |                default:
    |                    fmt.Println("Key:", key, "Price:", p.Price(0.2))
    |            }
    |        }
    |    }
    |________________________________________________________________

    Output:
    |    .\main.go:21:42: item.Name undefined (type store.ItemForSale has no field or method Name)
    |    .\main.go:21:66: item.Category undefined (type store.ItemForSale has no field or method
    |    Category)
    |_____________________________________________________________________________

    Using Separate case Statements in the main.go:
    |    package main
    |    import (
    |       "fmt"
    |       "composition/store"
    |    )
    |    func main() {
    |       products := map[string]store.ItemForSale {
    |          "Kayak": store.NewBoat("Kayak", 279, 1, false),
    |          "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
    |       }
    |       for key, p := range products {
    |          switch item := p.(type) {
    |             case *store.Product:
    |                fmt.Println("Name:", item.Name, "Category:", item.Category,
    |                   "Price:", item.Price(0.2))
    |             case *store.Boat:
    |                fmt.Println("Name:", item.Name, "Category:", item.Category,
    |                   "Price:", item.Price(0.2))
    |             default:
    |                fmt.Println("Key:", key, "Price:", p.Price(0.2))
    |          }
    |       }
    |    }
    |__________________________________________________________________________________

    Output:
    |    Name: Kayak Category: Watersports Price: 334.8
    |    Name: Soccer Ball Category: Soccer Price: 23.4
    |__________________________________________________________________

136.the Type Switch Limitation An alternative solution
    example:
    Defining an Interface in the product.go:
    |    package store
    |    type Product struct {
    |        Name, Category string
    |        price float64
    |    }
    |    func NewProduct(name, category string, price float64) *Product {
    |        return &Product{ name, category, price }
    |    }
    |    func (p *Product) Price(taxRate float64) float64 {
    |        return p.price + (p.price * taxRate)
    |    }
    |    type Describable interface  {
    |        GetName() string
    |        GetCategory() string
    |    }
    |    func (p *Product) GetName() string {
    |        return p.Name
    |    }
    |    func (p *Product) GetCategory() string {
    |        return p.Category
    |    }
    |__________________________________________________________________

    Using Interfaces in the main.go:
    |    package main
    |    import (
    |       "fmt"
    |       "composition/store"
    |    )
    |    func main() {
    |       products := map[string]store.ItemForSale {
    |          "Kayak": store.NewBoat("Kayak", 279, 1, false),
    |          "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
    |       }
    |       for key, p := range products {
    |          switch item := p.(type) {
    |             case store.Describable:
    |                fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(),
    |                   "Price:", item.(store.ItemForSale).Price(0.2))
    |             default:
    |                fmt.Println("Key:", key, "Price:", p.Price(0.2))
    |          }
    |       }
    |    }
    |____________________________________________________________________

    Output:
    |    Name: Kayak Category: Watersports Price: 334.8
    |    Name: Soccer Ball Category: Soccer Price: 23.4
    |_________________________________________________________________

137.Composing Interfaces
    Go allows interfaces to be composed from other interfaces
    
    example:
    Composing an Interface in the product.go:
    |    package store
    |    type Product struct {
    |        Name, Category string
    |        price float64
    |    }
    |    func NewProduct(name, category string, price float64) *Product {
    |        return &Product{ name, category, price }
    |    }
    |    func (p *Product) Price(taxRate float64) float64 {
    |        return p.price + (p.price * taxRate)
    |    }
    |    type Describable interface  {
    |        GetName() string
    |        GetCategory() string
    |        ItemForSale
    |    }
    |    func (p *Product) GetName() string {
    |        return p.Name
    |    }
    |    func (p *Product) GetCategory() string {
    |        return p.Category
    |    }
    |_______________________________________________________________

138.Goroutines and Channels
    What are they?
    Goroutines are lightweight threads created and managed by the Go runtime.
    Channels are pipes that carry values of a specific type.
    
    Why are they useful?
    Goroutines allow functions to be executed concurrently, without needing to deal
    with the complications of operating system threads. Channels allow goroutines to
    produce results asynchronously.

    How are they used?
    Goroutines are created using the go keyword. Channels are defined as data types.

    Are there any or limitations?
    pitfalls Care must be taken to manage the direction of channels. 
    Goroutines that share data require additional features.
    
    Are there any alternatives?
    Goroutines and channels are the built-in Go concurrency features, but some
    applications can rely on a single thread of execution, which is created by default to
    execute the main function.
    
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received
    

    Problem                                 Solution
    ---------------------------------       -------------------------
    Execute a function asynchronously       Create a goroutine
    
    Produce a result from a function        Use a channel 
    executed asynchronously

    Send and receive values                 Use arrow expressions 
    using a channel

    Indicate that no further values         Use the close function
    will be sent over a channel

    Enumerate the values received           Use a for loop with the range keyword 
    from a channel




139.How Go Executes Code
    All Go programs use at least one goroutine because this is how Go executes the code in the
    main function.

    When compiled Go code is executed, the runtime creates a goroutine that starts executing
    the statements in the entry point, which is the main function in the main package. 
    Each statement in the main function is executed in the order in which they are defined. 
    The goroutine keeps executing statements until
    it reaches the end of the main function, at which point the application terminates.

    The goroutine executes each statement in the main function synchronously, 
    which means that it waits
    for the statement to complete before moving on to the next statement. 
    The statements in the main function

140.Creating Additional Goroutines
    Go allows the developer to create additional goroutines, 
    which execute code at the same time as the main
    goroutine. Go makes it easy to create new goroutines

    example:
        package main
        import (
            "fmt"
            "time
        )

        func main(){
            fmt.Println("first statement")
            fmt.Println("second statement")
            time.Sleep(time.Second * 2)

            go fmt.Println("first statement")
            
            fmt.Println("first statement")

        }

141.Returning Results from Goroutines
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received

    Getting a result from a function that is being executed asynchronously 
    can be complicated because it requires coordination between the goroutine 
    that produces the result and the goroutine that consumes the result.
    To address this issue, Go provides channels, which are 
    conduits through which data can be sent and received.

    Defining a Channel:
    example:
        var channel chan float64 = make(chan float64)


142.Sending a Result Using a Channel
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received

    example:
        resultChannel <- total

143.Receiving a Result Using a Channel
    Receiving from a channel is a blocking operation, 
    meaning that execution will not continue until a value
    has been received
    
    example:
        storeTotal += <- channel

144.using adapters to execute functions asynchronously
    It isn't always possible to rewrite existing functions or methods to use channels, but it is a simple matter
    to execute synchronous functions asynchronously in a wrapper, like this:
    The syntax is a little awkward because the arguments used to invoke the function are expressed
    immediately following the function definition. But the result is the same, which is that a synchronous
    function can be executed by a goroutine with the result being sent through a channel.

    example:
    |    ...
    |    calcTax := func(price float64) float64 {
    |        return price + (price * 0.2)
    |    }
    |    wrapper := func (price float64, c chan float64)  {
    |        c <- calcTax(price)
    |    }
    |    resultChannel := make(chan float64)
    |    go wrapper(275, resultChannel)
    |    result := <- resultChannel
    |    fmt.Println("Result:", result)
    |    ...
    |    The wrapper function receives a channel, which it uses to send the value received from executing the
    |    calcTax function synchronously. This can be expressed more concisely by defining a function without
    |    assigning it to a variable, like this:
    |    ...
    |    go func (price float64, c chan float64) {
    |        c <- calcTax(price)
    |    }(275, resultChannel)
    |    ...
    |_____________________________________________________________________________


145.Coordinating Channels
    By default, sending and receiving through a channel are blocking operations. This means a goroutine that
    sends a value will not execute any further statements until another goroutine receives the value from the
    channel. If a second goroutine sends a value, it will be blocked until the channel is cleared, causing a queue
    of goroutines waiting for values to be received. This happens in the other direction, too, so that goroutines
    that receive values will block until another goroutine sends one.

    example:
    |    If Bob has a message for Alice, the default channel behavior requires 
    |    Alice and Bob to agree on a meeting place, and
    |    whoever gets there first will wait for the other to arrive. 
    |    Bob will only give the message to Alice when they are
    |    both present. When Charlie also has a message for Alice, he will form a queue behind Bob. 
    |    Everyone waits patiently, messages are transferred only when the sender 
    |    and receiver are both available, and messages are
    |    processed sequentially.
    |_________________________________________________________________________________________

146.Buffered Channel
    The default channel behavior can lead to bursts of activity as goroutines do their work, followed by a long
    idle period waiting for messages to be received. This doesn't have an impact on the example application
    because the goroutines finish once their messages are received, but in a real project goroutines often have
    repetitive tasks to perform, and waiting for a receiver can cause a performance bottleneck.

    An alternative approach is to create a channel with a buffer, 
    which is used to accept values from a
    sender and store them until a receiver becomes available. 
    This makes sending a message a nonblocking
    operation, allowing a sender to pass its value to the channel and continue working without having to wait
    for a receiver. 
    
    example:
    |    This is similar to Alice having an inbox on her desk. Senders come to Alice's office and put
    |    their message into the inbox, leaving it for Alice to read when she is ready. But, if the inbox is full, then they
    |    will have to wait until she has processed some of her backlog before sending a new message.
    |_____________________________________________________________________________________________________

    Creating a Buffered Channel
    example:
    |    var channel chan float64 = make(chan float64, 2)
    |___________________________________________________________

    Result explain:
    |    For this example, I have set the size of the buffer to 2, meaning that two senders will be able to send
    |    values through the channel without having to wait for them to be received.
    |____________________________________________________________________



147.Inspecting a Channel Buffer














































































































    package main = first executable file main.go
func main() {} = every execute files in "package main"
    fmt = for printing in CLI
import = for importing another package
    string
e.g
    Format String like: fmt.Printf("%v %s %f", variable, string, float32)
    Scope
    Go is lexically scoped using blocks.

    Defining Multiple Variables
    fmt.Scanf("%f", &input)
    We use another function from the fmt package to readthe   user   input   (Scanf).
    for 
    loop
    if
    else
    switch
    default
    array
    An array is a numbered sequence of elements of a sin-gle type with a fixed length.
    target
    value
    range
    len(x)
    length

27.
    slice
    A slice is a segment of an array. Like arrays slices areindexable and have a length. 
    Unlike arrays this lengthis allowed to change.

28.
    make
    f you want to create a slice you should use the built-inmake function

29.
    slice functions:
    append
    add to slice
    [low : high]
    Another way to create slices is to use the [low : high]expression

30.
    map
    A map is an unordered collection of key-value pairs.
    Also known as an associative array, a hash table or dictionary, 
    maps are used to look up a value by its associated key.

31.
    slices maps 

32.
    delete function
    We can also delete items from a map using the built-indelete function.

33.
    make map

34.
    function
    A function is an independent section of code 
    that mapszero or more input parameters to zero or more outputparameters.   
    Functions (also known as procedures orsubroutines) 
    are often represented as a black box: (theblack box represents the function)
    func FunctionName(ParameterName ParameterType) ReturnType { body of function}
    e.g
    func getUserName(UserInput string) string { return UserInput }

35.
    returning multiple values
    Go is also capable of returning multiple values from afunction

36.
    variadic functions
    func add(args ...int) int {}

37.
    closure
    It is possible to create functions inside of functions

38.
    increment adds 1 to the variable x which is defined inthe main function's scope. 
    
39.
    return function from a function

40.
    recursion
    Finally a function is able to call itself. 
    Here is one way to compute the factorial of a number

41.
    defer
    Go has a special statement called defer which schedules 
    a function call to be run after the function completes.
    is often used when resources need to be freed insome way. 
    For example when we open a file we need to make sure to close it later. With defer:

42.
    panic
    the panicfunction to cause a run time error.  

43.
    recover
    We can handle arun-time panic with the built-in recover function.
    recover stops the panic and returns the value that waspassed to the call to panic.

44.
    Pointers    
    * &

45.
    new
    Another way to get a pointer is to use the built-in newfunction.

46.
    type
    In the provided Go code, "type" is a keyword used 
    to define a new named type or alias. 
    It allows you to create custom types with specific behaviors and characteristics.
    By using the "type" keyword, 
    you can create your own named types or aliases that can be used throughout your code, 
    providing better abstraction and type safety.

47.
    struct
    A struct can mix regular and embedded field types
    A struct is a type which contains named fields.
    the keyword structto indicate that we are defining 
    a struct type and alist of fields inside of curly braces.
    Each field has a name and a type. 
    For a struct zero means each of the fields is set to their correspond-ing zero value:
    (0 for ints, 0.0 for floats, "" for strings,nil for pointers, ...)
48.
    new function struct
    e.g:
    c := new(Circle)
    This allocates memory for all the fields,
    sets each of them to their zero value and returns a pointer.
    (*Circle) More often we want to give each of the fieldsa value.
49.
    method
    function receiver VS. function parameter
    The receiver can be of two types: 
    value receiver (func (t T) methodName()) 
    or pointer receiver (func (t *T) methodName()), 
    depending on whether the method modifies the receiver or not. 
50.
    interface
51.
    Concurrency
52.
    Goroutines
53.
    Channels
    chan
54.
    Channel Direction
55.
    Select
    case
56.
    goto
57.
    buffered channels
58.
    packages
    creating packages
59.
    testing
60.
    the core packages
61.
    iput / output
    Go's io package.
62.
    files & folders
    To open a file in Go use the Open function from the os package.
63.
    errors
64.
    containers & sort
    Go has several more collections available underneath the container package.
65.
    hashes & cryptography
    Hash functions in Go are broken into two categories:
    cryptographic and non-crypto-graphic.
    The non-cryptographic hash functions 
    can be found underneath the hash package and include adler32, crc32, crc64 and fnv.
66.
    servers
67.
    HTTP servers
    RPC
    The  net/rpc (remote procedure call) 
    and net/rpc/jsonrpc packages provide an easy way to expose methods 
    so they can be invoked over a network. (rather than just in the program running them)
68.
    parsing command line arguments
69.
    synchronization primitives
70.
    Mutexes
    A mutex (mutal exclusive lock) locks a section of code 
    to a single thread at a time and is used to protect 
    shared resources from non-atomic operations.
71.
    Collections
72.
    function recover()
73.
    function
        min(1,2)
        max(1,2)
74.
    Currying
    Function currying is the practice of writing a function 
    that takes a function (or functions) as input, and returns a new function.
75.
    Replace(s)
    This method returns a string for which all the replacements specified with the
    constructor have been performed on the string s.
76.
    WriteString(writer, s)	
    This method is used to perform the replacements specified with the constructor
    and write the results to an io.Writer

77.
    TrimSpace(s)			
    This function returns the string s without leading or trailing whitespace
    characters.
78.
    Trim(s, set)
    This function returns a string from which any leading or trailing characters
    contained in the string set are removed from the string s.
79.
    TrimLeft(s, set)
    This function returns the string s without any leading character contained
    in the string set. This function matches any of the specified characters—use
    the TrimPrefix function to remove a complete substring.
80.
    TrimRight(s, set)
    This function returns the string s without any trailing character contained
    in the string set. This function matches any of the specified characters—use
    the TrimSuffix function to remove a complete substring.
81.
    TrimPrefix(s, prefix)
    function returns the string s after removing the specified prefix string.
    This function removes the complete prefix string—use the TrimLeft
    function to remove characters from a set.
82.
    TrimSuffix(s, suffix)
    function returns the string s after removing the specified suffix string.
    This function removes the complete suffix string—use the TrimRight
    function to remove characters from a set.
83.
    TrimFunc(s, func)
    This function returns the string s from which any leading or trailing
    character for which a custom function returns true are removed.
84.
    TrimLeftFunc(s, func)
    function returns the string s from which any leading character for
    which a custom function returns true are removed.
85.
    TrimRightFunc(s, func)
    function returns the string s from which any trailing character for
    which a custom function returns true are removed.
86.
    Join(slice, sep)
    function combines the elements in the specified string slice, with the specified
    separator string placed between elements.
87.
    Repeat(s, count)
    function generates a string by repeating the string s for a specified number
    of times.
88.

89.

90.

91.

92.

93.

94.

95.

96.

97.

98.

99.

100.

101.

102.

103.

104.

105.

106.

107.

108.

109.

110.

111.

112.

113.

114.

115.
    `,
}

