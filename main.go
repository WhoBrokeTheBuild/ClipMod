package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/atotto/clipboard"
	"github.com/kortschak/zalgo"
	"github.com/lithammer/dedent"
)

func main() {
	if len(os.Args) == 0 {
		return
	}
	mode := os.Args[1]

	buf, err := clipboard.ReadAll()
	if err != nil {
		return
	}

	if mode == "uppercase" {
		buf = strings.ToUpper(buf)
	} else if mode == "lowercase" {
		buf = strings.ToLower(buf)
	} else if mode == "spongebob" {
		up := false
		newBuf := []string{}
		for i := range buf {
			if !unicode.IsLetter(rune(buf[i])) {
				newBuf = append(newBuf, string(buf[i]))
			} else {
				if up {
					newBuf = append(newBuf, strings.ToUpper(string(buf[i])))
				} else {
					newBuf = append(newBuf, strings.ToLower(string(buf[i])))
				}
				up = !up
			}
		}
		buf = strings.Join(newBuf, "")
	} else if mode == "spaced" {
		buf = strings.Join(strings.Split(buf, ""), " ")
	} else if mode == "zalgo" {
		b := bytes.Buffer{}
		z := zalgo.NewCorrupter(io.Writer(&b))
		z.Up = complex(5, 0.3)
		z.Middle = complex(5, 0.1)
		z.Down = complex(5, 0.7)

		z.Write([]byte(buf))
		buf = string(b.Bytes())
	} else if mode == "code" {
		buf = "```\n" + dedent.Dedent(buf) + "\n```"
	}

	err = clipboard.WriteAll(buf)
	if err != nil {
		return
	}
}
