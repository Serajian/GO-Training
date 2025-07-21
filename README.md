GO-train
GO-train is a command-line application written in Go that provides a collection of programming exercises to practice string and array manipulation. It includes 45 exercises (25 string-based and 20 array-based) to help developers enhance their problem-solving skills in Go.
Table of Contents

Overview
Features
Project Structure
Installation
Usage
Exercises
String-Based Exercises
Array-Based Exercises


Contributing
License

Overview
GO-train is an interactive tool designed for learning and practicing Go programming through a set of coding challenges. Users can select an exercise by number, provide the required input, and view the output. The program supports continuous interaction, allowing users to run multiple exercises in a single session.
Features

45 coding exercises divided into string and array categories.
Interactive CLI interface to select exercises and provide inputs.
Clear descriptions for each exercise.
Option to continue or exit after each exercise.
Input validation and error handling for robust user experience.

Project Structure
GO-train/
├── docs.md              # Documentation or additional notes
├── go.mod               # Go module file
├── main.go              # Main application entry point
└── trains/              # Package containing exercise implementations
    ├── 1.go             # Exercise 1: Reverse a string
    ├── 2.go             # Exercise 2: Check palindrome
    ├── 3.go             # Exercise 3: Remove duplicates from string
    ├── 4.go             # Exercise 4: First non-repeating character
    └── ...              # Other exercise files (5.go to 45.go)

Installation

Ensure you have Go installed (version 1.16 or later recommended).
Clone the repository:git clone <repository-url>
cd GO-train


Initialize the Go module (if not already initialized):go mod init go-train


Install dependencies (if any):go mod tidy



Usage

Run the program:go run main.go


Follow the prompts:
Enter an exercise number (1 to 45).
Read the exercise description.
Provide the required input (e.g., a string or space-separated numbers).
View the result.
Choose to continue (y or yes) or exit (n or no).



Example session:
Welcome to GO-train exercises!
Please enter the exercise number (1-45): 1

Exercise description:
Exercise 1: Reverse a given string — without using built-ins.
Please enter a string: hello

Exercise result:
olleh

Do you want to continue? (y/n): y

Exercises
The program includes 45 exercises, categorized into string-based and array-based problems.
String-Based Exercises

Reverse a given string — without using built-ins.
Check if a string is a palindrome.
Remove duplicates from a string — efficiently.
Find the first non-repeating character — who stands alone?
Count how many times each character appears.
Flip the words in a sentence, not the letters.
Are two strings anagrams? Prove it.
Longest substring without repeats — sliding window style.
Build your own atoi — string to integer.
Compress strings with run-length encoding.
Most frequent character — who dominates?
List all possible substrings of a string.
Is one string a rotation of another?
Strip all white spaces from a string.
Is this a valid shuffle of two strings?
Convert text to Title Case — properly.
Find the longest common prefix among words.
Break a string into a char array — without confusion.
Replace spaces with %20 — classic URL trick.
Turn full sentences into acronyms.
Check if the string is all digits — no alphabets allowed.
Count how many words are in the string.
Remove a specific character — cleanly.
Find the shortest word in a sentence.
Longest palindromic substring — two-pointer style.

Array-Based Exercises

Reverse an array in-place.
Find the largest and smallest element.
Check for duplicates in an array.
Remove duplicates — return only unique values.
Find the missing number from 1 to N.
Move all zeros to the end — keep order.
Rotate the array left/right by K positions.
Find the Kth largest/smallest element.
Merge two sorted arrays — without using extra space.
Find the intersection of two arrays.
Sort 0s, 1s, and 2s without using sort().
Find subarrays with a given sum.
Detect if a subarray sums to 0.
Find the longest increasing subsequence.
Kadane’s Algorithm — maximum subarray sum.
Check if array is sorted and rotated.
Rearrange array in max-min order alternately.
Find leaders in an array (no greater element to the right).
Calculate frequency of all elements in O(n).
Product of all elements except self.

Contributing
Contributions are welcome! If you want to add new exercises, improve existing ones, or fix bugs:

Fork the repository.
Create a new branch (git checkout -b feature/your-feature).
Make your changes and commit (git commit -m "Add your feature").
Push to the branch (git push origin feature/your-feature).
Open a pull request.

Please ensure your code follows Go best practices and includes appropriate comments.
License
This project is licensed under the MIT License. See the LICENSE file for details.