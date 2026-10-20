package homework

// PairInput соответствует объекту nums/target из условия контеста.
type PairInput struct {
	Nums   []int
	Target int
}

// longestPrefix возвращает самый длинный общий префикс строк.
func longestPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}
	firstWord := []rune(words[0])
	for i := range firstWord {
		c := firstWord[i]
		for j := 1; j < len(words); j++ {
			r := []rune(words[j])
			if i >= len(r) || r[i] != c {
				return string(firstWord[:i])
			}
		}
	}
	return words[0]
	// TODO: реализовать поиск префикса по кодовым точкам Unicode.
}

// pairIndices возвращает индексы двух разных элементов по возрастанию.
func pairIndices(input PairInput) [2]int {
	seen := make(map[int]int, len(input.Nums))
	for i, v := range input.Nums {
		if j, ok := seen[input.Target-v]; ok {
			return [2]int{j, i}
		}
		seen[v] = i
	}
	// TODO: найти единственную пару с суммой input.Target.
	return [2]int{}
}

// hasDuplicate сообщает, встречается ли какое-либо значение повторно.
func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = true
	}
	// TODO: проверить наличие повторов.
	return false
}

// compactSorted создаёт независимый слайс уникальных значений.
func compactSorted(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	seen := make(map[int]bool, len(nums))
	res := make([]int, 0)
	for i, v := range nums {
		if _, ok := seen[v]; !ok {
			res = append(res, nums[i])
		}
		seen[v] = true
	}
	// TODO: сохранить порядок и не изменять исходный слайс.
	return res
}

// balancedBrackets проверяет пары и порядок закрытия скобок.
func balancedBrackets(text string) bool {
	// TODO: обработать строку из символов ()[]{}.
	storage := make([]rune, 0, len(text))
	for _, v := range text {
		switch v {
		case '{':
			storage = append(storage, '}')
		case '(':
			storage = append(storage, ')')
		case '[':
			storage = append(storage, ']')
		default:
			if len(storage) == 0 || storage[len(storage)-1] != v {
				return false
			}
			storage = storage[:len(storage)-1]
		}
	}
	return len(storage) == 0
}
