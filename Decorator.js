// Definir la clase Car para renta
class Car{
    constructor(model, rentalPrice){
        this.model = model;
        this.rentalPrice = rentalPrice;
        this.isAvailable = true;
    }

    rent(){
        if (this.isAvailable){
            this.isAvailable = false;
            console.log("Rented "+this.model+" for $"+this.rentalPrice);
        } else{
            console.log("Car "+this.model+" is not available for rent");
        }
    }

    returnCar(){
        this.isAvailable = true;
        console.log("Returned "+this.model);
    }
}

// Definir decorators para funciones adicionales
class GPSNavigation{
    constructor(car){
        this.car    = car
        this.price  = 10; // Precio adicional por GPS
    }

    rent(){
        this.car.rent();
        console.log("Added GPS Navigation for an extra $"+this.price);
    }
}

class ChildSeat{
    constructor(car){
        this.car = car;
        this.price = 5; // Precio adicional por asionto de be3bé
    }

    rent(){
        this.car.rent();
        console.log("Added Child Seat for an extra $"+this.price);
    }
}

class Insurance{
    constructor(car){
        this.car = car;
        this.price = 25; // Cobro adicional por seguro
    }

    rent(){
        this.car.rent();
        console.log("Added Insurance for an extra $"+this.price);
    }
}

// Ejemplo de uso ---------------------
let car1 = new Car("Toyota Camry", 50);
let car2 = new Car("Honda Civic", 40);

car1 = new GPSNavigation(car1); // Adding GPS navigation to car1
car2 = new ChildSeat(car2); // Adding child seat to car2
car2 = new Insurance(car2); // Adding insurance to car2

                // Output:
car1.rent();    // Rented Toyota Camry for $50 + GPS Navigation for an extra $10
car2.rent();    // Rented Honda Civic for $40 + Child Seat for an extra $5 + Insurance for an extra $20
car1.returnCar(); // Returned Toyota Camry
car2.rent();    // Rented Honda Civic for $40 + Child Seat for an extra $5 + Insurance for an extra $20
