package iniciante.algorithms.palindrome;

import iniciante.algorithms.palindome.Palindome;
import iniciante.conditionals.b.Toll;
import iniciante.conditionals.b.Vehicle;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class PalindromeTest
{
    private final Palindome palindrome = new Palindome();

    @Test
    public void simpleTest()
    {
        Boolean expected = true;
        String text = "arara";

        Boolean result = palindrome.isPalindrome(text);
        assertEquals(expected, result);
    }

    @Test
    public void upperCasedTest()
    {
        Boolean expected = true;
        String text = "Osso";

        Boolean result = palindrome.isPalindrome(text);
        assertEquals(expected, result);
    }

    @Test
    public void withSpacesTest()
    {
        Boolean expected = true;
        String text = "a grama é amarga";

        Boolean result = palindrome.isPalindrome(text);
        assertEquals(expected, result);
    }

    @Test
    public void withNumbersTest()
    {
        Boolean expected = true;
        String text = "20/02/2002";

        Boolean result = palindrome.isPalindrome(text);
        assertEquals(expected, result);
    }

    @Test
    public void complexPhraseTest()
    {
        Boolean expected = true;
        String text = "Aí, Lima falou: 'Olá, família!'";

        Boolean result = palindrome.isPalindrome(text);
        assertEquals(expected, result);
    }


    @Test
    public void notPalindromeTest()
    {
        Boolean expected = false;
        String text = "Não sou palíndromo";

        Boolean result = palindrome.isPalindrome(text);
        assertEquals(expected, result);
    }
}