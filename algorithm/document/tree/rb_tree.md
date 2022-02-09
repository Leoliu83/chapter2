#### 红黑树
R-B Tree，全称是Red-Black Tree，又称为“红黑树”，它一种特殊的二叉查找树。红黑树的每个节点上都有存储位表示节点的颜色，可以是红(Red)或黑(Black)。

##### 红黑树的特点
1. 节点是红色或黑色。
2. 根节点是黑色。
3. 每个叶节点（NIL节点，空节点）是黑色的。
4. 每个红色节点的两个子节点都是黑色。(从每个叶子到根的所有路径上不能有两个连续的红色节点)
5. 从任一节点到其每个叶子的路径上包含的黑色节点数量都相同。




##### java中红黑树的解析
##### Java中hashmap源码中的红黑树学习
``` java
int MIN_TREEIFY_CAPACITY = 64
/* 
  将指定hash桶下的链表转化为红黑树
  @param tab  hash桶
  @param hash hash值
*/
final void treeifyBin(Node<K,V>[] tab, int hash) {
    int n, index; Node<K,V> e;
    // 如果桶的数量（也就是tab.length）小于最小树化的阈值（也就是64），则扩容
    // 这里的tab也就是hash桶
    if (tab == null || (n = tab.length) < MIN_TREEIFY_CAPACITY)
        resize(); // 扩容
    // 如果hash桶已经存在，并且当前的hash值，不在桶里，并且桶的大小已经>=MIN_TREEIFY_CAPACITY阈值
    // (n - 1) & hash 这个就是经典的取余操作，也就是 当 n为 2的幂次的时候，hash%n = (n - 1) & hash
    else if ((e = tab[index = (n - 1) & hash]) != null) {
        // 下面的步骤，就是将hash桶中e这个hash值所挂的链表，全部转换为红黑树
        // hd 表示 head，tl(tree leaf)记录了上一次循环所遍历到的TreeNode
        TreeNode<K,V> hd = null, tl = null;
        //迭代链表,将链表转换为红黑树
        do {
            // 将链表的结点 e 转化为红黑树结点
            TreeNode<K,V> p = replacementTreeNode(e, null);
            // 如果tl为空，则将p设置为head
            if (tl == null)
                hd = p;
            else { // 如果tl不为空，即记录上一次循环的TreeNode，也就是说，不是循环的第一次
                // 则将 tl设置为p的前一个结点，而p设置为tl的下一个结点
                // 也就是将上一次遍历的TreeNode.next设置为当前遍历到的结点，将当前遍历的TreeNode.prev设置为上一次遍历的TreeNode
                // 也就是互相持有对方的引用
                p.prev = tl;
                tl.next = p;
            }
            // 将当前TreeNode记录到变量tl
            tl = p;
        } while ((e = e.next) != null); // 循环取出链表中的元素，如果e.next == null 则表示链表结束
        if ((tab[index] = hd) != null) // 把 head 记录到 hash桶，BTW: hash值在哪个桶由 index 决定
            hd.treeify(tab); // 转换
    }
}
```

``` java
// 红黑树结点的属性如下：
TreeNode<K,V> parent;  // red-black tree links
TreeNode<K,V> left;    // 左枝
TreeNode<K,V> right;   // 右枝
TreeNode<K,V> prev;    // needed to unlink next upon deletion
boolean red;
// 红黑树的方法如下：
// 1. 移动root到最前面，确保root是树的第一个结点
void moveRootToFront(Node<K,V>[] tab, TreeNode<K,V> root)
/**
  * 2. 从root结点开始查找指定hash值和key的元素
  * kc 在第一次键比较时，缓存了 comparableClassFor(key) 的结果，缓存代码如下
  * (kc != null || (kc = comparableClassFor(k)) != null) 
  *        && (dir = compareComparables(kc, k, pk)) != 0 // 执行 (Comparable)k).compareTo(pk)
  */
TreeNode<K,V> find(int h, Object k, Class<?> kc)
/*
  3. 根据hash值寻找指定k的值
  h: hash值
  k: key
*/
TreeNode<K,V> getTreeNode(int h, Object k)
/*
  返回root结点，通过 parent 反复查找，直到parent==null  返回
*/
TreeNode<K,V> root()
/*
  
*/
void treeify(Node<K,V>[] tab)
```



###### moveRootToFront
```java
/**
  Ensures that the given root is the first node of its bin.
  确保桶中的元素一定是给定元素，并且是树的根元素
  该函数在 balanceInsertion 函数获取root后调用

1. 
+-------+
|       |
+-------+
| first | → fnext
+-------+

prev → root → next

2. 
+-------+
|       |
+-------+
|  root |
+-------+

3.
+-------+
|       |
+-------+   prev → next
|       | ↗ 
|  root |
|       | ↘
+-------+   first → fnext
             

*/
static <K,V> void moveRootToFront(Node<K,V>[] tab, TreeNode<K,V> root) {
    int n;
    if (root != null && tab != null && (n = tab.length) > 0) {
        int index = (n - 1) & root.hash; // 根据root的hash值计算桶的下标
        TreeNode<K,V> first = (TreeNode<K,V>)tab[index]; // 取出桶中的第一个元素
        if (root != first) { // 如果 第一个元素，不是给定的root
            Node<K,V> rn;
            tab[index] = root; // 把桶的结点设置为root
            TreeNode<K,V> rp = root.prev; // rp 设置为root的前一个结点
            if ((rn = root.next) != null) // 如果 root的下一个结点不为空
                ((TreeNode<K,V>)rn).prev = rp; // 将root的下一个结点的前一个结点设置为 root的原前一个结点
            if (rp != null) // 如果root的前一个结点不为空
                rp.next = rn; // 将root的前一个结点的下一个结点，设置为root的下一个结点
            if (first != null) // 如果原来的桶中的结点不为空
                first.prev = root; // 则将它的前一个结点设置为root
            root.next = first;
            root.prev = null;
        }
        assert checkInvariants(root);
    }
}

```
###### identityHashCode的例子

