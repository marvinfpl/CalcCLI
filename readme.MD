Calc CLI - Scientific Calculator In The Terminal

A fast and extensible command line calculator written in Go.
Supports basic arithmetic, scientific functions and a CLI interface powered by Cobra.

Features

Basic arithmetic
• Addition
• Subtraction
• Multiplication
• Division

Scientific functions
• Cosine
• Sine
• Exponential
• Logarithm (with custom base)

Installation

```bash 
# Clone the repository
    git clone https://github.com/marvinfpl/CalcCLI.git
    cd CalcCLI

# uild the project
    go build -o calc

# Run it
    ./calc
```

Usage

```bash
# Addition
calc add 2 3
# output: 5

# Subtraction
calc sub 5 2
# output: 3

# Multiplication
calc mul 4 3
# output: 12

# Division
calc div 10 2
# output: 5

# Cosine
calc cos 45
# output: √2/2
calc cos π/4 --rad
# output: √2/2

# Sine
calc sin 45
# output: √2/2
calc sin π/4 --rad
# output: √2/2

# Exponential
calc exp 1
# output: 2.71828182...

# Logarithm
calc log 2.71828182
# output: 1
calc log 8 --base 2
# output: 3
```