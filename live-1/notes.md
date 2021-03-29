# Ambiente dos sonhos no Windows com WSL2 e VSCode

## WSL2 (Windows Subsystem Linux)

* Melhorias em relação à V1
    * Kernel completo do Linux embarcado
    * Grande compatibilidade com programas e ferramentas Linux
    * Mais rápido que o WSL 1
    * Grande desempenho executando dentro do Linux
    * Suporte ao Docker, Kubernetes e ao padrão de containers
    * Uso do Virtual Machine Plataform

* Requisitos para rodas o WSL 2
    * Windows 10 Home ou Pro
    * Windows 10 versão >= 19.03
    * 4 GB RAM (Recomendado: 8GB RAM)
    * Virtual Machine Plataform

---

## Docker

* Docker Toolbox (DEPRECIADO)
    * Roda com VirtualBox
    * É muito lento

* Docker Desktop
    * Roda com Hyper-V
    * Precisa de Windows Pro
    * Exige mais recursos da máquina
    * Desempenho superior do Docker Toolbox

* [Tutorial](https://github.com/codeedu/wsl2-docker-quickstart)

---

## Windows Terminal

* É um terminal, não é um shell

* Funciona com WSL, DOS, PowerShell

* Instalar: [oh my zsh no WSL 2](https://dev.to/erickrock80/pt-br-instalando-oh-my-zsh-no-windows-terminal-3a8l)
    * Autocomplete
    * highlights
    * correção de comandos
    * Temas e cores
    * Gerência avançada do histórico
    * Vários Plug-ins

    * Tema no zsh: sobole, spaceship

---

## PowerToys

* Seletor de cor (em Hexadecimal)
* FancyZones
* Complementos do Windows Explorer
* Preview de arquivos .md e .svg
* Redimensionador de imagem
* Gerenciador de Teclado
* Renomear avançados, guia de atalhos, mudo em video conferêcia
* PowerToys Run

* DOWNLOAD: [CLIQUE AQUI](https://github.com/microsoft/PowerToys)

---

## VSCode (Visual Studio Code)

* Suporte a multi-linguagens
* Continuamente melhorado e atualizado
* Zelo da Microsoft em disponibilizar várias extensões úteis
* Loja de extensãoes imensa e comunidade de desenvolvimento muito ativa
* Leve e rápido quando comparado a outras IDEs
* Suporte a desenvolvimento remoto e Docker

* extensões
    * Settings Sync (https://gist.github.com/argentinaluiz/f134417fa4753064820b6ee2af194e55)
    * Remote Development
    * Remote - SSH
    * Peacock (cor da IDE por tecnologia)