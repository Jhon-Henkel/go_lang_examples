Por convenção, quando fizer o go mod init, deve-se colocar como nome do pacote a url do repositório, por que se alguém for baixar esse módulo, o go vai baixar diretamente da url.

Exemplo:
```
go mod init github.com/Jhon-Henkel/go_lang_examples/05-packaging/01-introduzindo-o-go-mod
```

A pasta cmd é onde ficam os executáveis, como por exemplo o arquivo main.

para importar vai ter que ter o nome do pacote e o nome do arquivo, por exemplo: github.com/Jhon-Henkel/go_lang_examples/05-packaging/01-introduzindo-o-go-mod/math