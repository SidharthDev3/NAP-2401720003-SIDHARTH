let safeDiv a b =
    if b = 0 then None
    else Some (a/b)

match safeDiv 10 2 with
    | Some r -> printfn "%d" r
    | None -> printfn "no result"