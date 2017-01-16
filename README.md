# BaiyeJVM

GO语言实现的JVM;

基于《自己动手写Java虚拟机》，原作者发布了完整的项目在GitHub上，可以搜索jvm.go看到;

原作者还缺少一些指令没有实现，代码中有一些Hack的处理，准备补全这些地方;

TODO:
实现Class文件的验证(4.10);
MethodType MethodHandle 实现;
接口的静态方法和默认方法;
invokedynamic?;
invokestatic?;
ACC_SUPER-FLAG 特殊处理;
vtable为invokevirtual增加优化;
多个Catch Block的支持;
嵌套的Try-Catch Block的支持;
Finally的支持;
FileOutputStream处理文件;
