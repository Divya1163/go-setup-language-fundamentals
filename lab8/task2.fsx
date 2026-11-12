
// Lab 8 - Task 2
// CRUD Operations using SQLite and MongoDB

#r "nuget: Microsoft.Data.Sqlite, 9.0.0"
#r "nuget: MongoDB.Driver, 2.30.0"

open Microsoft.Data.Sqlite
open MongoDB.Bson
open MongoDB.Driver

// ==========================================
// PART 1: CRUD OPERATIONS USING SQLITE
// ==========================================

printfn "========== SQLITE DATABASE =========="

// Connect to SQLite database
use sqliteConnection =
    new SqliteConnection("Data Source=lab8_students.db")

sqliteConnection.Open()

// Create Student table
let createTable = sqliteConnection.CreateCommand()
createTable.CommandText <-
    "CREATE TABLE IF NOT EXISTS Students
     (id TEXT PRIMARY KEY, name TEXT, age INTEGER);"
createTable.ExecuteNonQuery() |> ignore

// Clear old test data to allow repeated execution
let clearTable = sqliteConnection.CreateCommand()
clearTable.CommandText <- "DELETE FROM Students;"
clearTable.ExecuteNonQuery() |> ignore

// Function to display all students
let printSQLiteStudents() =
    let command = sqliteConnection.CreateCommand()
    command.CommandText <- "SELECT * FROM Students;"

    use reader = command.ExecuteReader()

    printfn "Student Records:"

    if not (reader.HasRows) then
        printfn "No records found."

    while reader.Read() do
        printfn "ID: %s, Name: %s, Age: %d"
            (reader.GetString(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

// CREATE operation
printfn "\n1. CREATE - Adding a student"

let insertCommand = sqliteConnection.CreateCommand()
insertCommand.CommandText <-
    "INSERT INTO Students VALUES ($id, $name, $age);"

insertCommand.Parameters.AddWithValue("$id", "S001") |> ignore
insertCommand.Parameters.AddWithValue("$name", "Aarav") |> ignore
insertCommand.Parameters.AddWithValue("$age", 20) |> ignore

insertCommand.ExecuteNonQuery() |> ignore
printSQLiteStudents()

// READ operation
printfn "\n2. READ - Reading student details"

let readCommand = sqliteConnection.CreateCommand()
readCommand.CommandText <-
    "SELECT * FROM Students WHERE id = $id;"

readCommand.Parameters.AddWithValue("$id", "S001") |> ignore

use readReader = readCommand.ExecuteReader()

while readReader.Read() do
    printfn "Found: ID: %s, Name: %s, Age: %d"
        (readReader.GetString(0))
        (readReader.GetString(1))
        (readReader.GetInt32(2))

readReader.Close()

// UPDATE operation
printfn "\n3. UPDATE - Updating student details"

let updateCommand = sqliteConnection.CreateCommand()
updateCommand.CommandText <-
    "UPDATE Students SET name = $name, age = $age WHERE id = $id;"

updateCommand.Parameters.AddWithValue("$name", "Aarav Sharma") |> ignore
updateCommand.Parameters.AddWithValue("$age", 21) |> ignore
updateCommand.Parameters.AddWithValue("$id", "S001") |> ignore

updateCommand.ExecuteNonQuery() |> ignore
printSQLiteStudents()

// DELETE operation
printfn "\n4. DELETE - Removing student"

let deleteCommand = sqliteConnection.CreateCommand()
deleteCommand.CommandText <-
    "DELETE FROM Students WHERE id = $id;"

deleteCommand.Parameters.AddWithValue("$id", "S001") |> ignore

deleteCommand.ExecuteNonQuery() |> ignore
printSQLiteStudents()


// ==========================================
// PART 2: CRUD OPERATIONS USING MONGODB
// ==========================================

printfn "\n========== MONGODB DATABASE =========="

// Connect to MongoDB
let mongoClient =
    MongoClient("mongodb://localhost:27017")

let mongoDatabase =
    mongoClient.GetDatabase("Lab8Database")

let studentCollection =
    mongoDatabase.GetCollection<BsonDocument>("Students")

// Clear old test data to allow repeated execution
studentCollection.DeleteMany(
    Builders<BsonDocument>.Filter.Empty
) |> ignore

// Function to display all student documents
let printMongoStudents() =
    let students =
        studentCollection.Find(
            Builders<BsonDocument>.Filter.Empty
        ).ToList()

    printfn "Student Records:"

    if students.Count = 0 then
        printfn "No records found."

    for student in students do
        printfn "ID: %s, Name: %s, Age: %d"
            (student.["_id"].AsString)
            (student.["name"].AsString)
            (student.["age"].AsInt32)

// CREATE operation
printfn "\n1. CREATE - Adding a student"

let newStudent = BsonDocument()
newStudent.Add("_id", BsonString("S001")) |> ignore
newStudent.Add("name", BsonString("Aarav")) |> ignore
newStudent.Add("age", BsonInt32(20)) |> ignore

studentCollection.InsertOne(newStudent)
printMongoStudents()

// READ operation
printfn "\n2. READ - Reading student details"

let filter =
    Builders<BsonDocument>.Filter.Eq<string>("_id", "S001")

let foundStudent =
    studentCollection.Find(filter).FirstOrDefault()

if not (isNull foundStudent) then
    printfn "Found: ID: %s, Name: %s, Age: %d"
        (foundStudent.["_id"].AsString)
        (foundStudent.["name"].AsString)
        (foundStudent.["age"].AsInt32)

// UPDATE operation
printfn "\n3. UPDATE - Updating student details"

let update =
    Builders<BsonDocument>.Update
        .Set("name", "Aarav Sharma")
        .Set("age", 21)

studentCollection.UpdateOne(filter, update) |> ignore
printMongoStudents()

// DELETE operation
printfn "\n4. DELETE - Removing student"

studentCollection.DeleteOne(filter) |> ignore
printMongoStudents()

printfn "\nAll CRUD operations completed successfully!"
