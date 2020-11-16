# Comprehensive tests of protocol message processing.
# See example.star for a simpler example.

load("assert.star", "assert")

# Load a protocol message descriptor from some file.
# The host application determines how this name resolves to a FileDescriptorProto.
test = proto.file("testdata/test/test.proto")

def to_dict(msg):
    """to_dict returns a dictionary containing the defined fields of a message.
    Messages in the result are not recursively converted."""
    dict = {}
    for key in dir(msg):
        if key != "descriptor" and proto.has(msg, key):
            dict[key] = getattr(msg, key)
    return dict

# -------- Descriptors --------

# FileDescriptor
assert.eq(type(test), "proto.FileDescriptor")
assert.eq(str(test), "testdata/test/test.proto")
assert.eq(dir(test), ["Coin", "D", "KeyValuePair", "M", "Outer", "Suit", "ext"])

def set_filedesc():
    test.x = 1

assert.fails(set_filedesc, "can't assign to .x field of proto.FileDescriptor")
assert.fails(lambda: test.Nonesuch, "no .Nonesuch field")

# equality
assert.eq(test, test)  # file
assert.ne(test, test.M)

# MessageDescriptor
assert.eq(type(test.M), "proto.MessageDescriptor")
assert.eq(str(test.M), "test.M")
assert.eq(dir(test.M), ["Inner"])

def set_desc():
    test.M.x = 1

assert.fails(set_desc, "can't assign to .x field of proto.MessageDescriptor")
assert.fails(lambda: test.M.Nonesuch, "no .Nonesuch field")

# equality
assert.eq(test.M, test.M)
assert.ne(test.M, test.Outer)

# nested MessageDescriptor
assert.eq(type(test.M.Inner), "proto.MessageDescriptor")
assert.eq(str(test.M.Inner), "test.M.Inner")
assert.eq(dir(test.M.Inner), [])

# EnumDescriptor
assert.eq(type(test.Coin), "proto.EnumDescriptor")
assert.eq(str(test.Coin), "test.Coin")
assert.eq(dir(test.Coin), ["DIME", "NICKEL", "PENNY", "QUARTER"])

# EnumDescriptor(x) converts a Starlark value to an enum value.
assert.eq(test.Coin(5), test.Coin.NICKEL)
assert.fails(lambda: test.Coin(4), "invalid number 4 for Coin enum")  # a hole
assert.fails(lambda: test.Coin(123), "invalid number 123 for Coin enum")  # out of range
assert.fails(lambda: test.Coin(True), "cannot convert bool to Coin enum")
assert.eq(test.Coin("NICKEL"), test.Coin.NICKEL)
assert.fails(lambda: test.Coin("wontfix"), 'invalid name "wontfix" for Coin enum')
assert.eq(test.Coin(test.Coin.NICKEL), test.Coin.NICKEL)
assert.fails(lambda: test.Coin(test.Suit.HEARTS), "invalid value Suit.HEARTS for Coin enum")

# equality:
assert.eq(test.Coin, test.Coin)
assert.ne(test.Coin, test.Suit)

# immutability:
def set_enum():
    test.Coin.NICKEL = 1

assert.fails(set_enum, "can't assign to .NICKEL field of proto.EnumDescriptor")
assert.fails(lambda: test.Suit.Nonesuch, "no .Nonesuch field")

# EnumValueDescriptor
assert.eq(type(test.Coin.PENNY), "proto.EnumValueDescriptor")
assert.eq(str(test.Coin.PENNY), "Coin.PENNY")
assert.eq(test.Coin.NICKEL.number, 5)
assert.eq(test.Coin.NICKEL.index, 1)
assert.eq(test.Coin.NICKEL.name, "NICKEL")
assert.eq(test.Coin.NICKEL.type, test.Coin)
assert.eq(dir(test.Coin.PENNY), ["index", "name", "number", "type"])
assert.fails(lambda: test.PENNY, "no .PENNY field")  # enum type is mandatory
assert.eq(test.Coin.PENNY, test.Coin.PENNY)
assert.ne(test.Coin.PENNY, test.Coin.NICKEL)
assert.eq(test.Coin.PENNY.number, test.Suit.SPADES.number)
assert.ne(test.Coin.PENNY, test.Suit.SPADES)

