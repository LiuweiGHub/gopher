package main

/**
23种设计模式

GoF <<设计模式>>
*/

///////////////////////////////////////创建型设计模式家族 1-5//////////////////////////////////////////////////
//总结成一句话就是，当软件系统中需要对象的创建、组合或聚集时，你就可以考虑使用创建型模式“家族”来帮助你提升代码的灵活性。

/**
1.单例模式

强制实现了有限、唯一对象的生产。
*/

/**
2.建造者模式（Builder模式）

将复杂对象的构造与其表示分离，以便同一构造过程可以创建不同的表示

  -- Product
  -- Builder  -->  ConcreteBuilder
  -- Director

使用场景：
1、需要生成的对象包含多个成员属性
2、需要生成的对象的属性相互依赖，需要指定其生成顺序
3、对象的创建过程独立于创建该对象的类
4、需要隔离复杂对象的创建和使用，并使得相同的创建过程可以创建不同的产品

目的： 分离创建与使用
个人理解：比如创建一个苹果手机，那么很多部件的生产都是委托给第三方的，期构建过程是不可见的，最终苹果公司只需拿来用于手机组装
经典应用：Locallog builder

实现了对象创建过程的多态
*/

/**
3.抽象工厂模式
重点和难点在于：如何找到正确的抽象   如何找到某一个类产品的正确共性功能
抽象工厂用来继承
目的： 升组件的复用性、提升代码的扩展性、解决跨平台带来的兼容性问题
工厂模式：
 	1、简单工厂
    2、工厂方法
    3、抽象工厂

保证被创建对象的工厂之间的一致性，常常用来解决跨平台的设计问题
*/

/**
4.工厂方法模式

目的很简单，就是封装对象创建的过程，提升创建对象方法的可复用性
工厂方法模式将对象的实例化操作延迟到了具体产品子类中去完成。

interface
实现类
Factory

- 把对象的创建和使用过程分开，降低代码耦合性
- 减少重复代码
- 统一管理创建对象的不同实现逻辑

push的各种client可以使用

实现了对象创建时的多态
*/

/**
5、原型模式： 使用原型实例指定创建对象的种类，然后通过拷贝这些原型来创建新的对象  （在内存中进行二进制流的拷贝）

原型模式的适用场景通常都是对已有的复杂对象或大型对象的创建上，在这样的场景中，创建对象通常是一件非常烦琐的事情，通过拷贝对象能快速地创建对象。

浅拷贝vs深拷贝
对象中注入了别的对象，就需要深拷贝
- 减少每次创建对象的资源消耗
- 降低对象创建的时间消耗
- 快速复制对象运行时状态
- 能保存原始对象的副本

实现了对象拷贝的多态。
*/

////////////////////////////////6-11  结构型模式 //////////////////////////////////////
/**
6、适配器模式： 将类的接口转换为客户期望的另一个接口，适配器可以让不兼容的两个类一起协同工作。

使用场景： 第一类就是原有接口功能不满足现有要求，需要在兼容老接口的同时做适当的扩展；第二类是有相似性的多个不同接口之间做功能的统一。

- 第一，原有接口无法修改但又必须快速兼容部分新功能。
- 第二，需要使用外部组件组合成新组件来提供功能，而又不想重复开发部分功能
- 第三，不同数据格式、不同协议需要转换


*/

/**
7、桥接模式： 将抽象部分与实现部分分离，使它们都可以独立变化


*/

/**
8、组合模式：将对象组合成树形结构以表示整个部分的层次结构
使用场景：
- 处理一个树形结构，比如，公司人员组织架构、订单信息等
- 跨越多个层次结构聚合数据，比如，统计文件夹下文件总数
- 统一处理一个结构中的多个对象，比如，遍历文件夹下所有 XML 类型文件内容
*/

