ifunc generateParenthesis(n int) []string {
    r := []string{}
    oc := 0
    cc := 0
    s := ""
    var f func ()

    f = func () {
        if oc == n && cc == n {
            r = append(r, s)
            return
        }

        if oc < n {
            s += "("; oc++
            f()
            s = s[0:len(s) - 1]; oc--
        }

        if cc < oc {
            s += ")"; cc++
            f()
            s = s[0:len(s) - 1]; cc--
        }
    }

    f()
    return r
}
