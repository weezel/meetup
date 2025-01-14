marp:
	docker run --rm \
		-e MARP_USER="$(id -u):$(id -g)" \
		-v "$(PWD):/home/marp/app" \
		marpteam/marp-cli \
		--theme-set themes/ \
		 --allow-local-files \
		--pdf \
		--engine-options='{"html":true}' \
		--output /home/marp/app/go102_and_design_patterns.pdf \
		go102_and_design_patterns.md

include ../../../Makefile
