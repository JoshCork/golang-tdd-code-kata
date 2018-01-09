# tdd-code-kata
Learning Go through a TDD Code Kata

## TDD Kata - String Calculator

Developed by Roy Osherove - http://osherove.com/tdd-kata-1

The goal of this exercise is to see how far you can get in 30 minutes. Write a test before implementing any code. Remember to solve things as simply as possible so that you force yourself to write tests you did not think about. Remember to refactor after each passing test

# TDD Code Kata


## TDD Kata - String Calculator

Developed by Roy Osherove - [http://osherove.com/tdd-kata-1](http://osherove.com/tdd-kata-1)

The goal of this exercise is to see how far you can get in 30 minutes. Write a test before implementing any code. Remember to solve things as simply as possible so that you force yourself to write tests you did not think about. Remember to refactor after each passing test

### Step 1: Create a simple String calculator with a single method:  int Add(string numbers)

The method can take 0, 1 or 2 numbers, and will return their sum. For example:

- Add(“”) should return 0
- Add(“2112”) should return 2112
- Add(“2,3”) should return 5

### Step 2: Allow the Add method to handle an unknown amount of numbers

### Step 3: Allow the Add method to handle new lines between numbers (instead of commas).

- The following input is ok:  “1\n2,3”  (will equal 6)
- The following input is NOT ok:  “1,\n” (no need to prove it - just clarifying)

### Step 4: Support different delimiters

- To change a delimiter, the beginning of the string will contain a separate line that looks like this: `//[delimiter]\n[numbers…]`
- For example `//;\n1;2” should return three where the default delimiter is ‘;’`
- All existing tests should still be supported

### Step 5: Calling Add with a negative number will throw an exception “negatives not allowed” and the negative that was passed

- If there are multiple negatives, show all of them in the exception message

### Step 6: Numbers bigger than 1000 should be ignored

- Adding 2 + 1001  = 2

### Step 7: Delimiters can be of any length with the following format:  “//[delimiter]\n”

- For example: `“//[***]\n1***2***3”` should return 6

### Step 8: Allow multiple delimiters like this:  “//[delim1][delim2]\n”

- For example: `“//[*][%]\n1*2%3”` should return 6
- Make sure you can also handle multiple delimiters with length longer than one char


## Progress by Day

### 1/9/18

#### Kata 2: 4:03pm - 4:36pm

Needed a few extra minutes.  Was able to get the test to pass for inputs of both "" and "1,2,3".

#### Kata 1

Good progress here.  A lot of the time was spent remember where I was at and what needed to be done.  Still working through things.

### 1/5/18


#### Kata 2

Great progress again.  My Add function got much more real but still didn't get it passing.  My input was returning a slice of length 1 with an empty string using the split function on a comma.  I just realized that I spent the rest of my time building logic for the case statement under this branching logic for the length of zero instead of one so it kept failing.  I need to go in and update that on the next go-around.

I was distracted a lot during this pomodoro/kata as I had a phone call from Kal's coach but it was the wrong coach. Drama ensued afer that and I had a hard time getting back into the right frame of mind.

#### Kata 1

Good progress here.  Got past my block on the packages stuff.  I was able to get into testing and make the test fail.  Ended looking at how to make the Add function work.

### 12/20/2017

#### Kata 1
Well I got a little further... sort of.  I was trying to follow a different pattern that Jeff was using and moving stuff to an outside package.  Didn't finish getting that working yet.

Deleted everything in controllers and main.
Checking in and starting over.


### 12/15/17

Not a lot of progress today.  Sadly a lot of interuptions and re-learning stuff.  Will pick up here next week.