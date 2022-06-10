package iniciante.conditionals.b;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class TollTest
{
    private final Toll toll = new Toll();

    @Test
    public void citizenBike()
    {
        double expectedPayment = 0.00;
        Vehicle citizenBike = new Vehicle(2, false);

        double actualPayment = toll.calculateTollPayment(citizenBike);
        assertEquals(expectedPayment, actualPayment);
    }

    @Test
    public void policeBike()
    {
        double expectedPayment = 0.00;
        Vehicle policeBike = new Vehicle(2, true);

        double actualPayment = toll.calculateTollPayment(policeBike);
        assertEquals(expectedPayment, actualPayment);
    }

    @Test
    public void familyCar()
    {
        double expectedPayment = 5.20;
        Vehicle familyCar = new Vehicle(4, false);

        double actualPayment = toll.calculateTollPayment(familyCar);
        assertEquals(expectedPayment, actualPayment);
    }

    @Test
    public void policeCar()
    {
        double expectedPayment = 0.00;
        Vehicle policeCar = new Vehicle(4, true);

        double actualPayment = toll.calculateTollPayment(policeCar);
        assertEquals(expectedPayment, actualPayment);
    }

    @Test
    public void nonOfficial6WheelsTruck()
    {
        double expectedPayment = 7.80;
        Vehicle nonOfficial6WheelTruck = new Vehicle(6, false);

        double actualPayment = toll.calculateTollPayment(nonOfficial6WheelTruck);
        assertEquals(expectedPayment, actualPayment, 0.001);
    }

    @Test
    public void official6WheelsTruck()
    {
        double expectedPayment = 0.00;
        Vehicle nonOfficial6WheelTruck = new Vehicle(6, true);

        double actualPayment = toll.calculateTollPayment(nonOfficial6WheelTruck);
        assertEquals(expectedPayment, actualPayment);
    }

    @Test
    public void nonOfficial8WheelsTruck()
    {
        double expectedPayment = 10.0;
        Vehicle nonOfficial6WheelTruck = new Vehicle(8, false);

        double actualPayment = toll.calculateTollPayment(nonOfficial6WheelTruck);
        assertEquals(expectedPayment, actualPayment);
    }

    @Test
    public void official8WheelsTruck()
    {
        double expectedPayment = 0.0;
        Vehicle nonOfficial6WheelTruck = new Vehicle(8, true);

        double actualPayment = toll.calculateTollPayment(nonOfficial6WheelTruck);
        assertEquals(expectedPayment, actualPayment);
    }

}