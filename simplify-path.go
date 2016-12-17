import "strings"

func simplifyPath(path string) string {
    path += "/"
    s := []string{}
    i := 0

    for j, c := range path {
        if c == '/' {
            if j - i >= 1 {
                p := path[i:j]

                if p == "." {
                } else if p == ".." {
                    if len(s) >= 1 {
                        s = s[:len(s) - 1]
                    }
                } else {
                    s = append(s, p)
                }
            }

            i = j + 1
        }
    }

    return "/" + strings.Join(s, "/")
}
