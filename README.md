 **weightedarrays**

To use this library you need to implement the Value interface.
```
type Fruit struct {  
 Name   string 
 Weight float64 
 Draws  int
}
   
func (f Fruit) GetWeight() float64 {  
 return Weight}  
 
items := weightedarray.Values{  
 &Fruit{Name: "Banana", Weight: 1}, 
 &Fruit{Name: "Lemon", Weight: 2}, 
 &Fruit{Name: "Apple", Weight: 4}, 
 &Fruit{Name: "Mellon", Weight: 8}, 
 &Fruit{Name: "Orange", Weight: 10}, 
 &Fruit{Name: "Pineapple", Weight: 7},
}
```

To retrieve a weighted value you call

```
 value =: items.Random()
```
