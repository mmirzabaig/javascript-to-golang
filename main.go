package main

import (
	"fmt"
	"math"
	s "strings"

	sl "./help"
)

func main() {
	// #1 isOldEnoughToDrink
	fmt.Println(isOldEnoughToDrink(21))

	// #2 getProperty
	objq2 := make(map[string]string)
	objq2["yes"] = "true"
	objq2["no"] = "false"
	fmt.Println(getProperty(objq2, "no"))

	// #3 addProperty
	objq3 := make(map[string]bool)
	fmt.Println(addProperty(objq3, "rain"))

	// #4 removeProperty
	objq4 := make(map[string]int)
	objq4["odd"] = 3
	objq4["even"] = 4
	fmt.Println(removeProperty(objq4, "even"))

	// #5 getLengthOfTwoWords
	fmt.Println(getLengthOfTwoWords("one", "two"))

	// #6 addArrayProperty
	objq6 := make(map[string][2]string)
	arrq6 := [2]string{"hello", "world"}
	fmt.Println(addArrayProperty(objq6, "arr", arrq6))

	// #7 getLastElement
	arrq8 := []string{"12", "34", "56"}
	fmt.Println(getLastElement(arrq8))

	// #8 addToFront
	sliceq8 := []int{1, 2, 3, 4, 5}
	fmt.Println(addToFront(sliceq8, 6))

	// Testing importing functions from another file
	sl.Demo()

	// #9 computePower
	fmt.Println(computePower(2, 3))

	// #10 computePerimeterOfACircle
	fmt.Println(computePerimeterOfACircle(12))

	// #11 joinSlices
	fmt.Println([]int{1, 2, 3}, []int{3, 4, 5})

	// #12 getElementsAfter
	fmt.Println(getElementsAfter([]string{"1", "2", "3", "4", "5", "6", "7", "8"}, 3))

	// #13 getElementsUpTo
	fmt.Println(getElementsUpTo([]string{"a", "b", "c", "d", "e"}, 3))

	// #14 countCharacter
	fmt.Println(countCharacter("what is going onn", "n"))

	// #15 getAllLetters
	fmt.Println(getAllLetters("hello"))

	// #16 countWords
	fmt.Println(countWords("ask a bunch get a bunch"))

	// #17 removeFromBack
	fmt.Println(removeFromBack([]int{1, 2, 3, 4}))

	// #18 extend
	obj1q18 := make(map[string]int)
	obj1q18["a"] = 1
	obj1q18["b"] = 2
	obj2q18 := make(map[string]int)
	obj2q18["b"] = 2
	obj2q18["c"] = 3
	fmt.Println(extend(obj1q18, obj2q18))

	// #19 removeNumbersLargerThan
	objq19 := make(map[string]int)
	objq19["a"] = 91
	objq19["b"] = 19
	objq19["c"] = 9
	objq19["d"] = 29
	objq19["e"] = 3
	objq19["f"] = 4
	fmt.Println(removeNumbersLargerThan(10, objq19))

	// #20 removeStringValuesLongerThan
	objq20 := make(map[string]string)
	objq20["a"] = "Ben"
	objq20["b"] = "Dawaj"
	objq20["c"] = "Julien"
	objq20["d"] = "Alex"
	objq20["e"] = "Christine"
	objq20["f"] = "Mariana"
	objq20["g"] = "Kyle"
	fmt.Println(removeStringValuesLongerThan(4, objq20))

	// #21 countAllCharacters
	fmt.Println(countAllCharacters("banana"))
}

// #1
// Write a function called “isOldEnoughToDrink”.
// Given a number, in this case an age, “isOldEnoughToDrink” returns whether a person of this given age is old enough to legally drink in the United States.
// Notes:
// * The legal drinking age in the United States is 21.

func isOldEnoughToDrink(age int) bool {
	if age >= 21 {
		return true
	}
	return false
}

// #2
// Write a function called “getProperty”.
// Given an object and a key, “getProperty” returns the value of the property at the given key.
// Notes:
// * If there is no property at the given key, it should return undefined.

// var obj = {
//   key: 'value'
// };
// var output = getProperty(obj, 'key');
// console.log(output); // --> 'value'

func getProperty(obj map[string]string, str string) string {
	// if statements in go can include both a condition and an intialization
	// val will receive the value and ok will be either true or false
	if val, ok := obj[str]; ok {
		if ok == true {
			fmt.Println(val)
		} else {
			return "undefined"
		}
	}
	return ""
}

