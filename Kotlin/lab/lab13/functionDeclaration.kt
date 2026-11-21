fun studentDetails(name: String = "Unknown", age: Int = 20) {
    println("Name: $name, Age: $age")
}

fun displayDetails(name: String, age: Int, city: String) {
    println("Name: $name, Age: $age, City: $city")
}

fun printNumbers(vararg numbers: Int) {
    println("Numbers:")
    for (num in numbers) {
        println(num)
    }
}

fun calculate(numbers: List<Int>, operation: (Int) -> Int) {
    val result = numbers.map(operation)
    println("Result: $result")
}

fun main() {


    println("Default Arguments:")
    studentDetails()
    studentDetails("Sid", 21)


    println("\nNamed Arguments:")
    displayDetails(city = "Delhi", age = 21, name = "Sid")

    
    println("\nVararg:")
    printNumbers(10, 20, 30, 40, 50)

    println("\nHigher-Order Function:")
    val numbers = listOf(1, 2, 3, 4, 5)
    calculate(numbers) { it * it }
}