package ascii

import (
	"asciiweb/ascii/funcs"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func CheckFileHashing(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}

	sum1 := sha256.Sum(nil)
	sum := hex.EncodeToString(sum1)

	shadow := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoy := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	standard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"

	if fileName == "data/standard.txt" && string(sum) == standard {
		return true
	} else if fileName == "data/thinkertoy.txt" && string(sum) == thinkertoy {
		return true
	} else if fileName == "data/shadow.txt" && string(sum) == shadow {
		return true
	}
	return false
}

func Asciiart(font string, word string) (string, error) {
	res := ""
	if err := funcs.IsValid(word, font); err != nil {
		return "", err
	}
	file := "data/" + font + ".txt"

	if !CheckFileHashing(file) {
		err := errors.New("ERROR: the file was changed!")
		return "", err
	}

	// word = strings.ReplaceAll(word, "\r\n", "\n")
	arr := strings.Split(word, "\r\n") //"\n"
	arr = funcs.SepNewLine(arr)

	text, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	mp := funcs.StoreInMap(string(text))
	for i := 0; i < len(arr); i++ {
		if arr[i] == string("") {
			res += "\n"
		} else {
			res += funcs.GetWord(arr[i], mp)
		}
	}
	return res, nil
}
