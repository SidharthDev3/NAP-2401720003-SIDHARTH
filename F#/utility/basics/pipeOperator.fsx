let nums = [1..6] 
                |> List.filter (fun x -> x % 2 = 0) 
                |> List.map(fun x -> x * x) 
                |> List.sum

printfn "%A" nums

