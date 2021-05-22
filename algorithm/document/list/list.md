#### <center>线性表（List）<center/>
##### 线性表的定义：
&emsp;&emsp;线性表是零个或者多个元素的有限序列。
&emsp;&emsp;用数学语言来进行定义如下：
&emsp;&emsp;若将线性表记为 $(a_1,\cdots,a_{i-1},a_i,a_{i+1},\cdots,a_n)$，则表中 $a_{i-1}$ 领先于 $a_{i}$，则称 $a_{i-1}$ 是 $a_{i}$ 的**直接前驱元素**， $a_{i+1}$ 是 $a_{i}$ 的**直接后继元素**。当${i}\le{n-1})$ 时 $a_i$只有一个直接后继元素，当 ${i}\ge{2}$ 时，$a_i$有且仅有一个直接前驱元素。
&emsp;&emsp;线性表的元素个数$n({n}\ge{0})$定义为线性表的长度，当$n=0$时，称为空表。
#### 线性表的抽象数据类型定义如下：
``` c
ADT线性表（List）
    Data
    Operation
        InitList(l *List) bool; // 初始化操作，建立一个新的L。
        ListEmpty(l List) bool; // 若线性表为空，返回true，否则返回false
        ClearList(l *List) bool; // 清空线性表
        GetElem(l List,i int,e *Element) bool; // 将线性表List中第i个元素返回给*Element
        LocateElem(l List,e Element) (int,bool); // 在线性表中查找与给定元素e相同的元素，如果查找成功，返回该元素在表中序号，如果失败，返回-1
        ListInsert(l *List,i int,e Element) bool; // 在线性表List中第i个位置插入新元素
        ListDelete(l *List,i int,e *Element) bool; // 删除线性表List中的第i个位置的元素，并返回其值到e
        ListLength(l List) (int,bool); // 返回线性表L的元素个数
end ADT
```
ADT表示**抽象数据类型（Abstract Data Type）**

---