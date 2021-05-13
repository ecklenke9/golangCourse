<h1>GOLANG:</h1> 

    - Efficient compilation
    - Efficient execution
    - Ease of programming
    

<h5>We create VALUES of a certain TYPE that are stored
in VARIABLES, and those VARIABLES have IDENTIFIERS. </h5>


<h1>Is Go an object-oriented language?</h1>

Yes and no. Although Go has types and methods and allows an object-oriented style of programming, 
there is no type hierarchy. The concept of “interface” in Go provides a different approach that we 
believe is easy to use and in some ways more general. There are also ways to embed types in other 
types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are 
more general than in C++ or Java: they can be defined for any sort of data, even built-in types such 
as plain, “unboxed” integers. They are not restricted to structs (classes).

Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages 
such as C++ or Java.

Go has OOP aspects. But, it’s all about TYPE. We create TYPES in Go; user-defined TYPES. We can then 
have VALUES of that type. We can assign VALUES of a user-defined TYPE to VARIABLES. Anecdote: makes 
me think of that song, “It’s all about the bass, all about the bass” except “it’s all about the TYPE, 
all about the TYPE.”

Padding & architectural alignment
Convention: logically organize your fields together. Readability & clarity trump performance as a 
design concern. Go will be performant. Go for readability first. However, if you are in a situation 
where you need to prioritize performance: lay the fields out from largest to smallest, 
eg, int 64, int64, float32, bool. 

