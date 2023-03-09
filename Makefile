export SHELL=/bin/bash
export GO111MODULE=on


# ------------------------------------------------------------------------------
# 定数定義
# ------------------------------------------------------------------------------

# カラーコード
C_RESET  := \033[m
C_RED    := \033[31m
C_GREEN  := \033[32m
C_YELLOW := \033[33m

# ログレベル
INFO  := printf "${C_GREEN}[INFO]  ${C_RESET}"  && echo -e
WARN  := printf "${C_YELLOW}[WARN]  ${C_RESET}" && echo -e
ERROR := printf "${C_RED}[ERROR] ${C_RESET}"   && echo -e


# ------------------------------------------------------------------------------
# コマンド一覧表示
# ------------------------------------------------------------------------------

.PHONY: list
list:
	@${INFO} 'select the number of the command.';\
		echo '';\
		select SELECT_VAL in $$(cat Makefile | grep -e '.PHONY:' | grep -v 'list' | sed 's!^\.PHONY\: *!!') 'CANCEL';\
		do\
			echo '';\
			if [ "$${SELECT_VAL}" = 'CANCEL' ]; then\
				${INFO} "'CANCEL' selected. abort the process...";\
				exit 0;\
			fi;\
			if [ -z $${SELECT_VAL} ]; then\
				${WARN} 'that selection does not exist. abort the process...';\
				exit 0;\
			fi;\
			echo -e ">>> make $${SELECT_VAL}${C_RESET}";\
			make --no-print-directory "$${SELECT_VAL}";\
			break;\
		done;


# ------------------------------------------------------------------------------
# モジュール更新
# ------------------------------------------------------------------------------

.PHONY: update
update:
	@${INFO} 'start go get...'
	@cat go.mod | awk '/\t.+ v[0-9]+\.[0-9]+\.[a-z0-9\-\+]+$$/ { print $$1 }' | xargs -I {} go get -u -d {}
	@${INFO} 'start go mod tidy...'
	@go env GOVERSION | sed -r 's/^go([0-9]+\.[0-9]+).[0-9]+$$/\1/' | go mod tidy -compat=$$(cat)
	@${INFO} 'completed.'


# ------------------------------------------------------------------------------
# フォーマット
# ------------------------------------------------------------------------------

.PHONY: format
format:
	@${INFO} 'start format...'
	@gofmt -l -s -w ./src/
	@${INFO} 'completed.'


# ------------------------------------------------------------------------------
# 静的解析
# ------------------------------------------------------------------------------

.PHONY: lint
lint:
	@${INFO} 'start lint...'
	@set -a && source .env && set +a; \
		golangci-lint run --out-format=tab --color=always
	@${INFO} 'completed.'


# ------------------------------------------------------------------------------
# ユニットテスト
# ------------------------------------------------------------------------------

.PHONY: test
test:
	@${INFO} 'start test...'
	@rm -rf ./coverage
	@mkdir -p ./coverage
	@go clean -testcache
	@set -a && source .env && set +a; \
		go test ./... -coverprofile ./coverage/cover.out > ./coverage/test.out 2>&1; \
		EXIT_CODE=$$?; \
		cat ./coverage/test.out; \
		go tool cover -html=./coverage/cover.out -o ./coverage/cover.html; \
		go tool cover -func=./coverage/cover.out | grep total | tr -d '\t' | sed -e 's/(statements)/ /g'; \
		exit $$EXIT_CODE
	@${INFO} 'completed.'


# ------------------------------------------------------------------------------
# ビルド
# ------------------------------------------------------------------------------

.PHONY: build
build:
	@${INFO} 'start build...'
	@rm -rf ./build
	@cp -r ./public ./build
	@set -a && source .env && set +a && \
		go build -buildvcs=false -o ./build/case/main.wasm   ./src/case && \
		go build -buildvcs=false -o ./build/format/main.wasm ./src/format && \
		go build -buildvcs=false -o ./build/url/main.wasm    ./src/url && \
		go build -buildvcs=false -o ./build/random/main.wasm ./src/random && \
		go build -buildvcs=false -o ./build/hash/main.wasm   ./src/hash && \
		go build -buildvcs=false -o ./build/base/main.wasm   ./src/base && \
		go build -buildvcs=false -o ./build/jwt/main.wasm    ./src/jwt && \
		go build -buildvcs=false -o ./build/crypto/main.wasm ./src/crypto
	@${INFO} 'completed.'


# ------------------------------------------------------------------------------
# ローカル実行
# ------------------------------------------------------------------------------

.PHONY: start
start:
	@make build
	@${INFO} 'starting the development server...'
	@${INFO} 'you can now view app in the browser. http://localhost:3000'
	@goexec 'http.ListenAndServe(`:3000`, http.FileServer(http.Dir(`./build`)))'


# ------------------------------------------------------------------------------
# クリーン
# ------------------------------------------------------------------------------

.PHONY: clean
clean:
	@${INFO} 'start clean...'
	@rm -rf ./coverage
	@rm -rf ./build
	@${INFO} 'completed.'
