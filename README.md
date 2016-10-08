#Console prompts 

Go package which add some basic interactive console user interfaces

## Table of Contents

  1. [Getting started](#getting-started)
    1. [Installation](#install-the-package)
    2. [Basic usage](#basic-usage)
  2. [Prompts](#prompts)
  3. [Examples](#examples)
  4. [TODOS](#todos)
  5. [License](#license)
  

#[Getting started](#getting-started)

A quick overview  

##[Install the package](#install-the-package)

    go get -u github.com/julienroland/copro/prompt

##[Basic usage](#basic-usage)

    package main

    import (
        "fmt"

        "github.com/julienroland/copro/prompt"
    )

    func main() {
        ask := prompt.NewInput()
        ask.Question = "What's your name"
        name, _ := ask.Run()
        fmt.Println(name)
    }

#[Prompts](#prompt)

TODO

#[Examples](#examples)

You can find examples on the [http://github.com/julienroland/copro/tree/master/examples](github repository)

#[TODOS](#todos)

- [ ] Add custom validation for input type
- [ ] Add password field type
- [ ] Write tests
- [ ] Add auto-complete type

#[License](#license)

MIT License

Copyright (c) 2016 Julien Roland

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
