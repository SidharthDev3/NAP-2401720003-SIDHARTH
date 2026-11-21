fun main() {
    print("Enter department code: ")
    val deptCode = readLine()!!.uppercase()

    print("Enter years of experience: ")
    val years = readLine()!!.toInt()

    print("Enter employee salary: ")
    val salary = readLine()!!.toDouble()

    val bonusRate = when {
        deptCode == "IT" && years > 5 -> 0.20
        deptCode == "HR" && years > 3 -> 0.15
        deptCode == "FIN" && years > 4 -> 0.12
        else -> 0.10
    }

    val bonus = salary * bonusRate

    println("Bonus Rate: ${bonusRate * 100}%")
    println("Bonus Amount: $bonus")
}