package iniciante.loops.c;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class WhileLoopTest
{
    private final WhileLoop loop = new WhileLoop();

    @Test
    public void listOne()
    {
        int expected = 5;
        List<Integer> items = Arrays.asList(1, 2, 3, 4, 5);

        int result = loop.countItems(items.iterator());
        assertEquals(expected, result);
    }

    @Test
    public void listTwo()
    {
        int expected = 0;
        List<Integer> items = Collections.emptyList();

        int result = loop.countItems(items.iterator());
        assertEquals(expected, result);
    }

}