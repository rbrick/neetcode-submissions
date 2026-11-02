class KthLargest {

    private final PriorityQueue<Integer> heap;
    private final int k;

    public KthLargest(int k, int[] nums) {
        this.heap = new PriorityQueue<>();
        this.heap.addAll(IntStream.of(nums).boxed().toList());
        this.k = k;
    }

    public int add(int val) {
        this.heap.add(val);

        while (this.heap.size() > this.k) {
            this.heap.poll();
        }

        return this.heap.peek();
    }
}
