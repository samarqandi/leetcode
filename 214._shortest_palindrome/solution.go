func reverse(input string) string {
    runes := []rune(input)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func shortestPalindrome(s string) string {
    length := len(s)
    revS := reverse(s)

    for i := range length {
        left := s[0:length-i]
        right := revS[i:length]
        if (left == right) {
            toAppend := s[length-i:]
            return revS + toAppend
        }
    }
    return s + revS
}
