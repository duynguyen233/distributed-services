# Chapter 2: Structure Data with Protocol Buffers

**What is protocol buffer?**
Protocol buffers (also known as *protobuf*), which is Google's language and platform-neutral extensible mechanism for structuring and serializing data 

**Advantages:**
- Guarantees type-safety
- Prevents schema-violations;
- Enables fast serialization
- Offers backward compatibility

## Why use Protocol Buffers?
*Consistent schemas*: encode semantics once and use them across services.
*Versioning for free*: with protobuf, you number your fields on your messages to ensure you maintain backward compatibility as you roll out new features and changes to your protobuf. Easy to add new fields -> intermediate servers not use can simply parse and pass. With removing fields: compiler will conplain when anyone tries to use the deprecated field
*Less boilerplate*: libraries encode and decode
*Extensibility*: compiler supports extensions that can compile your protobuf into code using your own compilation logic.
*Language agnosticism*: multiple languages support
*Performance*: smaller payloads and serializes 6 times faster than JSON

gRPC use protocol buffers to define APIs and serialize messages
