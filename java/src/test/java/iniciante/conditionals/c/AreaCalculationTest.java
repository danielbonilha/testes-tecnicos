package iniciante.conditionals.c;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class AreaCalculationTest
{
    private final AreaCalculation area = new AreaCalculation();

    @Test
    public void squareArea()
    {
        double expectedArea = 200.0;
        Square square = new Square(10, 2);

        double result = area.calcArea(square);
        assertEquals(expectedArea, result);
    }

    @Test
    public void triangleArea()
    {
        double expectedArea = 250.0;
        Triangle square = new Triangle(10, 5);

        double result = area.calcArea(square);
        assertEquals(expectedArea, result);
    }

    @Test
    public void circleArea()
    {
        double expectedArea = 392.5;
        Circle square = new Circle(5, 5);

        double result = area.calcArea(square);
        assertEquals(expectedArea, result);
    }

    @Test
    public void invalidShape()
    {
        Square square = new Square(5, 0);
        assertThrows(IllegalArgumentException.class, () -> area.calcArea(square));
    }
}