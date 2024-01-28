* Use %+v in PrintF to super verbose printing

```go
package main

import "fmt"

type Player struct { 
    name string
    health int
    attackPower float64
}

func main() {
    player := Player{ 
        name: "Captain Jack",
        health: 100,
        attackPower: 45.1,
    }

    fmt.Printf("player: %+v\n", player)
}
```

