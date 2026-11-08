class Solution {

    public int findKthLargest(int[] nums, int k) {
        var heap = new PriorityQueue<Integer>((a, b) -> Integer.compare(b, a));

        heap.addAll(Arrays.stream(nums).boxed().toList());


        for (int i = 1; i < k; i++) {
            heap.poll();
        }

        return  heap.peek();
    }
}
