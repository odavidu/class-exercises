# Week 1: Installing Go + "Hello World+"

Welcome to CS351! In this in-class exercise we will set up Go and play around with writing a few simple programs.

As you work through this exercise, note how Go is similar to other languages that you might know and how it is different!

## Installing Go

1. Download and install Go following the directions here: https://go.dev/doc/install
2. If you use Ubuntu, these are good instructions:
https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04
3. Windows users: install WSL (Windows subsystem for Linux). You should use Ubuntu.
This is a good guide: https://learn.microsoft.com/en-us/windows/wsl/install
To launch the subsystem, run “wsl” in a shell (command prompt or powershell). Then
install Go on you Ubuntu system:
https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04
- You may be able to run go natively on Windows, however, note that the autograder is a Linux-based system, so please verify the autograder output when you submit!

## Seting up an Editor

1. In your IDE of choice, we highly encourage VS Code, go to extensions (left hand side
bar) and search Go and download.
2. In VS Code, install the Code Runner extension (same technique). You can use this or
the command line to run your programs.
3. In VS Code, you might have to “trust your project” in order to get the little run carrots.

## Your First Go Program

Create a new file called `week1.go`.

As is tradition with a new language, let's first say "Hello!".

```golang
// Every executable Go Program should contain a package called main.
// This tells the Go compiler to compile the package into an executable
// program rather than a shared library.
package main

import (
	"fmt"
)

// The entry point of a Go program should be the main function of main package.
// When the executable is run, main() is automatically called.
func main() {
	fmt.Println("Hello World\n")
}
```

Copy the above program into `week1.go` and run the command:
```go run week1.go```

You should see the output:
```
Hello World
```

### About Importing Packages

Notice the following portion of the go program:
```
import (
	"fmt"
)
```

The printing functionality within go is included via the "fmt" package.
The `import` statement dictates which packages will be included in our code file. In this case, we only need a single package: "fmt". 

Later, we will add to this import statement and include additional packages like: `time`, `sync`, `strconv`, and `log`.

## Redirecting Output to a File

As developers, it is often useful to redirect our program's output to a file. 

In this class, our most difficult assignments will create long **logs** of events. It is common for our servers to log every message and event that occurs in our system for debugging purposes. These logs can be **10,000s** of lines long!

As you can imagine, these logs will overwhelm our terminal. A good tool to have under your belt is the ability to **redirect** a program's output to a file.

You can redirect output via the `>` character.

Now run:
```
go run week1.go > output.txt
```

Open the output.txt file, what is contained in the file?

## Go Language Basics

Now we are going to jump in and write a simple program! 


Read the pages (2-3 minutes each):
- https://gobyexample.com/variables
- https://gobyexample.com/for
- https://gobyexample.com/if-else
- https://gobyexample.com/slices
- https://gobyexample.com/strings-and-runes


Copy the following line into the main function of your `week1.go` program.
This line initializes and assigns a new variable `input` to hold a long string.
```
input := "There once was a cat named Barry. He was a very good cat. This cat lived in Boston. He loved doing Boston-related activities (that were good for cats). He walked the esplanade. He shopped on Newbury. He ate at Tatte. He sometimes even went to TD Garden. Did you know that cats are not allowed in TD Garden?"
```

Now, write a program that searches `input` for all occurrences of the word `"cat"`.
We want to print out a message for every found "cat" and also the **index** of where this "cat" was found within the string.

When you find an occurrence, print it out in the following format, where `i` is the index:
```
fmt.Printf("found cat @ %v\n",i)
```

You should see the output:
```
found cat @ 17
found cat @ 53
found cat @ 63
found cat @ 145
found cat @ 272
```

## Reading in a File

Now, let's search for the word "cat" in a much longer file!

Please read:
- https://gobyexample.com/reading-files

Notice you have a file called "dictionary.txt". 

Comment out the line that defines `input` to be a hard-coded string. 

Instead, set `input` to be the contents of `dictionary.txt`.

Read in the file "dictionary.txt" and search this file for all occurrences of the word "cat".

You should see a lot of output! The end of this output should be:
```
...
found cat @ 4220934
found cat @ 4220944
found cat @ 4220953
found cat @ 4227664
found cat @ 4227673
found cat @ 4227760
```

## Creating Functions

Let's move our word-searching code into a helper function. 

Copy over the function header into `week1.go` and move your logic that opens a file into this helper.
```
func searchForWord(filepath string, target string) {
    // implementation here
}
```

Let's also _generalize_ this function by having it search for **any** target string, not just "cat".

Now, our main function can call this helper on a new word:

```
func main() {
	fmt.Println("Hello World\n")
	searchForWord("dictionary.txt", "fish")
}
```

## The 'go' keyword

Now, we are going to dip our toe into the ocean of **multi-threaded code**.

What if we want to write a program that searches for the word "fish" and the word "dog" **at the same time**?

Let's create two new `goroutines` one that searches for the word "fish" and one that searches for "dog".

Please read:
- https://gobyexample.com/goroutines

Once you have read the article, change your main method to the following:

```golang
func main() {
	go searchForWord("dictionary.txt", "fish")
	go searchForWord("dictionary.txt", "dog")
	time.Sleep(2 * time.Second)
}
```

Notice, we have a line `time.Sleep(2 * time.Second)` which just tells the main thread to sleep. Try removing this line. What happens?

In future lessons, we will use better methods to **synchronize** multiple threads of execution.

