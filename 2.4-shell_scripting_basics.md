
# Shell Scripting Basics

## Introduction
Shell scripting allows automation of tasks in Linux. A shell script is a file containing commands executed in sequence.

---

## Creating a Shell Script
```bash
#!/bin/bash
# My first script
echo "Hello, World!"
```
Save the file with a `.sh` extension.

Make it executable:
```bash
chmod +x script.sh
```
Run the script:
```bash
./script.sh
```

---

## Variables
```bash
name="John"
echo "Hello, $name"
```

---

## User Input
```bash
echo "Enter your name:"
read name
echo "Hello, $name!"
```

---

## Conditional Statements

### If-Else
```bash
num=10
if [ $num -gt 5 ]; then
    echo "Number is greater than 5"
else
    echo "Number is 5 or less"
fi
```

### Case Statement
```bash
echo "Enter a number:"
read num
case $num in
    1) echo "One" ;;
    2) echo "Two" ;;
    *) echo "Other number" ;;
esac
```

---

## Operations and Conditions

### Arithmetic Operations
```bash
a=10
b=5
sum=$((a + b))
echo "Sum: $sum"
```

### Logical Conditions
```bash
x=10
y=20
if [[ $x -gt 5 && $y -lt 30 ]]; then
    echo "Both conditions are true"
fi
```

### String Comparisons
```bash
str1="hello"
str2="world"
if [ "$str1" != "$str2" ]; then
    echo "Strings are not equal"
fi
```

---

## Loops

### For Loop
```bash
for i in 1 2 3 4 5; do
    echo "Number: $i"
done
```

### While Loop
```bash
count=1
while [ $count -le 5 ]; do
    echo "Count: $count"
    count=$((count + 1))
done
```

---

## Functions
```bash
say_hello() {
    echo "Hello, $1"
}
say_hello "Alice"
```

---

## File Operations

### Check if a File Exists
```bash
if [ -f "file.txt" ]; then
    echo "File exists"
else
    echo "File does not exist"
fi
```

### Read a File Line by Line
```bash
while read line; do
    echo "$line"
done < file.txt
```

---

## Scheduling Scripts
Run a script every minute using cron:
```bash
* * * * * /path/to/script.sh
```

---

## Debugging
Run script in debug mode:
```bash
bash -x script.sh
```

---

## Operators

### Comparison Operators
| Operator | Meaning                    |
|----------|----------------------------|
| -eq      | Equal to                   |
| -ne      | Not equal to               |
| -gt      | Greater than               |
| -ge      | Greater than or equal to   |
| -lt      | Less than                  |
| -le      | Less than or equal to      |

### Logical Operators
| Operator | Meaning      |
|----------|--------------|
| &&       | Logical AND  |

### Arithmetic Operators
```bash
sum=$((num1 + num2))
diff=$((num1 - num2))
prod=$((num1 * num2))
div=$((num1 / num2))
mod=$((num1 % num2))
```

### File Test Operators
| Operator | Meaning                                 |
|----------|-----------------------------------------|
| -f       | Check if a file exists and is a regular file |
| -d       | Check if a directory exists             |
| -e       | Check if a file or directory exists     |
| -s       | Check if a file exists and is not empty |
| -r       | Check if a file has read permission     |
| -w       | Check if a file has write permission    |
| -x       | Check if a file has execute permission  |
| -L       | Check if a file is a symbolic link     |
```


