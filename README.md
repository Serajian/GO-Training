# GO-train

GO-train is a command-line application written in Go that provides a collection of programming exercises to practice string and array manipulation. It includes 45 exercises (25 string-based and 20 array-based) to help developers enhance their problem-solving skills in Go.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Exercises](#exercises)
  - [String-Based Exercises](#string-based-exercises)
  - [Array-Based Exercises](#array-based-exercises)
- [Contributing](#contributing)
- [License](#license)

## Overview
GO-train is an interactive tool designed for learning and practicing Go programming through a set of coding challenges. Users can select an exercise by number, provide the required input, and view the output. The program supports continuous interaction, allowing users to run multiple exercises in a single session.

## Features
- 45 coding exercises divided into string and array categories.
- Interactive CLI interface to select exercises and provide inputs.
- Clear descriptions for each exercise.
- Option to continue or exit after each exercise.
- Input validation and error handling for robust user experience.

## Project Structure
```
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
```

## Installation
1. Ensure you have [Go](https://golang.org/doc/install) installed (version 1.16 or later recommended).
2. Clone the repository:
   ```bash
   git clone <repository-url>
   cd GO-train
   ```
3. Initialize the Go module (if not already initialized):
   ```bash
   go mod init go-train
   ```
4. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

## Usage
1. Run the program:
   ```bash
   go run main.go
   ```
2. Follow the prompts:
   - Enter an exercise number (1 to 45).
   - Read the exercise description.
   - Provide the required input (e.g., a string or space-separated numbers).
   - View the result.
   - Choose to continue (`y` or `yes`) or exit (`n` or `no`).

Example session:
```
Welcome to GO-train exercises!
Please enter the exercise number (1-45): 1

Exercise description:
Exercise 1: Reverse a given string — without using built-ins.
Please enter a string: hello

Exercise result:
olleh

Do you want to continue? (y/n): y
```

## Exercises
The program includes 45 exercises, categorized into string-based and array-based problems.

### String-Based Exercises
1. Reverse a given string — without using built-ins.
2. Check if a string is a palindrome.
3. Remove duplicates from a string — efficiently.
4. Find the first non-repeating character — who stands alone?
5. Count how many times each character appears.
6. Flip the words in a sentence, not the letters.
7. Are two strings anagrams? Prove it.
8. Longest substring without repeats — sliding window style.
9. Build your own atoi — string to integer.
10. Compress strings with run-length encoding.
11. Most frequent character — who dominates?
12. List all possible substrings of a string.
13. Is one string a rotation of another?
14. Strip all white spaces from a string.
15. Is this a valid shuffle of two strings?
16. Convert text to Title Case — properly.
17. Find the longest common prefix among words.
18. Break a string into a char array — without confusion.
19. Replace spaces with %20 — classic URL trick.
20. Turn full sentences into acronyms.
21. Check if the string is all digits — no alphabets allowed.
22. Count how many words are in the string.
23. Remove a specific character — cleanly.
24. Find the shortest word in a sentence.
25. Longest palindromic substring — two-pointer style.

### Array-Based Exercises
26. Reverse an array in-place.
27. Find the largest and smallest element.
28. Check for duplicates in an array.
29. Remove duplicates — return only unique values.
30. Find the missing number from 1 to N.
31. Move all zeros to the end — keep order.
32. Rotate the array left/right by K positions.
33. Find the Kth largest/smallest element.
34. Merge two sorted arrays — without using extra space.
35. Find the intersection of two arrays.
36. Sort 0s, 1s, and 2s without using sort().
37. Find subarrays with a given sum.
38. Detect if a subarray sums to 0.
39. Find the longest increasing subsequence.
40. Kadane’s Algorithm — maximum subarray sum.
41. Check if array is sorted and rotated.
42. Rearrange array in max-min order alternately.
43. Find leaders in an array (no greater element to the right).
44. Calculate frequency of all elements in O(n).
45. Product of all elements except self.

## Contributing
Contributions are welcome! If you want to add new exercises, improve existing ones, or fix bugs:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please ensure your code follows Go best practices and includes appropriate comments.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.