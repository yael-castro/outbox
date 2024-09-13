.PHONY: http relay proto

http:
	bash scripts/build.bash http

relay:
	bash scripts/build.bash relay

proto:
	bash scripts/proto.bash

database:
	bash scripts/database.bash