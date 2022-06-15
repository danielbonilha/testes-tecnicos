package iniciante.string.operations.c;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class StartsEndsWithTest
{
    private final StartsEndsWith manipulation = new StartsEndsWith();

    @Test
    public void aStart()
    {
        Boolean expected = true;

        Boolean result = manipulation.startsWith("abacaxi");
        assertEquals(expected, result);
    }

    @Test
    public void bStart()
    {
        Boolean expected = false;

        Boolean result = manipulation.startsWith("banana");
        assertEquals(expected, result);
    }
    
    @Test
    public void uStart()
    {
        Boolean expected = true;

        Boolean result = manipulation.startsWith("uva");
        assertEquals(expected, result);
    }

    @Test
    public void emptyStart()
    {
        Boolean expected = false;

        Boolean result = manipulation.startsWith("");
        assertEquals(expected, result);
    }

    @Test
    public void aEnd()
    {
        Boolean expected = true;

        Boolean result = manipulation.endsWith("banana");
        assertEquals(expected, result);
    }

    @Test
    public void rEnd()
    {
        Boolean expected = false;

        Boolean result = manipulation.endsWith("comer");
        assertEquals(expected, result);
    }

    @Test
    public void uEnd()
    {
        Boolean expected = true;

        Boolean result = manipulation.endsWith("cupuaçu");
        assertEquals(expected, result);
    }

    @Test
    public void emptyEnd()
    {
        Boolean expected = false;

        Boolean result = manipulation.endsWith("");
        assertEquals(expected, result);
    }
}