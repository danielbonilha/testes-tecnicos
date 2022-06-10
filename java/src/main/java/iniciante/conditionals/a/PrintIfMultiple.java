package iniciante.conditionals.a;

public class PrintIfMultiple {
    public void printMultiples(int n) {
        // TODO
    }
}

/*
# Problema

Dado um número inteiro n:
- Imprima "Multiplo de 3" se o número for múltiplo de 3, somente.
- Imprima "Multiplo de 5" se o número for múltiplo de 5, somente.
- Imprima "Multiplo de 3 e 5" se o número for múltiplo de 3 e 5.

# Dicas

Múltiplos, em programação, são calculados através do resto de uma divisão.
Se a divisão de um número A por um outro número B tiver RESTO
igual à 0 (zero), então o número A é múltiplo do número B.

Em programação usa-se o operador % (modulus) para tal cálculo.
O modulus retorna sempre o RESTO de uma divisão, e não o resultado da divisão.

# Exemplos
15 % 3 retorna 0 (zero), pois o resultado da divisão de 15 por 3 tem ZERO de resto.
16 % 3 retorna 1 (um), pois o resultado da divisão de 16 por 3 tem UM de resto.

# Testar:
mvn '-Dtest=iniciante.conditionals.a.PrintIfMultipleTest' test
*/