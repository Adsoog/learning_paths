fn main() {
    let (x, y);
    (x, ..) = (3, 4);
    [.., y] = [1, 2];
    // fill the blank to make code work
    assert_eq!([x, y], [3, 2]);
    println!("Success!!!");
}
