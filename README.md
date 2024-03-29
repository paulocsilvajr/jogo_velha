# Jogo da velha(GOLANG)
### Desenvolvido no KDENeon User Edition 5.11(Ubuntu 16.04), Go1.9.2 linux/amd64.

Este repositório contém um exemplo de jogo da velha desenvolvido em Go.

### Pré-requisitos
Instalar a linguagem [Go](https://golang.org/dl/) e seguir os [passos](https://golang.org/doc/install) para adicionar o binário ao PATH do sistema.

Pode-se obter este repositório pelo utilitário go get:
```
$ go get -v github.com/paulocsilvajr/jogo_velha
```
Após o download, este repositório estará disponível em ~/go/src/github.com/paulocsilvajr/jogo_velha.

Para compilar o jogo, entre na pasta do jogo da velha e execute:
```
$ go build  # Gera o binário na pasta local
$ go install  # Instala o jogo da velha na pasta ~/go/bin
```
Ou compile via script './build.sh'

Para acessar a docstring, execute:
```
$ go doc github.com/paulocsilvajr/jogo_velha       # exibe pacotes
$ go doc github.com/paulocsilvajr/jogo_velha/jogo  # pacote jogo
$ go doc github.com/paulocsilvajr/jogo_velha/model # pacote model
$ go doc github.com/paulocsilvajr/jogo_velha/view  # pacote view
```

### Arquivos

```
jogo/jogador_vs_computador.go: Módulo do pacote jogo com funções relacionadas a jogadas do computador contra um jogador.
jogo/jogador_vs_jogador.go: Módulo do pacote jogo com funções relacionadas a jogadas de jogador contra jogador.
jogo/jogo.go: Módulo principal do pacote jogo que agrupa as funcionalidades do jogo da velha.
model/jogador.go: Módulo do pacote model que representa um jogador e suas ações.
model/tabuleiro.go: Módulo do pacote model que representa um tabuleiro e suas ações.
view/interface_usuario.go: Módulo do pacote view que concentra toda a interação do usuário com a interface do jogo.
jogo_velha.go: Agrupa todas as funcionalidades dos pacotes anteriores.
main.go: Módulo principal da aplicação.
build.sh: Script para compilar como o executável jogo_velha
```

### Licença

[Licença GPL](https://github.com/paulocsilvajr/jogo_velha-go/blob/master/license_gpl.txt), arquivo em anexo no repositório.

### Contato

Paulo Carvalho da Silva Junior - pauluscave@gmail.com
