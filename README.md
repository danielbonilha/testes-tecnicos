# Testes Técnicos

Testes técnicos para iniciantes.

A explicação dos testes é em português, mas as variáveis, os nomes dos métodos e
demais elementos dentro do software são em inglês.

## Como utilizar:
1. Crie uma nova branch a partir da `main`:
`git checkout -b minhas_respostas`
2. Leia as instruções do teste e complete o código nos arquivos que possuem a anotação `// TODO`
3. Rode os testes da linguagem conforme instruções de cada linguagem abaixo
4. Não precisa fazer o push da sua branch para o repositório principal. Deixe local.
5. Atualize sua branch regularmente com os novos testes sempre que houver
`git checkout main`
`git pull`
`git checkout minhas_repostas`
`git merge main`

## Linguagens suportadas:
- Java 11+

## Instruções para testes em Java

### Instalar Java 11 ou maior

Verifique se o java está instalado em seu computador:
`java -version`

Caso não esteja instalado, seguir as instruções em
> https://www.java.com/en/download/help/download_options.html

### Instalar maven:

Maven é uma ferramenta que auxilia no uso de dependências, e ajuda a rodar os testes.
Verifique se maven está instalado:
`mvn -version`

Caso não esteja instalado, seguir as instruções em
> https://maven.apache.org/install.html

### Testes

Para rodar os testes em java (a partir da raíz do projeto):
`mvn test -f java`

Os testes disponíveis são:
- Condicionais
- Loops
- Manipulação de strings

Complete o código nos arquivos abaixo.
Depois, para cada arquivo completado, rode os testes para verificar.

#### Condicionais
- src/java/iniciante/condicionals/a/PrintIfMultiple.java
- src/java/iniciante/condicionals/b/Toll.java
- src/java/iniciante/condicionals/c/AreaCalculation.java

#### Loops
- src/java/iniciante/loops/a/IncreasingLoop.java
- src/java/iniciante/loops/b/DecreasingLoop.java
- src/java/iniciante/loops/c/WhileLoop.java

#### Manipulação de Strings
- src/java/iniciante/string/operations/a/CountAndCase.java
- src/java/iniciante/string/operations/b/Substring.java
- src/java/iniciante/string/operations/c/StartsEndsWith.java