# equality:
assert.eq(test.Coin.NICKEL, test.Coin.NICKEL)
assert.ne(test.Coin.NICKEL, test.Coin.DIME)
assert.lt(test.Coin.NICKEL.number, test.Coin.DIME.number)
assert.fails(
    lambda: test.Coin.NICKEL < test.Coin.DIME,
    "proto.EnumValueDescriptor < proto.EnumValueDescriptor not implemented",
)
assert.ne(test.Coin.NICKEL, test.Suit.HEARTS)

# hashing:
coin_names = {
    test.Coin.PENNY: "penny",
    test.Coin.NICKEL: "nickel",
    test.Coin.DIME: "dime",
    test.Coin.QUARTER: "quarter",
}
assert.eq(coin_names[test.Coin.DIME], "dime")

# -------- Message instantiation --------

# Descriptor(x) instantiates a message.
m1 = test.M(
    int32 = 12345,
    coin = test.Coin.PENNY,
    string = "hello",
    strings = ["hello", "world"],
    kvs = [test.KeyValuePair(key = "greeting", value = "hello")],
)
proto.set_field(m1, test.Outer.ext, 1)
proto.set_field(m1, test.ext, 2)
assert.eq(type(m1), "proto.Message")
assert.eq(m1.descriptor, test.M)

# str:
assert.eq(str(m1), 'test.M(int32=12345, string="hello", coin=Coin.PENNY, strings=["hello", "world"], kvs=[test.KeyValuePair(key="greeting", value="hello")], test.ext=2, test.Outer.ext=1)')

# str elides types of submessages:
assert.eq(str(test.M(int32 = 1, inner = test.M.Inner(int32 = 2))), "test.M(int32=1, inner=test.M.Inner(int32=2))")

# other conversions
assert.fails(lambda: test.M(1), "M: got int, want dict or message")
assert.fails(lambda: test.M(nonesuch = 1), "no .nonesuch field")
assert.fails(lambda: test.M(kvs = 1), "got int.*want iterable")
assert.fails(lambda: test.M(kvs = [123]), "index 0: got int, want test.KeyValuePair")
assert.fails(lambda: test.M(inner = [1, 2, 3]), "in field inner: got list, want test.M.Inner")
assert.fails(lambda: test.M(coin = 123), "invalid number 123 for Coin enum")
assert.fails(lambda: test.M(coin = "nonesuch"), 'invalid name "nonesuch" for Coin enum')
assert.fails(lambda: test.M(coin = test.Suit.HEARTS), "invalid value Suit.HEARTS for Coin enum")

# Submessages may be expressed using dicts:
assert.eq(str(test.M(inner = {"int32": 123})), "test.M(inner=test.M.Inner(int32=123))")
assert.eq(str(test.M(inner = test.M.Inner(int32 = 123))), "test.M(inner=test.M.Inner(int32=123))")
assert.eq(test.M(int32 = 1, inner = test.M.Inner(int32 = 123)).inner.int32, 123)

# When nested dicts are used, the error message should make the field path clear.
# When descriptor calls are used, the errors are shorter since the location does the work.
assert.fails(
    lambda: test.M(inner = dict(inner = dict(x = 1))),
    #             ^-- failing operation
    "in field inner: in field inner: test.M.Inner has no .x field",
)
assert.fails(
    lambda: test.M(inner = test.M.Inner(inner = dict(x = 1))),
    #                                ^-- failing operation
    "in field inner: test.M.Inner has no .x field",
)
assert.fails(
    lambda: test.M(inner = test.M.Inner(inner = test.M.Inner(x = 1))),
    #                                                   ^-- failing operation
    "test.M.Inner has no .x field",
)

# proto.has reports whether a field is set.
assert.true(proto.has(m1, "kvs"))
assert.true(proto.has(m1, "int32"))
assert.true(not proto.has(m1, "bools"))
assert.true(not proto.has(m1, "floats"))

# proto.has allows fields to be specified by descriptor.
assert.true(proto.has(m1, test.M.kvs))
assert.true(proto.has(m1, test.M.int32))
assert.true(not proto.has(m1, test.M.bools))
assert.true(not proto.has(m1, test.M.floats))

# proto.has: descriptor for wrong message
assert.fails(lambda: proto.has(test.D(), test.M.floats), "test.D does not have field test.M.floats")

# With one positional argument, descriptor(x) makes a shallow copy.
orig = test.M(
    string = "hello",
    strings = ["hello", "world"],
    kvs = [test.KeyValuePair(key = "greeting", value = "hello")],
)
new = test.M(orig)
assert.eq(str(new), str(orig))

