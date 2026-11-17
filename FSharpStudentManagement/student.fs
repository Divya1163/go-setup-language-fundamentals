
module StudentManagement

open Microsoft.Data.Sqlite

// Domain model using a discriminated union
type StudentStatus =
    | Active
    | Inactive

// Student record
type Student = {
    Id: int
    Name: string
    Age: int
    Marks: float
    Status: StudentStatus
}

// Pattern matching to convert status into string
let statusToString status =
    match status with
    | Active -> "Active"
    | Inactive -> "Inactive"

// Pattern matching to assign grades
let getGrade marks =
    match marks with
    | m when m >= 90.0 -> "A"
    | m when m >= 75.0 -> "B"
    | m when m >= 50.0 -> "C"
    | _ -> "Fail"

// Higher-order function: map
let getStudentGrades students =
    students
    |> List.map (fun student ->
        student.Name, getGrade student.Marks)

// Higher-order function: filter
let getTopStudents students =
    students
    |> List.filter (fun student -> student.Marks >= 75.0)

// SQLite connection string
let connectionString = "Data Source=students.db"

// Create database table asynchronously
let createTable () = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <-
        """
        CREATE TABLE IF NOT EXISTS Students (
            Id INTEGER PRIMARY KEY,
            Name TEXT NOT NULL,
            Age INTEGER NOT NULL,
            Marks REAL NOT NULL,
            Status TEXT NOT NULL
        )
        """

    let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask
    printfn "Student table created successfully."
}

// CREATE operation
let addStudent (student: Student) = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <-
        """
        INSERT INTO Students (Id, Name, Age, Marks, Status)
        VALUES (@id, @name, @age, @marks, @status)
        """

    command.Parameters.AddWithValue("@id", student.Id) |> ignore
    command.Parameters.AddWithValue("@name", student.Name) |> ignore
    command.Parameters.AddWithValue("@age", student.Age) |> ignore
    command.Parameters.AddWithValue("@marks", student.Marks) |> ignore
    command.Parameters.AddWithValue("@status", statusToString student.Status) |> ignore

    let! _ = command.ExecuteNonQueryAsync() |> Async.AwaitTask
    printfn "Student added successfully: %s" student.Name
}

// READ operation
let getStudents () = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()
    command.CommandText <- "SELECT * FROM Students"

    let! reader = command.ExecuteReaderAsync() |> Async.AwaitTask

    use reader = reader
    let students = ResizeArray<Student>()

    let mutable hasRows = true

    while hasRows do
        let! hasNext = reader.ReadAsync() |> Async.AwaitTask
        hasRows <- hasNext

        if hasNext then
            let status =
                match reader.GetString(4) with
                | "Active" -> Active
                | _ -> Inactive

            students.Add({
                Id = reader.GetInt32(0)
                Name = reader.GetString(1)
                Age = reader.GetInt32(2)
                Marks = reader.GetDouble(3)
                Status = status
            })

    return students |> Seq.toList
}

// UPDATE operation
let updateStudent id newMarks = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <-
        "UPDATE Students SET Marks = @marks WHERE Id = @id"

    command.Parameters.AddWithValue("@marks", newMarks) |> ignore
    command.Parameters.AddWithValue("@id", id) |> ignore

    let! rows = command.ExecuteNonQueryAsync() |> Async.AwaitTask

    if rows > 0 then
        printfn "Student updated successfully."
    else
        printfn "Student not found."
}

// DELETE operation
let deleteStudent id = async {
    use connection = new SqliteConnection(connectionString)
    do! connection.OpenAsync() |> Async.AwaitTask

    use command = connection.CreateCommand()

    command.CommandText <-
        "DELETE FROM Students WHERE Id = @id"

    command.Parameters.AddWithValue("@id", id) |> ignore

    let! rows = command.ExecuteNonQueryAsync() |> Async.AwaitTask

    if rows > 0 then
        printfn "Student deleted successfully."
    else
        printfn "Student not found."
}