/**
9、装饰模式：允许动态地向一个现有的对象添加新的功能，同时又不改变其结构，相当于对现有的对象进行了一个包装。所以装饰模式又叫包装器模式。

适配器模式侧重于转换，而装饰模式侧重于动态扩展；桥接模式侧重于横向宽度的扩展，而装饰模式侧重于纵向深度的扩展

装饰模式本质上就是给已有不可修改的类附加新的功能，同时还能很方便地撤销

一般情况下不建议装饰器超过 10 个

对于没有上下逻辑的装饰器，也要尽量避免使用装饰模式
*/

/**
9、门面模式： 为子系统中的一组接口提供统一的接口。它定义了一个更高级别的接口，使子系统更易于使用。

应用非常广泛

当我们需要用更统一的标准方式来与系统交互时，就可以采用门面模式。比如，使用 Slf4j 日志框架来统一 log4j、log4j2、CommonLog 等日志框架。
再比如，在支付时通过扫描二维码来使用支付系统。对于用户来说，他们并不关心后台系统实现有多么复杂，只关心最终能否支付成功。

门面模式可能代理的是多个接口，而代理模式通常只是代理某一个接口

我们平时最常见的电脑开机按钮就是一个门面模式，点击按钮电脑就会启动，再点击按钮电脑就会关闭，至于电脑如何运行 CPU、启动内存、读取硬盘、点亮显示器，我们其实并不关心。我们只关心使用视角下的电脑，而不关心电脑本身是如何运行的。

本质就是简化外部系统使用内部多个子系统的使用方式

mobileapi 就是一个门面吧，对客户端简化操作，更多工作量给到服务端
门面模式最典型的一种应用就是 API 网关


降低了可靠性
容易导致子系统越来越复杂
*/

/**
10、享元模式：摒弃了在每个对象中保存所有数据的方式，通过共享多个对象所共有的相同状态，从而让我们能在有限的内存容量中载入更多对象
应用场景却相对狭窄

你会发现，享元模式本质上在使用时就是找到不可变的特征，并缓存起来，当类似对象使用时从缓存中读取，以达到节省内存空间的目的。
*/

/**
11、代理模式：让你能够提供对象的替代品或其占位符。代理控制着对于原对象的访问，并允许将请求提交给对象前后进行一些处理。


代理模式的应用却比装饰模式更为广泛 在很多框架和组件的设计里都能看到代理模式的身影，比如，JDK 的动态代理机制、Spring 的 AOP 机制、Dubbo 框架

PHP hdfRedis  __call 适配加中间件日志

应用却非常广泛，特别是在中间件领域随处可见代理模式，比如 Dubbo、gRPC 都是代理模式的优秀实践。
*/

///////////////////////////////////12-23行为型设计模式//////////////////////////////////////////////////////////////
/////行为型设计模式主要的关注点是对象内部算法及对象之间的职责和分配，比如，具体实现算法、选择策略、状态变化等抽象概念

//类行为型模式：使用继承的方式来关联不同类之间的行为。

//对象行为型模式：使用组合或聚合方式来分配不同类之间的行为。

/**
12、访问者模式: 允许在运行时将一个或多个操作应用于一组对象，将操作与对象结构分离。

访问者模式是以行为（某一个操作）作为扩展对象功能的出发点，在不改变已有类的功能的前提下进行批量扩展。

使用场景：
- 当对象的数据结构相对稳定，而操作却经常变化的时候
- 需要将数据结构与不常用的操作进行分离的时候
- 需要在运行时动态决定使用哪些对象和方法的时候

访问者模式在真实的应用中算得上是比较难
*/

/**
13、模板方法模式：在操作中定义算法的框架，将一些步骤推迟到子类中。模板方法让子类在不改变算法结构的情况下重新定义算法的某些步骤。
软件开发中也被广泛应用，但是因为使用继承机制，副作用往往盖过了主要作用，所以在使用时尤其要小心谨慎

- 多个类有相同的方法并且逻辑可以共用时
- 将通用的算法或固定流程设计为模板，在每一个具体的子类中再继续优化算法步骤或流程步骤时
- 重构超长代码时，发现某一个经常使用的公有方法


通常是对算法的特定步骤进行优化，而不是对整个算法进行修改
模板方法模式是一个比较常用的结构型设计模式，它能帮助我们快速构建一个通用模板，对于固定流程、通用协议而言，模板方法模式都是一个很好的选择。
*/