# scalar fields are copied
new.string = "goodbye"
assert.eq(new.string, "goodbye")
assert.ne(str(new), str(orig))
assert.eq(orig.string, "hello")  # unchanged

# repeated fields are shared
new.strings[0] = "goodbye"
assert.eq(new.strings[0], "goodbye")
assert.eq(orig.strings[0], "goodbye")  # changed

# The argument must be a message of the correct type.
assert.fails(
    lambda: test.M(test.KeyValuePair()),
    "M: got message of type test.KeyValuePair, want type test.M",
)

# A value of None unsets a field.
m6 = test.M(coin = 5)
assert.eq(str(m6), str(test.M(coin = 5)))
m6.coin = None
assert.eq(str(m6), str(test.M()))

# -------- Field access and conversion --------

# A message's dir includes all non-extension fields, even unset ones.
assert.eq(dir(test.M()), ["bool", "bools", "bytes", "byteses", "coin", "coins", "descriptor", "double", "doubles", "float", "floats", "inner", "inners", "int32", "int32s", "int64", "int64s", "kv", "kvs", "outer", "outers", "string", "strings", "uint32", "uint32s", "uint64", "uint64s"])

# Access to an unset field returns its default value;
# see type-specific assertions below.
# But access to an undefined field is an error.
assert.fails(lambda: test.M().nonesuch, "test.M has no .nonesuch field")

# Field conversions for all types.

# bool
assert.eq(test.M().bool, False)
assert.eq(test.M(bool = True).bool, True)
assert.eq(test.M(bool = False).bool, False)
assert.fails(lambda: test.M(bool = 1), "got int, want bool")  # no implicit truth-value conversion

# int32
assert.eq(test.M().int32, 0)
assert.eq(test.M(int32 = 123).int32, 123)
assert.eq(test.M(int32 = -1).int32, -1)
assert.fails(lambda: test.M(int32 = 1.0), "got float, want int")  # no implicit int conversion
assert.fails(lambda: test.M(int32 = 0x100000000), "invalid int32: 4294967296")

# uint32
assert.eq(test.M().uint32, 0)
assert.eq(test.M(uint32 = 123).uint32, 123)
assert.fails(lambda: test.M(uint32 = -1), "invalid uint32: -1")

# float
assert.eq(test.M().float, 0.0)
assert.eq(test.M(float = 1.25).float, 1.25)
assert.eq(test.M(float = 1).float, 1)  # implicit int conversion

# string
assert.eq(test.M().string, "")
assert.eq(test.M(string = "hello").string, "hello")
assert.eq(test.M(string = "\xff").string, "\xff")  # invalid UTF-8
assert.fails(lambda: test.M(string = 1), "in field string: got int, want string")

# bytes
assert.eq(str(test.M().bytes), "<0 bytes>")
assert.eq(str(test.M(bytes = "hello").bytes), "<5 bytes>")
assert.fails(lambda: test.M(bytes = 1), "in field bytes: got int, want bytes")

# messages
assert.eq(str(test.M().kv), str(test.KeyValuePair()))
kv = test.KeyValuePair(key = "greeting", value = "hello")

# TODO(adonovan): define and use equality comparisons of messages instead of to_dict.
assert.eq(to_dict(test.M(kv = dict(key = "greeting", value = "hello")).kv), to_dict(kv))
assert.eq(to_dict(test.M(kv = kv).kv), to_dict(kv))
assert.eq(to_dict(test.M(kv = test.KeyValuePair(to_dict(kv))).kv), to_dict(kv))
assert.fails(lambda: test.M(kv = 1), "in field kv: got int, want test.KeyValuePair")

# TODO(adonovan): test identity and aliasing in the above conversions.
assert.fails(lambda: test.M(kv = test.M.Inner()), "in field kv: got test.M.Inner, want test.KeyValuePair")

# enums.  Conversions are similar to EnumDescriptor(x).
assert.eq(test.M().coin, test.Coin.PENNY)
assert.eq(test.M(coin = 10).coin, test.Coin.DIME)
assert.eq(test.M(coin = test.Coin.PENNY).coin, test.Coin.PENNY)
assert.eq(test.M(coin = 1).coin, test.Coin.PENNY)
assert.eq(test.M(coin = "PENNY").coin, test.Coin.PENNY)
assert.fails(lambda: test.M(coin = test.Suit.HEARTS), "in field coin: invalid value Suit.HEARTS for Coin enum")