// #3
// Write a function called “addProperty”.
// Given an object, and a key, “addProperty” sets a new property on the given object with a value of true.

// var myObj = {};
// addProperty(myObj, 'myProperty');
// console.log(myObj.myProperty); // --> true

func addProperty(obj map[string]bool, key string) map[string]bool {
	obj[key] = true
	return obj
}

// #4
// Write a function called “removeProperty”.
// Given an object and a key, “removeProperty” removes the given key from the given object.

// var obj = {
//   name: 'Sam',
//   age: 20
// }
// removeProperty(obj, 'name');
// console.log(obj.name); // --> undefined

func removeProperty(obj map[string]int, even string) map[string]int {
	delete(obj, even)
	return obj
}

// #5
// Write a function called “getLengthOfTwoWords”.
// Given 2 words, “getLengthOfTwoWords” returns the sum of their lengths.

func getLengthOfTwoWords(str1 string, str2 string) int {
	return len(str1) + len(str2)
}

// #6
// Write a function called “addArrayProperty”.
// Given an object, a key, and an array, “addArrayProperty” sets a new property on the object at the given key, with its value set to the given array.

// var myObj = {};
// var myArray = [1, 3];
// addArrayProperty(myObj, 'myProperty', myArray);
// console.log(myObj.myProperty); // --> [1, 3]

func addArrayProperty(obj map[string][2]string, key string, array [2]string) map[string][2]string {
	obj[key] = array
	return obj
}

// #7
// Write a function called “getLastElement”.

// Given an array, “getLastElement” returns the last element of the given array.

// Notes:
// * If the given array has a length of 0, it should return ‘undefined’.

// var output = getLastElement([1, 2, 3, 4]);
// console.log(output); // --> 4

func getLastElement(arr []string) string {
	return arr[len(arr)-1]
}

// #8
// Write a function called “addToFront”.

// Given an array and an element, “addToFront” adds the given element to the front of the given array, and returns the given array.

// Notes:
// * It should be the SAME array, not a new array.
// var output = addToFront([1, 2], 3);
// console.log(output); // -> [3, 1, 2]

// input: array(slice), int 	return:	array

func addToFront(slice []int, num int) []int {
	// make antoher slice and append the original slice to it
	slice = append([]int{num}, slice...)
	return slice
}

// #9
// Write a function called “computePower”.

// Given a number and an exponent, “computePower” returns the given number, raised to the given exponent.

// var output = computePower(2, 3);
// console.log(output); // --> 8

func computePower(num float64, exp float64) float64 {
	result := math.Pow(num, exp)
	return result
}

// # 10
// Write a function called “computePerimeterOfACircle”.

// Given the radius of a circle, “computePerimeterOfACircle” returns its perimeter.

// var output = computePerimeterOfACircle(4);
// console.log(output); // --> 25.132741228718345

func computePerimeterOfACircle(radius float64) float64 {
	return (2 * math.Pi * radius)
}

// # 11
// Write a function called “joinArrays”.

// Given two arrays, “joinArrays” returns an array with the elements of “arr1” in order, followed by the elementsin “arr2”.

// var output = joinArrays([1, 2], [3, 4]);
// console.log(output); // --> [1, 2, 3, 4]

func joinSlices(slice1 []int, slice2 []int) []int {
	// append will always append the second argument passed in, to the first one
	slice := append(slice1, slice2...)
	return slice
}

// #12
// Write a function called “getElementsAfter”.

// Given an array and an index, “getElementsAfter” returns a new array with all the elements after (but not including) the given index.

func getElementsAfter(slice []string, index int) []string {
	// index = 3  		So it returns everthing after the 3 elemnt in the array
	return slice[index:]
}

// #13
// Write a function called “getElementsUpTo”.

// Given an array and a index, “getElementsUpTo”, returns an array with all the elements up until, but not including, the element at the given index.

// Notes:
// * In order to do this you should be familiar with the ‘slice’ method.

// var output = getElementsUpTo(['a', 'b', 'c', 'd', 'e'], 3)
// console.log(output); // --> ['a', 'b', 'c']

func getElementsUpTo(slice []string, index int) []string {
	return slice[:index]
}

// #14
// Write a function called “countCharacter”.

// Given a string input and a character, “countCharacter” returns the number of occurences of a given character in the given string.

// var output = countCharacter('I am a hacker', 'a');
// console.log(output); // --> 3

