Maps in Go Language
--------------------------------
Practice project to learn maps in go language

1. Created a main.go file to write source code

2. Initialised, declared & assigned a variable colors to a map with a key of type string and value of type string (map[string]string) with initial values and printed output.

3. Initialised & declared a second variable called colours with type map[string]string with no initial values. I.e. Zero Value applied

4. Initialised and declared a third variable called colorMap with the use of "make" with no initial values. I.e empty map

5. created a key value pair within an empty declared map. colorMap

6. Utilising the delete function to pass in colorMap with key identifier to delete key value pair

7. Created a function called printAccounts that accepts an arguement of type map[int]string which will be our accountsMap.

8. Created for loop to iterate over each key value pair and print to terminal

------------------------------------------------------------------
Things to know about maps
------------------------------------------------------------------

1. All keys must be the same type
2. All values must be the values must be of the same type
3. Keys are indexed - we can iterate over them
4. Used to represent a collection of related properties 
5. Don't need to know all the keys at compile time
6. Reference type - This means that when map actually points to a location in In memory (heap) in ram to where the physical array is stored. 

**Note: as a reference type. when passed in as argument and modified, the address that the map points to will change. Thus updating value for all maps that point to the same location in the heap.


