# Golang: Do zero ao Multithreading

### Golang

* Fortemente tipada
* Inferência de tipo

### Declaração de variáveis

```golang
type nome []string

var nome nome
nome = append(nome, "Wesley")
nome = append(nome, "Davi")

for k,v := range nome {
    fmt.Println(v)
}
```

### OO

```golang
type Pessoa struct {
    Nome string
    Idade int
}

func main() {
    wesley := Pessoa {
        Nome: "Wesley",
        Idade: 36,
    }

    fmt.Println(wesley.Nome)
}
```

### Função
```golang
package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func andou(pessoa Pessoa) {
	fmt.Println(pessoa.Nome, "andou")
}

func main() {
	wesley := Pessoa{
		Nome:  "Wesley",
		Idade: 36,
	}

	andou(wesley)
}
```

### Métodos para um "Objeto"

```golang
package main

import (
	"errors"
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

// (string, error) => retorno da função
func (p Pessoa) andou() (string, error) {
	if p.Nome != "Wesley" {
		return "", errors.New("Nome tem que ser Wesley")
	}

	return p.Nome + " andou", nil
}

func main() {
	wesley := Pessoa{
		Nome:  "Wesley",
		Idade: 36,
	}

	res, err := wesley.andou()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}

```

### Validação de erros
```golang
package main

import (
	"errors"
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func andou(pessoa Pessoa) (string, error) {
	if pessoa.Nome != "Wesley" {
		return "", errors.New("Nome tem que ser Wesley")
	}

	return pessoa.Nome + "andou", nil
}

func main() {
	wesley := Pessoa{
		Nome:  "Wesley",
		Idade: 36,
	}

	res, err := andou(wesley)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
```

### Endereço da variável na memória
```golang
func main() {
	nome := "Wesley"

    // &nomeDaVariavel = Endereço da variável na memória
	fmt.Println(&nome)
}
```

```golang
func main() {
	nome := "Wesley"
	var nome2 *string // ponteiro
	nome2 = &nome

	fmt.Println(*nome2)
}
```

### Exemplo de ponteiro

* Exemplo sem usar ponteiro
```golang
package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) setNome(nome string) {
	p.Nome = nome
	fmt.Println(p.Nome) // Gabriel
}

func main() {
	pessoa := Pessoa{
		Nome:  "Wesley",
		Idade: 36,
	}

	pessoa.setNome("Gabriel")
	fmt.Println(pessoa.Nome) // Wesley
}
```

* Exemplo usando ponteiro
```golang
package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) setNome(nome string) {
	p.Nome = nome
	fmt.Println(p.Nome) // Gabriel
}

func main() {
	pessoa := Pessoa{
		Nome:  "Wesley",
		Idade: 36,
	}

	pessoa.setNome("Gabriel")
	fmt.Println(pessoa.Nome) // Gabriel
}
```

### Tags em GO

```golang
type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
```

### Multithreading

* green thread (go routine) = threads gerenciadas pelo GO
* channel = conexão entre as threads
* keyword _go_ = geração de um nova thread (paralelo a thread principal)

<br>

* Exemplo
    ```golang
    package main

    import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "time"
    )

    func main() {
        requestId := make(chan int)
        concurrency := 2 // 2 workers

        for i := 0; i <= concurrency; i++ {
            go worker(requestId, 1) // go => geração de um nova thread (paralelo a thread principal)
        }

        for i := 0; i < 10; i++ {
            requestId <- i // atribuindo um valor a um channel
        }
    }

    func worker(requestId chan int, worker int) {
        for r := range requestId {
            res, err := http.Get("http://localhost:8585/product")

            if err != nil {
                log.Fatal("Erro")
            }

            defer res.Body.Close()

            content, _ := ioutil.ReadAll(res.Body)
            fmt.Printf("Worker %d. RequestId: %d Content: %s", worker, r, string(content))
            time.Sleep(time.Second * 2)
        }
    }
    ```
