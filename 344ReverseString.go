func reverseString(s []byte)  []byte{
    
    // if the string is empty of one character, safely return that string
    if len(s) == 0 || len(s) == 1 {
        return s
    }
    
    
    // use left right anchor to replace
    // initialize anchors
    left := 0
    right := len(s) - 1
    for ok := true; ok; ok = left < right {
        // perform the swap
        temp := s[left]
        s[left] = s[right]
        s[right] = temp
        
        left++
        right--
    }
    return s
}