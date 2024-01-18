// 🌟 A variable can be used only if it has been initialized.

// Fix the error below with least amount of modification to the code
fn main() {
    let x: i32 = 5; // Uninitialized but used, ERROR !
    let _y: i32; // Uninitialized but also unused, only a Warning !

    assert_eq!(x, 5);
    println!("Success!");
}

// Para resolver esto primero vemos que la variable x no tenia ningun valor n iestaba siendo inicializada le asignamos
// el valor de 5 y para evitar errores tambien pusimos _y para que si ignore la variable no usada
