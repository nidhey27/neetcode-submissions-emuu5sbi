func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }

    sFreq := make(map[string]int)
    tFreq := make(map[string]int)

    for i := 0; i < len(s); i++ {
        sFreq[string(s[i])] += 1
        tFreq[string(t[i])] += 1
    }

    for key, _ := range sFreq {
        if sFreq[key] != tFreq[key] {
            return false
        }
    }
    return true
}
