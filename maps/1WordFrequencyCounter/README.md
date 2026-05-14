# Word Frequency Counter: 
* Create a tool that reads a string and returns a map where keys are words and values are their occurrence counts.

In the context of the Word Frequency Counter, "occurrence counts" simply refers to the total number of times each specific word appears within a given text.When you use a map for this, the word acts as the unique key, and the count (an integer) acts as the value. As you iterate through the text, you check if the word is already in your map:If it is not there, you add it with a value of 1.If it is already there, you increment its existing value by 1.

## Visual Example
* Imagine you have the following sentence: "apple banana apple orange banana apple". The occurrence counts would be stored in a map like this:

|Key(Word)     |Value(Count)|
|--------------|------------|
|apple         |3           |
|banana        |2           |
|orange        |1           |

TWS
