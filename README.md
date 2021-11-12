<!-- Ctrl+Shift+V -->
# GO Basics
Learning Go language for my next goal, HL.

<a href="#user-content-day1">Day 1</a>　2021/11/08
* Setup
* Main Package
* Import
* Variables & Constants
* Functions
* For

<a href="#user-content-day2">Day 2</a>　2021/11/09
* If
* Switch
* Pointer
* Arrays, Slices
* Maps

<a href="#user-content-day3">Day 3</a>　2021/11/10
* Struct
* Methods
* Error & Handling
* Mini project
[
<a href="https://github.com/tyomhk2015/go_basics/tree/main/banking" target="_blank" rel="noopener">Banking Account</a>,
<a href="https://github.com/tyomhk2015/go_basics/tree/main/dictionary" target="_blank" rel="noopener">Dictionary</a>
]

<a href="#user-content-day4">Day 4</a>　2021/11/11
* goroutines
* Channels
* Mini project
[
<a href="https://github.com/tyomhk2015/go_basics/tree/main/urlChecker" target="_blank" rel="noopener">URL Checker</a>,
<a href="https://github.com/tyomhk2015/go_basics/tree/main/jobScrapper" target="_blank" rel="noopener">Job Scrapper</a>
]
* Troubleshooting

<a href="#user-content-day5">Day 5</a>　2021/11/12
* Troubleshooting
* Mini project
[
<a href="https://github.com/tyomhk2015/go_basics/tree/main/jobScrapper" target="_blank" rel="noopener">Job Scrapper</a>
]

<hr>

## Notes 📝

### **<a href="javascript:void(0);" id="day1">Day 1</a>** ☀️
2021/11/08

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

* (IMPORTANT) The GO code will look for the 'main' package and `main` function as the first priority, similar to Java 'main' method.

#### 💡 **Import**

* If you want to export some functions, the function name must start with an `uppercase` letter, a.k.a `exported` function' from a different package. (e.g. Println)
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

* To omit some returns from a function, use the `_`, an underscore sign. This ignores the current return value. 
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

### **<a href="javascript:void(0);" id="day2">Day 2</a>** ☀️
2021/11/09

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

### **<a href="javascript:void(0);" id="day3">Day 3</a>** ☀️
2021/11/10

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
<a href="https://github.com/tyomhk2015/go_basics/tree/main/banking" target="_blank" rel="noopener">Code</a>

#### Utilize Map: Dictionary Program, done. ✔️
<a href="https://github.com/tyomhk2015/go_basics/tree/main/dictionary" target="_blank" rel="noopener">Code</a>

### **<a href="javascript:void(0);" id="day4">Day 4</a>** ☀️
2021/11/11

💡 **Tip for making an empty map**

<pre>
  // make() can make empty 
  results := make(map[string]string)

  // This will be 'nil' and cannnot add any objects into it.
  var results map[string]string
</pre>


#### 💡 **goroutines**

* Enable functions to run `Concurrently`.

* By adding `go` keyword before calling a function, the function will run `concurrently` with the main func or the main thread.
<br>
Caution: `main func` do not wait for `go` routines to finish their jobs. Thus if all called functions have `go` keyword, and if the `main func` don't have anything else to do, the program will be finished. 
<br>
It is important to communicate between `go` routines and `main func`.

<pre>
// main func does not wait for the 'go' routines to finish.
func main() {
  go functionOne()
  go functionTwo()
}

// Cannot receive returns from goroutines.
</pre>

#### 💡 **Channels**

* For communicating `main func` and `go` routines.
<br>
Solves the problem mentioned above, the paragraph in `Go routines` Caution'.

<pre>
func channel_main() {
  // variable := make(channel data_type_for_communication)
  channel := make(chan bool)

  foods := [2]string{"Sushi", "Ekiben"}
  for _, food := range foods {
    // go routine for concurrency
    // Giving channel as an argument, the main() and eatFood() can communicate.
    go eatFood(food, channel)
  }

  // Get the value used in channel.
  // <- : Blocking operation, getting something out from another thing.
  fmt.Println(<-channel) // Waiting for a message.
  fmt.Println(<-channel)

  // If all 'go' routines are finished, the channel has nothing more todo.
  // Therefore, an error invokes.
  fmt.Println(<-channel)
}

// The 2nd arg: Takes a channel, and a bool for communicating with main(),
//              or data type you want to give to the channel to communicate.
func eatFood(food string, c chan bool) {
  fmt.Println(food)
  // Send bool to channel for communicating w/ the main func.
  c <- true
}

// Result:
// Ekiben
// true
// Sushi
// true
</pre>

* (c chan<- communcation_type): For one-direction sending, write-only.
<a href="https://blog.gopheracademy.com/advent-2019/directional-channels/" target="_blank" rel="noopener">Explanation Link</a>

<pre>
chan<- type : Put type into the channel.
<-chan      : Get something out from the channel.
</pre>

* Started using <a href="https://github.com/PuerkitoBio/goquery" target="_blank" rel="noopener">goquery</a> to scrap a website. 
<a href="https://github.com/tyomhk2015/go_basics/tree/main/jobScrapper" target="_blank" rel="noopener">Link to the mini project</a>.

#### ⚠️ **Troubleshooting**

<b>Problem</b>:

There was an error about not able to find goquery module.
The source of the problem was directory for installing 3rd party libraries were at `C:\Users\user_name\go`, the `GOPATH`, not `C:\Go\`.

<br>

<b>Solution</b>: ✔️

Followed the <a href="https://golang.org/doc/gopath_code#GOPATH" target="_blank" rel="noopener">guide</a>, which explains about creating a new directory and setting a customized `GOPATH`. Then changed the value of `GO111MODULE` variable of `go env` to false. (<a href="https://golang.org/doc/gopath_code#GOPATH" target="_blank" rel="noopener">Link</a>)
<br>
Finally, installed the goquery again. Then the error disappeared.

### **<a href="javascript:void(0);" id="day5">Day 5</a>** ☀️
2021/11/12

* Thanks to goquery library, I was able to utilize my jQuery experience and scrap the site without much trouble. 😊

* Continued working on the mini project, the `jobScrapper'.
[
<a href="https://github.com/tyomhk2015/go_basics/tree/main/jobScrapper" target="_blank" rel="noopener">Job Scrapper</a>
]

#### ⚠️ **Troubleshooting**

<b>Problem</b>:

Tried to take a struct, filled with some data, as an argument of a function, I've got an `empty struct` several times, which I did not expect to happen.

<br>

<b>Solution</b>: ✔️

Used `Pointer` feature and was able to get the desired struct, the one filled with some data.