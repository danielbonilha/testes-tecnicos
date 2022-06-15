package iniciante.string.operations.b;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class SubstringTest
{
    private final Substring manipulation = new Substring();

    @Test
    public void tenChars()
    {
        String expectedString = "erati";

        String result = manipulation.pieceOfString("generation");
        assertEquals(expectedString, result);
    }

    @Test
    public void zeroChars()
    {
        String expectedString = "";

        String result = manipulation.pieceOfString("");
        assertEquals(expectedString, result);
    }
}