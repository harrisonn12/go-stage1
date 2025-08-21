🏗️ Stage 1 Exercises (Foundations)
1. Hello Go

Print "Hello, Go!" to the terminal.

Use both var and := for variable declarations.

Print the zero value of: int, string, bool, and a custom struct.

✅ Concepts: zero values, short declaration, formatting.

2. Temperature Converter

Write a program that:

Reads a temperature in Celsius,

Converts it to Fahrenheit,

Prints both.

Store the conversion function in a separate package.

✅ Concepts: functions, packages & imports, variables.

3. Struct Playground

Define a Car struct with fields: Brand, Year, Electric.

Create two cars (one electric, one gas).

Print their details with fmt.Printf.

Add a method IsVintage() (year < 1990).

✅ Concepts: structs, methods (value receiver), printing.

4. Slice Manager

Start with an empty slice of strings.

Append names: "Alice", "Bob", "Charlie".

Print length and capacity.

Remove "Bob" from the slice.

Reverse the slice in-place.

✅ Concepts: slices, append, length vs capacity.

5. Map Dictionary

Create a map of string → string where keys are English words and values are Spanish translations.

Add at least 5 entries.

Look up a word; if not found, print "Word not found".

Delete one entry.

Iterate over all key-value pairs with range.

✅ Concepts: maps, checking for existence, range.

6. Rune & String Explorer

Write a function that:

Accepts a string,

Prints each character and its Unicode code point.

Test with "Go 🐹".

✅ Concepts: strings, runes, Unicode.

7. Enum with iota

Define constants for the days of the week using iota.

Write a function DayName(d int) string that returns the day’s name.

Print all 7 days using a loop.

✅ Concepts: const, iota, type safety.

8. Todo CLI (Mini Project)

Build a command-line todo list app that can:

Add a task (struct with ID, Title, Done),

List all tasks,

Mark a task as done,

Delete a task.

Store tasks in a slice.

✅ Concepts: structs, slices, maps, range, fmt, idiomatic style.
💡 This consolidates almost all Stage 1 concepts into one project.

📌 Extra Challenges

Formatting: Run gofmt on all your code and compare before/after.

Linting: Run go vet to see if Go warns you about unused variables.

Idiomatic practice: Rewrite code using short declarations (:=) wherever possible.