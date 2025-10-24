package sigma;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public class MenuRecommendationTest {

    private MenuRecommendation menu;

    @BeforeEach
    void setup() {
        menu = new MenuRecommendation();
    }

    @Test
    void testAddItemAndGetRecommendation() {
        menu.addItem(1, "Item1");

        MenuItem recommended = menu.getRecommendedItem();
        assertNotNull(recommended, "Should recommend newly added item");
        assertEquals(1, recommended.itemId);
        assertEquals("Item1", recommended.displayName);
    }

    @Test
    void testRateItemAffectsRecommendation() {
        menu.addItem(1, "Pizza");
        menu.addItem(2, "Burger");

        menu.rateItem(1, 5);
        menu.rateItem(2, 3);

        MenuItem recommended = menu.getRecommendedItem();
        assertEquals(1, recommended.itemId, "Pizza has higher rating and should be recommended");
    }

    @Test
    void testDealOfTheDayOverridesNormalRecommendation() {
        menu.addItem(1, "Pizza");
        menu.addItem(2, "Burger");
        menu.rateItem(1, 2);
        menu.rateItem(2, 5);

        // Make Pizza the deal of the day
        menu.makeDealOfTheDayItem(1);

        MenuItem recommended = menu.getRecommendedItem();
        assertEquals(1, recommended.itemId, "Deal of the day should override normal recommendation");
    }

    @Test
    void testDealOfTheDayRemovedIfOutOfStock() {
        menu.addItem(1, "Pizza");
        menu.addItem(2, "Burger");
        menu.rateItem(1, 5);
        menu.rateItem(2, 4);

        menu.makeDealOfTheDayItem(1);
        menu.outOfStockItem(1);

        MenuItem recommended = menu.getRecommendedItem();
        assertEquals(2, recommended.itemId, "Should fall back to next best item if deal of the day is out of stock");
    }

    @Test
    void testOutOfStockItemNotRecommended() {
        menu.addItem(1, "Pizza");
        menu.rateItem(1, 5);

        menu.outOfStockItem(1);

        assertNull(menu.getRecommendedItem(), "Out of stock item should not be recommended");
    }

    @Test
    void testRestockItemMakesItAvailableAgain() {
        menu.addItem(1, "Pizza");
        menu.outOfStockItem(1);
        menu.restockItem(1);

        assertNotNull(menu.getRecommendedItem(), "Restocked item should be recommended again");
    }

    @Test
    void testMultipleRatingsAffectAverageCorrectly() {
        menu.addItem(1, "Pizza");
        menu.rateItem(1, 5);
        menu.rateItem(1, 3);
        menu.rateItem(1, 4);

        MenuItem item = menu.getRecommendedItem();
        double avg = item.sumOfRating / (item.numberOfPeopleWhoRatedIt * 1.0);
        assertEquals(4.0, avg, 0.001, "Average rating should be correctly computed");
    }

    @Test
    void testTieBreakingByNumberOfRatings() {
        menu.addItem(1, "Pizza");
        menu.addItem(2, "Burger");

        menu.rateItem(1, 4);
        menu.rateItem(2, 4);
        menu.rateItem(2, 4); // Burger has more ratings

        MenuItem recommended = menu.getRecommendedItem();
        assertEquals(2, recommended.itemId, "Should prefer item with more ratings if average rating is equal");
    }

    @Test
    void testTieBreakingByItemId() {
        menu.addItem(1, "Pizza");
        menu.addItem(2, "Burger");

        menu.rateItem(1, 4);
        menu.rateItem(2, 4);

        MenuItem recommended = menu.getRecommendedItem();
        assertEquals(2, recommended.itemId, "Should prefer item with higher ID if all else equal");
    }

    @Test
    void testGetRecommendedItemReturnsNullWhenEmpty() {
        assertNull(menu.getRecommendedItem(), "Should return null when no items exist");
    }

    @Test
    void testRateItemOnNonExistingItemDoesNothing() {
        menu.rateItem(99, 5); // not added
        assertNull(menu.getRecommendedItem(), "Rating non-existing item should not throw or affect state");
    }

    @Test
    void testMakeDealOfTheDayOnlyWorksForStockedItem() {
        menu.addItem(1, "Pizza");
        menu.outOfStockItem(1);

        menu.makeDealOfTheDayItem(1);
        assertNull(menu.getRecommendedItem(), "Out of stock item cannot be made deal of the day");
    }
}
