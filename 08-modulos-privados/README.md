Para usar repositórios privados no projeto, rode: 
```bash
export GOPRIVATE=url dos repositórios privados específicos separados por vírgula
```
Caso o go não esteja credenciado no github, tem duas formas de se credenciar:
- Usando login e token (autenticação via http):
    ```bash
    vim ~/.netrc
    ```
    Adicione as seguintes linhas:
    ```
    machine github.com
    login <seu login>
    password <seu token>
    ```
- Usando chave ssh:
    ```bash
    vim ~/.gitconfig
    ```
  no arquivo adicione as seguintes linhas:
    ```
    [url "ssh://git@github.com/"]
        insteadOf = https://github.com/
    ```