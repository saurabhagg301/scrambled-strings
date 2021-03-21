# Problem Statement
Count how many of the words from a dictionary appear as substrings in a long string of
characters either in their original form or in their scrambled form. The scrambled form of the
dictionary word must adhere to the following rule: the first and last letter must be maintained
while the middle characters can be reorganised.   

The scrambled or original form of the dictionary word may appear multiple times but we only count it once since we only need to know whether it shows up at least once.
For example, if we had the word ​ this​ in the dictionary, the possible valid words which would be counted are ​ this​ (original version) and ​ tihs​ (scrambled version). ​ tsih​ , ​ siht​ and other variations are ​ not valid​ since they do not start with ​ t ​ and end with ​ s ​ . Also, ​ tis​ , ​ tiss​ , and ​ thiss​ are ​ not scrambled forms, because they are not reorderings of the original set of letters.

# Usage
Run below commands
```
# cd /path/to/project/directory
# go build
# ./scrmabled-strings --dictionary [PATH TO DICTIONARY FILE] --input [PATH TO INPUT FILE]
```
Example:
```
# ./scrambled-strings --dictionary=dict/dict.txt --input=input/input.txt
```

# Testing
Run below command
```
go test --input=input/input_test.txt --dictionary=dict/dict_test.txt --expectedValue=5
```
**Note:**   
* Current test case is written assuming one line of input in input_test.txt file (actual input.txt can have any number of line)
* For different inputs, just change the string in the input_test.txt file and mention the expected value in the above mentioned command


# Logging
Check logs in the **logs.txt** file which contains map of dictionary words   
key = starting character    
value = array of word's signature   
and in subsequent lines:   
Line number, the word being matched from input file and the counter value being incremented for the matched word

# Input
Your input will consist of:
1. a dictionary file, where each line comprises one dictionary word from which you cancreate your dictionary. E.g. “and”, “bath”, etc, but note the dictionary words do not needto be real words.
2. an input file that contains a list of long strings, each on a newline, that you will need touse to search for your dictionary words. E.g. “btahand”

# Output
Treating each line of the input file as one search string, your program should output a line ​Case#x:y​per input file string, where ​x​ is the line number (starting from 1) and ​y​ is the number ofwords from the dictionary that appear (in their original or scrambled form) as substrings of thegiven string.   
E.g.   
   Case #1: 2

# Validations 
* No two words in the dictionary are the same.
* Each word in the dictionary is between 2 and 105 letters long, inclusive.
* The sum of lengths of all words in the dictionary does not exceed 105.


# Stretch Goals
* Dockerfile that we can build and run the code   
* Documentation for public interfaces   
* Logging