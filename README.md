<!-- Ctrl+Shift+V -->
# GO Basics
Learning Go language for my next goal, HL.

<hr>

## Notes 📝

### **Day 1** ☀️
<small>(2021/11/08)</small>

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

<hr>

### **Day 2** ☀️
<small>(2021/11/09)</small>

#### 💡 **If**

* Variable Expression: A variable that can be create at the timing of `if` statement. The variable expression is exclusive for the `if`, which means the variable can only be used in the `if` statement scope.

<pre>
// age = 26
func areYouABoomer(age int) bool {
	// Normal If
	if age > 30 {
		return true
	}

	// Variable Expression
	if newAge := age + 5; newAge > 30 {
		fmt.Println("Variable Expression")
		return true
	}
	return false
}

// Return:
// Variable Expression
// true
</pre>

#### 💡 **Switch**

* Similar to the switch in Java, but has more flexibility.

<pre>
// Switch with flexibility
switch{
  case price > 800 && price < 2500:
    return "Restaurant sushi"
  case price >= 2500 && price < 5000:
    return "Superior restaurant sushi"
}

// Switch with variable expression
switch newPrice := price + 1000; newPrice {
  case 10000:
    return "Omakase sushi"
  case 30000:
    return "Premium omakase sushi"
}
</pre>

#### 💡 **Pointer**

* Enables developers to do low level programming while using the high level language. Useful for when handling heavy operations or performances that have to be taken seriously.

<pre>
// &: Get the memory address of the variable
// *: See what is inside the memory address

chiba := 272
otaku := 145
tokyo := 100
destination := &chiba
fmt.Println(destination, *destination)
chiba = otaku
fmt.Println(destination, *destination)
destination = &tokyo
fmt.Println(destination, *destination, chiba)

// Return:
// memoryAddress, 272
// memoryAddress, 145
// memoryAddress, 100, 145
</pre>

#### 💡 **Arrays, Slices,  append()**

* Slice: Creating an array without declaring its length. This dynamically gets bigger or smaller depening on its elements, similar to ArrayList in Java.

* append(): Receives two arguments, the first one is a 'slice' and the second one is the value you want to put in. This function does `not modify` the original array like slice() in JS. Instead, it `returns a new array`.

<pre>
// Slice
consoleGames := [] string {"Persona 4G", "DragonQuest XI S", "Biohazard Village"}
consoleGames = append(consoleGames, "Yakuza 0")

for index, value := range consoleGames {
  fmt.Println(index, value)
}

// Return:
// Persona 4G
// DragonQuest XI S
// Biohazard Village
// Yakuza 0
</pre>

#### 💡 **Maps**

* An object that consists of `key` and `values`, similar to the `object` in JS or `maps` in Java.

<pre>
variable = map[type_of_the_key] type_of_the_value { key: value, key: value ...}

cities := map[string]string {"South Korea": "Seoul", "Japan": "Tokyo", "Thailand": "Bangkok", "Austria": "Wien"}

// The order of newly added key-value is not guaranteed.
cities["France"] = "Paris"

// Iterate maps
for _, value := range cities {
  fmt.Println(value)
}

// Returns
// Paris
// Seoul
// Tokyo
// Bangkok
// Wien
</pre>

<hr>

### **Day 3** ☀️
<small>(2021/11/10)</small>

#### 💡 **Struct**

* Similar to `struct` in C. Unlike the GO `maps`, `struct` is useful for creating `customized objects`, e.g. an object that takes string type, int type, and string slice.
<pre>
staffs := [] string {"Kazuki Ōhashi", "Toshiya Ōno", "Kenichiro Suehiro"}

// One way of creating a struct; Omitting field name.
anime_one := anime{"Shadow House", 2021, staffs}

// Another way of creating a struct; Explicitly write field name.
anime_two := anime{title: "Shadow House", published_date: 2021, staffs: staffs}
</pre>

* `lowercase` variables or functions means they are `private` or cannot be exported to other modules.
* `uppercase` variables or functions means they are `public` or exportable to other modules.

<pre>
// This code is stored in the 'banking' package.
func CreateBankAccount (name string) *bankingAccount {
  bankingAccount := bankingAccount{owner: name, balance: 0}
  return &bankingAccount
}

func main() {
  // Creating an object.
  // Syntax: variable := package.ConstructorFunction()
  myBankAccount := banking.CreateBankAccount("ASH")
  fmt.Println(myBankAccount, *myBankAccount)
}

// Return: 
// &{ASH 0} {ASH 0}
</pre>

#### 💡 **Methods**

* The things in between `func` and `deposit` is called reciever. `Reciever` writing convention: The left one is initial of the struct, the right is the struct name. `Reciever` enables the object to be equipped with certain methods, like the `Deposite` method below.
* If the `Deposit` method's receiver did not use `*`, GO will return a new `bankingAccount` object. (e.g. {ASH 0}). 
<br>
By adding `*` to the receiver, GO will use existing `bankingAccount` object. (e.g. {ASH 20}).
<br>
※ GO tends to copy objects whenever you invoke a function or method, and return that copied obeject to the caller.

<pre>
// banking/banking.go
func (b *bankingAccount) Deposit(amount int) {
  b.balance += amount
}

// main.go
  myBankAccount := banking.CreateBankAccount("ASH")
  myBankAccount.Deposit(20)
  fmt.Println(*myBankAccount)

// Result: {ASH 20}
</pre>

#### 💡 **Error & Handling**

* In GO, there is no `try & catch` or `exception` etc. All the unexpected events must be handled manually.

<pre>
var errNoMoney = errors.New("Not enough money.")

func (b *bankingAccount) Withdraw(amount int) error {
  if b.balance < amount {
    return errNoMoney // Return the error message.
  }
  b.balance -= amount
  return nil // Go's syntax requirement.
}
</pre>

#### Utilize Struct: Banking Account Program, done. ✔️
#### Utilize Map: Dictionary Program, done. ✔️



