package funcs

import (
	"errors"
	"strings"
)

func StoreInMap(inp string) map[rune][]string {
	inp = CleanFile(inp)
	mp := make(map[rune][]string)
	i := 1
	for r := 32; r < 127 && i < len(inp); r++ {
		oneline := ""
		for l := 0; l < 8; l++ {
			for inp[i] != '\n' {
				oneline = oneline + string(inp[i])
				i++
			}
			mp[rune(r)] = append(mp[rune(r)], oneline)
			i++
			oneline = ""
		}
		for i < len(inp) && inp[i] != '\n' {
			i++
		}
		i++
	}
	return mp
}

func GetWord(word string, mp map[rune][]string) string {
	res := ""
	for j := 0; j < 8; j++ {
		for ind := 0; ind < len(word); ind++ {
			if ind != len(word)-1 {
				res += mp[rune(word[ind])][j]
			} else {
				res += mp[rune(word[ind])][j] + "\n"
			}
		}
	}
	return res
}

func CleanFile(inp string) string {
	res := ""
	for _, let := range inp {
		if let != 13 {
			res += string(let)
		}
	}
	return res
}

func IsValid(word string, font string) error {
	for _, let := range word {
		if let < 32 || let > 126 {
			if let != 10 && let != 13 {
				err := errors.New("ERROR: the string is invalid (NOT ASCII RANGE)")
				return err
			}
		}
	}
	if font != "shadow" && font != "thinkertoy" && font != "standard" {
		err := errors.New("ERROR: not valid font or didn't choose font")
		return err
	}
	return nil
}

func SepNewLine(s []string) []string {
	var res []string
	for i := 0; i < len(s); i++ {
		res = append(res, strings.Split(s[i], string('\n'))...)
	}
	return res
}
