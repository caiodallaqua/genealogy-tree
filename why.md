# Gen Tree Why File

**Why to generate code in `deserializers.go`?**

As I once read, "Golang is boring and that's awesome!". The deserializers' code are pretty much the same so we could:
- Write everything by hand;
- Try to use generics;
- Use meta programming (which ended up being our choice).

The first is highly error-prone. The second is elegant, but adds complexity where we don't need it. By using go templates we were able to keep things far less error-prone, readable and scalable - we can go ahead and add as many deserializers as we want, there will be no need to write tons of code and also no performance overhead. 