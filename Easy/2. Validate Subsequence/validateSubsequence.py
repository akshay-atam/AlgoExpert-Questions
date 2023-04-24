# O(n) time | O(1) space
def solution1(array, sequence):
    index = 0
    for i in array:
        if index == len(sequence):
            break
        if sequence[index] == i:
            index += 1
    return index == len(sequence)

# O(n) time | O(1) space
def solution2(array, sequence):
    array_index = 0
    sequence_index = 0
    while array_index < len(array) and sequence_index < len(sequence):
        if array[array_index] == sequence[sequence_index]:
            sequence_index += 1
        array_index += 1
    return sequence_index == len(sequence)

if __name__ == "__main__":
    array = [5, 1, 22, 25, 6, -1, 8, 10]
    sequence = [1, 6, -1, 10]

    print(f"Solution 1: {solution1(array, sequence)}")
    print(f"Solution 2: {solution2(array, sequence)}")
