package mergesort

// Merge kombiniert zwei sortierte Listen zu einer sortierten Liste.
func Merge(a1 []int, a2 []int) []int {
	// Erstelle ein neues Slice für das Ergebnis.
	// Verwende zwei Indizes, um durch beide Eingabelisten zu iterieren.
	// Hänge jeweils das kleinere Element an das Ergebnis an.
	result := []int{}
	i, j := 0, 0

	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			result = append(result, a1[i])
			i++
		} else {
			result = append(result, a2[j])
			j++
		}
	}

	result = append(result, a1[i:]...)
	result = append(result, a2[j:]...)

	return result
}

// MergeSort sortiert die übergebene Liste mittels des Merge-Sort-Algorithmus.
func MergeSort(arr []int) []int {
	// Basisfall: Wenn die Liste weniger als 2 Elemente hat, ist sie bereits sortiert.
	if len(arr) < 2 {
		return arr
	}

	// Teile die Liste in zwei Hälften, sortiere beide Hälften rekursiv und
	// führe die beiden sortierten Hälften mit Merge zusammen.
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}
