docker run -v /Users/arturo/git/shootr-backend:/ci -d --name java8-ci -t fav24/java8-ci bash

go install && \
ci-github-aws \
 --c test_and_build \
 --git-commit a1234b567 \
 --git-branch develops
