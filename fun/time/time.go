package time

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-module/carbon/v2"
)

func Main() {
	fmt.Println("Welcome")

	fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15

	carbon.SetDefault(carbon.Default{
		Layout: carbon.DateTimeLayout,
	})

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`

		Field1 carbon.Carbon `json:"field1"`
		Field2 carbon.Carbon `json:"field2"`
		Field3 carbon.Carbon `json:"field3"`
		Field4 carbon.Carbon `json:"field4"`

		Field5 carbon.Carbon `json:"field5"`
		Field6 carbon.Carbon `json:"field6"`
		Field7 carbon.Carbon `json:"field7"`
		Field8 carbon.Carbon `json:"field8"`
	}

	now := carbon.Parse("2020-08-05 13:14:15", carbon.PRC)
	person := Person{
		Name: "gouguoyin",
		Age:  18,

		Field1: now,
		Field2: now,
		Field3: now,
		Field4: now,
		Field5: now,
		Field6: now,
		Field7: now,
		Field8: now,
	}

	data, marshalErr := json.Marshal(person)
	if marshalErr != nil {
		// Error handle...
		log.Fatal(marshalErr)
	}
	fmt.Printf("%s", data)

	var personOne Person
	unmarshalErr := json.Unmarshal(data, &personOne)
	if unmarshalErr != nil {
		// Error handle...
		log.Fatal(unmarshalErr)
	}

	fmt.Printf("%s !!!!", person.Field1) // 2020-08-05 13:14:15
	fmt.Printf("%s", person.Field2)      // 2020-08-05 13:14:15
	fmt.Printf("%s", person.Field3)      // 2020-08-05 13:14:15
	fmt.Printf("%s", person.Field4)      // 2020-08-05 13:14:15

	fmt.Printf("%s", person.Field5) // 2020-08-05 13:14:15
	fmt.Printf("%s", person.Field6) // 2020-08-05 13:14:15
	fmt.Printf("%s", person.Field7) // 2020-08-05 13:14:15
	fmt.Printf("%s", person.Field8) // 2020-08-05 13:14:15

}
