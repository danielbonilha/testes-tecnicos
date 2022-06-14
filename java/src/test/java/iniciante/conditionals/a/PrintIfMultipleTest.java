package iniciante.conditionals.a;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.io.ByteArrayOutputStream;
import java.io.PrintStream;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class PrintIfMultipleTest {
    private final PrintStream standardOut = System.out;
    private final ByteArrayOutputStream outputStreamCaptor = new ByteArrayOutputStream();
    private final PrintIfMultiple printIfMultiple = new PrintIfMultiple();

    @BeforeEach
    public void setUp() {
        System.setOut(new PrintStream(outputStreamCaptor));
    }

    @AfterEach
    public void tearDown() {
        System.setOut(standardOut);
    }

    @Test
    public void multipleOf3Only()
    {
        printIfMultiple.printMultiples(6);
        assertEquals("Multiplo de 3", outputStreamCaptor.toString().trim());
    }

    @Test
    public void multipleOf5Only()
    {
        printIfMultiple.printMultiples(20);
        assertEquals("Multiplo de 5", outputStreamCaptor.toString().trim());
    }

    @Test
    public void multipleOf3And5()
    {
        printIfMultiple.printMultiples(30);
        assertEquals("Multiplo de 3 e 5", outputStreamCaptor.toString().trim());
    }

    @Test
    public void multipleOfNone()
    {
        printIfMultiple.printMultiples(11);
        assertEquals("", outputStreamCaptor.toString().trim());
    }
}
