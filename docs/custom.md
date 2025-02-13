### **Custom types**

Ok, let's start. We need to get gosli and to install it to our GOPATH/bin folder. All we need to do is run:
```
go get github.com/maxtroughear/gosli
```

Then, for example we have a file in our project `example/faketype.go`. And it contains a structure called `FakeType`.

Run:
```
GOPATH/gosli example/faketype.go FakeType
```
If you're using Windows, of course it should be `gosli.exe` instead of `gosli`.

This command will generate two files:
```
example/faketype_generated.go
example/faketype_equal.go
```

`faketype_generated.go` contains all the methods that gosli provides to use. This file shouldn't be edited manually.

In turn, `faketype_equal.go` has to be updated by you. After it's generated it contains the only method with such body:

```go
func (r *FakeType) equal(another *FakeType) bool {
	if r == nil && another == nil {
		return true
	}
	if (r == nil && another != nil) || (r != nil && another == nil) {
		return false
	}
	// `equal` method has to be implemented manually
}
```

In this method we should tell gosli how it should understand if two instances of a structure are equal. You can expect them to be equal, for example, if all of their fields are equal, or some `id` fields only matter. Anyway, you could implement this method whatever you like. The `equal` method will work in methods `Contains`, `GetUnion` and `InFirstOnly`.

Please put your implementation after the comment in this method (after `// 'equal' method`... line)

Example:
```go
func (r *FakeType) equal(another *FakeType) bool {
	if r == nil && another == nil {
		return true
	}
	if (r == nil && another != nil) || (r != nil && another == nil) {
		return false
	}
	// `equal` method has to be implemented manually
	return r.A == another.A &&
		r.B == another.B
}
```

That's it, now you can use the methods generated by gosli for your type. If our type's name is `FakeType` then two types will is generated:
* `type FakeTypeSlice []FakeType` is a type representing a slice of `FakeType` structures
* `type FakeTypePSlice []*FakeType` describes a slice of `FakeType` pointers slice (slice of `*FakeType`)


Let's look on the methods list. I will describe the methods of `FakeTypePSlice` but the idea will be the same both for `FakeTypePSlice` and `FakeTypeSlice`.

---
## **Methods**


* ### First
    Returns first item of a slice that is passed through a filter.

    If an item wasn't found, the method returns an error.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A == 2
    }
    res, err := FakeTypePSlice(sl).First(filter)

    //res = &FakeType{
    //    A: 2,
    //    B: "two",
    //} 
    ```

* ### FirstOrDefault
    Returns first item of a slice that is passed through a filter.

    If an item wasn't found, the result is nil for `FakeTypePSlice`.
    In the case of `FakeTypeSlice` it will be a default value of that type.

    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A == 2
    }
    res, err := FakeTypePSlice(sl).FirstOrDefault(filter)

    //res = &FakeType{
    //    A: 2,
    //    B: "two",
    //} 
    ```

* ### Where
    Returns all items of a slice that is passed through a filter.

    If items weren't found, the result is empty slice.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A >= 2
    }
    res, err := FakeTypePSlice(sl).Where(filter)

    //res = FakeTypePSlice{
	//    &FakeType{
    //        A: 2,
    //        B: "two",
    //    },
    //    &FakeType{
    //        A: 3,
    //        B: "three",
    //    },
    //} 
    ```

* ### Select
    Applies a function to every item of a slice and returns slice of results.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    f := func(t *FakeType) interface{} {
        return struct {
            Msg string
        }{
            Msg: t.B,
        }
    }
    res, err := FakeTypePSlice(sl).Select(f)

    //res = []struct {
	//    Msg string
	//}{
    //    {
    //        Msg: "one",
    //    },
    //    {
    //        Msg: "two",
    //    },
    //    {
    //        Msg: "three",
    //    },
	//} 
    ```

* ### Page
    Returns paginated slice according to given `number` (number of selected page) and `perPage` 
    (items per a page). `number` parameter should start with 1 (not 0).
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    res, err := FakeTypePSlice(sl).PerPage(1, 2)

    //res = FakeTypePSlice{
    //    &FakeType{
    //        A: 1,
    //        B: "one",
	//    },
    //    &FakeType{
    //        A: 2,
    //        B: "two",
    //    },
	//} 
    ```

* ### Any
    Returns `true` if any item of the slice is passed through a filter.

    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }
    
    filter = func(t *FakeType) bool {
        return t.A == 2
    }
    res, err := FakeTypePSlice(sl).Any(filter)

    //res = true
    ```

* ### Contains
    Returns `true` if a slice contains at least one item that is equal to the desired one.
    
    ```go
    sl := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }

    el := &FakeType{
        A: 2,
        B: "two",
    }
    
    res, err := FakeTypePSlice(sl).Contains(el)

    //res = true
    ```

* ### GetUnion
    Returns a slice that contains items that are contained in both given slices.
    
    ```go
    sl1 := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }

    sl2 := []*FakeType{
	    &FakeType{
            A: 2,
            B: "two",
	    },
        &FakeType{
            A: 3,
            B: "three",
        },
        &FakeType{
            A: 4,
            B: "four",
        },
    }

    res, err := FakeTypePSlice(sl1).GetUnion(sl2)

    //res = FakeTypePSlice{
	//    &FakeType{
    //        A: 2,
    //        B: "two",
	//    },
    //    &FakeType{
    //        A: 3,
    //        B: "three",
    //    },
    //} 
    ```

* ### InFirstOnly
    Returns elements that are contained only in a first slice and is not contained in a second one.
    
    ```go
    sl1 := []*FakeType{
	    &FakeType{
            A: 1,
            B: "one",
	    },
        &FakeType{
            A: 2,
            B: "two",
        },
        &FakeType{
            A: 3,
            B: "three",
        },
    }

    sl2 := []*FakeType{
	    &FakeType{
            A: 2,
            B: "two",
	    },
        &FakeType{
            A: 3,
            B: "three",
        },
        &FakeType{
            A: 4,
            B: "four",
        },
    }

    res, err := FakeTypePSlice(sl1).InFirstOnly(sl2)

    //res = FakeTypePSlice{
	//    &FakeType{
    //        A: 1,
    //        B: "one",
	//    },
    //} 
    ```

#### If the description looks unclear for you, please take a look at [`experiment` folder](https://github.com/maxtroughear/gosli/tree/master/experiment). You can find there unit test, benchmarks and some generated code that could describe the essent of the library much better than my poor English :)