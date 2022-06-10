package iniciante.conditionals.b;

public class Vehicle {

    // número de rodas
    private final int numberOfWheels;

    // se é ou não oficial
    private final boolean isOfficial;

    public Vehicle(int numberOfWheels, boolean isOfficial) {
        this.numberOfWheels = numberOfWheels;
        this.isOfficial = isOfficial;
    }

    public int getNumberOfWheels() {
        return numberOfWheels;
    }

    public boolean isOfficial() {
        return isOfficial;
    }
}