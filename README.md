<!-- Ctrl+Shift+V -->
# GO Basics
Learning Go language for my next goal, HL.

<hr>

## Notes 📝

### **★ Day 1** ☀️

#### Resource 📖

📘 https://golang.org/doc/ <br>
📗 https://go101.org/article/101.html

#### 💡 **Setup**

* The projects must be stored in `GO PATH`.

  <pre>
    Windows OS: 'C:\go\'
    Mac OS:     'user\local\go\'
  </pre>

  Personally, this practice could be more convenient for managing many different projects than other languages.
  (e.g. java or nodejs, which you can store the projects anywhere.)

* GO codes can be downloaded anywhere and will be stored in a seperated domain, which tells where you got the code from.

#### 💡 **Main Package**

* If you want to compile the project, there should be a GO file with the name 'main', just like 'main' method in Java.

* If you want to create a library, the file name does not have to main. Because the person who will use the library will complile with his/her 'main.go'.

* (IMPORTANT) The GO code will look for the 'main' package and '`main`' function as the first priority, similar to Java 'main' method.

#### 💡 **Import**

* If you want to export some functions, the function name must start with an `uppercase` letter, a.k.a '`exported` function' from a different package. (e.g. Println)
Functions start with `lowercase` letter are `private` functions.

#### 💡 **Variables & Constants**

<pre>
Variables: values you can change.
Constants: values you cannot change.
</pre>

* Go is a type language simimlar to Java or Typescript etc.

* Variables can be declared with shorthand but constants aren't.

#### 💡 **Functions**

* Just like Java, arguments and returns must have types delcared.

* Functions can have multiple arguments (similar to a spread operator in JS), and can return multiple values.

* To get receive `multiple` returns, write an independent variable for each returns. (e.g. Destruct from JS).
<pre>
returnOne, returnTwo, returnThree := multipleReturn(arg ...type)
</pre>

* To omit some returns from a function, use the '`_`', an underscore sign. This ignores the current return value. 
<pre>
// The first, and the third return values are store in the corresponding variables.
returnOne, _, returnThree := multipleReturn(arg ...type)
</pre>

* Naked Function
Returning variables are declared at the `first line` of its function, and return keyword is all by itself.
<pre>
                                ↓------------------------↓
func orderUberEats(menu string) (food string, payment int) {
food = menu
payment = len(menu) * 1000
return
</pre>

* Defer Function
A function that will be `excecuted after` when the previous function finishes its job. (Similar to the defer or the callback function in JS.)
<pre>
defer fmt.Println("The delivery has arrived at the destination.");
fmt.Println("The delivery is on the way.");

// Result:
// The delivery is on the way.
// The delivery has arrived at the destination.
</pre>

#### 💡 **For**

* The only way to iterate in Go is to use `for`.
* The `range` keyword in `for` returns two values. The first return is the index of given array, the second return is the value.
<pre>
func accumulate(numbers ...int) int {
	total := 0
	for index, value := range numbers {
		total += value
		fmt.Println(index, value)
	}
	return total
}

// Return:
// 0 1
// 1 2
// 2 3
// 3 4
// 4 5
</pre>
