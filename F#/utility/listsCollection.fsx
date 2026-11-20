// Creating a list
let nums1 = [1;2;3;4;5]


// Creating a list from 1 to 100
let nums2 = [1..100]



// Creating an empty list
let num3 = []


// List Map transforms to a new list
let double x = x * 2
let newNums = List.map double nums1


// List Filter keep some based on a condition
// let filteredList = List.filter(fun x -> x % 2 = 0) nums1
let checkEven x = x % 2 = 0
let filteredList = List.filter checkEven nums1

printfn "%A" filteredList

// List Sum adds every element is the list and returns a value
let sum = List.sum nums2
printfn "%d" sum
