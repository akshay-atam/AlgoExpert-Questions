# O(n^2) time | O(1) space
def solution1(array, targetSum):
    for i in range(len(array)-1):
        firstNum = array[i]
        for j in range(i+1, len(array)):
            secondNum = array[j]
            if firstNum + secondNum == targetSum:
                return [firstNum, secondNum]
    return []

# O(n) time | O(n) space
def solution2(array, targetSum):
    nums = {}
    for x in array:
        y = targetSum - x
        if y in nums:
            return [y, x]
        else:
            nums[x] = True
    return []

# O(nlog(n)) time | O(1) space
def solution3(array, targetSum):
    array.sort()
    l = 0
    r = len(array) - 1

    while l<r:
        target = array[l] + array[r]
        if target == targetSum:
            return [array[l], array[r]]
        elif target < targetSum:
            l += 1
        elif target > targetSum:
            r -= 1
    return []

if __name__ == "__main__":
    array = [3, 5, -4, 8, 11, 1, -1, 6]
    targetSum = 10

    print(f"Solution 1: {solution1(array, targetSum)}")
    print(f"Solution 2: {solution2(array, targetSum)}")
    print(f"Solution 3: {solution3(array, targetSum)}")