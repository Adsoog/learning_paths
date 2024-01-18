// Fix the error with the use of define_x

fn main() {
    define_x()
}

fn define_x() {
    let x = "hello";
    println!("{}, world", x);
}

// Aqui solo necesitamos obtner el valor de la funcion define y usarla dentro de main
