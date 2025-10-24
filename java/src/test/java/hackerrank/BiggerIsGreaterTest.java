package hackerrank;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class BiggerIsGreaterTest {

    @Test
    void testSimpleCase() {
        assertEquals("abdc", BiggerIsGreater.biggerIsGreater("abcd"));
    }

    @Test
    void testAlreadyHighest() {
        assertEquals("no answer", BiggerIsGreater.biggerIsGreater("dcba"));
    }

    @Test
    void testMiddlePermutation() {
        assertEquals("hegf", BiggerIsGreater.biggerIsGreater("hefg"));
    }

    @Test
    void testSingleSwap() {
        assertEquals("dhkc", BiggerIsGreater.biggerIsGreater("dhck"));
    }

    @Test
    void testMultipleSameLetters() {
        assertEquals("hcdk", BiggerIsGreater.biggerIsGreater("dkhc"));
    }

    @Test
    void testAllSameLetters() {
        assertEquals("no answer", BiggerIsGreater.biggerIsGreater("aaaa"));
    }

    @Test
    void testLongerWord() {
        assertEquals("abdc", BiggerIsGreater.biggerIsGreater("abcd"));
    }

    @Test
    void testLastTwoSwap() {
        assertEquals("abdc", BiggerIsGreater.biggerIsGreater("abcd"));
    }

    @Test
    void testNoAnswer() {
        assertEquals("no answer", BiggerIsGreater.biggerIsGreater("zyx"));
    }
}
