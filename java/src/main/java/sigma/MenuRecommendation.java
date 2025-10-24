package sigma;

import java.io.*;
import java.util.*;

/**
 * Example Input:
 * 8
 * getRecommendedItem
 * addItem 1 Item1
 * rateItem 1 5
 * getRecommendedItem
 * outOfStockItem 1
 * rateItem 1 4
 * rateItem 1 4
 * getRecommendedItem
 *
 *
 *
 * 18
 * getRecommendedItem
 * addItem 1 Item1
 * addItem 2 Item2
 * addItem 3 Item3
 * makeDealOfTheDayItem 1
 * rateItem 1 5
 * rateItem 1 4
 * rateItem 2 4
 * rateItem 2 5
 * rateItem 2 5
 * rateItem 3 4
 * rateItem 3 5
 * rateItem 3 5
 * getRecommendedItem
 * outOfStockItem 1
 * rateItem 1 4
 * rateItem 1 4
 * getRecommendedItem
 */
class MenuRecommendation implements IMenuRecommendation {

    private Map<Integer, MenuItem> menu = new HashMap<>();
    private Set<Integer> stock = new HashSet<>();
    private MenuItem dealOfTheDay = null;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        PrintWriter out = new PrintWriter(System.out);
        int totalNumberOfRequests = Integer.parseInt(br.readLine().trim());

        String arr[];

        IMenuRecommendation menuRecommendation = new MenuRecommendation();

        for (int requestNumber = 1; requestNumber <= totalNumberOfRequests; requestNumber++) {
            arr = br.readLine().trim().split(" ");
            switch (arr[0]) {
                case "addItem":
                    int id = Integer.parseInt(arr[1]);
                    String displayName = arr[2];
                    menuRecommendation.addItem(id, displayName);
                    break;
                case "getRecommendedItem":
                    MenuItem menuItem = menuRecommendation.getRecommendedItem();
                    if (menuItem == null) {
                        out.println("N/A");
                        break;
                    }
                    double averageRating = (menuItem.numberOfPeopleWhoRatedIt == 0) ?
                            0 :
                            (menuItem.sumOfRating / (menuItem.numberOfPeopleWhoRatedIt * 1.0d));
                    out.println(menuItem.itemId + " " + menuItem.displayName + " Rating: " + averageRating);
                    break;
                case "outOfStockItem":
                    int itemId = Integer.parseInt(arr[1]);
                    menuRecommendation.outOfStockItem(itemId);
                    break;
                case "restockItem":
                    itemId = Integer.parseInt(arr[1]);
                    menuRecommendation.restockItem(itemId);
                    break;
                case "makeDealOfTheDayItem":
                    itemId = Integer.parseInt(arr[1]);
                    menuRecommendation.makeDealOfTheDayItem(itemId);
                    break;
                case "rateItem":
                    itemId = Integer.parseInt(arr[1]);
                    int rating = Integer.parseInt(arr[2]);
                    menuRecommendation.rateItem(itemId, rating);
                    break;
            }
        }

        out.flush();
        out.close();
    }

    @Override
    public void addItem(int id, String displayName) {
        var item = new MenuItem(id, displayName);

        menu.put(id, item);
        restockItem(id);
    }

    @Override
    public MenuItem getRecommendedItem() {
        MenuItem recommendation = null;

        if (dealOfTheDay != null) {
            if (stock.contains(dealOfTheDay.itemId)) {
                recommendation = dealOfTheDay;
            } else {
                dealOfTheDay = null;
            }
        }

        if (recommendation == null) {

            Optional<MenuItem> topRated = menu.values()
                    .stream()
                    .filter(item -> stock.contains(item.itemId))
                    .sorted((m1, m2) -> {
                        int compare = (int) ((m2.sumOfRating / m2.numberOfPeopleWhoRatedIt) - (m1.sumOfRating / m1.numberOfPeopleWhoRatedIt));

                        if (compare == 0) {
                            compare = m2.numberOfPeopleWhoRatedIt - m1.numberOfPeopleWhoRatedIt;
                        }

                        if (compare == 0) {
                            compare = m2.itemId - m1.itemId;
                        }

                        return compare;
                    }).findFirst();

            recommendation = topRated.isPresent() ? topRated.get() : null;
        }

        return recommendation != null && stock.contains(recommendation.itemId) ? recommendation : null;
    }

    @Override
    public void outOfStockItem(int itemId) {
        if (stock.contains(itemId))
            stock.remove(itemId);
    }

    @Override
    public void makeDealOfTheDayItem(int itemId) {
        if (stock.contains(itemId)) {
            this.dealOfTheDay = menu.get(itemId);
        }
    }

    @Override
    public void rateItem(int itemId, int rating) {
        menu.computeIfPresent(itemId, (id, item) -> {
            item.sumOfRating += rating;
            item.numberOfPeopleWhoRatedIt += 1;
            return item;
        });
    }

    @Override
    public void restockItem(int itemId) {
        stock.add(itemId);
    }

}
