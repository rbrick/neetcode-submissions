class Solution {
    public int evalRPN(String[] tokens) {
        final Stack<Integer> stack = new Stack<>();
        for (String token : tokens) {
            switch (token) {
                case "+": stack.push(stack.pop() + stack.pop());break;
                case "-": {
                    int a = stack.pop(); int b = stack.pop();
                    stack.push(b - a);
                    break;
                }
                case "*": stack.push(stack.pop() * stack.pop());break;
                case "/": {
                    int a = stack.pop(); int b = stack.pop();
                    stack.push(b / a);
                    break;
                }
                default:
                    stack.push(Integer.parseInt(token, 10));
            }
        }
        return stack.pop();
    }
}