/**
14、策略模式
*/

/**
15、状态模式
*/

/**
17、观察者模式：定义对象之间的一对多依赖关系，这样当一个对象改变状态时，它的所有依赖项都会自动得到通知和更新。

非常流行的设计模式，也常被叫作订阅-发布模式

被依赖的对象：被观察者
依赖的对象：观察者

A依赖B，A就是观察者，B就是被观察者
CI 就是吧，监控到提交release分支代码 就自动构建镜像。
*/

/**
18、备忘录模式：捕获并外部化对象的内部状态，以便以后可以恢复，所有这些都不会违反封装。
不常用
应用场景就更是比较明确和有限，一般应用于编辑器或会话上下文中防丢失、撤销、恢复等场景中。

- 需要保存一个对象在某一个时刻的状态时
- 不希望外界直接访问对象的内部状态时

也叫快照模式，具体来说，就是通过捕获对象在某一个时刻的对象状态，再将其保存到外部对象，以便在需要的时候恢复对象到指定时刻下的状态。
*/

/**
19、中介者模式：中介者对象封装了一组对象之间的交互，这组对象会将它们的交互委托给中介者对象，而不是直接交互。
使用场景：
- 系统中对象之间存在复杂的引用关系时，比如，聊天系统
- 通过一个中间对象来封装多个类中的共有行为时，比如，在分层架构中的 DAO 层和数据库 DB 层中间再引入一个读写分离和读写均衡的中间层；
- 不想生成太多的子类时

Dubbo 一类的 RPC 框架就是一个完整的中介者模式的体现。对于所有的 Java RPC 调用来说，只需要通过这个中间层来进行通信即可

视图层一般不会直接使用 DAO 层，因为一旦直接使用，DAO 层任何一个微小的变动都可能引起视图层的变化，这时通常会引入 Service 层作为中介者来进行请求的转发，以达到解耦的目的，

Maven 就是 Java 中引用 jar 包时的中介者，如果我们手动直接引用 jar 包，会容易造成非常混乱的引用关系，而使用 Maven 则能很方便地减少代码直接依赖 jar 包的问题。
*/

/**
20、迭代器模式：迭代器提供一种对容器对象中的各个元素进行访问的方法，而又不需要暴露该对象的内部细节。
很少用

又叫游标模式

比如 Redis 中的 rehash() 操作，就是迭代器模式的体现，而且 Redis 更进一步地还区分了安全迭代器和非安全迭代器
*/

/**
21、解释器模式:用于定义语言的语法规则表示，并提供解释器来处理句子中的语法。

使用频率不算高, 特定的领域被用到，比如编译器、规则引擎、正则表达式、SQL 解析等
*/

/**
22、命令模式：
使用频率不算太高
*/

/**
23、责任链模式：通过为多个对象提供处理请求的机会，避免将请求的发送者与其接收者耦合。链接接收对象并沿着链传递请求，直到对象处理它。
使用频率很高的模式

Handler
- 在运行时需要动态使用多个关联对象来处理同一次请求时。比如，请假流程、员工入职流程、编译打包发布上线流程等
- 不想让使用者知道具体的处理逻辑时。比如，做权限校验的登录拦截器
- 需要动态更换处理对象时。比如，工单处理系统、网关 API 过滤规则系统等

责任链模式就像工厂的流水线作业一样，按照某一个标准化的流程来执行，用于规则过滤、Web 请求协议解析等具备链条式的场景中，通过拆分不同的处理节点来完成整个流程的处理

，可以说只要是与流程相关的软件系统都能够使用责任链模式来构建，一方面可以用在代码中实现松散耦合，另一方面可以动态增删子处理流程
*/
