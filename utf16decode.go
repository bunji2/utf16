package utf16

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf16"
	"unicode/utf8"
)

// ReadFileUTF16 reads file, decodes in utf16, and returns string.
func ReadFileUTF16(filePath string) (r string, bs []byte, err error) {

	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	//var bs []byte
	bs, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}

	r, err = DecodeUTF16(bs)
	//fmt.Println(r)
	return
}

// DecodeUTF16 decodes bytes in utf16.
func DecodeUTF16(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	if b[0] == 0xFF && b[1] == 0xFE {
		return decodeUTF16LE(b[2:])
	} else if b[0] == 0xFE && b[1] == 0xFF {
		return decodeUTF16BE(b[2:])
	}

	return "", fmt.Errorf("BOM must be 0xFFFE or 0xFEFF")
}

// decodeUTF16LE decodes bytes in utf16 little endian
func decodeUTF16LE(b []byte) (string, error) {

	u16s := make([]uint16, 1)
	ret := &bytes.Buffer{}
	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}

// decodeUTF16BE decodes bytes in utf16 big endian
func decodeUTF16BE(b []byte) (string, error) {

	u16s := make([]uint16, 1)
	ret := &bytes.Buffer{}
	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i+1]) + (uint16(b[i]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}
