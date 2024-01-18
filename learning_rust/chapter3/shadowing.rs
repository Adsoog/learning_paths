//You can declare a new variable with the same name as a previous variable, here we can say the first one is shadowed by the second one.

fn main() {
    let x: i32 = 12;
    {
        let x = 5;
        assert_eq!(x, 5);
    }

    assert_eq!(x, 12);

    let x = 42;
    println!("x vale: {}", x);
}
