# Parentheses-checker
Implement a function that verifies if the given string is a balanceds sequence of brackets, i.e. each of the open brackets
must be closed by the same type of bracket in the right order:
```
[(]) - unbalanced
{[{}]()} - balanced
((1 + 2) * 3) - 4)/5 - balanced
```
Has a parentheses web <i>service</i> that generates a random sequence of parentheses of the length n. Use the standard router for http requests.
```
GET /generate?n={length}
```
<i>Client</i>  can call the service N = 1000 times for each of the strings of the length 2,4,8 and calculate the percent of balanced strings for each length and print the results to stdout.