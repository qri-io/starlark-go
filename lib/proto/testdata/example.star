# Example of protocol message processing.

load("assert.star", "assert")

# Load a protocol message descriptor from some file.
# The host application determines how this name resolves to a FileDescriptorProto.
bugtracker = proto.file("testdata/bugtracker/bugtracker.proto")

# FileDescriptor
assert.eq(type(bugtracker), "proto.FileDescriptor")
assert.eq(str(bugtracker), "testdata/bugtracker/bugtracker.proto")
assert.eq(dir(bugtracker), ["AllTypes", "Issue", "KeyValuePair", "Outer", "Status", "Type", "ext"])

# message Descriptor
assert.eq(type(bugtracker.Issue), "proto.MessageDescriptor")
assert.eq(str(bugtracker.Issue), "bugtracker.Issue")
assert.eq(dir(bugtracker.Issue), ["Nested"])

# nested message Descriptor
assert.eq(type(bugtracker.Issue.Nested), "proto.MessageDescriptor")
assert.eq(str(bugtracker.Issue.Nested), "bugtracker.Issue.Nested")
assert.eq(dir(bugtracker.Issue.Nested), [])

# EnumDescriptor
assert.eq(type(bugtracker.Status), "proto.EnumDescriptor")
assert.eq(str(bugtracker.Status), "bugtracker.Status")
assert.eq(dir(bugtracker.Status), ["ACCEPTED", "ASSIGNED", "FIXED", "NEW", "WONTFIX"])

# Calling an EnumDescriptor converts a Starlark value to an enum value.
assert.eq(bugtracker.Status(5), bugtracker.Status.WONTFIX)
assert.eq(bugtracker.Status("WONTFIX"), bugtracker.Status.WONTFIX)

# EnumValueDescriptor
assert.eq(type(bugtracker.Status.ASSIGNED), "proto.EnumValueDescriptor")
assert.eq(str(bugtracker.Status.ASSIGNED), "Status.ASSIGNED")
assert.eq(bugtracker.Status.WONTFIX.number, 5)
assert.eq(bugtracker.Status.WONTFIX.index, 4)
assert.eq(bugtracker.Status.WONTFIX.name, "WONTFIX")
assert.eq(bugtracker.Status.WONTFIX.type, bugtracker.Status)
assert.eq(dir(bugtracker.Status.ASSIGNED), ["index", "name", "number", "type"])

# FieldDescriptor
assert.eq(type(bugtracker.Issue.id), "proto.FieldDescriptor")
assert.eq(str(bugtracker.Issue.id), "bugtracker.Issue.id")
assert.eq(dir(bugtracker.Issue.id), [])

# extension FieldDescriptors
assert.eq(type(bugtracker.Outer.ext), "proto.FieldDescriptor")
assert.eq(str(bugtracker.Outer.ext), "bugtracker.Outer.ext")
assert.eq(type(bugtracker.ext), "proto.FieldDescriptor")
assert.eq(str(bugtracker.ext), "bugtracker.ext")

# Calling a message Descriptor instantiates a message.
issue = bugtracker.Issue(
    id = 12345,
    status = bugtracker.Status.ASSIGNED,
    reporter = "adonovan",
    note = ["A note", "Another note"],
    metadata = [bugtracker.KeyValuePair(key = "greeting", value = "hello")],
)

# Extensions must be set using set_field.
proto.set_field(issue, bugtracker.Outer.ext, 1)
proto.set_field(issue, bugtracker.ext, 2)

# Fields may be updated:
issue.type = bugtracker.Type.FEATURE

# Unset fields act like an immutable default value.
assert.eq(issue.nested.y.y.y.y.y.x, "")
assert.eq(bugtracker.Issue().status, bugtracker.Status.NEW)

# Calling a descriptor with a single positional argument makes a shallow copy.
issue2 = bugtracker.Issue(issue)
assert.eq(str(issue2), str(issue))
issue2.reporter = "nobody"
assert.eq(issue2.reporter, "nobody")
assert.eq(issue.reporter, "adonovan")  # unchanged

# Submessages can be expressed using dicts.
assert.eq(str(bugtracker.Issue(nested = {"x": "x"})), 'bugtracker.Issue(nested=bugtracker.Issue.Nested(x="x"))')

# The str function prints the entire structure.
assert.eq(str(issue), 'bugtracker.Issue(id=12345, status=Status.ASSIGNED, reporter="adonovan", type=Type.FEATURE, note=["A note", "Another note"], metadata=[bugtracker.KeyValuePair(key="greeting", value="hello")], bugtracker.ext=2, bugtracker.Outer.ext=1)')

# marshal, marshal_text encode a message.
data = proto.marshal(issue)
text = proto.marshal_text(issue)
assert.eq(str(data), "<64 bytes>")

# NOTE: prototext.Marshal emits a redundant colon in "metadata: {...}"
assert.eq(text, '''id: 12345
status: ASSIGNED
reporter: "adonovan"
type: FEATURE
note: "A note"
note: "Another note"
metadata: {
  key: "greeting"
  value: "hello"
}
[bugtracker.Outer.ext]: 1
[bugtracker.ext]: 2
''')

# unmarshal, unmarshal_text decode a message.
msg = proto.unmarshal(bugtracker.Issue, data)
assert.eq(str(msg), str(issue))  # messages are equal
