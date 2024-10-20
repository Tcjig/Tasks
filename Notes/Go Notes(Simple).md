# Go Notes(Simple)

## 变量作用域

在局部中局部变量比全局变量优先考虑



## 数组

特点：

+ 长度固定、类型相同
+ 可通过索引读取或修改

声明：``var arrayName [size]dataType``

创建数组：

```
//初始值为0
var number [5]int

//初始化列表
var number = [5]int{1,2,3,4,5}
number := [5]int{1,2,3,4,5}

//数组长度不固定,用...代替
number := [...]int{1,2,3}
```



## 指针

一个指针变量指向了一个值的内存地址(变量和常量)

声明：``var name *type``

判断是否为空指针：``ptr != nil``



## 结构体

创建结构体并声明变量：``type str_var_type struct {val,val...}``

访问结构体成员：``结构体.成员名``

将变量声明为结构体类型：``var val_name str_name``



## 切片(动态数组)

特点：

+ 长度不固定
+ 可追加元素，使容量增大

声明：``var slicename []type``

创建切片：

```
//make创建切片
slice := make([]type,len)

//也可指定容量
make([]Type,length,capacity)
```

### 切片截取

代码演示

```
number := []int{0,1,2,3,4,5,6,7,8}

//打印切片从索引1(包含)到索引4(不包含)
fmt.Println(number[1:4])

//打印切片从索引3开始
fmt.Println(number[3:])

//打印切片到索引3
fmt.Println(number[:3])
```

### 切片追加、复制

代码演示

```
var number []int

//添加三个元素1，2，3
number = append(number,1,2,3)

//复制切片,将number的内容加到number1
copy(number1,number)
```



## 范围

作用：用于在for循环中对slice、map、数组、字符串进行迭代循环

**同for循环的区别：**for循环只能遍历数组和切片

语法

```
 //key对应索引，value对应索引值
 for key,value := range oldMap{
 	newMap[key] = value
 }
 
 //在数组、切片中的应用，以pow为例
 for i, v := range pow
 
 //在字符串中的应用，以hello为例
 for i, c := range "hello"
 
 //在Map中的应用，以map1为例
 for key, value := range map1 //读取key和value
 for _, value := range map1 //只读取value
```



## 集合

特点：

+ 无序遍历，每次遍历的值排列顺序不同
+ 若该值不存在，则返回该类型的零值
+ Map为引用类型，对Map的修改会影响到所有引用它的变量

语法

```
//定义一个Map
map_var := make(map[keyType]valType,capacity)

//创建一个初始容量为10的Map
m := make(map[string]int,10)

//使用字面量创建Map
m := map[string]int{
	"apple": 1,
	"banana": 2,
	"orange": 3,
}

//获取元素(键值对)
v1 := m["apple"]
v2, ok := m["pear"] //若键不存在，两个变量的值都取零值

//修改元素(键值对)
m["apple"] = 5

//获取Map的长度
len := len(m)

//删除元素(键值对)
delete(m,"banana")
```



## 递归函数(设置退出条件，否则陷入无限循环)

含义：在本函数中调用自己

语法

```
//基准条件：函数在满足某个条件时会停止递归调用,返回最终结果。
//递归调用：在基准条件未满足时,函数会不断调用自身来推进问题的求解过程。
func func_name(para) return_type {
	if 基准条件 {
		return 基准条件返回值
	}
	return func_name(递归调用参数)
}
```



## 接口

理解：就是先定义一个接口里面包含方法，然后定义一个结构体里面包含变量，之后通过方法和结构体名构成函数并完善其函数功能，之后即可使用，这就称之为接口

语法 

```
//定义接口
type inter_name interface {
	method_name return_type
}

//定义结构体
type str_name struct {
	//定义变量
}

//实现接口方法
func (str_name_var str_name) meth_name() return_type {
	//方法实现
}
```

