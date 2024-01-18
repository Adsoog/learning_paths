// 🌟🌟 We can use a pattern with let to destructure a tuple to separate variables.

fn main() {
    let (mut x, y) = (1, 2);
    x += 2;

    assert_eq!(x, 3);
    assert_eq!(y, 2);

    println!("Success!")
}

// Aqui recordar que solo usando mut es que una variable puede cambiar su valor si hacer sadowing eso seria otra cosa