```java
Object s1 = "abcd";
Object s2 = new String("abcd");
Object s3 = "abcd";
System.out.println("identityHashCode : " + System.identityHashCode(s1) + " HashCode : " + s1.hashCode());
System.out.println("identityHashCode : " + System.identityHashCode(s2) + " HashCode : " + s2.hashCode());
System.out.println("identityHashCode : " + System.identityHashCode(s3) + " HashCode : " + s3.hashCode());

//输出结果:
identityHashCode : 2018699554 HashCode : 2987074
identityHashCode : 1311053135 HashCode : 2987074
identityHashCode : 2018699554 HashCode : 2987074 
```
总结：identityHashCode 由于对象的内存地址的不同而不同，identityHashCode() 方法由 HotSpot 虚拟机实现

###### putTreeVal 源码+注释
```java
/*
    Tree version of putVal.
    当存在hash碰撞的时候，且元素数量大于8个时候，就会以红黑树的方式将这些元素组织起来
    map 当前节点所在的HashMap对象
    tab 当前HashMap对象的元素数组
    h   指定key的hash值
    k   指定key
    v   指定key上要写入的值
    返回：指定key所匹配到的节点对象，针对这个对象去修改V（返回空说明创建了一个新节点）

    由于hash值求余到tab的下标时，发生了碰撞（也就是已经有了值），因此 该hash值也应该挂在当前桶(tab)下
    但是求余得到的下标相同，但是hash值本身不同，因此 这里的h传递的是 key 真实的 hash值
*/
final TreeNode<K,V> putTreeVal(HashMap<K,V> map, Node<K,V>[] tab, int h, K k, V v) {
    Class<?> kc = null;
    // searched 表示 是否已经检索了它所有的子树
    boolean searched = false;
    // 如果<当前结点> 的父结点不为空，则通过root()方法递归找到root结点
    // 如果<当前结点> 的父结点为空，则当前结点为root结点
    TreeNode<K,V> root = (parent != null) ? root() : this;
    // 循环查找，后续都在用变量p，p初始化为root，后续在循环中，p会不断变化
    for (TreeNode<K,V> p = root;;) {
        // dir 表示左还是右，如果为-1 则为左，1为右
        int dir, ph; K pk;
        if ((ph = p.hash) > h) // 如果参数传递的hash值（参数h） 小于 root结点的hash值，
            dir = -1;  // 记录，应该向左继续查找
        else if (ph < h) // 如果参数传递的hash值（参数h） 大于 root结点的hash值
            dir = 1;   // 记录，应该向右继续查找
        else if ((pk = p.key) == k || (k != null && k.equals(pk))) // 如果hash值正好相等，说明找到了key
            return p;   // 返回 当前结点，说明找到了
        else if ((kc == null &&
                 (kc = comparableClassFor(k)) == null) ||
                 (dir = compareComparables(kc, k, pk)) == 0) { 
            // 这个 else if 表示无法用 > 或者 < 来比较大小，因此需要通过 compareTo 来比较大小
            // comparableClassFor(k) 方法：如果k实现了 comparable 接口返回一个Class对象，否则返回null
            // compareComparables(kc, k, pk) 方法：如果 x 匹配 kc(也就是k所属的类) k.compareTo(pk) 

            // 如果还没有检索过子节点，那么久遍历 p 的所有子节点，包括儿子节点，孙子节点......，一直到 叶子结点
            // 遍历所有子节点由 find 方法实现，这里是 ch.find
            if (!searched) {
                TreeNode<K,V> q, ch;
                searched = true;
                // ch.find 表示基于ch结点，查找所有的子节点，是否有符合要求的结点
                // Finds the node starting at root p with the given hash and key.
                if (((ch = p.left) != null &&
                        (q = ch.find(h, k, kc)) != null) ||  
                    ((ch = p.right) != null &&
                        (q = ch.find(h, k, kc)) != null))
                    return q;
            }
            // 上一步遍历了所有子节点也没有找到和当前键equals相等的节点，也就是说，没有办法对于现有的key直接做值的替换，下面就需要新加结点了！！！
            // tieBreakOrder 方法的作用是：**当hashCodes相等且不可比较时!**，用于排序插入的工具。
            // 这里要的其实不是一个总的顺序，而是一个一致的插入规则，以保持在重新平衡中的等价性。
            // 在必要的情况下进一步打破平局可以简化测试。方法内部，当System.identityHashCode(k) == System.identityHashCode(pk) 时，也返回 -1，也就是说，进入左枝，identityHashCode 返回的是 系统默认的hashCode()方法，即使重写了，也不会调用重写过的hashCode方法
            // 至此，如果是否可比较，都可以确定dir为-1还是1
            dir = tieBreakOrder(k, pk);
        } // if 结束

        // 保存当前结点到 xp 变量，第一次循环时，xp为root
        TreeNode<K,V> xp = p;
        // 如果 dir<=0 则判断<左枝>是否为空，并保存到变量 p
        // 如果 dir>=0 则判断<右枝>是否为空，并保存到变量 p
        // *** 这里开始，p就发生了变化，这里是如果左枝或者是右枝为空(根据dir来决定判断左枝还是判断右枝)，如果不为空，就会进入下一次循环
        // 如果左或右都不为空则会进入下一次循环
        if ((p = (dir <= 0) ? p.left : p.right) == null) {
            // 将当前结点的next保存起来
            Node<K,V> xpn = xp.next;
            // 新增结点（putTreeVal方法就是插入结点，这里开始真正处理插入结点动作了）
            // h:   当前新节点hash值
            // k:   当前新节点键
            // v:   当前新节点值
            // xpn: 当前新节点next结点
            TreeNode<K,V> x = map.newTreeNode(h, k, v, xpn);
            // 根据dir(-1 左枝，1 右枝)设置当前结点的左枝或者右枝为新结点
            if (dir <= 0)
                xp.left = x;
            else
                xp.right = x;
            xp.next = x;
            // 设置 新节点的 父结点，前一结点为 当前结点
            x.parent = x.prev = xp;
            if (xpn != null) // 如果当前结点的下一个结点不为空
                ((TreeNode<K,V>)xpn).prev = x; // 设置原当前结点的下一个结点的前一个结点为新结点
            // 上面这段也就是相当于  将 新结点 加到 当前结点和当前结点的下一个结点之间
            moveRootToFront(tab, balanceInsertion(root, x));
            return null;
        }
    }
}
```

