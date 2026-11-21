
fun main() {

    // Task 1: Hello Kotlin
    println("Hello Kotlin!!!")

    // Task 2: Employee Bonus Calculation
    print("\nEnter department code (IT/HR/FIN): ")
    val deptCode = readLine()!!.uppercase()

    print("Enter years of experience: ")
    val experience = readLine()!!.toInt()

    print("Enter employee salary: ")
    val salary = readLine()!!.toDouble()

    val bonusPercentage = when {
        deptCode == "IT" && experience > 5 -> 20
        deptCode == "HR" && experience > 3 -> 15
        deptCode == "FIN" && experience > 4 -> 12
        else -> 10
    }

    val bonusAmount = salary * bonusPercentage / 100

    println("Bonus Percentage: $bonusPercentage%")
    println("Bonus Amount: Rs. $bonusAmount")

    // Task 3: Average of Digits
    print("\nEnter an integer: ")
    val number = readLine()!!.toInt()

    var n = kotlin.math.abs(number)
    var sum = 0
    var count = 0

    if (n == 0) {
        sum = 0
        count = 1
    } else {
        while (n > 0) {
            val digit = n % 10
            sum += digit
            count++
            n /= 10
        }
    }

    val average = sum.toDouble() / count

    println("Sum of digits: $sum")
    println("Number of digits: $count")
    println("Average of digits: $average")
}
