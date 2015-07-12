// chris 052315

package swar

// AsciiUpper converts ASCII characters to their upper-case counterparts
// in place.  Non-alphabetic characters are converted too, such as @ and
// `.
//
// See:
// https://blog.cloudflare.com/the-oldest-trick-in-the-ascii-book/
//
//	     0123456789ABCDEF0123456789ABCDEF
//	    +--------------------------------
//	0x20| !"#$%&'()*+,-./0123456789:;<=>?
//	0x40|@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_
//	0x60|`abcdefghijklmnopqrstuvwxyz{|}~
//
// Note that the upper and lower-case characters are exactly 0x20 apart.
// So to convert from lower to upper case, you just need to mask out the
// 0x20 bit.  To convert from upper to lower case, you just need to add
// it in.
func AsciiUpper(s string) string {
	b := []byte(s)
	for i := range b {
		b[i] &= 0xdf
	}
	return string(b)
}

// AsciiLower converts ASCII characters to their lower-case counterparts
// in place.  Non-alphabetic characters are converted too, such as @ and
// `.
//
// See:
// https://blog.cloudflare.com/the-oldest-trick-in-the-ascii-book/
//
//	     0123456789ABCDEF0123456789ABCDEF
//	    +--------------------------------
//	0x20| !"#$%&'()*+,-./0123456789:;<=>?
//	0x40|@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_
//	0x60|`abcdefghijklmnopqrstuvwxyz{|}~
//
// Note that the upper and lower-case characters are exactly 0x20 apart.
// So to convert from lower to upper case, you just need to mask out the
// 0x20 bit.  To convert from upper to lower case, you just need to add
// it in.
func AsciiLower(s string) string {
	b := []byte(s)
	for i := range b {
		b[i] |= 0x20
	}
	return string(b)
}
