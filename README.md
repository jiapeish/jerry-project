# jerry-project

## 怎么运行？

1. 直接运行
```shell
go run .
```

2. 先构建再运行
```shell
go build -o jerry-project .

./jerry-project
```

3. 运行测试
```shell
go test .
```

## 代码结构

segment.go 3个对外提供的方法
type.go 数据结构定义
utils.go 内部使用的方法


## 下一步优化的地方
1. 不支持并发调用，后续可以考虑加锁；
2. 支持更多数据类型
3. 




## Intensity Segments

### Guidelines
In this part of the interview process, we’d like you to come up with an algorithm to solve the problem as
described below. The problem itself is quite simple to solve. What we are mainly looking for in this test
(other than that the solution should work) is, how well you actually write the code. We want to see how you
**write production-quality code in a team setting** where multiple developers will be collaborating on the
codebase.

Specifically, we are looking for: simple, clean, readable and maintainable code, for example:
- Code organization and submission format. Things like code organization, readability, documentation,
testing and deliverability are most important here.
- Your mastery of idiomatic programming.

The solution is prefered to be in JavaScript. We understand that you may not have much experience with JS.
We encourage you to take some time to research modern JS and best practices, and try your best to apply
them when writing your test solution.

If you choose to use a programming language other than JS, please still make sure you stick to the idiomatic
way of that programming language.

Never use AI to write any part of your code.

### Problem Set
We are looking for a program that manages “intensity” by segments. Segments are intervals from -infinity to
infinity, we’d like you to implement functions that updates intensity by an integer amount for a given range.
All intensity starts with 0. Please implement these three functions:

```javascript
export class IntensitySegments {
    add(from, to, amount) {
        // TODO: implement this
    }
    set(from, to, amount) {
        // TODO: implement this
    }
    toString() {
        // TODO: implement this
    }
}

// Here is an example sequence:
// (data stored as an array of start point and value for each segment.)
const segments = new IntensitySegments();
segments.toString(); // Should be "[]"

segments.add(10, 30, 1);
segments.toString(); // Should be: "[[10,1],[30,0]]"

segments.add(20, 40, 1);
segments.toString(); // Should be: "[[10,1],[20,2],[30,1],[40,0]]"

segments.add(10, 40, -2);
segments.toString(); // Should be: "[[10,-1],[20,0],[30,-1],[40,0]]"

// Another example sequence:
const segments = new IntensitySegments();
segments.toString(); // Should be "[]"

segments.add(10, 30, 1);
segments.toString(); // Should be "[[10,1],[30,0]]"

segments.add(20, 40, 1);
segments.toString(); // Should be "[[10,1],[20,2],[30,1],[40,0]]"

segments.add(10, 40, -1);
segments.toString(); // Should be "[[20,1],[30,0]]"

segments.add(10, 40, -1);
segments.toString(); // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"
```