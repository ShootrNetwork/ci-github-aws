#docker kill java8-ci
#docker rm java8-ci

docker run -v /Users/arturo/git/shootr-backend:/ci -d --name java8-ci -t fav24/java8-ci bash

# aws credentials
#export AWS_ACCESS_KEY_ID=
#export AWS_SECRET_ACCESS_KEY=
source env_vars.sh

go install && \
ci-github-aws \
  --c upload_to_s3 \
 --git-commit a1234b567 \
 --git-branch develops

# --c test_and_build \
