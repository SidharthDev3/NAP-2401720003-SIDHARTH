fun main() {
    print("Enter an integer: ")
    val number = readLine()!!.toInt()

    val digits = number.toString()
        .filter { it.isDigit() }
        .map { it.digitToInt() }

    val sum = digits.sum()
    val average = sum.toDouble() / digits.size

    println("Sum of digits: $sum")
    println("Average of digits: $average")
}