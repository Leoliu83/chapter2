#### <center>线性表的顺序存储结构</center>
&emsp;&emsp;线性表的顺序存储结构指的是用一段连续的存储单元，依次存储线性表的数据元素。

##### 顺序存储方式
&emsp;&emsp;由于线性表中的每个元素的数据类型相同，因此可以使用**一维数组**来实现顺序存储结构。

###### 线性表顺序存储结构定义如下
``` c
# define MAXSIZE 20  /* 存储空间初始分配量 */
typedef int ElemType /* ElemType根据实际情况而定，这里假设是int */
typedef struct{
    ElemType data[MAXSIZE] /* 数组存储数据元素，最大为MAXSIZE */
    int length; /* 线性表当前长度 */
}SqList
```
描述顺序存储结构需要以下三个属性：
* 存储空间的起始位置：数组data，它的存储位置就是存储空间的存储位置
* 线性表的最大存储容量：数组长度 MAXSIZE
* 线性表的当前长度：length

##### 数据长度与线性表长度的区别
&emsp;&emsp;任意时刻，数据长度应该小于等于线性表长度

##### 地址计算方式
&emsp;&emsp;存储其中每一个存储单元的都有自己的编号，这个编号称之为地址，假设每个元素占用c个存储单元。那么线性表中第i+1个数据元素的地址等于第i个元素的位置（起始地址）+存储单元的大小:
$LOC(a_{i+1}) = LOC(a_{i})+c$
$LOC$ 表示获取元素位置函数

##### 顺序存储结构的插入与删除
###### 获取元素操作（GetElem）
``` c
#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0
typedef int Status;
/*Status 是函数的返回类型，其值是函数结果状态代码，入OK等*/
/*初始条件：顺序线性表L已经存在，1<i<=ListLength(L)*/
/*操作结果：用e返回L中第i个元素的值*/

Status GetElem(SqList l,int i,ElemType *e){
    if(L.length==0||i<1||i>L.length)
        return ERROR;
    *e=L.data[i-1];
    return OK;
}
```
###### 插入元素操作（ListInsert）
**元素插入算法思路：**
* 如果插入位置不合理，则抛出异常
* 如果线性表长度大于等于数组长度，则抛出异常或动态增加容量
* 从最后一个元素开始向前遍历到第i个位置，分别将他们向后移动一个位置
* 将要插入的元素填入位置i处
* 线性表长度+1

**实现代码如下：**
``` c
/*初始条件：顺序线性表L已经存在，1<i<=ListLength(L)*/
/*操作结果：在L中第i个位置之前插入新的元素e，L的长度加1*/
Status ListInsert(SqList *L,int i,ElemType e){
    int k;
    if(L->length == MAXSIZE){ /*顺序线性表已满*/
        return ERROR;
    }
    if(i<1||i>L->length-1){ /*当i不在范围内时*/
        return ERROR;
    }
    if(i<=L->length){ /*若插入数据位置不在表尾*/
        for(k=L->length-1;k>=i-1;i--){
            L->data[k+1]=L->data[k];
        }
    }
    L->data[i-1]=e;
    L->length++
    return OK;
}
```
###### 删除操作（ListDelete）
**删除算法思路：**
* 如果删除位置不合理，抛出异常
* 取出将被删除的元素
* 从删除位置开始遍历到最后一个元素，将每一个元素向前移动一个位置
* 线性表长度-1

``` c
Status ListDelete(SqList *L,int i,ElemType e){
    int k
    if(L->length==0){ /*线性表为空*/
        return ERROR;
    }
    if(i<1 || i>L->length){
        return ERROR;
    }
    *e=L->data[i-1]
    if(i<L->length){
        for(k=i;k<L->length;k--)
            L->data[k-1] = L->data[k]
        L->length--
        return OK;
    }
}
```

**插入和删除元素的时间复杂度**
* 最好的情况: 插入和删除的元素都是最后一个元素，时间复杂度为$O(1)$
* 最坏的情况: 插入和删除的元素都是第一个元素，时间复杂度为$O(n)$
* 一般情况: 插入和删除的元素在中间 $i$ 的位置，需要移动 $n-i$ 个元素，根据概率原理，每个插入或删除的元素的可能性是相同的，也就是说，位置靠前，移动元素多，位置靠后，移动元素少，最终的平均移动次数和最中间那个元素的移动次数相等，因此为 $\frac{n-1}{2}$，根据时间复杂度推导，时间复杂度仍然为$O(n)$

###### 线性表的优缺点
| 优点 | 缺点 |
| :--: | :--: |
| 无需为表示表中元素件的逻辑关系而增加额外的存储空间 | 删除和插入操作需要移动大量的数据  |
| 可以快速的取表中任意位置的元素 | 当线性表长度变化较大时，难以确定存储空间的容量 |
| &nbsp;&emsp; | 造成存储空间碎片 |
