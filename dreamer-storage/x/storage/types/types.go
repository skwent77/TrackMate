package types

/*

Data Structure

key  : address
value: map[time.Time]Data

*/

type Data struct {
	raw_data []byte
	meta     []byte
}
