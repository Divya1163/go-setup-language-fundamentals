
open StudentManagement

[<EntryPoint>]
let main argv =

    // Sample student records
    let students = [
        {
            Id = 1
            Name = "Aman"
            Age = 21
            Marks = 92.0
            Status = Active
        }
        {
            Id = 2
            Name = "Riya"
            Age = 22
            Marks = 82.0
            Status = Active
        }
        {
            Id = 3
            Name = "Karan"
            Age = 21
            Marks = 45.0
            Status = Inactive
        }
    ]

    printfn "\n--- Student Grades (MAP) ---"

    let grades = getStudentGrades students

    grades
    |> List.iter (fun (name, grade) ->
        printfn "%s : Grade %s" name grade)

    printfn "\n--- Top Students (FILTER) ---"

    let topStudents = getTopStudents students

    topStudents
    |> List.iter (fun student ->
        printfn "%s - %.1f marks" student.Name student.Marks)

    // Asynchronous database operations
    async {
        printfn "\n--- DATABASE OPERATIONS ---"

        // Create table
        do! createTable ()

        // CREATE
        for student in students do
            do! addStudent student

        // READ
        printfn "\nReading Student Records:"

        let! records = getStudents ()

        records
        |> List.iter (fun student ->
            printfn "ID: %d | Name: %s | Age: %d | Marks: %.1f"
                student.Id student.Name student.Age student.Marks)

        // UPDATE
        printfn "\nUpdating Student 1:"
        do! updateStudent 1 96.0

        // DELETE
        printfn "\nDeleting Student 3:"
        do! deleteStudent 3

        // READ after update and delete
        printfn "\nFinal Student Records:"

        let! finalRecords = getStudents ()

        finalRecords
        |> List.iter (fun student ->
            printfn "ID: %d | Name: %s | Marks: %.1f"
                student.Id student.Name student.Marks)
    }
    |> Async.RunSynchronously

    0
