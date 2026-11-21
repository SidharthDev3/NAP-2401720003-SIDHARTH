class Student(
    rollNo: Int,
    name: String,
    course: String,
    percentage: Double
) {
    var rollNo: Int = rollNo
        private set

    var name: String = name
        private set

    var course: String = course
        private set

    var percentage: Double = percentage
        private set

    fun acceptData() {
        print("Enter Roll Number: ")
        rollNo = readLine()!!.toInt()

        print("Enter Name: ")
        name = readLine()!!

        print("Enter Course: ")
        course = readLine()!!

        print("Enter Percentage: ")
        percentage = readLine()!!.toDouble()
    }

    fun displayDetails() {
        println("\n----- Student Details -----")
        println("Roll Number: $rollNo")
        println("Name: $name")
        println("Course: $course")
        println("Percentage: $percentage%")
    }
}

fun main() {

    val student1 = Student(1, "Unknown", "Unknown", 0.0)
    val student2 = Student(2, "Unknown", "Unknown", 0.0)

    println("Enter Student 1 Details:")
    student1.acceptData()
    student1.displayDetails()

    println("\nEnter Student 2 Details:")
    student2.acceptData()
    student2.displayDetails()
}