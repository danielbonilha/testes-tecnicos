package iniciante.loops.b;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class DecreasingLoopTest
{
    private final DecreasingLoop loop = new DecreasingLoop();

    @Test
    public void listOne()
    {
        List<Integer> expected = Arrays.asList(5, 4, 3, 2, 1);
        List<Integer> items = Arrays.asList(1, 2, 3, 4, 5);

        List<Integer> result = loop.reverseItems(items);
        assertEquals(expected, result);
    }

    @Test
    public void listTwo()
    {
        List<Integer> expected = Arrays.asList(-5, 10, -1, 0, 15, -100);
        List<Integer> items = Arrays.asList(-100, 15, 0, -1, 10, -5);

        List<Integer> result = loop.reverseItems(items);
        assertEquals(expected, result);
    }
}