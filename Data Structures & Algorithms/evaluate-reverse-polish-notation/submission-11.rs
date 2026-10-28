impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack: Vec<i32> = Vec::new();
        
       for t in tokens {
            let res = match t.as_str() {
                "+" => stack.pop().unwrap() + stack.pop().unwrap(),
                "-" => {
                    let a = stack.pop().unwrap();
                    let b = stack.pop().unwrap();
                    b - a
                },
                "/" => {
                    let a = stack.pop().unwrap();
                    let b = stack.pop().unwrap();
                    b / a
                },
                "*" => stack.pop().unwrap() * stack.pop().unwrap(),
                _ => {
                    let i: i32 = t.parse().unwrap();
                    i
                }
            };
            stack.push(res);
       }
        return stack.pop().unwrap();
    }
}
