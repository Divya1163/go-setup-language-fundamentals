
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

    // Function to accept and update student data
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

    // Function to display student details
    fun displayDetails() {
        println("\n----- Student Details -----")
        println("Roll Number: $rollNo")
        println("Name: $name")
        println("Course: $course")
        println("Percentage: $percentage%")
        println("--------------------------")
    }
}

fun main() {

    // First Student object
    val student1 = Student(0, "", "", 0.0)

    println("Enter details for Student 1:")
    student1.acceptData()
    student1.displayDetails()

    // Second Student object
    val student2 = Student(0, "", "", 0.0)

    println("\nEnter details for Student 2:")
    student2.acceptData()
    student2.displayDetails()
}
