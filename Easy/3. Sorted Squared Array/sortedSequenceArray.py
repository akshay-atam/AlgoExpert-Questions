# O(nlogn) time | O(n) space
def solution1(array):
    sortedArray = []
    for i in range(len(array)):
        sortedArray.append(array[i]**2)
    sortedArray.sort()
    return sortedArray

# O(n) time | O(n) space
def solution2(array):
    sortedArray = [0 for _ in array]
    smaller = 0
    larger = len(array) - 1

    for i in reversed(range(len(array))):
        smallerValue = array[smaller]
        largerValue = array[larger]
        
        if abs(smallerValue) > abs(largerValue):
            sortedArray[i] = smallerValue ** 2
            smaller += 1
        else:
            sortedArray[i] = largerValue ** 2
            larger -= 1

    return sortedArray

if __name__ == "__main__":
    array = [1, 2, 3, 5, 6, 8, 9]
    print(f"Solution 1: {solution1(array)}")
    print(f"Solution 2: {solution2(array)}")
