package play
// hello will return true if input.message and input.userid matches world and veeru respectively
default hello = false

hello {
    m := input.message
    q := input.userid
    m == "world"
    q == "veeru"
}
// magic will match if input.magic is expel expicitely 
default magic = false

magic {
	m := input.magic
    m == "expel"
}