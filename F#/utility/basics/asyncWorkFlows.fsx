let job n = async{
    do! Async.Sleep 2000
    let result = n * n
    printfn "%d" result
}

let jobList = [job 2;job 3;job 4]
                                |> Async.Parallel 
                                |> Async.RunSynchronously