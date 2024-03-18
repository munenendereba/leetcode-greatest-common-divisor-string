# Greatest Common Divisor of Strings

## Problem

Solve the following problem in Leetcode [greatest-common-divisor-of-strings](https://leetcode.com/problems/greatest-common-divisor-of-strings) using Golang

For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""

Constraints:

1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.

## Running

Run using the following command: `go run .`

To run the tests `go test -v .`

To run the benchmark tests `go test -bench="." `

## Authors

[Munene Ndereba](https://github.com/munenendereba)

## License

This project is licensed under the MIT License.