###### balanceInsertion 源码+注释
```java
/*
    balanceInsertion 指的是红黑树的插入平衡算法，当树结构中新插入了一个节点后，要对树进行重新的结构化，以保证该树始终维持红黑树的特性。
    root 为<当前根结点>
    x    为<当前插入结点>
    返回值为<新的根节点>
*/
static <K,V> TreeNode<K,V> balanceInsertion(TreeNode<K,V> root, TreeNode<K,V> x) {
    x.red = true; // 新结点先设置为红
    /*
        xp   父亲结点
        xpp  爷爷结点
        xppl 爷爷结点的左枝（当xp是爷爷结点的右枝时，xppl叔叔结点，xppr就是父亲结点）
        xppr 爷爷结点的右枝（当xp是爷爷结点的左枝时，xppr叔叔结点，xppl就是父亲结点）
    */
    for (TreeNode<K,V> xp, xpp, xppl, xppr;;) {
        // 1. 如果父结点为空，即插入结点位置为根结点，不需要调整，直接设置为黑色，并返回当前的根结点 x
        if ((xp = x.parent) == null) { 
            x.red = false;
            return x;
        } // 2. 如果父亲节点是黑色 或者 爷爷结点为空，则不需要调整，那么直接返回root
        else if (!xp.red || (xpp = xp.parent) == null)
            return root;
        // 上述两种情况是不需要调整的情况，下面是需要调整的情况
        /*
        如果父节点(xp)是爷爷结点的左结点(xppl)，即xppr是叔叔结点
                 pp
                /  \
              xp    xppr
        */
        if (xp == (xppl = xpp.left)) {
            /* 
              3. 下面的if表示，如果叔叔结点不为空，且是红色，上面父节点为黑色的情况已经处理过了,
                 因此这里处理的情况是 父节点也为红色
                 根据红黑树的特征4：每个红色节点的两个子节点都是黑色。(从每个叶子到根的所有路径上不能有两个连续的红色节点)，可以推出 pp 一定是黑色
                         pp (black)
                        /  \
                (red) xp    xppr (red)
            */
            if ((xppr = xpp.right) != null && xppr.red) {
            /*
                对于这种情况，只需要将 xppr,xp 都设置成黑色，将 pp设置为红色
                           pp (red)
                          /  \
                (black) xp    xppr (black)
                但是 如果 pp的父节点也是红色，那就需要让爷爷结点作为新加入的结点左递归处理
            */
                xppr.red = false; // 叔叔结点设置为黑色
                xp.red = false;   // 父结点设置为黑色
                xpp.red = true;   // 爷爷结点设置为红色
                x = xpp;          // x 设置为爷爷结点（这里这么做的原因是：这时爷爷结点变红了，如果此时爷爷的父节点也是红的，那就必须把爷爷也看成一个新结点，继续处理）
            }
            /*
              4. 叔叔结点为空 或者 叔叔结点是黑色
                         pp (black)
                        /  \
                (red) xp    xppr (black or nil)
            */
            else {
                /*
                         pp (black)
                        /  \
                (red) xp    xppr (black or nil)
                        \
                         x (red)
                */
                if (x == xp.right) { // 如果当前新结点是父结点的右结点
                    /*
                      先左旋，旋转后如下所示
                             pp (black)
                            /  \
                     (red) x    xppr (black or nil)
                          /
                  (red) xp
                        这里设置 x = xp 就是将原来的xp变成了新节点
                             pp (black)
                            /  \
                    (red) xp    xppr (black or nil)
                          /
                   (red) x  <-- 这个原来是xp
                    */
                    root = rotateLeft(root, x = xp);
                    // 左旋之后，处理就和"新节点是父结点的左结点"做同样处理了
                    // 如果当前结点的父结点为空，则 爷爷结点为空，否则则直接获取爷爷结点
                    xpp = (xp = x.parent) == null ? null : xp.parent;
                }
                // 如果 父结点不为空（叔叔结点为空 或者 叔叔结点是黑色）
                /*
                    上面的步骤将 x是xp的右枝的情况都变成了左枝的情况
                    所以后续的所有操作都是针对下面这种形状
                             pp (black)
                            /  \
                    (red) xp    xppr (black or nil)
                          /
                   (red) x
                */
                if (xp != null) {
                    xp.red = false;     // 父节点设置为黑
                    if (xpp != null) {  // 如果爷爷结点不为空
                        xpp.red = true; // 爷爷结点设置为红色
                        root = rotateRight(root, xpp); // 右旋转
                    }
                }
            }
        }
        /*
                     pp (black)
                    /  \
           (red) xppl   xp (red)
        */
        else { // 如果父节点(xp)不是爷爷结点的左结点(xppl)，即xppl是叔叔结点
            if (xppl != null && xppl.red) { // 如果叔叔结点是红色
                xppl.red = false; // 叔叔结点设置为黑色
                xp.red = false;   // 父结点设置为黑色
                xpp.red = true;   // 爷爷结点设置为红色
                x = xpp;  // x 设置为爷爷结点（这里这么做的原因是：这时爷爷结点变红了，如果此时爷爷的父节点也是红的，那就必须把爷爷也看成一个新结点，继续处理）
            }
            // 如果叔叔结点是黑色
            else {
                /*
                         pp (black)
                        /  \
     (black or nil) xppl    xp (red)
                           /
                          x (red)
                */
                if (x == xp.left) { // 如果当前新结点是父结点的左结点
                    // 先右旋
                    /*
                                             pp (black)
                                            /  \
                        (black or nil) xppl     x (red)
                                                 \
                                                  xp (red)
                        同样，这里设置 x = xp 是将xp看做是新结点
                                             pp (black)
                                            /  \
                        (black or nil) xppl     xp (red)
                                                 \
                                                  x (red) <-- 原来的xp
                    */
                    root = rotateRight(root, x = xp);
                    // 右旋之后，处理就和"新节点是父结点的右结点"做同样处理了
                    // 如果当前结点的父结点为空，则 爷爷结点为空，否则则直接获取爷爷结点
                    xpp = (xp = x.parent) == null ? null : xp.parent;
                }
                // 如果 父结点不为空（叔叔结点为空 或者 叔叔结点是黑色）
                /*
                    上面的步骤将 x是xp的左枝的情况都变成了右枝的情况
                    所以后续的所有操作都是针对下面这种形状
                                        pp (black)
                                       /  \
                    (black or nil) xppl    xp (red)
                                            \
                                             x (red)
                */
                if (xp != null) {
                    xp.red = false; // 父节点设置为黑
                    if (xpp != null) { // 如果爷爷结点不为空
                        xpp.red = true; // 爷爷结点设置为红色
                        root = rotateLeft(root, xpp); // 左旋
                    }
                }
            }
        }
    }
}
```

