def nonConstructibleChange(coins):
	currentMaxChange = 0
	coins.sort()
	for i in range(len(coins)):
		if currentMaxChange + 1 < coins[i]:
			return currentMaxChange + 1
		currentMaxChange += coins[i]
	return currentMaxChange + 1

if __name__ == "__main__":
	coins = [5, 7, 1, 1, 2, 3, 22]
	print(f"Minimum amount of change that cannot be created: {nonConstructibleChange(coins)}")

	if 0:
		print("hi")