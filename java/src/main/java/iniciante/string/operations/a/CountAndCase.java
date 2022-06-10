package iniciante.string.operations.a;

public class CountAndCase {
    public String manipulate(String first, String second) {
        // TODO
        return "";
    }
}

/*
Strings são palavras ou textos, compostas pela combinação de 1 ou mais caracteres.
Case significa se a letra está maiúscula ou minúscula.
"Uppercase" significa todas as letras são maiúsculas.
"Lowercase" significa todas as letras são minúsculas.

# Problema

Dado uma String A e B, retorne uma ÚNICA string que concatene:
- A soma da contagem de caracteres de cada string.
- A primeira string com todas as letras maiúsculas.
- A segunda string com a todas letra minúsculas.

# Dicas

Concatenar é juntar uma string com a outra.
Concatenar "amora" com "abacaxi" resulta em "amoraabacaxi".

O tamanho de uma string é calculado pelo método .lenght().

Para converter todas as letras para maiúsculas utilize .toUpperCase().
Para converter todas as letras para maiúsculas utilize .toLowerCase().

# Exemplos

String 1: Amora
String 2: Abacaxi

Retorna: 12AMORAabacaxi

# Testar:
mvn '-Dtest=iniciante.string.operations.a.CountAndCase' test
 */