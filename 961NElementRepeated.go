func repeatedNTimes(A []int) int {
    
    // create the lookup for keeping track of element frequencies
    lookup := map[int]int{}
    
    // iterate through the array and update lookup
    for _, num := range A {
        
        //check if number exists as key in lookup
        _, ok := lookup[num]
        if !ok { // if it does not exist, then initialize to 1
            lookup[num] = 1
        } else { // if it does exist, then increment
            lookup[num] = lookup[num] + 1
        }
    }
    
    // initialize number of times to be repeated
    targetRepeats := len(A) / 2
    var toReturn int
    // iterate over lookup and return correct value
    for k, v := range lookup {
        if v == targetRepeats {
            toReturn = k
        }
    }
    return toReturn
}