class Solution {
    public boolean isValid(String s) {
        final StringBuilder builder = new StringBuilder();

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(' || c == '[' || c == '{') {
                builder.append(c);
            } else if (!builder.isEmpty()) {
                char peek = builder.charAt(builder.length() - 1);

                if (c == ')' && peek != '(') {
                    return false;
                } else if (c == '}' && peek != '{') {
                    return false;
                } else if (c == ']' && peek != '[') {
                    return false;
                } else {
                    builder.setLength(builder.length() - 1);
                }
            } else {
                builder.append(c);
            }
        }

        return
                builder.isEmpty();
    }

}