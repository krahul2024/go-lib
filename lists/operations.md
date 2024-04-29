1. **Constructor/Destructor:**
   - `list()`: Constructs an empty list.
   - `list(size_type count)`: Constructs the list with `count` default-inserted instances of T.
   - `list(size_type count, const T& value)`: Constructs the list with `count` copies of `value`.
   - `list(const list& other)`: Constructs the list with the copy of the contents of `other`.
   - `~list()`: Destructs the list.

2. **Assignment Operator:**
   - `operator=`: Assigns values from another list.

3. **Modifiers:**
   - `push_back(const T& value)`: Adds a new element at the end of the list. -> done 
   - `push_front(const T& value)`: Adds a new element at the beginning of the list. -> done 
   - `pop_back()`: Removes the last element of the list.
   - `pop_front()`: Removes the first element of the list.
   - `insert(iterator pos, const T& value)`: Inserts a new element before the element at the specified position.
   - `erase(iterator pos)`: Removes the element at the specified position.
   - `clear()`: Removes all elements from the list.
   - `swap(list& other)`: Swaps the contents of two lists.

4. **Element Access:**
   - `front()`: Returns a reference to the first element.
   - `back()`: Returns a reference to the last element.

5. **Capacity:**
   - `empty()`: Checks whether the list is empty.
   - `size()`: Returns the number of elements in the list.

6. **Iterators:**
   - `begin()`, `end()`: Returns iterators referring to the first and one past the last element of the list, respectively.
   - `rbegin()`, `rend()`: Returns reverse iterators.

7. **Operations:**
   - `splice(iterator pos, list& other)`: Transfers elements from another list into the specified position.
   - `splice(iterator pos, list& other, iterator it)`: Transfers an element from another list to the specified position.
   - `splice(iterator pos, list& other, iterator first, iterator last)`: Transfers elements from another list into the specified position.
   - `merge(list& other)`: Merges two sorted lists into one.
   - `sort()`: Sorts the elements in the list.
   - `reverse()`: Reverses the order of elements in the list.
