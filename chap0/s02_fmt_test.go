package chap0

import (
	"fmt"
	"testing"
)

// https://programming.guide/go/fmt-printf-reference-cheat-sheet.html
func TestGenericFormatting(t *testing.T) {
	val := []int{1, 2}
	fmt.Printf("Default format: %v \n", val)
	fmt.Printf("Go-synctax format: %#v \n", val)
	fmt.Printf("the type of the value: %T \n", val)
}

func TestIntegerFormatting(t *testing.T) {
	val := 109
	fmt.Printf("Base 10: %d \n", val)
	fmt.Printf("Always show sign: %+d \n", val)

	fmt.Printf("Width 4, right justify: %4d \n", val)
	fmt.Printf("Width 4, left justify: %-4d \n", val)
	fmt.Printf("Width 4, pad with zero: %04d \n", val)

	fmt.Printf("Character: %c \n", val)
	fmt.Printf("Quoted Character: %q \n", val)

	fmt.Printf("Base 2: %b \n", val)
	fmt.Printf("Base 8: %o \n", val)

	fmt.Printf("Base 16, lowercase: %x \n", val)
	fmt.Printf("Base 16, uppercase: %X \n", val)
	fmt.Printf("Base 16, with leading 0x: %#x \n", val)

	fmt.Printf("Unicode: %U\n", val)
	fmt.Printf("Unicode with character: %#U \n", val)
}

func TestBooleanFormatting(t *testing.T) {
	val := true
	fmt.Printf("boolean format: %t\n", val)
}

func TestPointerFormatting(t *testing.T) {
	val := 3
	fmt.Printf("Pointer format: %p\n", &val)
}

func TestFloatFormatting(t *testing.T) {
	val := 123.456
	fmt.Printf("Scientific notation: %e\n", val)

	fmt.Printf("Decimal point, no exponent: %f\n", val) // no exponent 没有指数
	fmt.Printf("Default width, precision 2: %.2f\n", val)
	fmt.Printf("Width8, precision 2: %8.2f\n", val)

	fmt.Printf("Exponent as needed, necessary digits only: %g\n", val)
}

func TestStringOrByteSliceFormatting(t *testing.T) {
	val := "café"
	fmt.Printf("Plain string: %s\n", val)
	fmt.Printf("Width 6, right justify: %6s\n", val)
	fmt.Printf("Width 6, left justify: %-6s\n", val)

	fmt.Printf("Quoted string: %q\n", val)

	fmt.Printf("Hex dump of byte values: %x\n", val)
	fmt.Printf("Hex dump with sapces: % x\n", val)
}

func TestSpecialValuesFormatting(t *testing.T) {
	fmt.Printf("U+0007 alert or bell: \a \n")
	fmt.Printf("U+0008 backspace: \b \n")
	fmt.Printf("U+000C form feed: \f \n")
	fmt.Printf("U+000A line feed or newline: \n \n")
	fmt.Printf("U+000D carriage return: \r \n")
	fmt.Printf("U+0009 horizontal tab: \t abc \n")
	fmt.Printf("U+000b vertical tab: \v abc \n")
	fmt.Printf("U+005c backslash: \\ \n")

	fmt.Println("====================")
	// backslash escapes
	// \x followed by exactly two hexadecimal digits;
	// \ followed by exactly three octal digits.
	// \u followed by exactly four hexadecimal digits;
	// \U followed by exactly eight hexadecimal digits;
	fmt.Println("\\caf\u00e9") // Output: \café
	fmt.Println("====================")
}
