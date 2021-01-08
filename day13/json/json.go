package main
//Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.



import (
    "encoding/json"
    "fmt"
    "os"
)

type response1 struct {
	//We’ll use these two structs to demonstrate encoding and decoding of custom types below.

    Page   int
    Fruits []string
}

type response2 struct {
	//Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.

    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

	//First we’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.

    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    res1D := &response1{
		//The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.

        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &response2{//You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of response2 above to see an example of such tags.
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.


	var dat map[string]interface{}
	//We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.



    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
	}
	//Here’s the actual decoding, and a check for associated errors.


    fmt.Println(dat)

    num := dat["num"].(float64)
	fmt.Println(num)
	//In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. For example here we convert the value in num to the expected float64 type.



    strs := dat["strs"].([]interface{})//Accessing nested data requires a series of conversions.
    str1 := strs[0].(string)
    fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	//We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.

    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	//In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out. We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.


    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}