###### 左旋源码+注释
```java
/* ------------------------------------------------------------ */
// Red-black tree methods, all adapted from CLR
/*
左旋针对下面三种情况处理
情况1. 
      pp
     /
    p
   / \
 pl   r
     / \
    rl  rr

情况2. 
  pp
    \
     p
    / \
  pl   r
      / \
     rl  rr

情况3.
     p
    / \
  pl   r
      / \
     rl  rr  
*/
/*
  左旋,这里返回的是新root，也就是旋转后，r变成了root，则返回r作为新root
  也就是下面这种情况,情况3
    p
   / \
 pl   r
     / \
    rl  rr
*/
static <K,V> TreeNode<K,V> rotateLeft(TreeNode<K,V> root, TreeNode<K,V> p) {
    TreeNode<K,V> r, pp, rl;
    // 如果 p不为空，且p的右结点r不为空
    if (p != null && (r = p.right) != null) {
        // 将r的左节点赋值给p的右结点
        /*
                p
               / \
             pl   rl
        */
        if ((rl = p.right = r.left) != null)
            rl.parent = p;
        /*
              pp
              |  
              r
            如果旋转之后，pp为null说明r为root
            这里只设置了parent，并没有处理pp的左右枝的情况，所以这里用竖线
        */
        if ((pp = r.parent = p.parent) == null)
            (root = r).red = false;
        /*
            如果原来是下面这样的
                pp
               /  
              p
            那么将变成：
               pp
               /  
              r
        */
        else if (pp.left == p)
            pp.left = r;
        /*
        如果原来是下面这样的
            pp
              \  
               p
        那么将变成：
            pp
              \
               r
        */
        else
            pp.right = r;
        /*
                  r
                 / \
                p   rr
               / \
             pl   rl
        */
        r.left = p;
        p.parent = r;
    }
    return root;
}
```

###### 右旋源码+注释
``` java
/*
右旋针对下面三种情况处理
情况1. 
      pp
     /
    p
   / \
  l   pr
 / \
ll  lr

情况2. 
 pp
   \
    p
   / \
  l   pr
 / \
ll  lr

情况3.
    p
   / \
  l   pr
 / \
ll  lr  
*/
static <K,V> TreeNode<K,V> rotateRight(TreeNode<K,V> root, TreeNode<K,V> p) {
        TreeNode<K,V> l, pp, lr;
        // p 不为空，且p的左节点不为空（如果两个有一个为空都不需要旋转）
        if (p != null && (l = p.left) != null) {
            /*
                p
               / \
             lr   pr
            */
            if ((lr = p.left = l.right) != null)
                lr.parent = p;
            /*
                pp
                |  
                l
                如果旋转之后，pp为null说明 l 为root，说明未旋转前十下面这种情况
                    p
                   / \
                  l   pr
                 / \
               ll   lr
                这里只设置了parent，并没有处理pp的左右枝的情况，所以这里用竖线
            */
            if ((pp = l.parent = p.parent) == null)
                (root = l).red = false;
            /*
               pp
                \
                 l
                / \
              ll   lr
            */
            else if (pp.right == p)
                pp.right = l;
            /*
                   pp
                  /
                 l
                / \
              ll   lr
            */
            else
                pp.left = l;
            /*
                 l
                / \
              ll   p
                  / \
                lr   pr
            */
            l.right = p;
            p.parent = l;
        }
        return root;
    }
```

