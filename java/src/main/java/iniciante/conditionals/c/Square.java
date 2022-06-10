package iniciante.conditionals.c;

public class Square implements Shape {
    private final int measure;
    private final int multiplier;

    public Square(int measure, int multiplier) {
        this.measure = measure;
        this.multiplier = multiplier;
    }

    @Override
    public int getMeasure() {
        return measure;
    }

    @Override
    public int getMultiplier() {
        return multiplier;
    }
}
