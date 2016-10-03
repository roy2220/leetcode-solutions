func isMatch(s string, p string) bool {
    return regexMatch(s, p)
}


func regexMatch(s string, p string) bool {
    si := 0
    pi := 0
    var doRegexMatch func () bool

    doRegexMatch = func () bool {
        if pi == len(p) {
            return si == len(s)
        }

        stared := pi + 1 < len(p) && p[pi + 1] == '*'
        matched := si < len(s) && (s[si] == p[pi] || p[pi] == '.')

        if stared {
            if matched {
                var m bool

                si += 1
                m = doRegexMatch()
                si -= 1

                if m {
                    return true
                }

                si += 1; pi += 2
                m = doRegexMatch()
                si -= 1; pi -= 2

                if m {
                    return true
                }
            }

            pi += 2
            m := doRegexMatch()
            pi -= 2

            if m {
                return true
            }
        } else {
            if matched {
                si += 1; pi += 1
                m := doRegexMatch()
                si -= 1; pi -= 1

                if m {
                    return true
                }
            }
        }

        return false
    }

    return doRegexMatch()
}
