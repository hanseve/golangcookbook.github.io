---
title: Expanding Variables in User Input
question: How do I fetch user input to variable
---

### Read from flag

* `flag.Int() flag.String()`
* `flag.IntVar()` .etc

{% include example.html example="variables" flag="-num=1 -num2=2"%}

### Read from Console (Stdin)

* `fmt.Scanln()`
* `fmt.Scanf()`

{% include example.html example="scan"%}

### Read from buffer

* `bufio.ReadString`
* `bufio.NewScanner`

{% include example.html example="buffer"%}

