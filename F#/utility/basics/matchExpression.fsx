let name n = 
    match n with
    | 1 -> "one"
    | 2 -> "two"
    | 3 -> "three"
    | 4 -> "four"
    | _ -> "many"

printfn "%s" (name 12)
