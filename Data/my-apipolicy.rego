
package myapi.policy
import data.myapi.acl
import input
import data.grants

default allow = false
allow {
        access = acl[input.user]
        access[_] == input.access
}
whocan[user] {
        access = acl[user]
        access[_] == input.access
}