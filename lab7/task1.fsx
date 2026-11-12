
// Lab 7 - Task 1
// Concurrent execution using F# asynchronous workflows

// First asynchronous task
let task1 =
    async {
        printfn "Task 1: Started"
        do! Async.Sleep(3000)
        printfn "Task 1: Completed"
        return "Result from Task 1"
    }

// Second asynchronous task
let task2 =
    async {
        printfn "Task 2: Started"
        do! Async.Sleep(2000)
        printfn "Task 2: Completed"
        return "Result from Task 2"
    }

// Third asynchronous task
let task3 =
    async {
        printfn "Task 3: Started"
        do! Async.Sleep(1000)
        printfn "Task 3: Completed"
        return "Result from Task 3"
    }

// Store all asynchronous workflows
let allTasks = [| task1; task2; task3 |]

printfn "Starting all asynchronous tasks...\n"

// Run all tasks concurrently
let results = Async.RunSynchronously (Async.Parallel allTasks)

// Display collected results
printfn "\nAll tasks have finished!"
printfn "Collected Results:"

results
|> Array.iter (printfn "%s")
