# KPerms

Generate the k permutations for n elements in golang. Relies on a modifed Johnson's SEPA algorithm as described in this great post https://alistairisrael.wordpress.com/2009/09/22/simple-efficient-pnk-algorithm/

## Usage

```
kp,err := kperm.New(3, 2)
...
kp.MaxIndex() // 5
kp.Perm()     // [0 1]
kp.Index()    // 0
kp.Done()     // false
kp.Next()
kp.Perm()     // [0 2]
kp.Index()    // 1
kp.Done()     // false
kp.Next()
kp.Perm()     // [1 0]
kp.Index()    // 2
kp.Done()     // false
kp.Next()
kp.Perm()     // [1 2]
kp.Index()    // 3
kp.Done()     // false
kp.Next()
kp.Perm()     // [2 0]
kp.Index()    // 4
kp.Done()     // false
kp.Next()
kp.Perm()     // [2 1]
kp.Index()    // 5
kp.Done()     // true
kp.Reset()
kp.Perm()     // [0 1]
kp.Index()    // 0
kp.Done()     // false
```

## Test

``` bash
$ go test ./... -v
```
