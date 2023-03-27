# LispiE: Lisp inspired smart-contract programming interface for Ethereum

This is an experiment, solidity syntax is so ugly that I rather (try) build something new from scratch

## Goals

1. **Simplicity**: We are developing a language that is easy to understand and use, with minimal syntax and an uncluttered design. The Lisp-like syntax, using parenthesized prefix notation (S-expressions), reduces the learning curve for developers familiar with Lisp or other functional programming languages.

2. **Security**: Our language prioritizes security, incorporating mechanisms that help developers write secure smart contracts. This includes protection against common smart contract vulnerabilities, explicit error handling, and support for formal verification to ensure the correctness of the code.

3. **Auditability**: We are designing the language to facilitate code readability and auditability, making it easier for developers and auditors to review smart contracts. The use of a Lisp-like syntax and a focus on simplicity contributes to this goal.

4. **Strong Typing**: Our language retains strong typing, similar to Rust, to minimize runtime errors and improve overall code safety. This feature promotes more robust smart contract development and reduces the likelihood of unexpected behavior.

5. **EVM Compatibility**: The language is being developed to target the Ethereum Virtual Machine (EVM), enabling seamless integration with the Ethereum ecosystem and making it suitable for creating smart contracts on the Ethereum network.

## Syntax and grammar (WIP)


### Variables and Constants

#### defvar

Declare global variables.

```lisp
(defvar x uint256)
```

### defconst

Declare constants.

```lisp
(defconst PI 3.14159)
```

### let

Declare local variables

```lisp
(let (x uint256 10)
  (* x 2))
```

## Functions

### defun

Delcare a function with visibility

```lisp
(defun :external add (a uint256 b uint256) -> uint256
  (+ a b))
```

## Logical Operations

### and 

Logical AND

```lisp
(and (> x 0) (< x 10))
```


### or 

Logical OR

```lisp
(or (= x 0) (= x 10))
```

### not 

Logical NOT

```lisp
(not (= x 0))
```


## Conditionals

### if

Conditional expressions

```lisp
(if (< x 0)
    (- x)
    x)
```

## Functional Programming Constructs

### map

Apply a function to each element of a list

```lisp
(map (fn (x uint256) -> uint256 (* x 2)) my_list)
```

### filter 

Filter a list based on a predicate a function

```lisp
(filter (fn (x uint256) -> bool (= (mod x 2) 0)) my_list)
```

## Event Loggin

### event

Declare an event

```lisp
(event Transfer (from to value))
```

### emit

Emit an event

```lisp
(emit Transfer from to value)
```

## Loops and Iteration (WIP)

### for

Iterate through a range

```lisp
(for (i uint256 0) (< i 10) (set i (+ i 1))
  (do-something i))
```

### while

Loop while a condition is true.

```lisp
(while (< x 10)
  (set x (+ x 1)))
```

## Interfaces and External Calls

### definterface 

Define an interface.

```lisp
(definterface IERC20
  (function transfer (to address value uint256) -> bool))
```

