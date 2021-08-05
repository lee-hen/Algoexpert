package degrees_of_separation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase2(t *testing.T) {
	friendsLists := FriendsGraph{
		"Aaron":  []string{"Lexi"},
		"Aditya":  []string{"Kelly", "Sam"},
		"Akshay":  []string{"Amanda", "Antoine", "Steve"},
		"Alex":  []string{"Mei"},
		"Amanda":  []string{"Eric", "Sam", "Akshay"},
		"Antoine":  []string{"Akshay", "Changpeng", "Justin"},
		"Ayushi":  []string{"Hannah"},
		"Beyonce":  []string{"Tony", "Lexi", "Steve"},
		"Brandon":  []string{"Clement", "Sandeep"},
		"Changpeng":  []string{"Antoine"},
		"Charlie":  []string{"Simon", "Penny", "Ryan"},
		"Clement":  []string{"Brandon", "Hannah", "Jean"},
		"Eric":  []string{"Amanda", "Pedro", "Sandeep"},
		"Hannah":  []string{"Ayushi", "Clement"},
		"James":  []string{},
		"Jean":  []string{"Clement", "Molly"},
		"Jon":  []string{"Kunal", "Molly", "Penny"},
		"Justin":  []string{"Antoine", "Molly", "Xi"},
		"Kelly":  []string{"Saurabh", "Aditya"},
		"Kunal":  []string{"Jon", "Lee", "Nirali", "Sam"},
		"Lee":  []string{"Kunal", "Shakira"},
		"Lexi":  []string{"Aaron", "Beyonce"},
		"Mei":  []string{"Alex"},
		"Molly":  []string{"Jean", "Jon", "Justin", "Penny"},
		"Muhammad":  []string{},
		"Nirali":  []string{"Kunal", "Pedro"},
		"Pedro":  []string{"Eric", "Nirali", "Sam"},
		"Penny":  []string{"Charlie", "Jon", "Molly"},
		"Ryan":  []string{"Charlie"},
		"Sam":  []string{"Aditya", "Kunal", "Simon", "Amanda", "Pedro"},
		"Sandeep":  []string{"Eric", "Brandon"},
		"Saurabh":  []string{"Steve", "Kelly"},
		"Shakira":  []string{"Lee"},
		"Simon":  []string{"Charlie", "Sam"},
		"Steve":  []string{"Akshay", "Beyonce", "Saurabh"},
		"Tony":  []string{"Beyonce"},
		"Xi":  []string{"Justin"},
	}
	personOne := "Pedro"
	personTwo := "Aditya"

	output := DegreesOfSeparation(friendsLists, personOne, personTwo)
	expected := "Pedro"
	require.Equal(t, expected, output)
}

