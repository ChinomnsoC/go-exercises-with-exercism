# Learning Golang on Exercism

Learning Golang from scratch using the exercism track. Documenting my learnings.

## Day 4 and 5

1. Learned more about using the built-in formatting package "fmt" to create strings that embed values or variables with special formatting.
2. Also learned about how packages are the unit of code reuse in Go.
3. Learned to use `%03d` to embed integers with extra zeros (3 in this case) to the left of the integer.
4. Learned to use `%s` to embed string values and `%.1f` to embed float64 integer looking numbers and round them to 1 decimal place.

## Day 3

1. Learned about string manipulation in golang.
2. Particularly learned about the `strings` package with contains many useful functions.
3. Learned to use `ReplaceAll`, `TrimSpace` and `Repeat`.
4. I think my favorite is `Repeat` because you can just tell it to repeat a string as many times as you want or specify.

## Day 1 and 2

1. Learned that in Golang, there are comments conventions for the package, the functions, and the variables. Package comments should describe what the package does, and should begin with "Package", then the name of the package.
2. Also learned that the comments have to end with full stop.
3. The variable and function comments should begin with the names of the variables and functions, respectively.
4. I learned that I can multiply two floats together, to ensure that I don't lose precision when converting to integers. If I need the function to return an integer and one of the variables is a float, it is better to convert after the math function is done.
5. Also learned about boolean variables, and how its default value is false, yet for example, if `var awake bool`, and i say something like `if awake` then i mean if awake is true.