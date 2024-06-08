.PHONY: new
new:
	hz new -I idl -idl idl/user/user.proto

.PHONY: update
update:
	hz update -I idl -idl idl/user/user.proto
