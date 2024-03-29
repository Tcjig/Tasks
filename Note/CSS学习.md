# CSS学习

## 基础知识

定义：CSS主要是用来美化内容

书写位置：title标签下方添加**style**双标签，**style**标签里面书写**CSS代码**

CSS引入方式：

![image-20240324203024508](C:\Users\86187\AppData\Roaming\Typora\typora-user-images\image-20240324203024508.png)

## 基础选择器

**通用写法：**``选择器 {属性名:属性值;}``

### 标签选择器

作用：使用**标签名**作为选择器，并且将该标签及同名标签的所有内容都设置成**相同**的格式

### 类选择器

作用：**差异化**设置标签的显示效果

步骤：

+ 定义类选择器 -> ``.类名``
+ 使用类选择器 -> 标签添加``class="类名"``
+ **示例**![image-20240324203905667](C:\Users\86187\AppData\Roaming\Typora\typora-user-images\image-20240324203905667.png)

**提示：**如果定义了多个选择器，并想让一个标签都将其使用，比如分别定义了**apple**和**orange**选择器，则可以按照``class="apple orange"``这样进行使用

### id选择器(一般与JS配合使用)

作用：**差异化**设置标签的显示效果

步骤：

+ 定义id选择器 -> ``#id``
+ 使用id选择器 -> 标签添加``id="id名"``

+ **示例**

  ![image-20240324205103862](C:\Users\86187\AppData\Roaming\Typora\typora-user-images\image-20240324205103862.png)

**提示：**同一个id选择器在一个页面只能使用一次

### 通配符选择器

作用：查找页面所有标签，设置相同格式

格式：``* {属性名:属性值;}``

**提示：**不需要调用，浏览器会自动查找页面所有标签，设置相同格式