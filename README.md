# Golang console interactive prompts 

Copro offer you interactive console user prompts like select, password,...

## Table of Contents

  1. [Getting started](#getting-started)
    1. [Installation](#install-the-package)
    2. [Basic usage](#basic-usage)
  2. [Prompts](#prompts)
     1. [Input](#input)
     2. [Password](#password)
     3. [Confirm](#confirm)
     4. [Select](#select)
     5. [Checkbox](#checkbox)
  3. [Examples](#examples)
  4. [TODOS](#todos)
  5. [License](#license)
  

# [Getting started](#getting-started)

A quick overview how simple is to use copro

## [Install the package](#install-the-package)

    go get -u github.com/julienroland/copro/prompt

## [Basic usage](#basic-usage)

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

# [Prompts](#prompt)

## [Input](#input)
    
    ask := prompt.NewInput()
    ask.Question = "What's your name"
    name, _ := ask.Run()
    fmt.Println(name)

## [Password](#password)

    ask := prompt.NewPassword()
	ask.Question = "Enter your password"
	ask.Validation = func(answer string) bool {
		return answer != "password"
	}
	ask.ErrorMessage = func(answer string) string {
		return "Wrong password"
	}
	ask.MaximumNumberOfTry = 2
	result, err := ask.Run()
	fmt.Println(result, err)

## [Confirm](#confirm)

    ask := prompt.NewConfirm()
	ask.Question = "Are you a gopher"
	ask.Default = "yes"
	result := ask.Run()
	fmt.Println(result)

## [Select](#select)

    ask := prompt.NewSelect()
	ask.Question = "[?] Choose what you want:"
	ask.Choices = []*prompt.Choice{
		{
			ID:    1,
			Label: "Apple",
		},
		{
			ID:       2,
			Label:    "Tomato",
			Selected: true,
		},
		{
			Label:       "Others",
			IsSeparator: true,
		},
		{
			ID:    3,
			Label: "Orange",
		},
		{
			ID:       4,
			Label:    "Watermelon",
			Selected: true,
		},
	}
	result, err := ask.Run()
	fmt.Println(result, err)

## [Checkbox](#checkbox)

    ask := prompt.NewCheckbox()
	ask.Question = "[?] Choose what you want:"
	ask.Choices = []*prompt.Choice{
		{
			Label:       "List",
			IsSeparator: true,
		},
		{
			ID:    1,
			Label: "Apple",
		},
		{
			ID:       2,
			Label:    "Tomato",
			Selected: true,
		},
	}
	result, err := ask.Run()
	fmt.Println(result, err)

# [Examples](#examples)

You can find examples on the [github repository](http://github.com/julienroland/copro/tree/master/examples)

# [TODOS](#todos)

- [x] Add custom validation for input type
- [x] Add password field type
- [ ] Write tests
- [ ] Add auto-complete type

# [License](#license)

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
