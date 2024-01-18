//🌟 A scope is the range within the program for which the item is valid.

// Fix the error below with least amount of modification
fn main() {
    let x: i32 = 10;
    let y: i32 = 5;
    {
        println!("The value of x is {} and value of y is {}", x, y);
    }
    println!("The value of x is {} and value of y is {}", x, y);
}

// Scope significa el ambito al cual pertenece la variable en el problema teniamos a y que estaba dentro de las llavas {} y por lo tanto no se exponia a la funcion principal, poniendo y en el ambito general no tenemos problemas
