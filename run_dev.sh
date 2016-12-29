#docker kill java8-ci
#docker rm java8-ci
#docker run -v /Users/arturo/git/shootr-backend:/ci -d --name java8-ci -t fav24/java8-ci bash
docker run -v /Users/arturo/git/shootr-backend:/ci -d --name java8-ci-slim -t fav24/java8-ci-slim bash

# aws credentials
#export AWS_ACCESS_KEY_ID=
#export AWS_SECRET_ACCESS_KEY=
source env_vars.sh

go install

cd /Users/arturo/git/shootr-backend

ci-github-aws --c test_and_build  --git-commit a1234b567 --git-branch 'random'
#ci-github-aws --c upload_to_s3    --git-commit a1234b567 --git-branch 'dev-'
#ci-github-aws --c docker_build    --git-commit a1234b567 --git-branch 'dev-'
#ci-github-aws --c docker_tag      --git-commit a1234b567 --git-branch 'dev-'
#ci-github-aws --c run_all          --git-commit a1234b567 --git-branch 'dev-' --pull-request=false --pull-request-str=false

cd /Users/arturo/git/go/src/github.com/shootrnetwork/ci-github-aws
# --c test_and_build \
