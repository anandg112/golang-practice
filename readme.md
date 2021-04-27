# Go Code

## Data types

- String  
- integer
- float
- array
- map

**Array** - Fixed length list of things  
**Slice** - An array that can grow or shrink. Every element in a slice must be of same type  
**Map** - Used to represent a collection of related properties. Dict-like object.
  - All keys must be of the same type
  - All values must be of the same type
  - Keys are indexed, can iterate over them
  - Don't need to know all the keys at compile time
  - Reference type  

**Structs** - Used to represent a "thing" with a lot of different properties
  - Values can be of different types
  - Keys don't support indexing
  - Need to know all the different fields at compile time
  - Value type  

A function with a receiver is like a "method" - a function that belongs to an instance

Interfaces are not generic types (other languages have generic types, Go does not)

Interfaces are implicit. We don't manually have to say that our custom type satisifies
some interface.

Interfaces are a contract to help us manage types. If our custom type's implementation of a function is broken then interfaces wont help us.

