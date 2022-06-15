package iniciante.string.operations.a;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class CountAndCaseTest
{
    private final CountAndCase manipulation = new CountAndCase();

    @Test
    public void twoWords()
    {
        String expectedString = "12AMORAabacaxi";

        String result = manipulation.manipulate("Amora", "Abacaxi");
        assertEquals(expectedString, result);
    }

    @Test
    public void twoEmptyWords()
    {
        String expectedString = "0";

        String result = manipulation.manipulate("", "");
        assertEquals(expectedString, result);
    }
}