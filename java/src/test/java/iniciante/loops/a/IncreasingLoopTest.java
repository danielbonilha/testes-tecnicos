package iniciante.loops.a;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class IncreasingLoopTest
{
    private final IncreasingLoop loop = new IncreasingLoop();

    @Test
    public void listOne()
    {
        int expected = 0;
        List<Integer> items = Arrays.asList(1, 2, 3, 4, 5);

        int result = loop.sumItems(items);
        assertEquals(expected, result);
    }

    @Test
    public void listTwo()
    {
        int expected = 150;
        List<Integer> items = Arrays.asList(10, 20, 30, 40, 50);

        int result = loop.sumItems(items);
        assertEquals(expected, result);
    }

    @Test
    public void listThree()
    {
        int expected = 10;
        List<Integer> items = Arrays.asList(-10, -5, 0, 5, 10);

        int result = loop.sumItems(items);
        assertEquals(expected, result);
    }

}