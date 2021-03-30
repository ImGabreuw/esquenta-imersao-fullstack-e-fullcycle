# React e TypeScript: Pontapé inicial do jeito certo

### O que é React

* Library baseado em JavaScript e mantida pelo Facebook.
* DOCS: [CLIQUE AQUI](https://pt-br.reactjs.org/docs/getting-started.html)

* Paradígmas mais forte nessa linguagem: Component

### O que é TypScript

* Framework baseado em JavaScript e mantida pela Microsoft.
* DOCS: [CLIQUE AQUI](https://www.typescriptlang.org/docs/)

* .ts = TypeScript
* .jsx = desenvolvimento de HTML e JS no mesmo aquivo

### O que é Angular

* Framework baseado em JavaScript e mantida pelo Google.
* Arquitetura de modules, injeção de dependência, formulários, HTTP

### Como criar um projeto usando create-react-app

* comando = npx create-react-app --template typescript _nome do projeto_

### Component

* Uma função que sempre retorna HTML.

    ```ts
    function App() {
        // Lógica dos eventos

        return (
            <div className="App">
                text
            </div>
        );
    }     
    ```

### DockerFile

* Preferível trabalhar com images non-root

### .dockerignore

* Pacotes que serão ignorados pelo docker