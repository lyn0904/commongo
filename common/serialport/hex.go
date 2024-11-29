package serialport

import "fmt"

const HEX = "0123456789ABCDEF"

// BytesToHeX bytes 转 hex
func BytesToHeX(data []byte) string {
	hexStr := ""
	for _, c := range data {
		hexStr += ByteToHeX(c)
	}
	return hexStr
}

// HexToString 将16进制字符串转换为[]byte
func HexToString(hexStr string) ([]byte, error) {
	if len(hexStr)%2 != 0 {
		return nil, fmt.Errorf("input string length must be even")
	}

	bytes := make([]byte, 0, len(hexStr)/2)
	for i := 0; i < len(hexStr); i += 2 {
		byte, err := HexToByte(hexStr[i : i+2])
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, byte)
	}

	return bytes, nil
}

func ByteToHeX(b byte) string {
	return string(HEX[b>>4]) + string(HEX[b&0x0F])
}

// HexToByte 将两个16进制字符转换为一个字节
func HexToByte(s string) (byte, error) {
	var b byte
	var err error

	for i, r := range s {
		if b, err = Find(r); err != nil {
			return 0, err
		}
		if i == 0 {
			b = b << 4
		}
	}

	return b, nil
}

// Find 返回字符在字符串中的位置，如果未找到则返回错误
func Find(c rune) (byte, error) {
	for i, char := range HEX {
		if char == c {
			return byte(i), nil
		}
	}
	return 0, fmt.Errorf("invalid hex character '%c'", c)
}
