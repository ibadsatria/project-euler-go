package problem5

// 2520 is the smallest number that can ben divided by each of the numbers from 1 to 10 without any reminder.
// What is the smallest positive number that is evently divisible by all of the numbers from 1 to 20?

// MinFactor find minimum factor of sequence number to limit
func MinFactor(limit int) int {
	num := 2
	for {
		if divisible(num, limit) {
			return num
		}
		num = num + 2
	}
}

func divisible(num, limit int) bool {
	for i := 2; i <= limit; i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}
