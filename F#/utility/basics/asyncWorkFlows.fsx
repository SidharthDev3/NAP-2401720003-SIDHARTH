let job n = async{
    do! Async.Sleep 2000
    return n * n
}

let jobList = [job 2;job 3;job 4]
                                |> Async.Parallel 
                                |> Async.RunSynchronously