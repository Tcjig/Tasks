# JS学习

## 简单了解

+ **HTML**定义了网页的内容
+ **CSS**描述了网页的布局
+ **JS**控制了网页的行为



## 基本用法

### \<script>标签(双标签)

作用：在HTML网页中插入JS

注意：

+ HTML 中的 Javascript 脚本代码必须位于 ``<script>`` 与 ``</script>`` 标签之间
+ avascript 脚本代码可被放置在 HTML 页面的 ``<body>`` 和 ``<head>`` 中

### 外部调用\<script>标签

标签名：``<script src="JS的文件名"></script>``

**提示：**外部的**JS**文件不能包含``<script>``标签，直接写**JS**的代码即可



## 基本语法

### JS输出

JS可以通过以下方式来输出数据：

+ 使用 ``window.alert()`` 弹出警告框
+ 使用 ``document.write()`` 方法将内容写到HTML文档中
+ 使用 ``innerHTML`` 写入到HTML元素
+ 使用 ``console.log()`` 写入到浏览器的控制台

### 操作HTML元素

**注意点：**

+ 若需从JS中访问某个HTML元素怒，可以使用 ``document.getElementByld(id)`` 方法

+ 使用 ``id`` 属性来标识HTML元素，并在 ``innerHTML`` 来获取或插入元素内容

示例：

![image-20240331183045480](C:\Users\86187\AppData\Roaming\Typora\typora-user-images\image-20240331183045480.png)