func TestCase1(t *testing.T) {
	friendsLists := FriendsGraph{
		"Aaron":     []string{"Paul"},
		"Akshay":    []string{},
		"Alex":      []string{"Antoine", "Clement", "Ryan", "Simon"},
		"Antoine":   []string{"Alex", "Clement", "Simon"},
		"Ayushi":    []string{"Lee"},
		"Changpeng": []string{"Kelly", "Sandeep"},
		"Clement":   []string{"Alex", "Antoine", "Sandeep", "Simon"},
		"Hannah":    []string{"Lexi", "Paul", "Sandeep"},
		"James":     []string{"Paul"},
		"Kelly":     []string{"Changpeng", "Molly"},
		"Lee":       []string{"Ayushi", "Molly"},
		"Lexi":      []string{"Hannah"},
		"Molly":     []string{"Kelly", "Lee"},
		"Paul":      []string{"Aaron", "James", "Hannah"},
		"Ryan":      []string{"Alex"},
		"Sandeep":   []string{"Changpeng", "Clement", "Hannah"},
		"Simon":     []string{"Alex", "Antoine", "Clement"},
	}
	personOne := "Antoine"
	personTwo := "Clement"

	output := degreesOfSeparation(friendsLists, personOne, personTwo)
	expected := "Clement"
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	friendsLists := FriendsGraph{
		"Aaron": []string{"Charlie", "Lexi"},
		"Aditya": []string{"Saurabh", "Ryan"},
		"Akshay": []string{"Amanda"},
		"Alex": []string{"Eric", "Kunal", "Lexi", "Charlie", "Molly"},
		"Amanda": []string{"Akshay", "Ayushi", "Steve"},
		"Antoine": []string{},
		"Ayushi": []string{"Amanda", "Mei"},
		"Beyonce": []string{"Jean", "Shakira", "Tony", "Kunal", "Muhammad", "Xi"},
		"Brandon": []string{"Changpeng"},
		"Changpeng": []string{"Pedro", "Brandon", "James"},
		"Charlie": []string{"Aaron", "Alex", "Saurabh"},
		"Clement": []string{},
		"Eric": []string{"Alex", "Sam"},
		"Hannah": []string{"Kelly", "Kunal", "Lee"},
		"James": []string{"Changpeng", "Justin", "Kelly", "Tony"},
		"Jean": []string{"Beyonce"},
		"Jon": []string{"Nirali"},
		"Justin": []string{"James", "Kelly"},
		"Kelly": []string{"James", "Justin", "Lee", "Hannah"},
		"Kunal": []string{"Beyonce", "Lee", "Simon", "Alex", "Hannah", "Ryan"},
		"Lee": []string{"Hannah", "Kelly", "Kunal"},
		"Lexi": []string{"Aaron", "Alex"},
		"Mei": []string{"Ayushi"},
		"Molly": []string{"Alex", "Tony"},
		"Muhammad": []string{"Beyonce", "Sandeep", "Steve"},
		"Nirali": []string{"Jon"},
		"Pedro": []string{"Simon", "Changpeng", "Sam"},
		"Penny": []string{},
		"Ryan": []string{"Aditya", "Kunal", "Tony", "Xi"},
		"Sam": []string{"Eric", "Pedro", "Steve"},
		"Sandeep": []string{"Muhammad", "Shakira"},
		"Saurabh": []string{"Charlie", "Aditya"},
		"Shakira": []string{"Sandeep", "Steve", "Tony", "Beyonce"},
		"Simon": []string{"Kunal", "Pedro"},
		"Steve": []string{"Muhammad", "Sam", "Xi", "Amanda", "Shakira"},
		"Tony": []string{"James", "Molly", "Ryan", "Beyonce", "Shakira"},
		"Xi": []string{"Beyonce", "Ryan", "Steve"},
	}
	personOne := "Ayushi"
	personTwo := "Aditya"

	output := DegreesOfSeparation(friendsLists, personOne, personTwo)
	expected := "Aditya"
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	friendsLists := FriendsGraph{
		"Aaron": []string{"Akshay", "Alex", "Ayushi", "Clement"},
		"Aditya": []string{"Lee", "Simon", "Penny"},
		"Akshay": []string{"Lexi", "Muhammad", "Aaron", "Changpeng", "James", "Penny"},
		"Alex": []string{"Aaron", "Sam", "Lexi"},
		"Amanda": []string{"Ayushi", "Lexi"},
		"Antoine": []string{},
		"Ayushi": []string{"Aaron", "Amanda", "Charlie", "Saurabh"},
		"Beyonce": []string{"Brandon", "Kunal"},
		"Brandon": []string{"Muhammad", "Beyonce"},
		"Changpeng": []string{"Akshay"},
		"Charlie": []string{"Ayushi", "Hannah"},
		"Clement": []string{"Aaron", "Kunal"},
		"Eric": []string{"Shakira"},
		"Hannah": []string{"Charlie", "Saurabh", "Shakira", "Steve", "James", "Penny"},
		"James": []string{"Akshay", "Hannah"},
		"Jean": []string{"Kunal"},
		"Jon": []string{"Sam"},
		"Justin": []string{},
		"Kelly": []string{},
		"Kunal": []string{"Jean", "Steve", "Beyonce", "Clement"},
		"Lee": []string{"Aditya", "Xi"},
		"Lexi": []string{"Alex", "Penny", "Akshay", "Amanda", "Pedro", "Sandeep"},
		"Mei": []string{"Simon"},
		"Molly": []string{},
		"Muhammad": []string{"Akshay", "Brandon", "Ryan"},
		"Nirali": []string{},
		"Pedro": []string{"Lexi"},
		"Penny": []string{"Aditya", "Akshay", "Hannah", "Lexi"},
		"Ryan": []string{"Muhammad"},
		"Sam": []string{"Jon", "Alex"},
		"Sandeep": []string{"Lexi", "Tony"},
		"Saurabh": []string{"Ayushi", "Hannah"},
		"Shakira": []string{"Eric", "Hannah", "Xi"},
		"Simon": []string{"Aditya", "Mei"},
		"Steve": []string{"Hannah", "Kunal"},
		"Tony": []string{"Sandeep"},
		"Xi": []string{"Lee", "Shakira"},
	}
	personOne := "Muhammad"
	personTwo := "Ayushi"

	output := DegreesOfSeparation(friendsLists, personOne, personTwo)
	expected := ""
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	friendsLists := FriendsGraph{
		"A":     []string{ "C", "B"},
		"B":    []string{"A", "C"},
		"C":      []string{"E"},
	}
	personOne := "A"
	personTwo := "B"

	output := degreesOfSeparation(friendsLists, personOne, personTwo)
	expected := ""
	require.Equal(t, expected, output)
}

