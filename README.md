# Testes Técnicos

Testes técnicos para iniciantes.

A explicação dos testes é em português, mas as variáveis, os nomes dos métodos e
demais elementos dentro do software são em inglês.

## Linguagens suportadas:
- Java 11+

## Pré-requisitos

### Instalar Java 11 ou maior
> https://www.java.com/en/download/help/download_options.html

### Instalar maven:
> https://maven.apache.org/install.html
mvn archetype:generate -DgroupId=java -DartifactId=iniciante -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false

## Testes

### Java

#### Condicionais
`mvn test -Dtest=iniciante.conditionals.PrintIfMultipleTest`