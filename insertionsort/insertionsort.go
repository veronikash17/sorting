package insertionsort

// MoveLeft bewegt das Element an Stelle i so lange nach links,
// bis es an der richtigen Stelle im bereits sortierten Teil der Liste ist.
func MoveLeft(arr []int, i int) {

	for j := i; j > 0 && arr[j] < arr[j-1]; j-- { 
	 arr[j], arr[j-1] = arr[j-1], arr[j]
	
    }
}

// InsertionSort sortiert die übergebene Liste mittels des Insertion-Sort-Algorithmus.
func InsertionSort(arr []int) {
	// TODO
	for i := 1; i < len(arr); i++ {
		MoveLeft(arr, i)
	}
}