func TestCase6(t *testing.T) {
	friendsLists := FriendsGraph{
		"Aaron":  []string{"Kelly", "Ryan", "Charlie", "Lexi", "Muhammad"},
		"Aditya":  []string{"Lexi"},
		"Akshay":  []string{"Brandon", "Clement", "Charlie", "Hannah", "Shakira"},
		"Alex":  []string{"Sam", "Pedro"},
		"Amanda":  []string{"James", "Justin", "Simon", "Changpeng", "Sandeep"},
		"Antoine":  []string{"Changpeng"},
		"Ayushi":  []string{"Jon", "Kelly", "Pedro", "Kunal", "Simon"},
		"Beyonce":  []string{"Shakira"},
		"Brandon":  []string{"Jean", "Nirali", "Akshay", "Hannah"},
		"Changpeng":  []string{"Amanda", "Antoine"},
		"Charlie":  []string{"Aaron", "Akshay", "Penny", "Lee", "Saurabh"},
		"Clement":  []string{"Akshay", "Mei", "Shakira"},
		"Eric":  []string{"Hannah", "Jon", "Simon", "Xi"},
		"Hannah":  []string{"Akshay", "Brandon", "Ryan", "Eric", "Shakira"},
		"James":  []string{"Amanda"},
		"Jean":  []string{"Brandon"},
		"Jon":  []string{"Ayushi", "Eric", "Lexi"},
		"Justin":  []string{"Amanda", "Simon"},
		"Kelly":  []string{"Aaron", "Ayushi"},
		"Kunal":  []string{"Ayushi"},
		"Lee":  []string{"Charlie", "Simon"},
		"Lexi":  []string{"Aaron", "Aditya", "Jon", "Xi", "Pedro"},
		"Mei":  []string{"Clement", "Muhammad"},
		"Molly":  []string{"Ryan"},
		"Muhammad":  []string{"Aaron", "Nirali", "Mei"},
		"Nirali":  []string{"Brandon", "Muhammad", "Saurabh"},
		"Pedro":  []string{"Alex", "Lexi", "Ayushi"},
		"Penny":  []string{"Charlie"},
		"Ryan":  []string{"Aaron", "Hannah", "Molly"},
		"Sam":  []string{"Alex"},
		"Sandeep":  []string{"Amanda"},
		"Saurabh":  []string{"Charlie", "Nirali"},
		"Shakira":  []string{"Akshay", "Clement", "Hannah", "Beyonce"},
		"Simon":  []string{"Ayushi", "Eric", "Justin", "Amanda", "Lee"},
		"Steve":  []string{},
		"Tony":  []string{},
		"Xi":  []string{"Eric", "Lexi"},
	}
	personOne := "Steve"
	personTwo := "Saurabh"

	output := degreesOfSeparation(friendsLists, personOne, personTwo)
	expected := "Saurabh"
	require.Equal(t, expected, output)
}