###### 查找源码+注释
``` java
/*
  Finds the node starting at root p with the given hash and key.
  The kc argument caches comparableClassFor(key) upon first use
  comparing keys.
*/
final TreeNode<K,V> find(int h, Object k, Class<?> kc) {
    TreeNode<K,V> p = this; // 把当前对象赋给p，表示当前节点
    do { // 循环
        int ph, dir; K pk; // 定义当前节点的hash值、方向（左右）、当前节点的键对象
        TreeNode<K,V> pl = p.left, pr = p.right, q; // 获取当前节点的左孩子、右孩子。定义一个对象q用来存储并返回找到的对象
        if ((ph = p.hash) > h) // 如果当前节点的hash值大于k得hash值h，那么后续就应该让k和左孩子节点进行下一轮比较
            p = pl; // p指向左孩子，紧接着就是下一轮循环了
        else if (ph < h) // 如果当前节点的hash值小于k得hash值h，那么后续就应该让k和右孩子节点进行下一轮比较
            p = pr; // p指向右孩子，紧接着就是下一轮循环了
        else if ((pk = p.key) == k || (k != null && k.equals(pk))) // 如果h和当前节点的hash值相同，并且当前节点的键对象pk和k相等（地址相同或者equals相同）
            return p; // 返回当前节点


        // 执行到这里说明 hash比对相同，但是pk和k不相等
        else if (pl == null) // 如果左孩子为空
            p = pr; // p指向右孩子，紧接着就是下一轮循环了
        else if (pr == null)
            p = pl; // p指向左孩子，紧接着就是下一轮循环了

        // 如果左右孩子都不为空，那么需要再进行一轮对比来确定到底该往哪个方向去深入对比
        // 这一轮的对比主要是想通过comparable方法来比较pk和k的大小     
        else if ((kc != null || (kc = comparableClassFor(k)) != null) && (dir = compareComparables(kc, k, pk)) != 0)
            p = (dir < 0) ? pl : pr; // dir小于0，p指向右孩子，否则指向右孩子。紧接着就是下一轮循环了

        // 执行到这里说明无法通过comparable比较  或者 比较之后还是相等
        // 从右孩子节点递归循环查找，如果找到了匹配的则返回    
        else if ((q = pr.find(h, k, kc)) != null) 
            return q;
        else // 如果从右孩子节点递归查找后仍未找到，那么从左孩子节点进行下一轮循环
            p = pl;
    } while (p != null); 
    return null; // 为找到匹配的节点返回null
}
```
###### 删除结点源码+注释
``` java
/**
* Removes the given node, that must be present before this call.
* This is messier than typical red-black deletion code because we
* cannot swap the contents of an interior node with a leaf
* successor that is pinned by "next" pointers that are accessible
* independently during traversal. So instead we swap the tree
* linkages. If the current tree appears to have too few nodes,
* the bin is converted back to a plain bin. (The test triggers
* somewhere between 2 and 6 nodes, depending on tree structure).
*/
final void removeTreeNode(HashMap<K,V> map, Node<K,V>[] tab,
                            boolean movable) {
    int n;
    if (tab == null || (n = tab.length) == 0)
        return;
    int index = (n - 1) & hash;
    TreeNode<K,V> first = (TreeNode<K,V>)tab[index], root = first, rl;
    TreeNode<K,V> succ = (TreeNode<K,V>)next, pred = prev;
    if (pred == null)
        tab[index] = first = succ;
    else
        pred.next = succ;
    if (succ != null)
        succ.prev = pred;
    if (first == null)
        return;
    if (root.parent != null)
        root = root.root();
    if (root == null || root.right == null ||
        (rl = root.left) == null || rl.left == null) {
        tab[index] = first.untreeify(map);  // too small
        return;
    }
    TreeNode<K,V> p = this, pl = left, pr = right, replacement;
    if (pl != null && pr != null) {
        TreeNode<K,V> s = pr, sl;
        while ((sl = s.left) != null) // find successor
            s = sl;
        boolean c = s.red; s.red = p.red; p.red = c; // swap colors
        TreeNode<K,V> sr = s.right;
        TreeNode<K,V> pp = p.parent;
        if (s == pr) { // p was s's direct parent
            p.parent = s;
            s.right = p;
        }
        else {
            TreeNode<K,V> sp = s.parent;
            if ((p.parent = sp) != null) {
                if (s == sp.left)
                    sp.left = p;
                else
                    sp.right = p;
            }
            if ((s.right = pr) != null)
                pr.parent = s;
        }
        p.left = null;
        if ((p.right = sr) != null)
            sr.parent = p;
        if ((s.left = pl) != null)
            pl.parent = s;
        if ((s.parent = pp) == null)
            root = s;
        else if (p == pp.left)
            pp.left = s;
        else
            pp.right = s;
        if (sr != null)
            replacement = sr;
        else
            replacement = p;
    }
    else if (pl != null)
        replacement = pl;
    else if (pr != null)
        replacement = pr;
    else
        replacement = p;
    if (replacement != p) {
        TreeNode<K,V> pp = replacement.parent = p.parent;
        if (pp == null)
            root = replacement;
        else if (p == pp.left)
            pp.left = replacement;
        else
            pp.right = replacement;
        p.left = p.right = p.parent = null;
    }

    TreeNode<K,V> r = p.red ? root : balanceDeletion(root, replacement);

    if (replacement == p) {  // detach
        TreeNode<K,V> pp = p.parent;
        p.parent = null;
        if (pp != null) {
            if (p == pp.left)
                pp.left = null;
            else if (p == pp.right)
                pp.right = null;
        }
    }
    if (movable)
        moveRootToFront(tab, r);
}
```



