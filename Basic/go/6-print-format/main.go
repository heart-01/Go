package main

import "fmt"

func main() {
	// Print command and data type checking
	// 1. fmt.Print() fmt.Println()
	fmt.Print("Hello, ")
	fmt.Println("World!")

	println("--------------------")

	// 3. fmt.Printf()
	fmt.Printf("Hello, %s\n", "World!") // %s format specifier - string
	fmt.Printf("Hello, %d\n", 100)      // %d format specifier - integer
	fmt.Printf("Hello, %f\n", 100.123)  // %f format specifier - float
	fmt.Printf("Hello, %t\n", true)     // %t format specifier - boolean
	fmt.Printf("Hello, %v\n", "World!") // %v format specifier - any value

	println("--------------------")

	// 4. fmt.Sprintf()
	_ = fmt.Sprint("Hello, ", "World!") // This command returns a string, but does not use the result.

	s := fmt.Sprintf("Hello, %s", "World!") // This command returns a string and uses the result.

	fmt.Println(s)

	println("--------------------")

	// รหัสของรูปแบบต่างๆ ของ format specifier
	// %v format specifier - any value
	// %T format specifier - type
	// %t format specifier - boolean
	// %d format specifier - integer
	// %b format specifier - binary
	// %o format specifier - octal
	// %x format specifier - hexadecimal
	// %X format specifier - hexadecimal uppercase
	// %f format specifier - float
	// %e format specifier - scientific notation
	// %E format specifier - scientific notation uppercase
	// %s format specifier - string
	// %q format specifier - quoted string
	// %p format specifier - pointer
	// %c format specifier - character
	// %U format specifier - Unicode
	// %v format specifier - any value

	// example
	fmt.Printf("%v\n", 100)             // 100
	fmt.Printf("%T\n", 100)             // int
	fmt.Printf("%t\n", true)            // true
	fmt.Printf("%d\n", 100)             // 100
	fmt.Printf("%b\n", 100)             // 1100100
	fmt.Printf("%o\n", 100)             // 144
	fmt.Printf("%x\n", 100)             // 64
	fmt.Printf("%X\n", 100)             // 64
	fmt.Printf("%f\n", 100.123)         // 100.123000
	fmt.Printf("%e\n", 100.123)         // 1.001230e+02
	fmt.Printf("%E\n", 100.123)         // 1.001230E+02
	fmt.Printf("%s\n", "Hello, World!") // Hello, World!
	fmt.Printf("%q\n", "Hello, World!") // "Hello, World!"
	fmt.Printf("%p\n", &s)              // 0xc0000b6010
	fmt.Printf("%c\n", 100)             // d
	fmt.Printf("%U\n", 100)             // U+0064

	println("--------------------")

	// การใช้งานรูปแบบการแสดงผลต่าง ๆ ด้วย Printf
	fmt.Printf("%d\n", 42)         // 42
	fmt.Printf("%+d\n", 42)        // +42
	fmt.Printf("%5d\n", 42)        //    42
	fmt.Printf("%-5d\n", 42)       // 42
	fmt.Printf("%05d\n", 42)       // 00042
	fmt.Printf("%b\n", 5)          // 101
	fmt.Printf("%o\n", 10)         // 12
	fmt.Printf("%#o\n", 10)        // 012
	fmt.Printf("%x\n", 255)        // ff
	fmt.Printf("%X\n", 255)        // FF
	fmt.Printf("%U\n", 'A')        // U+0041
	fmt.Printf("%#U\n", 'A')       // U+0041 'A'
	fmt.Printf("%c\n", 65)         // A
	fmt.Printf("%q\n", "Hello")    // "Hello"
	fmt.Printf("%t\n", true)       // true
	fmt.Printf("%f\n", 123.456)    // 123.456000
	fmt.Printf("%.2f\n", 123.456)  // 123.46
	fmt.Printf("%9.2f\n", 123.456) // 123.46
	fmt.Printf("%9.f\n", 123.456)  // 123
	fmt.Printf("%g\n", 123.456)    // 123.456
	fmt.Printf("%G\n", 123.456)    // 123.456
	fmt.Printf("%e\n", 123.456)    // 1.234560e+02
	fmt.Printf("%s\n", "Hello")    // Hello
	fmt.Printf("%10s\n", "Hello")  //      Hello
	fmt.Printf("%-10s\n", "Hello") // Hello
	fmt.Printf("%T\n", 123)        // int
	fmt.Printf("%v\n", 123)        // 123
	fmt.Printf("%%\n")
}
