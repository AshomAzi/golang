# Unique Character Checker
Use a map as a set to determine if a string contains all unique characters.

In this project, you are using a map to solve a common computer science problem: deduplication.

The goal is to determine if a string has any repeating characters (like "apple" has two 'p's) or if every character is unique (like "world").

## The "Map as a Set" Concept
In many languages, there is a data structure called a Set that only allows unique values. Go doesn't have a built-in Set type, so we use a map to mimic one.

When checking for unique characters, we don't actually care about a "count." We only care about existence. We use the character as the key and a simple boolean as the value.

## How the Logic Works
Initialize: Create an empty map, e.g., seen := make(map[rune]bool).

Iterate: Loop through every character in your string.

Check: For each character, check if it already exists in the seen map.

If it exists: You’ve found a duplicate! The string does not have all unique characters. You can stop and return false.

If it doesn’t exist: Add it to the map with a value of true.

Finish: If you get through the whole string without finding a duplicate, return true.

## Example: "Gopher"
'G': Is it in the map? No. Add seen['G'] = true.

'o': Is it in the map? No. Add seen['o'] = true.

'p': Is it in the map? No. Add seen['p'] = true.

...and so on. Result: Unique.

## Example: "Apple"
'a': Not in map. Add it.

'p': Not in map. Add it.

'p': Is it in the map? YES. Return false immediately.