``` java
/* ------------------------------------------------------------ */
// Tree bins

/**
  * Entry for Tree bins. Extends LinkedHashMap.Entry (which in turn
  * extends Node) so can be used as extension of either regular or
  * linked node.
  */
static final class TreeNode<K,V> extends LinkedHashMap.Entry<K,V> {
    TreeNode<K,V> parent;  // red-black tree links
    TreeNode<K,V> left;
    TreeNode<K,V> right;
    TreeNode<K,V> prev;    // needed to unlink next upon deletion
    boolean red;
    TreeNode(int hash, K key, V val, Node<K,V> next) {
        super(hash, key, val, next);
    }

    /**
      * Returns root of tree containing this node.
      */
    final TreeNode<K,V> root() {
        for (TreeNode<K,V> r = this, p;;) {
            if ((p = r.parent) == null)
                return r;
            r = p;
        }
    }

    /**
      * Ensures that the given root is the first node of its bin.
      */
    static <K,V> void moveRootToFront(Node<K,V>[] tab, TreeNode<K,V> root) {
        int n;
        if (root != null && tab != null && (n = tab.length) > 0) {
            int index = (n - 1) & root.hash;
            TreeNode<K,V> first = (TreeNode<K,V>)tab[index];
            if (root != first) {
                Node<K,V> rn;
                tab[index] = root;
                TreeNode<K,V> rp = root.prev;
                if ((rn = root.next) != null)
                    ((TreeNode<K,V>)rn).prev = rp;
                if (rp != null)
                    rp.next = rn;
                if (first != null)
                    first.prev = root;
                root.next = first;
                root.prev = null;
            }
            assert checkInvariants(root);
        }
    }

    /**
      * Finds the node starting at root p with the given hash and key.
      * The kc argument caches comparableClassFor(key) upon first use
      * comparing keys.
      */
    final TreeNode<K,V> find(int h, Object k, Class<?> kc) {
        TreeNode<K,V> p = this; // 把当前对象赋给p，表示当前节点
        do { // 循环
            int ph, dir; K pk; // 定义当前节点的hash值、方向（左右）、当前节点的键对象
            TreeNode<K,V> pl = p.left, pr = p.right, q; // 获取当前节点的左孩子、右孩子。定义一个对象q用来存储并返回找到的对象
            if ((ph = p.hash) > h) // 如果当前节点的hash值大于k得hash值h，那么后续就应该让k和左孩子节点进行下一轮比较
                p = pl; // p指向左孩子，紧接着就是下一轮循环了
            else if (ph < h) // 如果当前节点的hash值小于k得hash值h，那么后续就应该让k和右孩子节点进行下一轮比较
                p = pr; // p指向右孩子，紧接着就是下一轮循环了
            else if ((pk = p.key) == k || (k != null && k.equals(pk))) // 如果h和当前节点的hash值相同，并且当前节点的键对象pk和k相等（地址相同或者equals相同）
                return p; // 返回当前节点
    
    
            // 执行到这里说明 hash比对相同，但是pk和k不相等
            else if (pl == null) // 如果左孩子为空
                p = pr; // p指向右孩子，紧接着就是下一轮循环了
            else if (pr == null)
                p = pl; // p指向左孩子，紧接着就是下一轮循环了
    
            // 如果左右孩子都不为空，那么需要再进行一轮对比来确定到底该往哪个方向去深入对比
            // 这一轮的对比主要是想通过comparable方法来比较pk和k的大小     
            else if ((kc != null || (kc = comparableClassFor(k)) != null) && (dir = compareComparables(kc, k, pk)) != 0)
                p = (dir < 0) ? pl : pr; // dir小于0，p指向右孩子，否则指向右孩子。紧接着就是下一轮循环了
    
            // 执行到这里说明无法通过comparable比较  或者 比较之后还是相等
            // 从右孩子节点递归循环查找，如果找到了匹配的则返回    
            else if ((q = pr.find(h, k, kc)) != null) 
                return q;
            else // 如果从右孩子节点递归查找后仍未找到，那么从左孩子节点进行下一轮循环
                p = pl;
        } while (p != null); 
        return null; // 为找到匹配的节点返回null
    }

    /**
      * Calls find for root node.
      */
    final TreeNode<K,V> getTreeNode(int h, Object k) {
        return ((parent != null) ? root() : this).find(h, k, null);
    }

    /**
      * Tie-breaking utility for ordering insertions when equal
      * hashCodes and non-comparable. We don't require a total
      * order, just a consistent insertion rule to maintain
      * equivalence across rebalancings. Tie-breaking further than
      * necessary simplifies testing a bit.
      */
    static int tieBreakOrder(Object a, Object b) {
        int d;
        if (a == null || b == null ||
            (d = a.getClass().getName().
              compareTo(b.getClass().getName())) == 0)
            d = (System.identityHashCode(a) <= System.identityHashCode(b) ?
                  -1 : 1);
        return d;
    }

    /**
      * Forms tree of the nodes linked from this node.
      * @return root of tree
      */
    final void treeify(Node<K,V>[] tab) {
        TreeNode<K,V> root = null;
        for (TreeNode<K,V> x = this, next; x != null; x = next) {
            next = (TreeNode<K,V>)x.next;
            x.left = x.right = null;
            if (root == null) {
                x.parent = null;
                x.red = false;
                root = x;
            }
            else {
                K k = x.key;
                int h = x.hash;
                Class<?> kc = null;
                for (TreeNode<K,V> p = root;;) {
                    int dir, ph;
                    K pk = p.key;
                    if ((ph = p.hash) > h)
                        dir = -1;
                    else if (ph < h)
                        dir = 1;
                    else if ((kc == null &&
                              (kc = comparableClassFor(k)) == null) ||
                              (dir = compareComparables(kc, k, pk)) == 0)
                        dir = tieBreakOrder(k, pk);

                    TreeNode<K,V> xp = p;
                    if ((p = (dir <= 0) ? p.left : p.right) == null) {
                        x.parent = xp;
                        if (dir <= 0)
                            xp.left = x;
                        else
                            xp.right = x;
                        root = balanceInsertion(root, x);
                        break;
                    }
                }
            }
        }
        moveRootToFront(tab, root);
    }

    /**
      * Returns a list of non-TreeNodes replacing those linked from
      * this node.
      */
    final Node<K,V> untreeify(HashMap<K,V> map) {
        Node<K,V> hd = null, tl = null;
        for (Node<K,V> q = this; q != null; q = q.next) {
            Node<K,V> p = map.replacementNode(q, null);
            if (tl == null)
                hd = p;
            else
                tl.next = p;
            tl = p;
        }
        return hd;
    }

     

    /**
      * Removes the given node, that must be present before this call.
      * This is messier than typical red-black deletion code because we
      * cannot swap the contents of an interior node with a leaf
      * successor that is pinned by "next" pointers that are accessible
      * independently during traversal. So instead we swap the tree
      * linkages. If the current tree appears to have too few nodes,
      * the bin is converted back to a plain bin. (The test triggers
      * somewhere between 2 and 6 nodes, depending on tree structure).
      */
    final void removeTreeNode(HashMap<K,V> map, Node<K,V>[] tab,
                              boolean movable) {
        int n;
        if (tab == null || (n = tab.length) == 0)
            return;
        int index = (n - 1) & hash;
        TreeNode<K,V> first = (TreeNode<K,V>)tab[index], root = first, rl;
        TreeNode<K,V> succ = (TreeNode<K,V>)next, pred = prev;
        if (pred == null)
            tab[index] = first = succ;
        else
            pred.next = succ;
        if (succ != null)
            succ.prev = pred;
        if (first == null)
            return;
        if (root.parent != null)
            root = root.root();
        if (root == null || root.right == null ||
            (rl = root.left) == null || rl.left == null) {
            tab[index] = first.untreeify(map);  // too small
            return;
        }
        TreeNode<K,V> p = this, pl = left, pr = right, replacement;
        if (pl != null && pr != null) {
            TreeNode<K,V> s = pr, sl;
            while ((sl = s.left) != null) // find successor
                s = sl;
            boolean c = s.red; s.red = p.red; p.red = c; // swap colors
            TreeNode<K,V> sr = s.right;
            TreeNode<K,V> pp = p.parent;
            if (s == pr) { // p was s's direct parent
                p.parent = s;
                s.right = p;
            }
            else {
                TreeNode<K,V> sp = s.parent;
                if ((p.parent = sp) != null) {
                    if (s == sp.left)
                        sp.left = p;
                    else
                        sp.right = p;
                }
                if ((s.right = pr) != null)
                    pr.parent = s;
            }
            p.left = null;
            if ((p.right = sr) != null)
                sr.parent = p;
            if ((s.left = pl) != null)
                pl.parent = s;
            if ((s.parent = pp) == null)
                root = s;
            else if (p == pp.left)
                pp.left = s;
            else
                pp.right = s;
            if (sr != null)
                replacement = sr;
            else
                replacement = p;
        }
        else if (pl != null)
            replacement = pl;
        else if (pr != null)
            replacement = pr;
        else
            replacement = p;
        if (replacement != p) {
            TreeNode<K,V> pp = replacement.parent = p.parent;
            if (pp == null)
                root = replacement;
            else if (p == pp.left)
                pp.left = replacement;
            else
                pp.right = replacement;
            p.left = p.right = p.parent = null;
        }

        TreeNode<K,V> r = p.red ? root : balanceDeletion(root, replacement);

        if (replacement == p) {  // detach
            TreeNode<K,V> pp = p.parent;
            p.parent = null;
            if (pp != null) {
                if (p == pp.left)
                    pp.left = null;
                else if (p == pp.right)
                    pp.right = null;
            }
        }
        if (movable)
            moveRootToFront(tab, r);
    }

    /**
      * Splits nodes in a tree bin into lower and upper tree bins,
      * or untreeifies if now too small. Called only from resize;
      * see above discussion about split bits and indices.
      *
      * @param map the map
      * @param tab the table for recording bin heads
      * @param index the index of the table being split
      * @param bit the bit of hash to split on
      */
    final void split(HashMap<K,V> map, Node<K,V>[] tab, int index, int bit) {
        TreeNode<K,V> b = this;
        // Relink into lo and hi lists, preserving order
        TreeNode<K,V> loHead = null, loTail = null;
        TreeNode<K,V> hiHead = null, hiTail = null;
        int lc = 0, hc = 0;
        for (TreeNode<K,V> e = b, next; e != null; e = next) {
            next = (TreeNode<K,V>)e.next;
            e.next = null;
            if ((e.hash & bit) == 0) {
                if ((e.prev = loTail) == null)
                    loHead = e;
                else
                    loTail.next = e;
                loTail = e;
                ++lc;
            }
            else {
                if ((e.prev = hiTail) == null)
                    hiHead = e;
                else
                    hiTail.next = e;
                hiTail = e;
                ++hc;
            }
        }

        if (loHead != null) {
            if (lc <= UNTREEIFY_THRESHOLD)
                tab[index] = loHead.untreeify(map);
            else {
                tab[index] = loHead;
                if (hiHead != null) // (else is already treeified)
                    loHead.treeify(tab);
            }
        }
        if (hiHead != null) {
            if (hc <= UNTREEIFY_THRESHOLD)
                tab[index + bit] = hiHead.untreeify(map);
            else {
                tab[index + bit] = hiHead;
                if (loHead != null)
                    hiHead.treeify(tab);
            }
        }
    }

    /* ------------------------------------------------------------ */
    // Red-black tree methods, all adapted from CLR

    static <K,V> TreeNode<K,V> rotateLeft(TreeNode<K,V> root,
                                          TreeNode<K,V> p) {
        TreeNode<K,V> r, pp, rl;
        if (p != null && (r = p.right) != null) {
            if ((rl = p.right = r.left) != null)
                rl.parent = p;
            if ((pp = r.parent = p.parent) == null)
                (root = r).red = false;
            else if (pp.left == p)
                pp.left = r;
            else
                pp.right = r;
            r.left = p;
            p.parent = r;
        }
        return root;
    }

    static <K,V> TreeNode<K,V> rotateRight(TreeNode<K,V> root,
                                            TreeNode<K,V> p) {
        TreeNode<K,V> l, pp, lr;
        if (p != null && (l = p.left) != null) {
            if ((lr = p.left = l.right) != null)
                lr.parent = p;
            if ((pp = l.parent = p.parent) == null)
                (root = l).red = false;
            else if (pp.right == p)
                pp.right = l;
            else
                pp.left = l;
            l.right = p;
            p.parent = l;
        }
        return root;
    }
 

    static <K,V> TreeNode<K,V> balanceDeletion(TreeNode<K,V> root,
                                                TreeNode<K,V> x) {
        for (TreeNode<K,V> xp, xpl, xpr;;)  {
            if (x == null || x == root)
                return root;
            else if ((xp = x.parent) == null) {
                x.red = false;
                return x;
            }
            else if (x.red) {
                x.red = false;
                return root;
            }
            else if ((xpl = xp.left) == x) {
                if ((xpr = xp.right) != null && xpr.red) {
                    xpr.red = false;
                    xp.red = true;
                    root = rotateLeft(root, xp);
                    xpr = (xp = x.parent) == null ? null : xp.right;
                }
                if (xpr == null)
                    x = xp;
                else {
                    TreeNode<K,V> sl = xpr.left, sr = xpr.right;
                    if ((sr == null || !sr.red) &&
                        (sl == null || !sl.red)) {
                        xpr.red = true;
                        x = xp;
                    }
                    else {
                        if (sr == null || !sr.red) {
                            if (sl != null)
                                sl.red = false;
                            xpr.red = true;
                            root = rotateRight(root, xpr);
                            xpr = (xp = x.parent) == null ?
                                null : xp.right;
                        }
                        if (xpr != null) {
                            xpr.red = (xp == null) ? false : xp.red;
                            if ((sr = xpr.right) != null)
                                sr.red = false;
                        }
                        if (xp != null) {
                            xp.red = false;
                            root = rotateLeft(root, xp);
                        }
                        x = root;
                    }
                }
            }
            else { // symmetric
                if (xpl != null && xpl.red) {
                    xpl.red = false;
                    xp.red = true;
                    root = rotateRight(root, xp);
                    xpl = (xp = x.parent) == null ? null : xp.left;
                }
                if (xpl == null)
                    x = xp;
                else {
                    TreeNode<K,V> sl = xpl.left, sr = xpl.right;
                    if ((sl == null || !sl.red) &&
                        (sr == null || !sr.red)) {
                        xpl.red = true;
                        x = xp;
                    }
                    else {
                        if (sl == null || !sl.red) {
                            if (sr != null)
                                sr.red = false;
                            xpl.red = true;
                            root = rotateLeft(root, xpl);
                            xpl = (xp = x.parent) == null ?
                                null : xp.left;
                        }
                        if (xpl != null) {
                            xpl.red = (xp == null) ? false : xp.red;
                            if ((sl = xpl.left) != null)
                                sl.red = false;
                        }
                        if (xp != null) {
                            xp.red = false;
                            root = rotateRight(root, xp);
                        }
                        x = root;
                    }
                }
            }
        }
    }

    /**
      * Recursive invariant check
      */
    static <K,V> boolean checkInvariants(TreeNode<K,V> t) {
        TreeNode<K,V> tp = t.parent, tl = t.left, tr = t.right,
            tb = t.prev, tn = (TreeNode<K,V>)t.next;
        if (tb != null && tb.next != t)
            return false;
        if (tn != null && tn.prev != t)
            return false;
        if (tp != null && t != tp.left && t != tp.right)
            return false;
        if (tl != null && (tl.parent != t || tl.hash > t.hash))
            return false;
        if (tr != null && (tr.parent != t || tr.hash < t.hash))
            return false;
        if (t.red && tl != null && tl.red && tr != null && tr.red)
            return false;
        if (tl != null && !checkInvariants(tl))
            return false;
        if (tr != null && !checkInvariants(tr))
            return false;
        return true;
    }
}
```