package piscine

func GenerateAlphabet() []rune {
    var alphabet [26]rune
    a := 'a'
    for i := 0; i < 26; i++ {
        alphabet[i] = a + rune(i)
    }
    return alphabet[:]
}