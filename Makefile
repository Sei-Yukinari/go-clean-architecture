.PHONY: ci
ci:
	circleci config process .circleci/config.yml >.circleci/process.yml ; \
    circleci local execute -c .circleci/process.yml --job  build

.PHONY: cd
cd:
	circleci config process .circleci/config.yml >.circleci/process.yml ; \
    circleci local execute -c .circleci/process.yml --job  deploy