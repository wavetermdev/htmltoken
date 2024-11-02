// Copyright 2024 Command Line Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package htmltoken

import "fmt"

func (z *Tokenizer) parseBraceAttr() {
	braceCount := 1
	inString := false
	prevStrBackslash := false

	z.pendingAttr[1].start = z.raw.end
	for {
		ch := z.readByte()
		if z.err != nil {
			z.pendingAttr[1].end = z.raw.end
			return
		}
		if inString {
			if prevStrBackslash {
				prevStrBackslash = false
				continue
			}
			if ch == '\\' {
				prevStrBackslash = true
				continue
			}
			if ch == '"' {
				inString = false
				continue
			}
			continue
		}
		if ch == '{' {
			braceCount++
			continue
		}
		if ch == '"' {
			inString = true
			continue
		}
		if ch == '}' {
			braceCount--
			if braceCount == 0 {
				z.pendingAttr[1].end = z.raw.end - 1
				return
			}
			continue
		}
	}
}

func (z *Tokenizer) parseBraceAttrEx(input string) (string, error) {
	var result []rune
	braceCount := 0
	inString := false

	for i := 0; i < len(input); i++ {
		ch := rune(input[i])

		if inString {
			// Handle string escape sequences
			if ch == '\\' && i+1 < len(input) {
				result = append(result, ch, rune(input[i+1]))
				i++
				continue
			}
			if ch == '"' {
				inString = false
			}
			result = append(result, ch)
			continue
		}

		switch ch {
		case '{':
			braceCount++
		case '}':
			braceCount--
			if braceCount == 0 {
				return string(result), nil
			}
		case '"':
			inString = true
		}

		result = append(result, ch)
	}

	if braceCount != 0 {
		return "", fmt.Errorf("unbalanced braces")
	}

	return string(result), nil
}