func countCharacter(str string, char string) int {

	// SOLVE WITH OTHER METHODS AS WELL

	// for i := 0; i < len(str); i++ {
	// 	if str[i] == char {
	// 		return i
	// 	}
	// }
	// for _, value := range oldMap {
	// 	// newMap[key] = value

	// 	// value = unicode.ToLower(value)
	// 	fmt.Printf("%s", value)
	// 	// if value == char {

	// 	// }
	// 	// fmt.Printf(, i, c)
	// }

	// return 5
	return s.Count(str, char)
}

// #15
// Write a function called “getAllLetters”.

// Given a word, “getAllLetters” returns an array containing every character in the word.

// Notes:
// * If given an empty string, it should return an empty array.

// var output = getAllLetters('Radagast');
// console.log(output); // --> ['R', 'a', 'd', 'a', 'g', 'a', 's', 't']

func getAllLetters(word string) []string {
	slice := s.Split(word, "")
	return slice
}

// #16
// Write a function called “countWords”.

// Given a string, “countWords” returns an object where each key is a word in the given string, with its value being how many times that word appeared in the given string.

// Notes:
// * If given an empty string, it should return an empty object.

// var output = countWords('ask a bunch get a bunch');
// console.log(output); // --> {ask: 1, a: 2, bunch: 2, get: 1}

func countWords(str string) map[string]int {
	slice := s.Split(str, " ")
	obj := make(map[string]int)
	for i := 0; i < len(slice); i++ {
		if _, ok := obj[slice[i]]; ok {
			if ok == true {
				obj[slice[i]]++
			}
		} else {
			obj[slice[i]] = 1
		}
	}
	return obj
}

// #17
// Write a function called “removeFromBack”.

// Given an array, “removeFromBack” returns the given array with its last element removed.

// Notes:
// * You should be familiar with the method ‘pop’.

// var output = removeFromBack([1, 2, 3]);
// console.log(output); // --> [1, 2]

func removeFromBack(slice []int) []int {
	return slice[:len(slice)-1]
}

// #18
// Write a function called “extend”.

// Given two objects, “extend” adds properties from the 2nd object to the 1st object.

// Notes:
// * Add any keys that are not in the 1st object.
// * If the 1st object already has a given key, ignore it (do not overwrite the property value).
// * Do not modify the 2nd object at all.

// var obj1 = {
//   a: 1,
//   b: 2
// };
// var obj2 = {
//   b: 4,
//   c: 3
// };

// extend(obj1, obj2);

// console.log(obj1); // --> {a: 1, b: 2, c: 3}
// console.log(obj2); // --> {b: 4, c: 3}
func extend(obj1 map[string]int, obj2 map[string]int) map[string]int {
	for key, val := range obj1 {
		if _, ok := obj2[key]; !ok {
			obj2[key] = val
		}
	}
	return obj2
}

// #19
// Write a function called “removeNumbersLargerThan”.

// Given a number and an object, “removeNumbersLargerThan” removes any properties whose values are numbers greater than the given number.

// var obj = {
//   a: 8,
//   b: 2,
//   c: 'montana'
// }
// removeNumbersLargerThan(5, obj);
// console.log(obj); // --> { b: 2, c: 'montana' }

func removeNumbersLargerThan(num int, obj map[string]int) map[string]int {
	for key, value := range obj {
		if value > num {
			delete(obj, key)
		}
	}
	return obj
}

// #20
// Write a function called “removeStringValuesLongerThan”.

// Given an number and an object, “removeStringValuesLongerThan” removes any properties on the given object whose values are strings longer than the given number.

// var obj = {
//   name: 'Montana',
//   age: 20,
//   location: 'Texas'
// };
// removeStringValuesLongerThan(6, obj);
// console.log(obj); // { age: 20, location: 'Texas' }

func removeStringValuesLongerThan(num int, obj map[string]string) map[string]string {
	for key, value := range obj {
		if len(value) > num {
			delete(obj, key)
		}
	}
	return obj
}

// #21
// Write a function called “countAllCharacters”.

// Given a string, “countAllCharacters” returns an object where each key is a character in the given string. The value of each key should be how many times each character appeared in the given string.

// Notes:
// * If given an empty string, countAllCharacters should return an empty object.

// var output = countAllCharacters('banana');
// console.log(output); // --> {b: 1, a: 3, n: 2}

func countAllCharacters(str string) map[string]int {
	slice := s.Split(str, "")
	obj := make(map[string]int)
	for i := 0; i < len(slice); i++ {
		if _, ok := obj[slice[i]]; ok {
			obj[slice[i]]++
		} else {
			obj[slice[i]] = 1
		}
	}
	return obj
}
