# Usefull Golang funcs


## Sample Usage

``` go
package main

import (
	"fmt"

	e "github.com/derivaro/golibri"
)

func main() {
	bases := e.SetBases(".")
	dt, sterr := e.Dsql((*bases)["databasealias1"], "select * from table limit 10;")
	if sterr != "" {
		fmt.Println(sterr)
	}
	for r := 0; r < dt.RowsCount; r++ {
		fmt.Println((*dt.Rows)[r].FI[1])
	}
}
```


## Config file (config.yaml) for databases alias:

```yaml

envName:
  keys:
	- VARENV1
	- VARENV2
  databases:
    - name: databasealias1
      url: >-
        host=$VARENV1 port=5431 user=userName
        password=*** dbname=test sslmode=disable
      typ: postgres
   

```

