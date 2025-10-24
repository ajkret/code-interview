package sigma;

public class MenuItem {
    int itemId;
    String displayName;
    long sumOfRating;
    int numberOfPeopleWhoRatedIt;

    public MenuItem(int id, String displayName) {
        this.itemId = id;
        this.displayName = displayName;
    }
}
