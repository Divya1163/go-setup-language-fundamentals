
fun greet(name: String = "Student", message: String = "Hello") {
    println("$message, $name!")
}

// Function with vararg
fun printNumbers(vararg numbers: Int) {
    println("Numbers:")
    for (number in numbers) {
        print("$number ")
    }
    println()
}

// Higher-order function
fun applyOperation(numbers: List<Int>, operation: (Int) -> Int) {
    val result = numbers.map(operation)
    println("Original list: $numbers")
    println("Modified list: $result")
}

fun main() {

    // Task 1: Default parameters
    println("--- Default Parameters ---")
    greet()
    greet("Divya")
    greet("Divya", "Welcome")

    // Task 2: Named arguments (out of order)
    println("\n--- Named Arguments ---")
    greet(message = "Good Morning", name = "Divya")

    // Task 3: Vararg
    println("\n--- Vararg Function ---")
    printNumbers(10, 20, 30, 40, 50)

    // Task 4: Higher-order function with lambda
    println("\n--- Higher-Order Function ---")
    val numbers = listOf(1, 2, 3, 4, 5)

    applyOperation(numbers) { it * it }
}
