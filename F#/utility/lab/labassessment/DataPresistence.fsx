type Student(Id,name,marks) = 
    member this.Id = id
    member this.Name = name
    member this.Marks = marks
    member this.Greet() = sprintf "Hi %s" name 

let students = [
    Student(1, "Sidharth", 90)
    Student(2, "Tom", 65)
    Student(3, "Elizabeth", 85)
]

let names = students |> List.map (fun s -> s.Name)
let toppers = students |> List.filter (fun s -> s.Marks >= 80)
let result = toppers |> List.iter (fun s -> s.Name )

printfn "%s" result

let getResult marks =
    match marks with
    | m when m >= 80 -> "Distinction"
    | m when m >= 60 -> "First Class"
    | m when m >= 40 -> "Pass"
    | _ -> "Fail"

printfn "\nStudent Results:"
students |> List.iter (fun s ->
    printfn "%s : %s" s.Name (getResult s.Marks)
)