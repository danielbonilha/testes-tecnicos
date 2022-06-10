package iniciante.conditionals.c;

public class AreaCalculation {
    private final double pi = 3.14;
    public double calcArea(Shape shape) {
        // TODO
        return -1.00;
    }
}

/*
Shape é uma forma.

Shape é definida como uma interface, que possui dois métodos:
- int getMultiplier, que retorna um multiplicador.
- int getMeasure, que retorna uma medida (um lado, no caso do quadrado e do
triângulo, e um raio no caso do círculo).

# Problema:

Retorne área da forma vezes seu multiplicador.
Se o multiplicador for menor ou igual à zero, lançar uma exception do
tipo IllegalArgumentException.

# Dicas:

A forma (shape) pode ser um quadrado, um círculo ou um triângulo.
Cada qual tem seu cálculo da área, baseado em sua medida (lado ou raio).

A área de um quadrado é medida pela medida * 2.
A área de um triângulo é medida pela medida * medida / 2.
A área de um círculo é medida pela medida elevada ao quadrado * Pi.

O valor de Pi já foi declarado para você.

Em java, o valor de A ao quadrado Pode ser calculado pelo pacote Math.pow().

Para descobrir qual o tipo da classe que representa a forma vc deve usar
o "instanceof".

# Testar:
mvn '-Dtest=iniciante.conditionals.c.AreaCalculationTest' test
 */
