package main

// LEEDCODE: 17
// Medium

// time: O(4^n), space: O(n)
func letterCombinations(digits string) []string {

	if len(digits) < 1 {
		return []string{}
	}

	letters := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	combinations := []string{}
	var dig func(int, string)

	dig = func(index int, temp string) {
		if len(temp) == len(digits) {
			combinations = append(combinations, temp)
			return
		}

		possibleLetters := letters[string(digits[index])]
		for _, letter := range possibleLetters {
			temp = temp + string(letter)
			dig(index+1, temp)
			temp = temp[:len(temp)-1]
		}
	}

	dig(0, "")

	return combinations
}

/*

python

def letterCombinations(self, digits):

	letters = {
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	def dig(index, temp):
		if len(temp) == len(digits):
				combinations.append(temp)
				return

				possible_letters = letter[digits[index]]
				for letter in possible_letters:
						temp = temp + letter
						dig(index+1, temp)
						temp = temp[:-1]

    combinations = []
	dig(0, "")
	return combinations
*/
