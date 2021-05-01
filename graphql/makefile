q:
	curl localhost:8080/graphql -d '{ "query":"{hello,morning}" }' -H 'Content-Type: application/json'

neo4j:
	docker run \
	--rm \
	--publish=7474:7474 --publish=7687:7687 \
	--volume=${PWD}/neo4j/data:/data \
	--env=NEO4J_AUTH=none \
	neo4j

browse_4j:
	cmd.exe /c start http://localhost:7474/

browse_ql:
	cmd.exe /c start http://localhost:8080/graphql?query={hello}

.PHONY: neo4j