# Docker e Portainer do zero: Containers em desenvolvimento

### O que é Docker

* Documentação: [CLIQUE AQUI](https://docs.docker.com/)

* Docker também é um container runtime.

* Container = 1 processo
* Image = tudo que tem dentro do container
* Images são imutáveis

### Docker Hub

* "GitHub" para armazenar as imagens do docker

### Docker Compose

* Subir o ambiente com apenas 1 comando

### Comandos essenciais

* docker run ubuntu => baixar a Image do Ubuntu

* docker ps => listar todos os containers rodando no momento
* docker ps -a => listar todos os containers (rodando ou não)

* docker run -p _porta:porta_ _nome da imagem_ => redirecionar de uma porta para outra
* docker run -d -p _porta:porta_ _nome da image_ => destravar o terminal, mesmo rodando um container

* docker exec _nome do container_ _comando_ => execcutar um comando dentro do container
* docker exec -it _nomo da imagem_ bash => entrar no container como root

* docker images => listar todas as images (rodando ou não)

* docker push _nome da imagem_ => upload para o Docker Hub

* docker-compose up -d => rodar o container sem travar o terminal
* docker-compose up => rodar o container porém trava o terminal

* docker-compose down => matar o container

### Variáveis de ambiente

* dockerFile
    ```dockerFile
    ENV key=value
    ```

* docker-compose.yml
    ```yml
    mysql:
        environment:
        - variável de ambiente 1
        - variável de ambiente 2
    ```

### networks

* Por padrão os container são interligados, ou seja, na mesma network.

* Para separa os containers é necessário separar as "networks" com nomes diferentes.
    ```yml
    version: '3'

    services: 
        web:
            networks: 
            - x

        mysql:
            networks: 
            - y
    ```

### portainer.io

* Ferramenta para gerenciar o Docker.

### OBS

* caso não for informado o nome do container, é gerado um nome aleatório.

* cada container para cada responsabilidade

* docker run hello-world
    ```bash
    Hello from Docker!
    This message shows that your installation appears to be working correctly.    

    To generate this message, Docker took the following steps:
    1. The Docker client contacted the Docker daemon.
    2. The Docker daemon pulled the "hello-world" image from the Docker Hub.     
        (amd64)
    3. The Docker daemon created a new container from that image which runs the  
        executable that produces the output you are currently reading.
    4. The Docker daemon streamed that output to the Docker client, which sent it
        to your terminal.

    To try something more ambitious, you can run an Ubuntu container with:        
    $ docker run -it ubuntu bash

    Share images, automate workflows, and more with a free Docker ID:
    https://hub.docker.com/

    For more examples and ideas, visit:
    https://docs.docker.com/get-started/
    ```

* docker run nginx (container fica de "pé")