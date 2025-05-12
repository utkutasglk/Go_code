package intermediate

import "fmt"


func main() {

	// --- General Formatting Verbs
	// %v     Prints value in the default format
	// %#v    Prints value in Go-syntax format
	// %T     Prints type of value
	// %%     Prints the % sign

	i := 1505.5
	str := "Hello"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)


	fmt.Printf("%v\n", str)
	fmt.Printf("%T\n", str)
	fmt.Printf("%#v\n", str)

	// ----- Integer Formatting Verbs
	// %b     Base 2
	// %d     Base 10
	// %+d    Base 10 and always show sign
	// %o     Base 8 
	// %O     Base 8, with leading 0o
	// %x     Base 16, lowercase
	// %X     Base 16, uppercase
	// %#x    Base 16, with leading 0x
	// %4d    Pad with spaces 
	// %-4d   Pad with spaces 
	// %04d   Pad with zeroes

	int := 8 
	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int)
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	// ----- String formatted text
	// %s    Prints the value as plain string
	// %q    Prints the value as a double-quoted string
	// %8s   Prints the value as a plain string
	// %-8s  Prints the value as a plain string
	// %x    Prints the value as hex dump of byte values
	// % x   Prints the value as hex dump with spaces 

	txt := "World"

	fmt.Printf("%s\n",txt)
	fmt.Printf("%q\n",txt)
	fmt.Printf("%8s\n",txt)
	fmt.Printf("%-8s\n",txt)
	fmt.Printf("%x\n",txt)
	fmt.Printf("% x\n",txt)

	// ------ Boolean formatting verbs
	// %t    Value of boolean operator in true or false format  (same as using %v)

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", f)
	fmt.Printf("%t\n", f)

	// -------- Float formatting verbs
	// Verb description
	// %e    Scientific notation with 'e' as exponent
	// %f    Decimal point, no exponent
	// %.2f  Default width, precision 2
	// %6.2f width 6, precision 2
	// %g   Exponent as needed, only necessary digits 

	flt := 9.18

	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)







 



}