package iniciante.string.operations.b;

public class Substring {
    public String pieceOfString(String text) {
        // TODO
        return null;
    }
}

/*
Strings são compostos por chars.
Substrings partes de uma String, ou seja, um subconjunto de chars.

# Problema

Dado uma String A, retorne uma nova string sem as 3 primeiras letras, e sem as 2 últimas letras.

# Dicas

Cada char na String possui um índex. A contagem inicia no índex zero.
A palavra "uva" possui 3 chars,
sendo o char 0 correspondente à letra "u", o char 1 à letra "v" e o char 2 à letra "a".

O método substring pode receber 1 ou 2 parâmetros.
O primeiro parâmetro é o index inicial (incluso), e o segundo o índex final (não incluso).

Se um índice for utilizado e a palavra não conter o índice (por ser pequena), uma
java.lang.StringIndexOutOfBoundsException será lançada.

# Exemplos

"uva".substring(1) retorna "va", pois o índice 1 indica o início na letra "v".
"uva".substring(1,2) retorna "v", pois o índice 1 indica o início na letra "v",
e o índice 2 indica o final da palavra na letra "u", mas de forma exlusiva (exclui a letra u).
 */