# Extensions
assert.eq(proto.get_field(m1, test.Outer.ext), 1)
assert.eq(proto.get_field(m1, test.ext), 2)

# An iterable (even a singleton) cannot be converted to a scalar field.
assert.fails(lambda: test.M(int32 = [1]), "in field int32: got list, want int32")

# Non-trivial field defaults.
d = test.D()
assert.eq(d.bool, True)
assert.eq(d.int32, 123)
assert.eq(d.float, 1.25)
assert.eq(d.string, "hello")
assert.eq(d.coin, test.Coin.QUARTER)

# ---- Field assignment ----
#
# All the same conversions apply as in the constructor.

m = test.M()

m.bool = True
assert.eq(m.bool, True)
m.bool = False
assert.eq(m.bool, False)

def set_bool(msg, x):
    msg.bool = x

assert.fails(lambda: set_bool(m, 1), "in field bool: got int, want bool")

# unset:
m.bool = None
assert.true(not proto.has(m, "bool"))
assert.eq(m.bool, False)

m.string = "hello"
assert.eq(m.string, "hello")
m.string = "goodbye"
assert.eq(m.string, "goodbye")

def set_string(msg, x):
    msg.string = x

assert.fails(lambda: set_string(m, 1), "in field string: got int, want string")

# unset:
m.string = None
assert.true(not proto.has(m, "string"))
assert.eq(m.string, "")

def set_kv(msg, x):
    msg.kv = x

assert.fails(lambda: set_kv(m, test.M.Inner()), "in field kv: got test.M.Inner, want test.KeyValuePair")

# and so on...

# -------- Repeated fields --------

# bools
assert.eq(type(test.M().bools), "proto.repeated<bool>")
assert.eq(list(test.M().bools), [])
assert.eq(list(test.M(bools = []).bools), [])
assert.eq(list(test.M(bools = [True, False]).bools), [True, False])
assert.eq(type(test.M(bools = []).bools), "proto.repeated<bool>")

# strings
assert.eq(type(test.M().strings), "proto.repeated<string>")
assert.eq(list(test.M().strings), [])
assert.eq(type(test.M(strings = []).strings), "proto.repeated<string>")
assert.eq(list(test.M(strings = []).strings), [])
assert.eq(list(test.M(strings = ["a", "b"]).strings), ["a", "b"])

# messages (Inner)
assert.eq(type(test.M().inners), "proto.repeated<test.M.Inner>")
assert.eq(list(test.M().inners), [])
assert.eq(str(test.M(inners = [test.M.Inner(int32 = 1)]).inners), "[test.M.Inner(int32=1)]")
assert.eq(test.M().inner.int32, 0)  # zero value
assert.eq(test.M().inner.inner.inner.inner.inner.int32, 0)

# messages (KeyValuePair)
mkv = test.M(kvs = [test.KeyValuePair(key = "greeting", value = "hello")])
assert.eq(len(mkv.kvs), 1)
assert.eq(type(mkv.kvs[0]), "proto.Message")
assert.eq(mkv.kvs[0].descriptor, test.KeyValuePair)
assert.eq(mkv.kvs[0].key, "greeting")
assert.eq(mkv.kvs[0].value, "hello")

# A scalar cannot be converted to a repeated field.
# And Starlark strings are not iterable.
assert.fails(lambda: test.M(int32s = 1), "got int for .int32s field, want iterable")
assert.fails(lambda: test.M(strings = "abc"), "got string for .strings field, want iterable")

# field assignment
mm = test.M(bools = [True, False])
assert.eq(list(mm.bools), [True, False])

def set_bools0(msg, x):
    msg.bools[0] = x

assert.fails(lambda: set_bools0(mm, 1), "setting element of repeated field: got int, want bool")
mm.bools[0] = False
assert.eq(list(mm.bools), [False, False])

# Accessing a repeated field yields a proto.repeated<T>.
# Element updates to it are visible in the original message.
m3 = test.M(strings = ["hello", "world"])
assert.eq(list(m3.strings), ["hello", "world"])
assert.eq(type(m3.strings), "proto.repeated<string>")
m3.strings[0] = "goodbye"
m3strings = m3.strings  # an alias, not a copy
m3strings[1] = "everyone"
assert.eq(list(m3.strings), ["goodbye", "everyone"])
assert.eq(dir(m3.strings), [])  # no append/extend/etc methods, for now
assert.fails(
    lambda: m3.strings + m3.strings,
    "unknown binary op: proto.repeated<string> \\+ proto.repeated<string>",
)  # nor +, for now

