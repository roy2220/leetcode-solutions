func minWindow(s string, t string) string {
    r := ""
    rl := len(s) + 1
    u := t + s
    ps := map[byte][]int{}

    for p := 1; p <= len(t); p++ {
        if ps[u[p - 1]] == nil {
            ps[u[p - 1]] = []int{p}
        } else {
            ps[u[p - 1]] = append(ps[u[p - 1]], p)
        }
    }

    sp := 1

    for p := len(t) + 1; p <= len(u); p++ {
        if ps[u[p - 1]] != nil {
            ps[u[p - 1]] = append(ps[u[p - 1]][1:], p)

            if u[p - 1] == u[sp - 1] {
                for {
                    sp++

                    if ps[u[sp - 1]] != nil && ps[u[sp - 1]][0] == sp {
                        break
                    }
                }

                if sp > len(t) {
                    if p - sp + 1 < rl {
                        r = u[sp - 1:p]
                        rl = p - sp + 1
                    }
                }
            }
        }
    }

    return r
}

// func minWindow(s string, t string) string {
//     r := ""
//     rl := len(s) + 1
//     u := t + s
//     ps := map[byte]int{}

//     for p := 1; p <= len(t); p++ {
//         ps[u[p - 1]] = p
//     }

//     sp := 1

//     for p := len(t) + 1; p <= len(u); p++ {
//         if ps[u[p - 1]] != 0 {
//             ps[u[p - 1]] = p

//             if u[p - 1] == u[sp - 1] {
//                 for {
//                     sp++

//                     if ps[u[sp - 1]] == sp {
//                         break
//                     }
//                 }

//                 if sp > len(t) {
//                     if p - sp + 1 < rl {
//                         r = u[sp - 1:p]
//                         rl = p - sp + 1
//                     }
//                 }
//             }
//         }
//     }

//     return r
// }
