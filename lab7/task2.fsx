
// Lab 7 - Task 2
// Demonstrating Higher-Order Functions in F#

// Two different lists
let list1 = [1; 2; 3; 4; 5]
let list2 = [10; 15; 20; 25; 30]

// 1. List.map - Square each element
let squareList1 = List.map (fun x -> x * x) list1
let squareList2 = List.map (fun x -> x * x) list2

printfn "1. List.map - Squaring elements:"
printfn "Original List 1: %A" list1
printfn "Squared List 1: %A" squareList1
printfn "Original List 2: %A" list2
printfn "Squared List 2: %A" squareList2

// 2. List.filter - Select even numbers
let evenList1 = List.filter (fun x -> x % 2 = 0) list1
let evenList2 = List.filter (fun x -> x % 2 = 0) list2

printfn "\n2. List.filter - Selecting even numbers:"
printfn "Original List 1: %A" list1
printfn "Even Numbers 1: %A" evenList1
printfn "Original List 2: %A" list2
printfn "Even Numbers 2: %A" evenList2

// 3. List.fold - Calculate sum of elements
let sumList1 = List.fold (fun acc x -> acc + x) 0 list1
let sumList2 = List.fold (fun acc x -> acc + x) 0 list2

printfn "\n3. List.fold - Calculating sum:"
printfn "List 1: %A" list1
printfn "Sum of List 1: %d" sumList1
printfn "List 2: %A" list2
printfn "Sum of List 2: %d" sumList2
