func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1

    // t = 7
    // i = [1,2,3,4]
    //      ^     ^
    for left < right {
        a, b := numbers[left], numbers[right]

        sum := a+b
        if sum > target {
            right--
        } else if sum < target {
            left++
        }  else {
            break
        }

    }


    return []int{left+1, right+1}
}
