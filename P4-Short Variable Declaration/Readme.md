Sad variable declaration:

var mySkillIssues int
mySkillIssues = 42

GOATed variable declaration:

mySkillIssues := 42

The walrus operator, :=, declares a new variable and assigns a value to it in one line. Go can infer that mySkillIssues is an int because of the 42 value. Yay type inference!

When to Use the Walrus Operator
The :=, (walrus operator) should be used instead of var style declarations basically anywhere possible. The limitation is that := can't be used outside of a function (in the global/package scope which we'll talk about later).

Type inference is based on the value being assigned.

An int:

mySkillIssues := 42

A float64:

pi := 3.14159

A string:

message := "Hello, world!"

A bool:

isGoat := true

Assignment
A common use case for Textio is to send birthday messages.

Complete the main function. It should print: "Happy birthday! You are now 21 years old!".
Create a string variable messageStart with the text "Happy birthday! You are now"
Create an integer variable age set to 21
Create another string variable messageEnd with the text "years old!"
The provided fmt.Println statement will print the full message on a single line separated by spaces.