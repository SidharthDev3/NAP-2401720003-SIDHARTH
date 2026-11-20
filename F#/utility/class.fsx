type Person(name,age) = 
    member this.Name = name
    member this.Age = age
    member this.Greet() = sprintf "Hi %s" name 


let p = Person("Sidharth",20)
p.Greet()