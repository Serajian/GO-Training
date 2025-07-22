package main

import (
	"bufio"
	"fmt"
	"go-train/trains"
	"os"
	"strings"
)

// exerciseDescriptions maps exercise numbers to their descriptions
var exerciseDescriptions = map[int]string{
	// String-Based Questions
	1:  "Exercise 1: Reverse a given string — without using built-ins.",
	2:  "Exercise 2: Check if a string is a palindrome.",
	3:  "Exercise 3: Remove duplicates from a string — efficiently.",
	4:  "Exercise 4: Find the first non-repeating character — who stands alone?",
	5:  "Exercise 5: Count how many times each character appears.",
	6:  "Exercise 6: Flip the words in a sentence, not the letters.",
	7:  "Exercise 7: Are two strings anagrams? Prove it.",
	8:  "Exercise 8: Longest substring without repeats — sliding window style.",
	9:  "Exercise 9: Build your own atoi — string to integer.",
	10: "Exercise 10: Compress strings with run-length encoding.",
	11: "Exercise 11: Most frequent character — who dominates?",
	12: "Exercise 12: List all possible substrings of a string.",
	13: "Exercise 13: Is one string a rotation of another?",
	14: "Exercise 14: Strip all white spaces from a string.",
	15: "Exercise 15: Is this a valid shuffle of two strings?",
	16: "Exercise 16: Convert text to Title Case — properly.",
	17: "Exercise 17: Find the longest common prefix among words.",
	18: "Exercise 18: Break a string into a char array — without confusion.",
	19: "Exercise 19: Replace spaces with %20 — classic URL trick.",
	20: "Exercise 20: Turn full sentences into acronyms.",
	21: "Exercise 21: Check if the string is all digits — no alphabets allowed.",
	22: "Exercise 22: Count how many words are in the string.",
	23: "Exercise 23: Remove a specific character — cleanly.",
	24: "Exercise 24: Find the shortest word in a sentence.",
	25: "Exercise 25: Longest palindromic substring — two-pointer style.",

	// Array-Based Questions
	26: "Exercise 26: Reverse an array in-place.",
	27: "Exercise 27: Find the largest and smallest element.",
	28: "Exercise 28: Check for duplicates in an array.",
	29: "Exercise 29: Remove duplicates — return only unique values.",
	30: "Exercise 30: Find the missing number from 1 to N.",
	31: "Exercise 31: Move all zeros to the end — keep order.",
	32: "Exercise 32: Rotate the array left/right by K positions.",
	33: "Exercise 33: Find the Kth largest/smallest element.",
	34: "Exercise 34: Merge two sorted arrays — without using extra space.",
	35: "Exercise 35: Find the intersection of two arrays.",
	36: "Exercise 36: Sort 0s, 1s, and 2s without using sort().",
	37: "Exercise 37: Find subarrays with a given sum.",
	38: "Exercise 38: Detect if a subarray sums to 0.",
	39: "Exercise 39: Find the longest increasing subsequence.",
	40: "Exercise 40: Kadane’s Algorithm — maximum subarray sum.",
	41: "Exercise 41: Check if array is sorted and rotated.",
	42: "Exercise 42: Rearrange array in max-min order alternately.",
	43: "Exercise 43: Find leaders in an array (no greater element to the right).",
	44: "Exercise 44: Calculate frequency of all elements in O(n).",
	45: "Exercise 45: Product of all elements except self.",
}

func main() {
	fmt.Println("Welcome to GO training exercises! 🧑‍💻")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("🍻 Please enter the exercise number (1-45): ")

		var exerciseNum int
		_, err := fmt.Scan(&exerciseNum)
		if err != nil {
			fmt.Println("🚨 Error reading input:", err)
			os.Exit(1)
		}

		// Validate exercise number
		if _, exists := exerciseDescriptions[exerciseNum]; !exists {
			fmt.Println("🚨 Invalid exercise number! Must be between 1 and 45. 🚨")
			os.Exit(1)
		}

		// Display exercise description
		fmt.Println("\n🔖 Exercise description:")
		fmt.Println(exerciseDescriptions[exerciseNum])

		result := funcRunner(exerciseNum, scanner)

		// Print result
		fmt.Println("\n✨ Exercise result:")
		fmt.Println(result)

		// Ask if the user wants to continue
		fmt.Print("\n♻️ If you want to continue write (y/yes): ")
		var continueInput string
		_, _ = fmt.Scanln(&continueInput)
		continueInput = strings.ToLower(strings.TrimSpace(continueInput))
		if continueInput != "y" && continueInput != "yes" {
			fmt.Println("🚧  Exiting program. Goodbye! 🚧")
			break
		}
		fmt.Println("🚀 Let's GO...")
	}
}

func funcRunner(num int, scanner *bufio.Scanner) interface{} {
	var result interface{}
	switch num {
	case 1:
		fmt.Print("✏️ Please enter a string: ")
		scanner.Scan()
		input := scanner.Text()
		result = trains.ReverseString(input)

	case 2:
		fmt.Print("✏️ Please enter a string: ")
		scanner.Scan()
		input := scanner.Text()
		result = trains.IsPalindrome(input)

	case 3:
		fmt.Print("✏️ Please enter text: ")
		scanner.Scan()
		input := scanner.Text()
		result = trains.RemoveDuplicates(input)

	case 4:
		fmt.Print("✏️ Please enter a string: ")
		scanner.Scan()
		input := scanner.Text()
		result = trains.FirstNoneRepeating(input)

	case 5:
		fmt.Print("✏️ Please enter a string: ")
		scanner.Scan()
		input := scanner.Text()
		result = trains.HowManyAppears(input)
	case 6:
		fmt.Print("✏️ Please enter a text: ")
		scanner.Scan()
		input := scanner.Text()
		result = trains.FlipTheWords(input)
	}

	return result
}
