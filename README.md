<div align="center">

# Learn `Golang`

![go-lang-intro-image](https://user-images.githubusercontent.com/194400/88049550-3f1ce500-cb4d-11ea-888b-2abad0f78446.png)

Learn `Golang` to build simple, reliable, and efficient software.

</div>

## Contents

- [Why?](#why)
- [What?](#what)
- [How?](#how)
- [Recommended Reading](#recommended-reading)
- [Popularity](#popularity)
- [Why Not?](#why-not)
- [Basic types](#basic-types)

## _Why_?

`Go` is a great programming language
if you want a good balance between simplicity,
beginner-friendliness (_being able to find/train new team members_)
and high performance.

## _What_?

Go is a statically typed, compiled programming language
designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.
Go is syntactically similar to C,
but with memory safety,
garbage collection,
structural typing, and CSP-style concurrency
The language is often referred to as "Golang" because of its domain name, golang.org, but the proper name is Go.
~ https://en.wikipedia.org/wiki/Go_(programming_language)

## _Who_?

Anyone who wants to broaden their programming skills should learn `Go`.
The syntax is immediately familiar to anyone who has written a `C`-style
language such as `JavaScript`.

## _How_?

You can learn the basics of `Go` online: https://play.golang.org
But when it comes time to actually _run_ a program
on your computer, you will need to install the compiler.
Visit:
https://golang.org/doc/install
and select the installer for your operating system.

### Hello World!

> Following: https://golang.org/doc/install#testing

Create a new file called `hello.go`
and type the following code in it:

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

Once you've saved the file,
in your terminal,
compile the programming
with the following command:

```sh
go build hello.go
```

That will create an _executable_ file called `hello` (_with no extension_)
To _execute_ the program,
run the command:

```
./hello
```

In your terminal, you should see:

```
hello, world
```

Congratulations you just wrote your first Go program!

<br />

### Learning Resources

If you just want to know the basics, watch Jake Wright's
"**Learn Go in 12 Minutes**": https://youtu.be/C8LgvuEBraI
[![image](https://user-images.githubusercontent.com/194400/88054566-83ac7e80-cb55-11ea-841c-d5fa72f3c8a5.png)](https://youtu.be/C8LgvuEBraI)

If you want to learn `Go` in more depth,
FreeCodeCamp Learn Go Programming
Tutorial for Beginners:
https://youtu.be/YS4e4q9oBaU
[![image](https://user-images.githubusercontent.com/194400/88054399-2f090380-cb55-11ea-943c-cc230ef0e6d7.png)](https://youtu.be/YS4e4q9oBaU)

The tutorial is almost 7 hours long but covers quite a lot.

## Recommended Reading

- Why Go: https://nathany.com/why-go (10 mins)
- The Case For Go:
  https://gist.github.com/ungerik/3731476 (many links; 30+ mins)
- Why should you learn Go?
  https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65 (8 mins)
- What’s so great about Go?
  https://stackoverflow.blog/2020/11/02/go-golang-learn-fast-programming-languages (5 mins)
- Why choose Go for your next project? (7 mins)
  https://www.softkraft.co/why-choose-go-for-your-next-project
- Getting started with Go:
  https://medium.com/rungo/working-in-go-workspace-3b0576e0534a
- How to Write Go Code (_official_):
  https://golang.org/doc/code.html

## Popularity?

Despite the simplicity and performance of `Go`,
it's nowhere _near_ as used/known
as Google would like you to believe.
For example it's not in the Top 10 languages on GitHub:
https://octoverse.github.com/#top-languages
<img width="1171" alt="github-most-popular-languages" src="https://user-images.githubusercontent.com/194400/87814203-2d8cc200-c85b-11ea-8229-c5d1b4dae2d4.png">

And it's used by fewer than 9% of developers (_who answered the survey_)
on StackOverflow.
https://insights.stackoverflow.com/survey/2020#technology
<img width="802" alt="stackoverflow-dev-survey-languages" src="https://user-images.githubusercontent.com/194400/87814783-4c3f8880-c85c-11ea-9b2a-c7e1c068491e.png">

Don't let these stats dissuade you.
The biggest reason `Go` is not as popular as `JavaScript`, `Python` or `PHP`
is simply because the other languages got a _massive_ head start.

Go usually scores well on synthetic benchmarks:
https://www.techempower.com/benchmarks
We know from experience that Postgres is usually the bottleneck in our Apps,
so the programming language doesn't matter quite as much you think.
What matters is having a sane approach to concurrency and channels
for building the features your people want.

### Why _Not_?

Common complaints/gripes developers have with `Go` include:

- Error handling - this can be quite tedious
  see: https://medium.com/rungo/error-handling-in-go-f0125de052f0
- Code duplication - some people complain about the lack of **inheritance** leading to code duplication.
- Lack of [metaprogramming](https://en.wikipedia.org/wiki/Metaprogramming) can mean less extensibility.
- Operator overloading - a lack of
  [operator overloading](http://www.golangpatterns.info/object-oriented/operators)
  can lead to more terse code.

We feel all the issues (_with the exception of error handling_)
are actually _features_ of Go that lead to simpler more maintainable code.

> As noted in ["Why Elixir"](https://github.com/dwyl/learn-elixir/issues/102) > `Go` is our 3rd Choice for programming languages.
> We use `Elixir` because Phoenix Channels makes Real Time _easy_
> and `Nerves` makes building, deploying and maintaining IoT a _breeze_.
> If `Elixir` didn't exist it would be a choice between `Rust` and `Go`.
> `Go` is a _lot_ simpler and beginner-friendly.
> `Rust` is "safer" and more performant.


## Basic types

Go's basic types are:

- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32, represents a Unicode code point
- float32 float64
- complex64 complex128

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type. 
