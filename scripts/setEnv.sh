function setEnv {
    export $(cat .env)
    echo "done"
}