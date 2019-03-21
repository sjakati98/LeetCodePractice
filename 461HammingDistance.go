func hammingDistance(x int, y int) int {
    
    // calculate the XOR of the two integer bit representations
    differentBits := x ^ y
    
    // count the occurences of definitive differences
    var counter int
    for ok := true; ok; ok = differentBits > 0 {
        maskedBits := differentBits & 1
        if maskedBits != 0 {
            counter++
        }
        differentBits = differentBits >> 1
    }
    return counter
}