# Fields are part of the message, not separate lists that may be shared:
m4 = test.M()
m4.strings = m3strings
m3.strings[0] = "welcome"
assert.eq(m3.strings[0], "welcome")
assert.eq(m4.strings[0], "goodbye")

# But you can't use a RepeatedField of one type with a field of a different type,
# even if the values would be ok.
m5 = test.M(int32s = [1, 2, 3])
assert.eq(type(m5.uint32s), "proto.repeated<uint32>")
assert.eq(list(m5.uint32s), [])

def set_uint32s(msg, x):
    msg.uint32s = x

#assert.fails(lambda: set_uint32s(m5, m5.int32s), "in field uint32s: got repeated<int32>, want repeated<uint32>") # FIXME no error
# You cannot assign None to an element of a repeated field,
# either numeric:
def set_int32s0(msg, x):
    msg.int32s[0] = x

assert.fails(
    lambda: set_int32s0(m5, None),
    "setting element of repeated field: got NoneType, want int32",
)

# or message:
def set_kvs0(msg, x):
    msg.kvs[0] = x

assert.fails(
    lambda: set_kvs0(test.M(kvs = [{}]), None),
    "setting element of repeated field: got NoneType, want test.KeyValuePair",
)

# But storing None to a repeated field unsets it:
m5.int32s = None
assert.eq(type(m5.int32s), "proto.repeated<int32>")
assert.eq(list(m5.int32s), [])

# The default value returned for an unset field is immutable.
def set_default():
    m1.inner.int32 = 1

assert.fails(set_default, "cannot set .int32 field of frozen test.M.Inner message")

# ---- Field assignment ----

# -------- Marshal & unmarshal --------

# TODO(adonovan): test encoding/decoding of all field types.

# marshal, marshal_text
data = proto.marshal(m1)
text = proto.marshal_text(m1)
assert.eq(str(data), "<54 bytes>")
assert.eq(text, '''int32: 12345
string: "hello"
coin: PENNY
strings: "hello"
strings: "world"
kvs: {
  key: "greeting"
  value: "hello"
}
[test.Outer.ext]: 1
[test.ext]: 2
''')

# Invalid UTF-8 strings can be marshalled and unmarshalled.
assert.eq(proto.unmarshal(test.M, proto.marshal(test.M(string = "\xff"))).string, "\377")

# unmarshal returns a message that produces identical str and marshal_text output.
m2 = proto.unmarshal(test.M, data)
assert.eq(str(m2), str(m1))
assert.eq(proto.marshal_text(m1), proto.marshal_text(m2))

assert.fails(lambda: proto.marshal(123), "got int, want proto.Message")
assert.fails(lambda: proto.unmarshal(1, ""), "got int, want proto.MessageDescriptor")
assert.fails(lambda: proto.unmarshal(test.Coin, ""), "got proto.EnumDescriptor")
assert.fails(lambda: proto.unmarshal(test.M, 123), "want bytes")
# TODO(adonovan): add a test of failed unmarshal when we have a way to construct byte strings.
# assert.fails(lambda: proto.unmarshal(test.M, b"not a message"), "unmarshalling test.M failed")

# TODO(adonovan): detect missing required fields during marshal.
assert.eq(str(proto.marshal(test.M())), "<0 bytes>")
assert.eq(str(proto.unmarshal(test.M, data[:0])), "test.M()")

m12 = test.M(int32 = 1, inner = test.M.Inner(int32 = 123))
m12_text = proto.marshal_text(m12)
assert.eq(m12_text, """int32: 1
inner: {
  int32: 123
}
""")

# unmarshal_text
m12b = proto.unmarshal_text(test.M, m12_text)
assert.eq(m12b.int32, 1)
assert.eq(m12b.inner.int32, 123)

assert.eq("""int32: 1
inner: {
  int32: 123
}
""", proto.marshal_text(test.M(int32 = 1, inner = dict(int32 = 123))))

# test of bytes type
# TODO: move to Starlark core.
abytes = proto.marshal(m12)[:4]
assert.eq(type(abytes), "bytes")
assert.eq(str(abytes), "<4 bytes>")
assert.eq(len(abytes), 4)
assert.eq(str(abytes[1:]), "<3 bytes>")
